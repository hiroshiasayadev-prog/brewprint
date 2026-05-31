# reference

ResolvedProject上のsemantic object間の直接参照。

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| source | yaml/mcp/model/reference.yaml |

### Fields

| field | type | note |
|---|---|---|
| kind | str | enum値集合: param_model / return_model / produces_asset / consumes_asset / reads / writes / store_of / field_type / field_fk / transition_event / transition_from / transition_to / transition_action / event_payload / event_actor / event_watches / scenario_state_file / scenario_step_transition |
| direction | str | enum値集合: out / in。query対象から見た方向 |
| from | object_ref | 参照元object |
| to | object_ref | 参照先object |
| source | source_location | このreferenceの定義元。任意 |
| doc | str | reference補足。任意 |

