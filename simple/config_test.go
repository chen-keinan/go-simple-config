package simple

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	assert.Equal(t, len(c.config), 0)
}

func TestConfig_GetValueString(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
	assert.Equal(t, c.GetStringValue("PARAMS.retention"), "4")
	assert.Equal(t, c.GetStringValue("PARAMS.test"), "true")
}

func TestConfig_LoadJson(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}

func TestConfig_LoadJsonError(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default1.json")
	assert.Error(t, err)
}

func TestConfig_LoadJsonBadJson(t *testing.T) {
	c := New()
	err := c.Load("./fixture/bad.config.default.json")
	assert.Error(t, err)
}

func TestConfig_LoadJsonBadYaml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/bad.config.default.yaml")
	assert.Error(t, err)
}

func TestConfig_LoadYaml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}
func TestConfig_LoadYml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}

func TestConfig_LoadProperties(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.properties")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}

func TestConfig_LoadIni(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.ini")
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}

func TestConfig_LoadEnv(t *testing.T) {
	err := os.Setenv("SERVER.host", "127.0.0.1")
	assert.NoError(t, err)
	c := New()
	err = c.Load()
	assert.NoError(t, err)
	assert.Equal(t, c.GetStringValue("SERVER.host"), "127.0.0.1")
}
