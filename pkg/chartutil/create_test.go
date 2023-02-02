package chartutil

import "testing"

func TestCreateFromTest1(t *testing.T) {

	if err := CreateFromTest1(); err != nil {
		t.Fatal(err)
	}
}
