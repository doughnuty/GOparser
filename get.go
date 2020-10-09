package GOparser

func (yaml Yaml) Get(path ...string) Mod {
	pEmpty := Property{
		Mod: "",
		Val: nil,
	}

	p := pEmpty

	if yaml.Map == nil {
		return pEmpty
	}

	for i, key := range path {
		p = yaml.Map[key]
		if i < len(path)-1 && p.Mod != MAP_MOD {
			p = pEmpty
			break
		} else if i == len(path)-1 {
			return yaml.Map[key]
		} else if p.Mod == MAP_MOD {
			yaml = p.Val.(Yaml)
		}
	}
	return p
}
