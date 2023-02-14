package common

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func UnmarshalYamlBytes(bytes []byte, out interface{}) {
	err := yaml.Unmarshal(bytes, out)
	Check(err)
}

func UnmarshalYaml(file string, out interface{}) {
	ymlFile, err := ioutil.ReadFile(file)

	if err != nil {
		log.Printf("yamlFile.Get err   #%v\n", err)
	}
	ymlFile = []byte(os.ExpandEnv(string(ymlFile)))
	err = yaml.Unmarshal(ymlFile, out)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
