package simple

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if l := len(c.config); l != 0 {
		t.Fatal(fmt.Sprintf("%d not equal to %d", l, 0))
	}
}

func TestConfig_GetValueString(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
	if value := c.GetStringValue("PARAMS.retention"); value != "4" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "4"))
	}
	if value := c.GetStringValue("PARAMS.test"); value != "true" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "true"))
	}
}

func TestConfig_LoadJson(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}

func TestConfig_LoadJsonError(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default1.json")
	if err == nil {
		t.Fatal(err)
	}
}

func TestConfig_LoadJsonBadJson(t *testing.T) {
	c := New()
	err := c.Load("./fixture/bad.config.default.json")
	if err == nil {
		t.Fatal(err)
	}
}

func TestConfig_LoadJsonBadYaml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/bad.config.default.yaml")
	if err == nil {
		t.Fatal(err)
	}
}

func TestConfig_LoadYaml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}
func TestConfig_LoadYml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}

func TestConfig_LoadProperties(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.properties")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}

func TestConfig_LoadIni(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.ini")
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}

func TestConfig_LoadEnv(t *testing.T) {
	err := os.Setenv("SERVER.host", "127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}
	c := New()
	err = c.Load()
	if err != nil {
		t.Fatal(err)
	}
	if value := c.GetStringValue("SERVER.host"); value != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("%s not equal to %s", value, "127.0.0.1"))
	}
}
