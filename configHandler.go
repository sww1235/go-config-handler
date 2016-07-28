package configHandler

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

//Configuration stores the configuration that is read in and out from a file
type Configuration interface{}

//ReadConfigJSON reads a JSON encoded configuration file into a
//Configuration struct
func ReadConfigJSON(filename string, configType interface{}) (Configuration, error) {
	readConfigTypeTemp := reflect.TypeOf(configType)
	readConfigType := readConfigTypeTemp.Kind()
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return readConfigType, err
	}
	var config readConfigType
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return readConfigType, err
	}
	return config, nil
}

//WriteConfigJSON writes a configuration struct into the specified file
func WriteConfigJSON(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}
