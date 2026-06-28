from __future__ import annotations

import importlib.util
import io
import sys
import tempfile
import unittest
from pathlib import Path
from unittest import mock

MODULE_PATH = (
    Path(__file__).resolve().parents[2]
    / "src"
    / "tools"
    / "generate_design_records_package.py"
)
SPEC = importlib.util.spec_from_file_location("generator", MODULE_PATH)
assert SPEC is not None and SPEC.loader is not None
generator = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = generator
SPEC.loader.exec_module(generator)


class GeneratorTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp = tempfile.TemporaryDirectory()
        self.root = Path(self.temp.name)
        self.source = self.root / "source"
        self.destination = self.root / "bin" / "design-records"
        self.source.mkdir(parents=True)

    def tearDown(self) -> None:
        self.temp.cleanup()

    def write_source(self, files: dict[str, str | bytes]) -> None:
        for relative, content in files.items():
            path = self.source / relative
            path.parent.mkdir(parents=True, exist_ok=True)
            if isinstance(content, bytes):
                path.write_bytes(content)
            else:
                path.write_text(content, encoding="utf-8", newline="")

    def package_ids(self) -> dict[str, str]:
        return {
            "index.md": (
                "# Index: Root\n\n"
                "- **id**: `spec:product.design_records`\n"
                "- **status**: accepted\n"
                "- **date**: 2026-06-26\n"
                "- **parent**: root\n\n"
                "See `spec:product.design_records.child`.\n"
            ),
            "child.md": (
                "# Reference: Child\n\n"
                "- **id**: `spec:product.design_records.child`\n"
                "- **status**: accepted\n"
                "- **date**: 2026-06-26\n"
                "- **parent**: `spec:product.design_records`\n"
            ),
        }

    def warning_classes(self, result) -> set[str]:
        return {warning.warning_class for warning in result.warnings}

    def test_rewrite_exact_root_suffix_and_common_locations(self) -> None:
        text = (
            "- **id**: `spec:product.design_records`\n"
            "- **parent**: `spec:product.design_records.parent`\n"
            "| Child | Reference | `spec:product.design_records.child` | x |\n"
            "| `spec:product.design_records.related` | relation |\n"
            "Body `spec:product.design_records.body`.\n"
        )
        rewritten = generator.rewrite_canonical_refs(text)
        self.assertIn("`spec:design_records`", rewritten)
        self.assertIn("`spec:design_records.parent`", rewritten)
        self.assertIn("`spec:design_records.child`", rewritten)
        self.assertIn("`spec:design_records.related`", rewritten)
        self.assertIn("`spec:design_records.body`", rewritten)
        self.assertNotIn("spec:product.design_records", rewritten)

    def test_rewrite_leaves_non_targets_unchanged(self) -> None:
        text = (
            "ordinary product prose\n"
            "PRODUCT-REQ-SPEC-003\n"
            "product/records/spec/design-records/index.md\n"
            "spec:drmcp.design_records_mcp.tools\n"
            "spec:product.design_records-extra\n"
            "Xspec:product.design_records\n"
        )
        self.assertEqual(text, generator.rewrite_canonical_refs(text))

    def test_expected_bytes_preserve_crlf_and_binary(self) -> None:
        source = b"a\r\nspec:product.design_records.child\r\n"
        expected = b"a\r\nspec:design_records.child\r\n"
        self.assertEqual(expected, generator.expected_generated_bytes(source))
        binary = b"\xff\x00spec:product.design_records"
        self.assertEqual(binary, generator.expected_generated_bytes(binary))

    def test_whole_tree_copy_nested_markdown_non_markdown_and_stale_removal(self) -> None:
        self.write_source(
            {
                **self.package_ids(),
                "nested/guide.md": "spec:product.design_records.child\n",
                "assets/data.bin": b"\x00\x01\x02",
            }
        )
        self.destination.mkdir(parents=True)
        (self.destination / "stale.txt").write_text("stale", encoding="utf-8")

        result = generator.generate_package(self.source, self.destination)

        self.assertEqual(0, result.exit_code, result.errors)
        self.assertFalse((self.destination / "stale.txt").exists())
        self.assertTrue((self.destination / "nested/guide.md").is_file())
        self.assertEqual(
            b"\x00\x01\x02",
            (self.destination / "assets/data.bin").read_bytes(),
        )
        self.assertEqual(
            generator._relative_file_set(self.source),
            generator._relative_file_set(self.destination),
        )

    def test_copy_check_detects_missing_file(self) -> None:
        self.write_source({"a.md": "a", "b.md": "b"})
        generated = self.root / "generated"
        generated.mkdir()
        (generated / "a.md").write_text("a", encoding="utf-8")
        with self.assertRaises(generator.OperationalError) as context:
            generator._check_copy_completeness(self.source, generated)
        self.assertIn("missing=b.md", str(context.exception))

    def test_copy_check_detects_unexpected_file(self) -> None:
        self.write_source({"a.md": "a"})
        generated = self.root / "generated"
        generated.mkdir()
        (generated / "a.md").write_text("a", encoding="utf-8")
        (generated / "extra.md").write_text("x", encoding="utf-8")
        with self.assertRaises(generator.OperationalError) as context:
            generator._check_copy_completeness(self.source, generated)
        self.assertIn("unexpected=extra.md", str(context.exception))

    def test_rewrite_integrity_detects_unaccepted_change(self) -> None:
        self.write_source({"a.md": "spec:product.design_records\nunchanged\n"})
        generated = self.root / "generated"
        generated.mkdir()
        (generated / "a.md").write_text(
            "spec:design_records\nchanged\n",
            encoding="utf-8",
        )
        with self.assertRaises(generator.OperationalError):
            generator._check_rewrite_integrity(
                self.source,
                generated,
                frozenset({Path("a.md")}),
            )

    def test_duplicate_warning(self) -> None:
        generated = self.root / "generated"
        generated.mkdir()
        text = "- **id**: `spec:design_records.same`\n"
        (generated / "a.md").write_text(text, encoding="utf-8")
        (generated / "b.md").write_text(text, encoding="utf-8")
        findings = generator._scan_warnings(
            generated,
            frozenset({Path("a.md"), Path("b.md")}),
        )
        self.assertIn(
            "duplicate canonical ref",
            {finding.warning_class for finding in findings},
        )

    def test_unresolved_external_unrewritten_and_boundary_warnings(self) -> None:
        generated = self.root / "generated"
        generated.mkdir()
        (generated / "a.md").write_text(
            "- **id**: `spec:design_records.a`\n"
            "spec:design_records.missing\n"
            "spec:drmcp.tools\n"
            "spec:product.design_records.left\n"
            "Brewprint migration wiring PRODUCT-WORK-SPEC-001\n",
            encoding="utf-8",
        )
        findings = generator._scan_warnings(
            generated,
            frozenset({Path("a.md")}),
        )
        classes = {finding.warning_class for finding in findings}
        self.assertTrue(
            {
                "unresolved internal ref",
                "external canonical ref",
                "unrewritten source ref",
                "source-authoring boundary finding",
            }.issubset(classes)
        )

    def test_semantic_warnings_do_not_fail_generation(self) -> None:
        self.write_source(
            {
                "index.md": (
                    "- **id**: `spec:product.design_records`\n"
                    "spec:drmcp.tools\n"
                    "Brewprint migration PRODUCT-WORK-SPEC-001\n"
                )
            }
        )
        result = generator.generate_package(self.source, self.destination)
        self.assertEqual(0, result.exit_code, result.errors)
        self.assertIn("external canonical ref", self.warning_classes(result))
        self.assertIn(
            "source-authoring boundary finding",
            self.warning_classes(result),
        )

    def test_missing_source_is_operational_failure(self) -> None:
        result = generator.generate_package(
            self.root / "missing",
            self.destination,
        )
        self.assertEqual(1, result.exit_code)
        self.assertTrue(result.errors)

    def test_temporary_tree_creation_failure(self) -> None:
        self.write_source({"a.md": "a"})
        with mock.patch.object(
            generator,
            "_create_unique_temporary_tree",
            side_effect=OSError("no temp"),
        ):
            result = generator.generate_package(self.source, self.destination)
        self.assertEqual(1, result.exit_code)
        self.assertIn("create temporary tree", result.errors[0])

    def test_copy_failure_preserves_existing_destination(self) -> None:
        self.write_source({"a.md": "a"})
        self.destination.mkdir(parents=True)
        old = self.destination / "old.md"
        old.write_text("old", encoding="utf-8")
        with mock.patch.object(
            generator,
            "_copy_source_files",
            side_effect=OSError("copy"),
        ):
            result = generator.generate_package(self.source, self.destination)
        self.assertEqual(1, result.exit_code)
        self.assertEqual("old", old.read_text(encoding="utf-8"))

    def test_rewrite_failure(self) -> None:
        self.write_source({"a.md": "a"})
        with mock.patch.object(
            generator,
            "_rewrite_tree",
            side_effect=OSError("rewrite"),
        ):
            result = generator.generate_package(self.source, self.destination)
        self.assertEqual(1, result.exit_code)
        self.assertIn("rewrite canonical refs", result.errors[0])

    def test_prepublication_check_failure(self) -> None:
        self.write_source({"a.md": "a"})
        with mock.patch.object(
            generator,
            "_check_rewrite_integrity",
            side_effect=generator.OperationalError(
                "check rewrite integrity",
                self.source,
                "bad",
            ),
        ):
            result = generator.generate_package(self.source, self.destination)
        self.assertEqual(1, result.exit_code)
        self.assertFalse(self.destination.exists())

    def test_warning_scan_operational_failure(self) -> None:
        self.write_source({"a.md": "a"})
        with mock.patch.object(
            generator,
            "_scan_warnings",
            side_effect=OSError("scan"),
        ):
            result = generator.generate_package(self.source, self.destination)
        self.assertEqual(1, result.exit_code)
        self.assertIn("scan semantic warnings", result.errors[0])

    def test_backup_move_failure_preserves_existing_destination(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")

        with mock.patch.object(
            generator,
            "_move_path",
            side_effect=OSError("backup"),
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertIn("backup destination", result.errors[0])
        self.assertEqual(
            "old",
            (self.destination / "old.md").read_text(encoding="utf-8"),
        )
        self.assertFalse((self.destination / "new.md").exists())

    def test_publish_failure_rolls_back_existing_destination(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        real_move = generator._move_path
        calls = 0

        def fail_publish(source: Path, destination: Path) -> None:
            nonlocal calls
            calls += 1
            if calls == 2:
                raise OSError("publish")
            real_move(source, destination)

        with mock.patch.object(
            generator,
            "_move_path",
            side_effect=fail_publish,
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertEqual(
            "old",
            (self.destination / "old.md").read_text(encoding="utf-8"),
        )
        self.assertFalse((self.destination / "new.md").exists())

    def test_rollback_failure_retains_backup(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        real_move = generator._move_path
        calls = 0

        def fail_publish_and_rollback(source: Path, destination: Path) -> None:
            nonlocal calls
            calls += 1
            if calls in {2, 3}:
                raise OSError(f"move {calls}")
            real_move(source, destination)

        with mock.patch.object(
            generator,
            "_move_path",
            side_effect=fail_publish_and_rollback,
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertTrue(
            any("rollback failed" in error for error in result.errors)
        )
        backups = list(
            self.destination.parent.glob(".design-records.backup-*")
        )
        self.assertEqual(1, len(backups))
        self.assertTrue((backups[0] / "old.md").is_file())

    def test_publication_confirmation_failure_restores_previous_destination(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        with mock.patch.object(
            generator,
            "_confirm_publication",
            side_effect=generator.OperationalError(
                "confirm publication",
                self.destination,
                "cannot enumerate",
            ),
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertEqual(
            "old",
            (self.destination / "old.md").read_text(encoding="utf-8"),
        )
        self.assertFalse((self.destination / "new.md").exists())

    def test_publication_confirmation_rollback_failure_retains_recovery_artifacts(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        real_move = generator._move_path
        calls = 0

        def fail_restore(source: Path, destination: Path) -> None:
            nonlocal calls
            calls += 1
            if calls == 4:
                raise OSError("restore")
            real_move(source, destination)

        with mock.patch.object(
            generator,
            "_confirm_publication",
            side_effect=generator.OperationalError(
                "confirm publication",
                self.destination,
                "cannot enumerate",
            ),
        ), mock.patch.object(
            generator,
            "_move_path",
            side_effect=fail_restore,
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertTrue(
            any("rollback failed" in error for error in result.errors)
        )
        backups = list(
            self.destination.parent.glob(".design-records.backup-*")
        )
        failed_publications = list(
            self.destination.parent.glob(".design-records.failed-*")
        )
        self.assertEqual(1, len(backups))
        self.assertEqual(1, len(failed_publications))
        self.assertTrue((backups[0] / "old.md").is_file())
        self.assertTrue((failed_publications[0] / "new.md").is_file())
        self.assertFalse(self.destination.exists())

    def test_publication_confirmation_failure_without_backup_removes_new_tree(self) -> None:
        self.write_source({"new.md": "new"})
        with mock.patch.object(
            generator,
            "_confirm_publication",
            side_effect=generator.OperationalError(
                "confirm publication",
                self.destination,
                "cannot enumerate",
            ),
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(1, result.exit_code)
        self.assertFalse(self.destination.exists())

    def test_success_cleans_temporary_and_backup_artifacts(self) -> None:
        self.write_source(self.package_ids())
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        result = generator.generate_package(self.source, self.destination)
        self.assertEqual(0, result.exit_code, result.errors)
        siblings = [path.name for path in self.destination.parent.iterdir()]
        self.assertEqual(["design-records"], siblings)

    def test_cleanup_only_failure_after_valid_publication_is_warning(self) -> None:
        self.write_source({"new.md": "new"})
        self.destination.mkdir(parents=True)
        (self.destination / "old.md").write_text("old", encoding="utf-8")
        real_remove = generator._remove_tree

        def fail_backup_cleanup(path: Path) -> None:
            if ".backup-" in path.name:
                raise OSError("cleanup")
            real_remove(path)

        with mock.patch.object(
            generator,
            "_remove_tree",
            side_effect=fail_backup_cleanup,
        ):
            result = generator.generate_package(self.source, self.destination)

        self.assertEqual(0, result.exit_code, result.errors)
        self.assertIn("cleanup", self.warning_classes(result))
        self.assertTrue((self.destination / "new.md").is_file())

    def test_public_cli_has_no_path_override(self) -> None:
        with mock.patch.object(sys, "stderr", new=io.StringIO()) as stderr:
            exit_code = generator.main(["--source", "other"])
        self.assertEqual(1, exit_code)
        self.assertIn(
            "accepts no source or destination",
            stderr.getvalue(),
        )

    def test_repository_root_uses_script_location(self) -> None:
        script = (
            self.root
            / "repo"
            / "product"
            / "src"
            / "tools"
            / "generate_design_records_package.py"
        )
        self.assertEqual(
            (self.root / "repo").resolve(),
            generator.resolve_repository_root(script),
        )

    def test_no_drmcp_or_registry_dependency(self) -> None:
        source = MODULE_PATH.read_text(encoding="utf-8")
        self.assertNotIn("import drmcp", source)
        self.assertNotIn("app_registry", source)
        self.assertNotIn("Path.cwd()", source)


if __name__ == "__main__":
    unittest.main()
