package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
}`

//Konfig config onject
type Konfig struct {
	config map[string]interface{}
}

//New return new config object
func New() *Konfig {
	return &Konfig{}
}

//Load load config to Object
// accept fs path
func (k *Konfig) Load(path string) error {
	/*b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}*/
	var i map[string]interface{}
	err := json.NewDecoder(strings.NewReader(name)).Decode(&i)
	if err != nil {
		return err
	}
	k.config = i
	return nil
}

//GetValueString return config value by key
//accept key and return value
func (k *Konfig) GetValueString(key string) string {
	if v := os.Getenv(key); len(v) > 0 {
		return v
	}
	return k.getValueFromConfig(key)
}

func (k *Konfig) getValueFromEnv(key string) string {
	return os.Getenv(key)
}

func (k *Konfig) getValueFromConfig(key string) string {
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

func main() {
	c := New()
	err := c.Load("./config.default.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c.GetValueString("POLICY.retention"))
}
