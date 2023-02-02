package chartutil

import "testing"

func TestLoadChartTest1(t *testing.T) {

	if err := LoadChartTest1(); err != nil {
		t.Fatal(err)
	}
}
