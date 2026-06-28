from __future__ import annotations

import ast
import importlib.util
import os
import re
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

REPOSITORY_ROOT = Path(__file__).resolve().parents[3]
MODULE_PATH = (
    REPOSITORY_ROOT
    / "product"
    / "src"
    / "tools"
    / "generate_design_records_package.py"
)
VERIFY_PATH = REPOSITORY_ROOT / "scripts" / "verify.bat"

SPEC = importlib.util.spec_from_file_location("generator_release", MODULE_PATH)
assert SPEC is not None and SPEC.loader is not None
generator = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = generator
SPEC.loader.exec_module(generator)


class ReleaseFixtureTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp = tempfile.TemporaryDirectory()
        self.root = Path(self.temp.name)

    def tearDown(self) -> None:
        self.temp.cleanup()

    def write_files(
        self,
        root: Path,
        files: dict[str, str | bytes],
    ) -> None:
        for relative, content in files.items():
            path = root / relative
            path.parent.mkdir(parents=True, exist_ok=True)
            data = content.encode("utf-8") if isinstance(content, str) else content
            path.write_bytes(data)

    def snapshot(self, root: Path) -> dict[str, bytes]:
        return {
            path.relative_to(root).as_posix(): path.read_bytes()
            for path in sorted(root.rglob("*"))
            if path.is_file()
        }

    def warning_semantics(self, finding) -> tuple[
        str,
        str,
        tuple[str, ...],
        tuple[str, ...],
        tuple[str, ...],
    ]:
        refs = tuple(
            sorted(set(generator.CANONICAL_REF_RE.findall(finding.detail)))
        )
        declaring_paths = tuple(
            sorted(
                set(
                    re.findall(
                        r"(?<![A-Za-z0-9_./-])"
                        r"([A-Za-z0-9_./-]+\.md)"
                        r"(?![A-Za-z0-9_./-])",
                        finding.detail,
                    )
                )
            )
        )
        boundary_groups = tuple(
            sorted(
                label
                for label, _pattern in generator.BOUNDARY_PATTERNS
                if label in finding.detail
            )
        )
        return (
            finding.warning_class,
            finding.path,
            refs,
            declaring_paths,
            boundary_groups,
        )

    def create_fake_repository(
        self,
        files: dict[str, str | bytes] | None,
    ) -> tuple[Path, Path]:
        repository = self.root / "fake-repository"
        script = (
            repository
            / "product"
            / "src"
            / "tools"
            / "generate_design_records_package.py"
        )
        script.parent.mkdir(parents=True, exist_ok=True)
        script.write_bytes(MODULE_PATH.read_bytes())

        if files is not None:
            source = (
                repository
                / "product"
                / "records"
                / "spec"
                / "design-records"
            )
            source.mkdir(parents=True, exist_ok=True)
            self.write_files(source, files)

        return repository, script

    def run_public_generator(
        self,
        script: Path,
        cwd: Path,
    ) -> subprocess.CompletedProcess[str]:
        environment = os.environ.copy()
        environment["BREWPRINT_REPO_ROOT"] = str(self.root / "wrong-root")
        environment["DRMCP_ROOT"] = str(self.root / "wrong-drmcp")
        environment["APP_REGISTRY"] = str(self.root / "wrong-registry")
        return subprocess.run(
            [sys.executable, "-I", str(script)],
            cwd=cwd,
            env=environment,
            capture_output=True,
            text=True,
            encoding="utf-8",
            errors="replace",
            timeout=30,
            check=False,
        )

    def test_repeated_generation_is_byte_reproducible_and_replaces_stale_output(
        self,
    ) -> None:
        source = self.root / "source"
        destination = self.root / "bin" / "design-records"
        source.mkdir(parents=True)
        self.write_files(
            source,
            {
                "index.md": (
                    "# Index: Root\n\n"
                    "- **id**: `spec:product.design_records`\n"
                    "- **status**: accepted\n"
                    "- **date**: 2026-06-28\n\n"
                    "See `spec:product.design_records.child`.\n"
                    "See `spec:drmcp.tools`.\n"
                    "Brewprint migration PRODUCT-WORK-SPEC-001.\n"
                ),
                "child.md": (
                    "# Reference: Child\n\n"
                    "- **id**: `spec:product.design_records.child`\n"
                    "- **status**: accepted\n"
                    "- **date**: 2026-06-28\n"
                    "- **parent**: `spec:product.design_records`\n"
                ),
                "assets/data.bin": b"\x00\xff\x01",
            },
        )

        first = generator.generate_package(source, destination)
        first_snapshot = self.snapshot(destination)

        self.assertEqual(0, first.exit_code, first.errors)
        self.assertTrue(first.warnings)

        (destination / "stale.txt").write_text("stale", encoding="utf-8")
        (destination / "child.md").write_text("tampered", encoding="utf-8")

        second = generator.generate_package(source, destination)
        second_snapshot = self.snapshot(destination)

        self.assertEqual(0, second.exit_code, second.errors)
        self.assertEqual(first_snapshot, second_snapshot)
        self.assertEqual(first.warnings, second.warnings)
        self.assertEqual(
            {"assets/data.bin", "child.md", "index.md"},
            set(second_snapshot),
        )
        self.assertNotIn("stale.txt", second_snapshot)

        forbidden_names = (
            ".design-records.tmp-",
            ".design-records.backup-",
            ".design-records.failed-",
        )
        for relative, data in second_snapshot.items():
            for forbidden in forbidden_names:
                self.assertNotIn(forbidden, relative)
                self.assertNotIn(forbidden.encode("ascii"), data)

    def test_prefix_rewrite_fixture_preserves_non_targets_and_bytes(self) -> None:
        source = self.root / "source"
        destination = self.root / "bin" / "design-records"
        source.mkdir(parents=True)

        source_markdown = (
            b"# Reference: Fixture\r\n\r\n"
            b"- **id**: `spec:product.design_records`\r\n"
            b"- **parent**: `spec:product.design_records.parent`\r\n"
            b"| Child | Reference | `spec:product.design_records.child` | x |\r\n"
            b"| `spec:product.design_records.related` | relation |\r\n"
            b"Body `spec:product.design_records.body`.\r\n"
            b"ordinary product prose\r\n"
            b"PRODUCT-REQ-SPEC-003\r\n"
            b"product/records/spec/design-records/index.md\r\n"
            b"spec:drmcp.tools\r\n"
            b"spec:product.design_records-extra\r\n"
            b"Xspec:product.design_records\r\n"
        )
        expected_markdown = (
            b"# Reference: Fixture\r\n\r\n"
            b"- **id**: `spec:design_records`\r\n"
            b"- **parent**: `spec:design_records.parent`\r\n"
            b"| Child | Reference | `spec:design_records.child` | x |\r\n"
            b"| `spec:design_records.related` | relation |\r\n"
            b"Body `spec:design_records.body`.\r\n"
            b"ordinary product prose\r\n"
            b"PRODUCT-REQ-SPEC-003\r\n"
            b"product/records/spec/design-records/index.md\r\n"
            b"spec:drmcp.tools\r\n"
            b"spec:product.design_records-extra\r\n"
            b"Xspec:product.design_records\r\n"
        )
        binary = b"\xff\x00spec:product.design_records\r\n"
        lf_only = b"alpha\nbeta\n"

        self.write_files(
            source,
            {
                "fixture.md": source_markdown,
                "assets/raw.bin": binary,
                "assets/lf.txt": lf_only,
            },
        )

        result = generator.generate_package(source, destination)

        self.assertEqual(0, result.exit_code, result.errors)
        self.assertEqual(expected_markdown, (destination / "fixture.md").read_bytes())
        self.assertEqual(binary, (destination / "assets/raw.bin").read_bytes())
        self.assertEqual(lf_only, (destination / "assets/lf.txt").read_bytes())

    def test_warning_fixture_has_exact_shape_order_and_no_content_change(
        self,
    ) -> None:
        generated = self.root / "generated"
        generated.mkdir()
        self.write_files(
            generated,
            {
                "a.md": (
                    "- **id**: `spec:design_records.same`\n"
                    "spec:design_records.missing\n"
                    "spec:drmcp.tools\n"
                    "spec:product.design_records.left\n"
                    "Brewprint migration wiring PRODUCT-WORK-SPEC-001\n"
                ),
                "b.md": "- **id**: `spec:design_records.same`\n",
                "ok.md": "- **id**: `spec:design_records.ok`\n",
            },
        )
        files = frozenset({Path("a.md"), Path("b.md"), Path("ok.md")})
        before = self.snapshot(generated)

        first = generator._scan_warnings(generated, files)
        second = generator._scan_warnings(generated, files)

        expected_semantics = (
            (
                "duplicate canonical ref",
                "a.md",
                ("spec:design_records.same",),
                ("a.md", "b.md"),
                (),
            ),
            (
                "duplicate canonical ref",
                "b.md",
                ("spec:design_records.same",),
                ("a.md", "b.md"),
                (),
            ),
            (
                "external canonical ref",
                "a.md",
                ("spec:drmcp.tools",),
                (),
                (),
            ),
            (
                "external canonical ref",
                "a.md",
                ("spec:product.design_records.left",),
                (),
                (),
            ),
            (
                "source-authoring boundary finding",
                "a.md",
                (),
                (),
                ("app-local", "migration", "project-tracking", "wiring"),
            ),
            (
                "unresolved internal ref",
                "a.md",
                ("spec:design_records.missing",),
                (),
                (),
            ),
            (
                "unrewritten source ref",
                "a.md",
                ("spec:product.design_records.left",),
                (),
                (),
            ),
        )

        self.assertEqual(
            expected_semantics,
            tuple(self.warning_semantics(finding) for finding in first),
        )
        self.assertEqual(first, second)
        self.assertEqual(before, self.snapshot(generated))

    def test_duplicate_and_unresolved_defects_are_localized_and_nonblocking(
        self,
    ) -> None:
        source = self.root / "source"
        destination = self.root / "bin" / "design-records"
        source.mkdir(parents=True)
        self.write_files(
            source,
            {
                "a.md": (
                    "- **id**: `spec:product.design_records.same`\n"
                    "spec:product.design_records.missing\n"
                ),
                "b.md": "- **id**: `spec:product.design_records.same`\n",
                "ok.md": (
                    "- **id**: `spec:product.design_records.ok`\n"
                    "unrelated content\n"
                ),
            },
        )

        result = generator.generate_package(source, destination)

        expected_warning_semantics = (
            (
                "duplicate canonical ref",
                "a.md",
                ("spec:design_records.same",),
                ("a.md", "b.md"),
                (),
            ),
            (
                "duplicate canonical ref",
                "b.md",
                ("spec:design_records.same",),
                ("a.md", "b.md"),
                (),
            ),
            (
                "unresolved internal ref",
                "a.md",
                ("spec:design_records.missing",),
                (),
                (),
            ),
        )

        self.assertEqual(0, result.exit_code, result.errors)
        self.assertEqual(
            expected_warning_semantics,
            tuple(
                self.warning_semantics(finding)
                for finding in result.warnings
            ),
        )
        self.assertEqual(
            {"a.md", "b.md", "ok.md"},
            set(self.snapshot(destination)),
        )
        self.assertIn(
            b"spec:design_records.missing",
            (destination / "a.md").read_bytes(),
        )
        self.assertEqual(
            b"- **id**: `spec:design_records.same`\n",
            (destination / "b.md").read_bytes(),
        )
        self.assertEqual(
            b"- **id**: `spec:design_records.ok`\nunrelated content\n",
            (destination / "ok.md").read_bytes(),
        )

    def test_public_command_runs_from_external_cwd_with_fixed_paths(self) -> None:
        repository, script = self.create_fake_repository(
            {
                "index.md": (
                    "- **id**: `spec:product.design_records`\n"
                    "spec:drmcp.tools\n"
                )
            }
        )
        outside = self.root / "outside"
        outside.mkdir()

        completed = self.run_public_generator(script, outside)

        destination = repository / "bin" / "design-records"
        self.assertEqual(0, completed.returncode, completed.stderr)
        self.assertTrue((destination / "index.md").is_file())
        self.assertFalse((outside / "bin" / "design-records").exists())
        self.assertIn(
            "Generated portable Design Records package: 1 file(s)",
            completed.stdout,
        )
        self.assertIn("WARNING [external canonical ref] index.md", completed.stderr)
        self.assertIn(
            b"spec:design_records",
            (destination / "index.md").read_bytes(),
        )

    def test_public_command_operational_failure_returns_one(self) -> None:
        _repository, script = self.create_fake_repository(None)
        outside = self.root / "outside"
        outside.mkdir()

        completed = self.run_public_generator(script, outside)

        self.assertEqual(1, completed.returncode)
        self.assertIn("ERROR ", completed.stderr)
        self.assertIn("path does not exist", completed.stderr)
        self.assertEqual("", completed.stdout)

    def test_generator_uses_only_standard_library_and_no_host_state_inputs(
        self,
    ) -> None:
        source = MODULE_PATH.read_text(encoding="utf-8")
        tree = ast.parse(source)
        imported_roots: set[str] = set()

        for node in ast.walk(tree):
            if isinstance(node, ast.Import):
                imported_roots.update(
                    alias.name.split(".", 1)[0] for alias in node.names
                )
            elif isinstance(node, ast.ImportFrom) and node.module:
                imported_roots.add(node.module.split(".", 1)[0])

        standard_library = set(sys.stdlib_module_names) | {"__future__"}
        self.assertLessEqual(imported_roots, standard_library)
        self.assertTrue(
            imported_roots.isdisjoint({"socket", "urllib", "http", "requests"})
        )
        self.assertNotIn("os.environ", source)
        self.assertNotIn("os.getenv", source)
        self.assertNotIn("Path.cwd()", source)
        self.assertNotIn("app_registry", source)
        self.assertNotIn("import drmcp", source)

    def test_verify_batch_discovers_release_tests_and_propagates_failure(
        self,
    ) -> None:
        verify = VERIFY_PATH.read_text(encoding="utf-8").replace("\r\n", "\n")
        self.assertIn(
            'python -X utf8 -m unittest discover -s product\\tests\\tools '
            '-p "test_*.py" -v\n'
            "if errorlevel 1 exit /b 1",
            verify,
        )
        self.assertIn(
            "python -X utf8 product\\src\\tools\\"
            "generate_design_records_package.py\n"
            "if errorlevel 1 exit /b 1",
            verify,
        )


if __name__ == "__main__":
    unittest.main()
