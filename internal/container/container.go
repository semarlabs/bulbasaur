package container

import (
	"bulbasaur/internal/config"
	"bulbasaur/internal/handler"
	"bulbasaur/internal/reports"
	"bulbasaur/internal/resources"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
)

func Init(cfgPath string) {
	cfg, err := config.ReadConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	newHandler := handler.New()
	for _, resourceCfg := range cfg.Resources {
		resource, err := resources.Create(resourceCfg)
		if err != nil {
			panic(err)
		}

		newHandler.RegisterResource(resourceCfg, resource)
	}

	godog := godog.TestSuite{
		Name: "bulbasaur",
		Options: &godog.Options{
			Format:        "progress",
			Output:        colors.Colored(reports.New()),
			Paths:         cfg.FeaturesPaths,
			StopOnFailure: cfg.StopOnFailure,
		},
		ScenarioInitializer: newHandler.RegisterScenario,
	}

	status := godog.Run()
	if status != 0 {
		fmt.Println("Test Failed")
	}

	os.Exit(status)
}
