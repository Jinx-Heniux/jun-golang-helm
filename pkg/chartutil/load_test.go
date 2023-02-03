package chartutil

import (
	"testing"
)

func TestLoadChartFromFS(t *testing.T) {

	src := "/home/zhs2si/projects/go/jun-golang-helm/testdata/sourcecharts/bci-umbrella-chart"

	schart, err := LoadChartFromFS(src)
	if err != nil {
		t.Fatal(err)
	}
	_ = schart
}
