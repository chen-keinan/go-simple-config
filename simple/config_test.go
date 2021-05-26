package simple

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_GetValueString(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	assert.NoError(t, err)
	assert.Equal(t, c.GetValueString("SERVER.host"), "127.0.0.1")
	assert.Equal(t, c.GetValueString("PARAMS.retention"), "4")
	assert.Equal(t, c.GetValueString("PARAMS.test"), "true")
}

func TestConfig_LoadJson(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.json")
	assert.NoError(t, err)
	assert.Equal(t, c.GetValueString("SERVER.host"), "127.0.0.1")
}
func TestConfig_LoadYaml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	assert.NoError(t, err)
	assert.Equal(t, c.GetValueString("SERVER.host"), "127.0.0.1")
}
func TestConfig_LoadYml(t *testing.T) {
	c := New()
	err := c.Load("./fixture/config.default.yaml")
	assert.NoError(t, err)
	assert.Equal(t, c.GetValueString("SERVER.host"), "127.0.0.1")
}

func TestConfig_LoadEnv(t *testing.T) {
	err := os.Setenv("SERVER.host", "127.0.0.1")
	assert.NoError(t, err)
	c := New()
	err = c.Load()
	assert.NoError(t, err)
	assert.Equal(t, c.GetValueString("SERVER.host"), "127.0.0.1")
}
