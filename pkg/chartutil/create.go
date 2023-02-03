package chartutil

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
)

func CreateChartFrom() error {

	cf := &chart.Metadata{
		APIVersion: chart.APIVersionV1,
		Name:       "foo2",
		Version:    "0.1.0",
	}

	destdir := "testdata/outputs/"
	// srcdir := "testdata/sourcecharts/mariner"
	srcdir := "testdata/sourcecharts/boilerplate-based-on-utility-toolkit"

	chartutil.CreateFrom(cf, destdir, srcdir)
	if err := chartutil.CreateFrom(cf, destdir, srcdir); err != nil {
		return err
	}

	return nil
}
