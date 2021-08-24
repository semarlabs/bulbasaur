package httpclient

import (
	"bulbasaur/internal/resources"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/pkg/errors"
	"strings"
)

type clientHandler struct {
	resource map[string]resources.ClientResource
}

func New(resource map[string]resources.ClientResource) *clientHandler {
	return &clientHandler{resource: resource}
}

func (h *clientHandler) InitializeScenario(s *godog.ScenarioContext) {
	s.Step(`^"([^"]*)" send request to "([^"]*)"$`, h.sendRequest)
	s.Step(`^"([^"]*)" send request to "([^"]*)" with body$`, h.sendRequestWithBody)
	s.Step(`^"([^"]*)" send request to "([^"]*)" with body from file "([^"]*)"$`, h.sendRequestWithBodyFromFile)
	s.Step(`^"([^"]*)" set request header key "([^"]*)" with value "([^"]*)"$`, h.setRequestHeader)
	s.Step(`^"([^"]*)" response code should be (\d+)$`, h.checkResponseCode)
}

func (h *clientHandler) sendRequest(resourceName, target string) error {
	return h.sendRequestWithBody(resourceName, target, nil)
}

func (h *clientHandler) sendRequestWithBody(resourceName, target string, body *godog.DocString) error {
	r, ok := h.resource[resourceName]
	if !ok {
		return errorNotFound(resourceName)
	}

	str := strings.Split(target, " ")
	if len(str) != 2 {
		return fmt.Errorf("invalid target format: %s,  should follow `[METHOD] [PATH]`", target)
	}

	var requestBody []byte
	if body != nil {
		requestBody = []byte(strings.TrimSpace(body.Content))
	}
	return r.Request(str[0], str[1], requestBody)
}

func (h *clientHandler) sendRequestWithBodyFromFile(resourceName, target, fileName string) error {
	r, ok := h.resource[resourceName]
	if !ok {
		return errorNotFound(resourceName)
	}

	str := strings.Split(target, " ")
	if len(str) != 2 {
		return errorNotFound(resourceName)
	}

	return r.RequestFromFile(str[0], str[1], fileName)
}

func (h *clientHandler) setRequestHeader(resourceName, key, value string) error {
	r, ok := h.resource[resourceName]
	if !ok {
		return errorNotFound(resourceName)
	}

	if err := r.SetRequestHeader(key, value); err != nil {
		return err
	}

	return nil
}

func (h *clientHandler) checkResponseCode(resourceName string, code int) error {
	r, ok := h.resource[resourceName]
	if !ok {
		return errorNotFound(resourceName)
	}

	statusCode, _, _, err := r.GetResponse()
	if err != nil {
		return err
	}

	if code != statusCode {
		return errors.Errorf("response code is %s not %s", statusCode, code)
	}

	return nil
}

func errorNotFound(resourceName string) error {
	return fmt.Errorf("%s not found", resourceName)
}
