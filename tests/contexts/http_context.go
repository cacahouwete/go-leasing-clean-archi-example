package contexts

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/cucumber/godog"
)

const minNbCellForKeyValue = 2

type HTTPContext struct {
	RequestBody        []byte
	ResponseBody       []byte
	ResponseStatusCode int
	RequestHeader      http.Header
	ResponseHeader     http.Header
	BaseURL            string
}

func NewHTTPContext() *HTTPContext {
	return &HTTPContext{BaseURL: "http://leasing.localhost"}
}

func (hc *HTTPContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		hc.RequestBody = nil
		hc.ResponseBody = nil
		hc.ResponseStatusCode = 0
		hc.RequestHeader = nil
		hc.ResponseHeader = nil

		return ctx, nil
	})

	ctx.Step(`^I have following request body:$`, hc.iHaveFollowingRequestBody)
	ctx.Step(`^I have following request headers:$`, hc.iHaveFollowingRequestHeaders)
	ctx.Step(`^I send "(OPTIONS|GET|HEAD|POST|PUT|DELETE|TRACE|CONNECT)" request to "([^"]*)"$`, hc.iSendRequestTo)
	ctx.Step(`^I send "(OPTIONS|GET|HEAD|POST|PUT|DELETE|TRACE|CONNECT)" request to "([^"]*)" with parameters:$`, hc.iSendRequestToWithParameters)
	ctx.Step(`^I send "(OPTIONS|GET|HEAD|POST|PUT|DELETE|TRACE|CONNECT)" request to "([^"]*)" with body:$`, hc.iSendRequestToWithBody)
	ctx.Step(`^the response should have "([^"]*)" header$`, hc.theResponseShouldHaveHeader)
	ctx.Step(`^the response status code should be (\d+)$`, hc.theResponseStatusCodeShouldBe)
	ctx.Step(`^the response payload should be:$`, hc.theResponsePayloadShouldBe)
	ctx.Step(`^the response should have following headers:$`, hc.theResponseShouldHaveFollowingHeaders)
	ctx.Step(`^the response payload should match pattern:$`, hc.theResponsePayloadShouldMatchPattern)
	ctx.Step(`^I print response body$`, hc.iPrintResponseBody)
	ctx.Step(`^the response payload should be empty$`, hc.theResponsePayloadShouldBeEmpty)
	ctx.Step(`^the response payload should not be empty$`, hc.theResponsePayloadShouldNotBeEmpty)

	ctx.Step(`^I send "(OPTIONS|GET|HEAD|POST|PUT|DELETE|TRACE|CONNECT)" request to "([^"]*)" with body in "([^"]*)"$`, hc.iSendRequestToWithBodyInFile)
}

func (hc *HTTPContext) iPrintResponseBody() error {
	fmt.Println(string(hc.ResponseBody))

	return nil
}

func (hc *HTTPContext) iHaveFollowingRequestBody(arg1 *godog.DocString) error {
	hc.RequestBody = []byte(arg1.Content)

	return nil
}

func (hc *HTTPContext) iHaveFollowingRequestHeaders(arg1 *godog.Table) error {
	for i, row := range arg1.Rows {
		if len(row.Cells) < minNbCellForKeyValue {
			return fmt.Errorf("row %d have less than 2 cell", i)
		}

		hc.RequestHeader.Set(row.Cells[0].Value, row.Cells[1].Value)
	}

	return nil
}

func (hc *HTTPContext) iSendRequestTo(arg1, arg2 string) error {
	if hc.BaseURL != "" {
		arg2 = fmt.Sprintf("%s%s", hc.BaseURL, arg2)
	}

	var req *http.Request

	var err error

	if len(hc.RequestBody) > 0 {
		reader := bytes.NewReader(hc.RequestBody)
		req, err = http.NewRequest(arg1, arg2, reader)
	} else {
		req, err = http.NewRequest(arg1, arg2, nil)
	}

	if err != nil {
		return err
	}

	if hc.RequestHeader != nil {
		req.Header = hc.RequestHeader
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	hc.ResponseStatusCode = res.StatusCode
	hc.ResponseHeader = res.Header

	defer res.Body.Close()
	hc.ResponseBody, err = io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return nil
}

func (hc *HTTPContext) iSendRequestToWithParameters(arg1, arg2 string, arg3 *godog.Table) error {
	if arg2 == "" {
		return errors.New("uri should not be empty")
	}

	query := ""

	for i, row := range arg3.Rows {
		if len(row.Cells) < minNbCellForKeyValue {
			return fmt.Errorf("row %d have less than 2 cell", i)
		}

		if query != "" {
			query += "&"
		}

		query = fmt.Sprintf("%s%s=%s", query, row.Cells[0], row.Cells[1])
	}

	query = fmt.Sprintf("%s?%s", arg2, query)

	return hc.iSendRequestTo(arg1, query)
}

func (hc *HTTPContext) iSendRequestToWithBody(arg1, arg2 string, arg3 *godog.DocString) error {
	hc.RequestBody = []byte(arg3.Content)

	if arg2 == "" {
		return errors.New("uri should not be empty")
	}

	return hc.iSendRequestTo(arg1, arg2)
}

func (hc *HTTPContext) theResponseShouldHaveHeader(arg1 string) error {
	if hc.ResponseHeader != nil && hc.ResponseHeader.Get(arg1) != "" {
		return nil
	}

	return fmt.Errorf("could not find header %s", arg1)
}

func (hc *HTTPContext) theResponseStatusCodeShouldBe(arg1 int) error {
	if hc.ResponseStatusCode != arg1 {
		return fmt.Errorf("expected http status code %d but got %d", arg1, hc.ResponseStatusCode)
	}

	return nil
}

func (hc *HTTPContext) theResponsePayloadShouldBe(arg1 *godog.DocString) error {
	if hc.ResponseBody == nil {
		return errors.New("the response payload is empty")
	}

	if string(hc.ResponseBody) != arg1.Content {
		return fmt.Errorf("expected response payload to be %q, got %q", arg1.Content, string(hc.ResponseBody))
	}

	return nil
}

func (hc *HTTPContext) theResponseShouldHaveFollowingHeaders(arg1 *godog.Table) error {
	for i, row := range arg1.Rows {
		if len(row.Cells) < minNbCellForKeyValue {
			return fmt.Errorf("row %d have less than 2 cell", i)
		}

		key := row.Cells[0].Value
		value := row.Cells[1].Value

		headerValue := hc.ResponseHeader.Get(key)
		if headerValue == "" {
			return fmt.Errorf("could not find header %s", key)
		}

		if headerValue != value {
			return fmt.Errorf("incorrect value for header %s. expected %s but got %s", key, value, headerValue)
		}
	}

	return nil
}

func (hc *HTTPContext) theResponsePayloadShouldMatchPattern(arg1 *godog.DocString) error {
	if hc.ResponseBody == nil {
		return errors.New("the response payload is empty")
	}

	r := regexp.MustCompile(arg1.Content)

	if !r.Match(hc.ResponseBody) {
		return fmt.Errorf(
			"expected response payload %q to match pattern %q, but it did not",
			hc.ResponseBody,
			arg1.Content,
		)
	}

	return nil
}

func (hc *HTTPContext) theResponsePayloadShouldBeEmpty() error {
	if hc.ResponseBody == nil {
		return nil
	}

	return errors.New("the response payload is not empty")
}

func (hc *HTTPContext) theResponsePayloadShouldNotBeEmpty() error {
	if hc.ResponseBody == nil {
		return errors.New("the response payload is empty")
	}

	return nil
}

func (hc *HTTPContext) iSendRequestToWithBodyInFile(method, url, filePath string) error {
	if hc.BaseURL != "" {
		url = fmt.Sprintf("%s%s", hc.BaseURL, url)
	}

	requestBody, errFile := os.ReadFile(filePath)
	if errFile != nil {
		return errFile
	}

	hc.RequestBody = requestBody

	var req *http.Request

	var err error

	if len(requestBody) > 0 {
		reader := bytes.NewReader(requestBody)
		req, err = http.NewRequest(method, url, reader)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	if hc.RequestHeader != nil {
		req.Header = hc.RequestHeader
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	hc.ResponseStatusCode = res.StatusCode
	hc.ResponseHeader = res.Header

	defer res.Body.Close()
	hc.ResponseBody, err = io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	return nil
}
