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

func TestStringConfigTable(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		key         string
		expectError bool
		want        string
	}{
		{name: "get value json good ip", configPath: "./fixture/config.default.json", key: "SERVER.host", expectError: false, want: "127.0.0.1"},
		{name: "get value json good flat ip", configPath: "./fixture/config.default.flat.json", key: "host", expectError: false, want: "127.0.0.1"},
		{name: "get value json good flat port", configPath: "./fixture/config.default.flat.json", key: "port", expectError: false, want: "8080"},
		{name: "get value json file not exist", configPath: "./fixture/config.default1.json", key: "SERVER.host", expectError: true, want: "127.0.0.1"},
		{name: "get value json bad", configPath: "./fixture/bad.config.default.json", key: "SERVER.host", expectError: true, want: "127.0.0.1"},
		{name: "get value yaml good", configPath: "./fixture/config.default.yaml", key: "SERVER.host", expectError: false, want: "127.0.0.1"},
		{name: "get value yml good", configPath: "./fixture/config.default.yml", key: "SERVER.host", expectError: false, want: "127.0.0.1"},
		{name: "get value properties good", configPath: "./fixture/config.default.properties", key: "SERVER.host", expectError: false, want: "127.0.0.1"},
		{name: "get value properties good", configPath: "./fixture/config.default.ini", key: "SERVER.host", expectError: false, want: "127.0.0.1"},
		{name: "get value yaml bad", configPath: "bad.config.default.yaml", key: "SERVER.host", expectError: true, want: "127.0.0.1"},
		{name: "get non string value", configPath: "./fixture/config.default.json", key: "PARAMS.test", expectError: false, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigStringValue(tt.configPath, tt.key)
			if (err == nil && tt.expectError) || (err != nil && !tt.expectError) {
				t.Errorf("GetConfigStringValue() = %v", err)
				return
			}
			if got != tt.want && !tt.expectError {
				t.Errorf("GetConfigStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntConfigTable(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		key         string
		expectError bool
		want        int
	}{
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.retention", expectError: false, want: 4},
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.test", expectError: false, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigIntValue(tt.configPath, tt.key)
			if (err == nil && tt.expectError) || (err != nil && !tt.expectError) {
				t.Errorf("GetConfigIntValue() = %v", err)
				return
			}
			if got != tt.want && !tt.expectError {
				t.Errorf("GetConfigIntValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileFormatNotSupported(t *testing.T) {
	_, err := GetConfigStringValue("./fixture/config.default.flat.notSupported", "aaa")
	if err == nil {
		t.Errorf("file should not be supported")
	}
}

func TestFloatConfigTable(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		key         string
		expectError bool
		want        float64
	}{
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.partial", expectError: false, want: 4.0},
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.test", expectError: false, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigFloat32Value(tt.configPath, tt.key)
			if (err == nil && tt.expectError) || (err != nil && !tt.expectError) {
				t.Errorf("GetConfigIntValue() = %v", err)
				return
			}
			if got != tt.want && !tt.expectError {
				t.Errorf("GetConfigIntValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestBoolConfigTable(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		key         string
		expectError bool
		want        bool
	}{
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.test", expectError: false, want: true},
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.retention", expectError: false, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigBoolValue(tt.configPath, tt.key)
			if (err == nil && tt.expectError) || (err != nil && !tt.expectError) {
				t.Errorf("GetConfigBoolValue() = %v", err)
				return
			}
			if got != tt.want && !tt.expectError {
				t.Errorf("GetConfigBoolValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringArrayConfigTable(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		key         string
		expectError bool
		want        []string
	}{
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.keys", expectError: false, want: []string{"a", "b"}},
		{name: "get value json good retention", configPath: "./fixture/config.default.json", key: "PARAMS.retention", expectError: false, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfigArrayStringValue(tt.configPath, tt.key)
			if (err == nil && tt.expectError) || (err != nil && !tt.expectError) {
				t.Errorf("GetConfigBoolValue() = %v", err)
				return
			}
			if len(got) != len(tt.want) && !tt.expectError {
				t.Errorf("GetConfigArrayStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func GetConfigStringValue(configPath string, key string) (string, error) {
	c := New()
	err := c.Load(configPath)
	if err != nil {
		return "", err
	}
	return c.GetStringValue(key), nil
}

func GetConfigArrayStringValue(configPath string, key string) ([]string, error) {
	c := New()
	err := c.Load(configPath)
	if err != nil {
		return []string{}, err
	}
	return c.GetStringArrayValue(key), nil
}

func GetConfigIntValue(configPath string, key string) (int, error) {
	c := New()
	err := c.Load(configPath)
	if err != nil {
		return 0, err
	}
	return c.GetIntValue(key), nil
}

func GetConfigFloat32Value(configPath string, key string) (float64, error) {
	c := New()
	err := c.Load(configPath)
	if err != nil {
		return 0, err
	}
	return c.GetFloat64Value(key), nil
}

func GetConfigBoolValue(configPath string, key string) (bool, error) {
	c := New()
	err := c.Load(configPath)
	if err != nil {
		return false, err
	}
	return c.GetBoolValue(key), nil
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
