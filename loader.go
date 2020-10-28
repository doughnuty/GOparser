package GOparser

import "errors"

type Source struct {
	mod        sourceMod
	properties interface{}
}

type sourceMod int

const (
	empty sourceMod = iota
	yam
	env
)

func NewYamlSource(path ...string) Source {
	return Source{mod: yam, properties: path}
}

func NewEnvSource(pref ...Prefix) Source {
	return Source{mod: env, properties: pref}
}

func (yaml *Yaml) Load(sources ...Source) error {
	var err error
	for _, s := range sources {
		if s.mod == yam {
			err = yaml.parseFiles(s.properties.([]string))
		} else if s.mod == env {
			err = yaml.parseDotEnv(s.properties.([]Prefix))
		} else if s.mod == empty {
			err = errors.New("empty source")
		}
	}

	return err
}
