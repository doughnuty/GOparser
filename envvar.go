package GOparser

import (
	"errors"
	"os"
	"strings"
)

type Prefix struct {
	prefix   string
	stripped string
}

func concat(opts []string) string {
	ans := ""
	for _, j := range opts {
		ans += j + "_"
	}
	return ans
}

func WithPrefix(opts ...string) Prefix {
	p := concat(opts)

	return Prefix{
		prefix:   p,
		stripped: "",
	}
}

func WithStrippedPrefix(opts ...string) Prefix {
	p := concat(opts)

	return Prefix{
		prefix:   "",
		stripped: p,
	}
}

// parse .env file and save to the structure
func (yaml *Yaml) parseDotEnv(pref []Prefix) error {
	for _, env := range os.Environ() {
		if len(pref) > 0 {
			notFound := true
			if _, ok := matchPrefix(pref, env); ok {
				notFound = false
			}

			if match, ok := matchPrefix(pref, env); ok {
				env = strings.TrimPrefix(env, match)
				notFound = false
			}

			if notFound {
				continue
			}
		}

		keys, val := getValFromEnv(env)
		err := yaml.FillYaml(keys, val)
		if err != nil {
			return err
		}
	}

	return nil
}

// compare env with each prefix
// return
func matchPrefix(pref []Prefix, env string) (string, bool) {
	for _, p := range pref {
		if len(p.prefix) > 0 {
			if strings.HasPrefix(env, p.prefix) {
				return "", true
			}
		}

		if len(p.stripped) > 0 {
			if strings.HasPrefix(env, p.stripped) {
				return p.stripped, true
			}
		}
	}

	return "", false
}

func getValFromEnv(env string) ([]string, string) {
	pair := strings.SplitN(env, "=", 2)
	if len(pair) != 2 { // not sure that this is needed
		return nil, ""
	}

	val := pair[1]
	keys := strings.Split(strings.ToLower(pair[0]), "_")

	return keys, val
}

func (yaml *Yaml) FillYaml(keys []string, val string) error {
	//slc := make([]Yaml, 0, 10)

	for i, key := range keys {
		var newYaml Yaml

		if yaml.Map[key].mod != MAP_MOD && i != len(keys)-1 { // which one - empty or map? what to do if
			// if there's two envs A_B = X && A = X ?
			newYaml = NewYaml()
			if newYaml.Map == nil {
				return errors.New("error initializing new yaml")
			}
			yaml.Map[key] = Property{
				mod: MAP_MOD,
				val: newYaml,
			}
		} else if i == len(keys)-1 {
			/*if yaml.Map[key].mod != EMPTY_MOD {
				return errors.New("reassignment of a value is prohibited")
			}*/

			if yaml.Map == nil {
				return errors.New("assignment to a nil map is prohibited")
			}

			yaml.Map[key] = Property{
				mod: VAL_MOD,
				val: val,
			}
			return nil

		} else if yaml.Map[key].mod == MAP_MOD {
			newYaml = yaml.Map[key].val.(Yaml)

			yaml.Map[key] = Property{
				mod: MAP_MOD,
				val: newYaml,
			}
		}

		yaml = &newYaml
	}

	return nil
}
