package chartutil

import (
	"fmt"

	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
)

func CreateFromTest1() {

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
		fmt.Println(err)
	}
}

func CreateFromTest2(chartfile *chart.Metadata, dest, src string) error {

	schart, err := loader.Load(src)
	if err != nil {
		return errors.Wrapf(err, "could not load %s", src)
	}

	schartName := schart.Name()
	schartAPIVersion := schart.AppVersion()
	schartFullPath := schart.ChartFullPath()
	schartPath := schart.ChartPath()

	fmt.Printf("Source Chart Name: %s\n", schartName)
	fmt.Printf("Source Chart AppVersion: %s\n", schartAPIVersion)
	fmt.Printf("Source Chart ChartFullPath: %s\n", schartFullPath)
	fmt.Printf("Source Chart ChartPath: %s\n", schartPath)

	fmt.Println()

	schartMetadata := schart.Metadata
	fmt.Printf("Source Chart Metadata: %v\n", schartMetadata)
	fmt.Printf("\tSource Chart Name: %s\n", schartMetadata.Name)
	fmt.Printf("\tSource Chart APIVersion: %s\n", schartMetadata.APIVersion)
	fmt.Printf("\tSource Chart Description: %s\n", schartMetadata.Description)
	fmt.Printf("\tSource Chart Version: %s\n", schartMetadata.Version)
	fmt.Printf("\tSource Chart Type: %s\n", schartMetadata.Type)

	return nil
}
