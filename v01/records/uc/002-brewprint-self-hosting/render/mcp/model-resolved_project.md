# resolved_project

QueryService / MCP wrapper が読む解決済みproject contextの概念model。
実際の保持先は store kind: context の resolved_project_store として表す想定。
Raw YAML ASTではなく、validation / name resolution / derived model build / index build後のsemantic modelを指す。
内部indexや任意mapはv1 modelでは厳密表現できないためanyで暫定表現する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/resolved_project.yaml |

### Fields

| field | type | note |
|---|---|---|
| semantic_objects | any | 解決済みnode / view / transition / asset / field / file等のregistry。<br/>具体的なindex構造はGo実装内部のResolvedProject責務であり、MCP公開contractでは直接公開しない。 |
| reference_indexes | any | referencesBySource / referencesByTarget等のsemantic reference index。<br/>v1 modelでは任意のmap / union value / 実装内部index shapeを表せないためany。 |
| render_context | any | render_index.yaml解決後のgroup / output mapping等。<br/>analyze_impact のrender_output_files等の材料になるが、公開schemaとしては固定しない。 |
| diagnostics | diagnostic_list | semantic build時に得られたdiagnostic。任意 |

