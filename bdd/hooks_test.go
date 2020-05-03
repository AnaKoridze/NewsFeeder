package tests

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"gopkg.in/resty.v1"
	"os"
	"strings"
)

type api struct {
	baseURL string
	resp    *resty.Response
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (a *api) aRunningApi() (err error) {
	a.baseURL = getEnv("TESTS_API_BASE_URL", "http://localhost:8080")
	return nil
}

func (a *api) iSendARequestTo(action string, endpoint string) (err error) {
	ds := &messages.PickleStepArgument_PickleDocString{}
	ds.Content = "{}"
	return a.iSendARequestToWithJson(action, endpoint, ds)
}

func (a *api) iSendARequestToWithJson(action string, endpoint string, body *messages.PickleStepArgument_PickleDocString) (err error) {

	var data interface{}
	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
		return
	}
	if _, err = json.Marshal(data); err != nil {
		return
	}

	// execute the request
	resp, err := resty.
		R().
		SetBody(strings.NewReader(body.Content)).
		Execute(action, a.baseURL+endpoint)

	a.resp = resp

	return err

}

func (a *api) theResponseCodeShouldBe(code int) (err error) {

	if code != a.resp.StatusCode() {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d with body: %s", code, a.resp.StatusCode(), string(a.resp.Body()))
	}

	return nil
}

func (a *api) theResponseShouldMatchJson(body *messages.PickleStepArgument_PickleDocString) (err error) {

	var expected, actual []byte
	var exp, act interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &exp); err != nil {
		return
	}
	if expected, err = json.MarshalIndent(exp, "", "  "); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body(), &act); err != nil {
		return
	}
	if actual, err = json.MarshalIndent(act, "", "  "); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"expected json length: %d does not match actual: %d:\n%s",
			len(expected),
			len(actual),
			string(actual),
		)
	}

	for i, b := range actual {
		if b != expected[i] {
			return fmt.Errorf(
				"expected JSON does not match actual, showing up to last matched character:\n%s",
				string(actual[:i+1]),
			)
		}
	}
	return

}

func FeatureContext(s *godog.Suite) {

	api := &api{}

	s.Step(`^a running api$`, api.aRunningApi)
	s.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, api.iSendARequestTo)
	s.Step(`^I send a "([^"]*)" request to "([^"]*)" with json:$`, api.iSendARequestToWithJson)
	s.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	s.Step(`^the response should match json:$`, api.theResponseShouldMatchJson)

}
