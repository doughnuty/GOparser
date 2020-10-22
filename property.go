package GOparser

type Yaml struct {
	Map     map[string]Property
	Spacing int
}

type Property struct {
	Mod PropertyMod
	Val interface{}
}

type PropertyMod int

const (
	EMPTY_MOD PropertyMod = iota
	VAL_MOD
	MAP_MOD
	ARR_MOD
)
