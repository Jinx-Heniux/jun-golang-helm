package chartutil

import (
	"fmt"

	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func LoadChartTest1() error {
	src := "testdata/sourcecharts/bci-umbrella-chart"

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
