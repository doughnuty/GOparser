package GOparser

func (yaml Yaml) Get(path ...string) Mod {
	pEmpty := Property{
		mod: EMPTY_MOD,
		val: nil,
	}

	p := pEmpty

	if yaml.Map == nil {
		return pEmpty
	}

	for i, key := range path {
		p = yaml.Map[key]
		if i < len(path)-1 && p.mod != MAP_MOD {
			p = pEmpty
			break
		} else if i == len(path)-1 {
			return yaml.Map[key]
		} else if p.mod == MAP_MOD {
			yaml = p.val.(Yaml)
		}
	}
	return p
}
