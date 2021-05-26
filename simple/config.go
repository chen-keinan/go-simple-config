package simple

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//Config config oblject
type Config struct {
	config map[string]interface{}
}

//New return new config object
func New() *Config {
	return &Config{config: make(map[string]interface{})}
}

//Load load config to Object
// accept fs path
func (k *Config) Load(path ...string) error {
	if len(path) == 0 {
		return nil
	}
	b, err := ioutil.ReadFile(path[0])
	if err != nil {
		return err
	}
	fileExtension := filepath.Ext(path[0])
	switch fileExtension {
	case ".yaml", ".yml":
		return k.ParseYaml(b)
	case ".json":
		return k.ParseJSON(b)
	case ".properties":
		return k.ParseProperties(b)
	default:
		return k.ParseJSON(b)
	}
}

//ParseJSON parse json file to map[string]interface
//accept path to json file and return error
func (k *Config) ParseJSON(b []byte) error {
	err := json.NewDecoder(strings.NewReader(string(b))).Decode(&k.config)
	if err != nil {
		return err
	}
	return nil
}

//ParseYaml parse json file to map[string]interface
//accept path to yaml file and return error
func (k *Config) ParseYaml(b []byte) error {
	err := yaml.NewDecoder(strings.NewReader(string(b))).Decode(&k.config)
	if err != nil {
		return err
	}
	return nil
}

//ParseProperties parse properties file to map[string]interface
//accept path to properties file and return error
func (k *Config) ParseProperties(b []byte) error {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				keys := strings.Split(key, ".")
				var r map[string]interface{}
				tempMap := k.config
				var p interface{}
				for i := 0; i < len(keys)-1; i++ {
					if _, ok := tempMap[keys[i]]; !ok {
						p = make(map[string]interface{})
					} else {
						p = tempMap[keys[i]]
					}
					tempMap[keys[i]] = p
				}
				r, ok := p.(map[string]interface{})
				if !ok {
					return fmt.Errorf("failed to parse properties file")
				}
				fkey := keys[len(keys)-1]
				r[fkey] = value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

//GetValueString return config value by key
//accept key and return value
func (k *Config) GetValueString(key string) string {
	if v := os.Getenv(key); len(v) > 0 {
		return v
	}
	return k.getValueFromConfig(key)
}

//getValueFromConfig return value by key from config file
// accept key and return value
func (k *Config) getValueFromConfig(key string) string {
	keys := strings.Split(key, ".")
	tempMap := k.config
	for _, ck := range keys {
		if v, ok := tempMap[ck]; ok {
			if nm, ok := v.(map[string]interface{}); ok {
				tempMap = nm
				continue
			}
			if s, ok := v.(string); ok {
				return s
			}
			if s, ok := v.(float64); ok {
				return strconv.Itoa(int(s))
			}
			if s, ok := v.(bool); ok {
				return strconv.FormatBool(s)
			}
		}
	}
	return ""
}
