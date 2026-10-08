from __future__ import annotations

import ast
import unittest
from pathlib import Path


SERVER_PATH = Path(__file__).resolve().parents[2] / "usdm-mcp" / "server.py"


class McpWrapperHierarchyTests(unittest.TestCase):
    def test_scope_coverage_exposes_and_forwards_include_warnings(self) -> None:
        module = ast.parse(SERVER_PATH.read_text(encoding="utf-8"))
        function = next(
            node
            for node in module.body
            if isinstance(node, ast.FunctionDef)
            and node.name == "check_usdm_scope_coverage"
        )

        argument_names = [argument.arg for argument in function.args.args]
        self.assertIn("include_warnings", argument_names)

        default_by_name = dict(
            zip(
                argument_names[-len(function.args.defaults) :],
                function.args.defaults,
                strict=True,
            )
        )
        self.assertIsInstance(default_by_name["include_warnings"], ast.Constant)
        self.assertIs(default_by_name["include_warnings"].value, True)

        delegated_call = next(
            node
            for node in ast.walk(function)
            if isinstance(node, ast.Call)
            and isinstance(node.func, ast.Attribute)
            and node.func.attr == "check_usdm_scope_coverage"
        )
        keyword_by_name = {
            keyword.arg: keyword.value
            for keyword in delegated_call.keywords
            if keyword.arg is not None
        }
        forwarded = keyword_by_name["include_warnings"]
        self.assertIsInstance(forwarded, ast.Name)
        self.assertEqual(forwarded.id, "include_warnings")


if __name__ == "__main__":
    unittest.main()
