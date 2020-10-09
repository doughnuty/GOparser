package GOparser

type Yaml struct {
	Map     map[string]Property
	Spacing int
}

type Property struct {
	Mod string
	Val interface{}
}

const (
	VAL_MOD string = "value"
	MAP_MOD string = "map"
	ARR_MOD string = "array"
)
