package parser

import (
	"GOparser"
	"reflect"
	"testing"
	"time"
)

func TestErrors(t *testing.T) {
	yaml := GOparser.NewYaml()

	err := yaml.Parse("file.yaml")
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

	//Bool(def bool) bool
	tempBool := yaml.Get("student", "stats").Bool(false)
	if reflect.TypeOf(tempBool) != reflect.TypeOf(false) {
		t.Errorf("Conflicting types. Expected %s found %v", "bool", reflect.TypeOf(tempBool))
	}
	if tempBool != false {
		t.Errorf("Bad value. Expected %s found %v", "false", tempBool)
	}

	//Int(def int) int
	tempInt := yaml.Get("student", "personal", "HoursActive").Int(0)
	if reflect.TypeOf(tempInt) != reflect.TypeOf(1) {
		t.Errorf("Conflicting types. Expected %s found %v", "int", reflect.TypeOf(tempInt))
	}
	if tempInt != 0 {
		t.Errorf("Bad value. Expected %s found %v", "0", tempInt)
	}

	//String(def string) string
	tempStr := yaml.Get("student", "personal").String("Name")
	if reflect.TypeOf(tempStr) != reflect.TypeOf("Name") {
		t.Errorf("Conflicting types. Expected %s found %v", "string", reflect.TypeOf(tempStr))
	}
	if tempStr != "Name" {
		t.Errorf("Bad value. Expected %s found %v", "Name", tempStr)
	}

	//Float64(def float64) float64
	tempFlt := yaml.Get("student", "personal", "").Float64(0.0)
	if reflect.TypeOf(tempFlt) != reflect.TypeOf(0.0) {
		t.Errorf("Conflicting types. Expected %s found %v", "float64", reflect.TypeOf(tempFlt))
	}
	if tempFlt != 0.0 {
		t.Errorf("Bad value. Expected %s found %v", "0.0", tempFlt)
	}

	//Duration(def time.Duration) time.Duration
	testDur, _ := time.ParseDuration("0")
	tempDur := yaml.Get("student", "personal", "ID").Duration(testDur)
	if reflect.TypeOf(tempDur) != reflect.TypeOf(testDur) {
		t.Errorf("Conflicting types. Expected %s found %v", "duration", reflect.TypeOf(tempDur))
	}
	if tempDur != testDur {
		t.Errorf("Bad value. Expected %v found %v", testDur, tempDur)
	}

	//StringSlice(def []string) []string
	testSlcSize := 3
	testSlc := make([]string, testSlcSize)
	testSlc[0] = "Exhibition"
	testSlc[1] = "New Album"
	tempSlc := yaml.Get("student", "clubs", "Art").StringSlice(nil)
	if reflect.TypeOf(tempSlc) != reflect.TypeOf(testSlc) {
		t.Errorf("Conflicting types. Expected %s found %v", "string slice", reflect.TypeOf(tempSlc))
	}
	if tempSlc != nil {
		t.Errorf("Successful parse of erroneous path")
	}

	//StringMap(def map[string]string) map[string]string
	testMap := make(map[string]string, 2)
	testMap["ID"] = "123456"
	testMap["Name"] = "B"
	tempMap := yaml.Get("worker").StringMap(nil)
	if reflect.TypeOf(tempMap) != reflect.TypeOf(testMap) {
		t.Errorf("Conflicting types. Expected %s found %v", "map", reflect.TypeOf(tempMap))
	}
	if tempMap != nil {
		t.Errorf("Successful parse of erroneous path")

		t.Errorf("Parsed as %v", yaml.Get("worker"))
	}

	//Bytes() []byte
	var testBts []byte
	tempBts := yaml.Get("").Bytes()
	if reflect.TypeOf(tempBts) != reflect.TypeOf(testBts) {
		t.Errorf("Conflicting types. Expected %s found %v", "", reflect.TypeOf(tempBts))
	}
}
