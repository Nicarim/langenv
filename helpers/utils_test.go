package helpers

import "testing"

func TestFilterListForMinorVersions(t *testing.T) {
	list := []string {
		"v0.0.1",
		"v0.0.2",
		"v0.0.3",
		"v1.0.1",
		"v1.0.2",
		"v2.1.1",
		"v2.1.2",
		"v2.2.0",
		"v2.3.0",
	}
	expectedList := []string {
		"v0.0.3",
		"v1.0.2",
		"v2.1.2",
		"v2.2.0",
		"v2.3.0",
	}
	newList := FilterListForMinorVersions(&list)
	if len(newList) != len(expectedList) {
		t.Errorf("Lists are inequal size.")
	}
	for i, element := range newList {
		if expectedList[i] != element {
			t.Errorf("Invalid array output, expected %s, got %s", expectedList[i], element)
		}
	}


}