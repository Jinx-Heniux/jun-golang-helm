package chartutil

import "testing"

func TestWriteDependencies(t *testing.T) {

	if err := WriteDependencies(); err != nil {
		t.Fatal(err)
	}
}
