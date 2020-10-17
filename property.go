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
	VAL_MOD PropertyMod = iota
	MAP_MOD
	ARR_MOD
	EMPTY_MOD
)
