package GOparser

func NewYaml() Yaml {
	return Yaml{Map: make(map[string]Property), Spacing: 0}
}
