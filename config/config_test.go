package config

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	yodaTest := GetEnv("TEST_YODA")
	if yodaTest != "Tests you should do" {
		t.Fatal(yodaTest + "Check your config.json and config.go file")
	}
}
