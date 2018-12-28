package config

import "testing"

func TestJsonConfiguration_GetString(t *testing.T) {
	cfg := createTestConfig()

	str, err := cfg.GetString("strName")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if str == "" {
		t.Error("str was empty, expected hello")
		t.Fail()
		return
	}

	if str != "hello" {
		t.Error("str was " + str + ", expected hello")
		t.Fail()
		return
	}
}

func TestJsonConfiguration_GetStringLayered(t *testing.T) {
	cfg := createTestConfig()

	str, err := cfg.GetString("mapper.someKey")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if str == "" {
		t.Error("str was empty, expected hello")
		t.Fail()
		return
	}

	if str != "someValue" {
		t.Error("str was " + str + ", expected someValue")
		t.Fail()
		return
	}
}

func TestJsonConfiguration_GetInt(t *testing.T) {
	cfg := createTestConfig()

	str, err := cfg.GetInt("number")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if str == 0 {
		t.Error("str was 0, expected hello")
		t.Fail()
		return
	}
}

func createTestConfig() Configuration {
	mapping := make(map[string]interface{})

	mapping["number"] = 10
	mapping["boolean"] = true
	mapping["strName"] = "hello"

	strList := make([]string, 2)
	strList[0] = "test123"
	strList[1] = "456data"
	mapping["strList"] = strList

	mapper := make(map[string]string)
	mapper["someKey"] = "someValue"
	mapper["another"] = "something"
	mapping["mapper"] = mapper

	newMap1 := make(map[string]string)
	newMap2 := make(map[string]string)

	newMap1["key"] = "value"
	newMap2["place"] = "birth"
	innerMap := make(map[string]map[string]string)
	innerMap["extra1"] = newMap1
	innerMap["happy2"] = newMap2
	mapping["innerMap"] = innerMap

	return &jsonConfiguration{mapping: mapping}
}
