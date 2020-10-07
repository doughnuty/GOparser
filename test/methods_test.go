package parser

import (
	"reflect"
	"testing"
	"theRealParser/parser"
	"time"
)

func TestMethods(t *testing.T) {
	yaml := parser.NewYaml()

	err := yaml.Parse("file.yaml")
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

	//Bool(def bool) bool
	tempBool := yaml.Get("student", "stats", "hasTime").Bool(false)
	if reflect.TypeOf(tempBool) != reflect.TypeOf(false) {
		t.Errorf("Conflicting types. Expected %s found %v", "bool", reflect.TypeOf(tempBool))
	}
	if tempBool != true {
		t.Errorf("Bad value. Expected %s found %v", "true", tempBool)
	}

	//Int(def int) int
	tempInt := yaml.Get("student", "personal", "ID").Int(0)
	if reflect.TypeOf(tempInt) != reflect.TypeOf(1) {
		t.Errorf("Conflicting types. Expected %s found %v", "int", reflect.TypeOf(tempInt))
	}
	if tempInt != 654321 {
		t.Errorf("Bad value. Expected %s found %v", "654321", tempInt)
	}

	//String(def string) string
	tempStr := yaml.Get("student", "personal", "Name").String("Name")
	if reflect.TypeOf(tempStr) != reflect.TypeOf("Name") {
		t.Errorf("Conflicting types. Expected %s found %v", "string", reflect.TypeOf(tempStr))
	}
	if tempStr != "A" {
		t.Errorf("Bad value. Expected %s found %v", "A", tempStr)
	}

	//Float64(def float64) float64
	tempFlt := yaml.Get("student", "personal", "HeightHistory").Float64(0.0)
	if reflect.TypeOf(tempFlt) != reflect.TypeOf(0.0) {
		t.Errorf("Conflicting types. Expected %s found %v", "float64", reflect.TypeOf(tempFlt))
	}
	if tempFlt != 103.6 {
		t.Errorf("Bad value. Expected %s found %v", "103.6", tempFlt)
	}

	//Duration(def time.Duration) time.Duration
	testDur, _ := time.ParseDuration("0")
	tempDur := yaml.Get("student", "personal", "HoursActive").Duration(testDur)
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
	tempSlc := yaml.Get("student", "clubs", "Art", "To_Do").StringSlice(nil)
	if reflect.TypeOf(tempSlc) != reflect.TypeOf(testSlc) {
		t.Errorf("Conflicting types. Expected %s found %v", "string slice", reflect.TypeOf(tempSlc))
	}
	if tempSlc == nil {
		t.Errorf("Unseccessful parse")
	}
	for i := range tempSlc {
		if len(testSlc) < len(tempSlc) {
			t.Errorf("Bad value. Expected %s found %v", "", tempSlc)
		}
		if tempSlc[i] != testSlc[i] {
			t.Errorf("Bad value. Expected %s found %v", testSlc[i], tempSlc[i])
		}
	}

	//StringMap(def map[string]string) map[string]string
	testMap := make(map[string]string, 2)
	testMap["ID"] = "123456"
	testMap["Name"] = "B"
	tempMap := yaml.Get("worker", "personal").StringMap(nil)
	if reflect.TypeOf(tempMap) != reflect.TypeOf(testMap) {
		t.Errorf("Conflicting types. Expected %s found %v", "map", reflect.TypeOf(tempMap))
	}
	if tempMap == nil {
		t.Errorf("Unseccessful parse")

		t.Errorf("Parsed as %v", yaml.Get("worker"))
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
	tempBts := yaml.Get("worker", "stats").Bytes()
	if reflect.TypeOf(tempBts) != reflect.TypeOf(testBts) {
		t.Errorf("Conflicting types. Expected %s found %v", "", reflect.TypeOf(tempBts))
	}
}
