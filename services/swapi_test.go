package swapi

import (
	"strings"
	"testing"
)

func TestGetRequest(t *testing.T) {
	expected := "terrain"
	expected2 := "rotation_period"
	result := GetRequest("planets/?search=Hoth")
	if !strings.Contains(result, expected) || !strings.Contains(result, expected2) {
		t.Errorf("Returned unexpected body check API")
	}

}
