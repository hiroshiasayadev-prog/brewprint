# analyze_impact_request

MCP tool `analyze_impact` のrequest model。
docs/spec/mcp/tools/analyze-impact.md §2 のInput、および §3 のchange discriminated objectに対応する。
`change.kind` ごとのpayload制約:
- rename: new_id必須
- remove: 追加payload不可
- change_type: new_type必須
- change_contract: note任意、payloadなしでも有効
- change_transition_target: new_to / new_action の少なくとも一方が必要
- add: added_id必須
kindとpayloadの不正な組み合わせは invalid_change_payload tool error。
v1 modelではoptional / enum / discriminated object / payload相関を型として表せないため、changeはany + noteで保持する。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/analyze_impact_request.yaml |

### Fields

| field | type | note |
|---|---|---|
| selector | object_selector | 必須。影響分析対象のObject selector。 |
| change | any | 必須。kindをdiscriminatorとする変更種別object。rename/remove/change_type/change_contract/change_transition_target/add。v1 modelではdiscriminated objectを表せないためany。 |
| scope_modules | any | 任意。分析範囲を絞るmodule list。文字列配列だが専用list modelを作らずanyで暫定表現する。 |
| max_impacts | int | 任意。impact返却上限。省略時200。 |

