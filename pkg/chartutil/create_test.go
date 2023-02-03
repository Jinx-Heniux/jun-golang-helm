package chartutil

import "testing"

func TestCreateChartFrom(t *testing.T) {

	if err := CreateChartFrom(); err != nil {
		t.Fatal(err)
	}
}
