package chartutil

import (
	"helm.sh/helm/v3/pkg/action"
)

func WriteDependenciesTest1() error {
	actionConfig := new(action.Configuration)

	// os.Stdout
	client := action.NewShowWithConfig(action.ShowAll, actionConfig)

	_ = client

	return nil
}
