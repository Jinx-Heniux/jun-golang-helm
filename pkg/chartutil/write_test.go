package chartutil

import "testing"

func TestWriteDependenciesTest1(t *testing.T) {

	if err := WriteDependenciesTest1(); err != nil {
		t.Fatal(err)
	}
}
