package GOparser

type Yaml struct {
	Map     map[string]Property
	Spacing int
}

type Property struct {
	mod propertyMod
	val interface{}
}

type propertyMod int

const (
	EMPTY_MOD propertyMod = iota
	VAL_MOD
	MAP_MOD
	ARR_MOD
)
