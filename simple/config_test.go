package simple

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_GetValueString(t *testing.T) {

}

func TestConfig_Load(t *testing.T) {
	c:=New()
	c.Load("./fixture/config.default.json")
	assert.Equal(t)
}
