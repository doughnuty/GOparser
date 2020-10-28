package parser

import (
	parser "github.com/doughnuty/GOparser"
	"reflect"
	"testing"
	"time"
)

func TestCorrect(t *testing.T) {

	config := parser.NewYaml()

	yaml := parser.NewYamlSource("file.yaml")
	env := parser.NewEnvSource(parser.WithPrefix("CGO"))

	err := config.Load(yaml, env)
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

	//Bool(def bool) bool
	tempBool := config.Get("student", "stats", "hasTime").Bool(false)
	if reflect.TypeOf(tempBool) != reflect.TypeOf(false) {
		t.Errorf("Conflicting types. Expected %s found %v", "bool", reflect.TypeOf(tempBool))
	}
	if tempBool != true {
		t.Errorf("Bad value. Expected %s found %v", "true", tempBool)
	}

	//Int(def int) int
	tempInt := config.Get("student", "personal", "ID").Int(0)
	if reflect.TypeOf(tempInt) != reflect.TypeOf(1) {
		t.Errorf("Conflicting types. Expected %s found %v", "int", reflect.TypeOf(tempInt))
	}
	if tempInt != 654321 {
		t.Errorf("Bad value. Expected %s found %v", "654321", tempInt)
	}

	//String(def string) string
	tempStr := config.Get("student", "personal", "Name").String("Name")
	if reflect.TypeOf(tempStr) != reflect.TypeOf("Name") {
		t.Errorf("Conflicting types. Expected %s found %v", "string", reflect.TypeOf(tempStr))
	}
	if tempStr != "M A" {
		t.Errorf("Bad value. Expected %s found %v", "M A", tempStr)
	}

	//Float64(def float64) float64
	tempFlt := config.Get("student", "personal", "HeightHistory").Float64(0.0)
	if reflect.TypeOf(tempFlt) != reflect.TypeOf(0.0) {
		t.Errorf("Conflicting types. Expected %s found %v", "float64", reflect.TypeOf(tempFlt))
	}
	if tempFlt != 103.6 {
		t.Errorf("Bad value. Expected %s found %v", "103.6", tempFlt)
	}

	//Duration(def time.Duration) time.Duration
	testDur, _ := time.ParseDuration("0")
	tempDur := config.Get("student", "personal", "HoursActive").Duration(testDur)
	if reflect.TypeOf(tempDur) != reflect.TypeOf(testDur) {
		t.Errorf("Conflicting types. Expected %s found %v", "duration", reflect.TypeOf(tempDur))
	}
	if tempDur == testDur {
		t.Errorf("Bad value. Expected %s found %v", "1h2m3s", tempDur)
	}

	//StringSlice(def []string) []string
	testSlcSize := 3
	testSlc := make([]string, testSlcSize)
	testSlc[0] = "Exhibition"
	testSlc[1] = "New Album"
	tempSlc := config.Get("student", "clubs", "Art", "To_Do").StringSlice(nil)
	if reflect.TypeOf(tempSlc) != reflect.TypeOf(testSlc) {
		t.Errorf("Conflicting types. Expected %s found %v", "string slice", reflect.TypeOf(tempSlc))
	}
	if tempSlc != nil {
		for i := range tempSlc {
			if len(testSlc) < len(tempSlc) {
				t.Errorf("Bad value. Expected %s found %v", "", tempSlc)
				return
			}
			if tempSlc[i] != testSlc[i] {
				t.Errorf("Bad value. Expected %s found %v", testSlc[i], tempSlc[i])
			}
		}
	} else {
		t.Errorf("Unseccessful parse")
	}

	tempSlc = config.Get("worker", "pets").StringSlice(nil)
	if tempSlc != nil {
		if tempSlc[0] != "dog" || tempSlc[1] != "cat" {
			t.Errorf("Bad value. Temp slice is %v", tempSlc)
		}
	} else {
		t.Errorf("Unseccessful parse")
	}

	//StringMap(def map[string]string) map[string]string
	testMap := make(map[string]string, 2)
	testMap["ID"] = "123456"
	testMap["Name"] = "B"
	tempMap := config.Get("worker", "personal").StringMap(nil)
	if reflect.TypeOf(tempMap) != reflect.TypeOf(testMap) {
		t.Errorf("Conflicting types. Expected %s found %v", "map", reflect.TypeOf(tempMap))
	}
	if tempMap == nil {
		t.Errorf("Unseccessful parse")

		t.Errorf("Parsed as %v", config.Get("worker"))
	}
	for i := range tempMap {
		if len(testMap) < len(tempMap) {
			t.Errorf("Bad value. Expected %s found %v", testMap[i], tempMap)
		}
		if tempMap[i] != testMap[i] {
			t.Errorf("Bad value. Expected %s found %v", testMap[i], tempMap[i])
		}
	}

	//Bytes() []byte
	var testBts []byte
	tempBts := config.Get("noone").Bytes()
	if reflect.TypeOf(tempBts) != reflect.TypeOf(testBts) {
		t.Errorf("Conflicting types. Expected %s found %v", "", reflect.TypeOf(tempBts))
	}
}
