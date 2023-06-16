## Functions

### run

▸ **run**(`src`: *string*, `args`: *...string*): *any*

run Go code `src` and take `args` as rest parameters for `Run` function.
`args` slice can be accessed with index with conversion to target type i.e. `s[0].(string)`.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `src` | *string* | Go code. |
| `args` | *...string* | rest parameters for `Run` function. |

**Returns:** *any*

Return value as defined in `src`

___