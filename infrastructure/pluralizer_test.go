package infrastructure

import "testing"

func TestThatThePluralOfCatIsCats(t *testing.T) {
	result := PluralOf("cat", 2)
	expected := "cats"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
