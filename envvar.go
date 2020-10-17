package GOparser

import (
	"errors"
	"fmt"
	"os"
)

// Parse given environmental variables and save to the structure
func (yaml *Yaml) ParseEnv(env ...string) error {
	var err error = nil
	errstr := "no value for key"
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

// parse .env file and save to the structure
/*func (yaml *Yaml) ParseDotEnv(env string) error {

}*/
