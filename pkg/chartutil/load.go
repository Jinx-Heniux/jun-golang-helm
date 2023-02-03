package chartutil

import (
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func LoadChartFromFS(src string) (*chart.Chart, error) {
	chart, err := loader.Load(src)
	if err != nil {
		return nil, errors.Wrapf(err, "could not load %s", src)
	}
	return chart, nil
}
