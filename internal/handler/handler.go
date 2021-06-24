package handler

import (
	"bulbasaur/internal/config"
	clientHandler "bulbasaur/internal/handler/httpclient"
	"bulbasaur/internal/resources"
	"bulbasaur/internal/shared/constants"
	"github.com/cucumber/godog"
)

type handler struct {
	resources           map[string]resources.Resource
	httpClientResources map[string]resources.ClientResource
}

func New() *handler {
	return &handler{
		resources:           make(map[string]resources.Resource),
		httpClientResources: make(map[string]resources.ClientResource),
	}
}

func (h *handler) RegisterScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(func(_ *godog.Scenario) {
		h.reset()
	})

	clientHandler.New(h.httpClientResources).InitializeScenario(s)
}

func (h *handler) RegisterResource(cfg *config.Resource, res resources.Resource) {
	h.resources[cfg.Name] = res

	switch cfg.Type {
	case constants.ResHttpClient:
		h.httpClientResources[cfg.Name] = res.(resources.ClientResource)
	}

}

func (h *handler) reset() {
	for _, resource := range h.resources {
		resource.Reset()
	}
}
