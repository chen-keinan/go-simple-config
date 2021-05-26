package simple

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
const name = `{

  "POSTGRES": {
    "host": "127.0.0.1",
    "port": "5432",
    "db_name": "postgres",
    "user": "postgres",
    "password": "mysecretpassword"
  },
  "POLICY": {
    "batch_size": "50",
    "retention": 4
  }
}
*/

//Config config onject
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
		return k.ParseJson(b)
	default:
		return k.ParseJson(b)
	}
}

func (k *Config) ParseJson(b []byte) error {
	err := json.NewDecoder(strings.NewReader(string(b))).Decode(&k.config)
	if err != nil {
		return err
	}
	return nil
}
func (k *Config) ParseYaml(b []byte) error {
	err := yaml.NewDecoder(strings.NewReader(string(b))).Decode(&k.config)
	if err != nil {
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

func (k *Config) getValueFromEnv(key string) string {
	return os.Getenv(key)
}

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
