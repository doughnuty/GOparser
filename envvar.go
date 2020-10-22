package GOparser

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Prefix struct {
	prefix   string
	stripped string
}

// Parse given environmental variables and save to the structure
func (yaml *Yaml) ParseEnv(env ...string) error {
	var err error = nil
	errstr := "no value for the key"
	for _, key := range env {
		val, ok := os.LookupEnv(key)
		if ok {
			yaml.Map[key] = Property{
				Mod: VAL_MOD,
				Val: val,
			}
		} else {
			errstr = fmt.Sprintf("%s %s", errstr, key)
			err = errors.New(errstr)
		}
	}
	return err
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
func (yaml *Yaml) ParseDotEnv(pref ...Prefix) error {
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
	newYaml := NewYaml()

	for i, key := range keys {
		if yaml.Map[key].Mod == EMPTY_MOD && i != len(keys)-1 {
			yaml.Map[key] = Property{
				Mod: MAP_MOD,
				Val: newYaml,
			}
		} else if i == len(keys)-1 {
			if yaml.Map[key].Mod != EMPTY_MOD {
				return errors.New("reassignment of a value is prohibited")
			}

			yaml.Map[key] = Property{
				Mod: VAL_MOD,
				Val: val,
			}

		} else if yaml.Map[key].Mod == MAP_MOD {
			newYaml = yaml.Map[key].Val.(Yaml)
		}

		yaml = &newYaml
	}

	return nil
}
