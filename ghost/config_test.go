package ghost

import (
	"testing"
)

// Test config with empty parameters
func TestConfigEmptyParameters(t *testing.T) {
	config := Config{
		User:     "",
		Password: "",
		URL:      "",
	}

	if _, err := config.Client(); err == nil {
		t.Fatalf("expected error, but got nil")
	}
}

// Test config with invalid url
func TestConfigInvalidUrl(t *testing.T) {
	config := Config{
		User:     "myuser",
		Password: "mypwd",
		URL:      "invalid.url",
	}

	if _, err := config.Client(); err == nil {
		t.Fatalf("expected error, but got nil")
	}
}

// Test config with valid parameters
func TestConfigValidParameters(t *testing.T) {
	config := Config{
		User:     "myuser",
		Password: "mypwd",
		URL:      "https://www.valid.url",
	}

	if _, err := config.Client(); err != nil {
		t.Fatalf("expected no error, but got %s", err)
	}
}
