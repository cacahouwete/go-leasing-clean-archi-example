package contexts

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/martinohmann/jsoncompare"
	"github.com/tidwall/gjson"
)

type JSONContext struct {
	HTTPContext *HTTPContext
	JSONObject  string
}

// Errors messages

func (jc *JSONContext) NewErrorJSONNodeDoesNotExist(arg1 string) error {
	return fmt.Errorf("the json node %s does not exist", arg1)
}

// Contexts handlers

func (jc *JSONContext) theJSONShouldBeValid() error {
	if gjson.Valid(jc.JSONObject) {
		return nil
	}

	return errors.New("the object is not a valid json format")
}

func (jc *JSONContext) theJSONShouldBeAnArray() error {
	result := gjson.Parse(jc.JSONObject)

	if result.IsArray() {
		return nil
	}

	return errors.New("the json object is not an array")
}

func (jc *JSONContext) theJSONShouldBeEqualTo(arg1 *godog.DocString) error {
	comparator := jsoncompare.NewComparator(jsoncompare.MatchStrict)

	return comparator.Compare([]byte(jc.JSONObject), []byte(arg1.Content))
}

func (jc *JSONContext) theJSONShouldContain(arg1 *godog.DocString) error {
	comparator := jsoncompare.NewComparator(jsoncompare.MatchSubtree)

	return comparator.Compare([]byte(jc.JSONObject), []byte(arg1.Content))
}

func (jc *JSONContext) theJSONNodeShouldBeAnArrayOfLength(arg1 string, arg2 int) error {
	err := jc.theJSONNodeShouldBeAnArray(arg1)
	if err != nil {
		return err
	}

	result := gjson.Get(jc.JSONObject, arg1)
	if len(result.Array()) == arg2 {
		return nil
	}

	return fmt.Errorf("expected json node array to be of length %d, but was actually of length %d", arg2, len(result.Array()))
}

func (jc *JSONContext) theJSONShouldBeAnArrayOfLength(arg1 int) error {
	err := jc.theJSONShouldBeAnArray()
	if err != nil {
		return err
	}

	result := gjson.Parse(jc.JSONObject)
	if len(result.Array()) == arg1 {
		return nil
	}

	return fmt.Errorf("expected json array to be of length %d, but was actually of length %d", arg1, len(result.Array()))
}

func (jc *JSONContext) theJSONNodeShouldMatchRegex(arg1, arg2 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	switch result.Type {
	case gjson.String:
		match, err := regexp.MatchString(arg2, result.String())
		if err != nil {
			return err
		}

		if !match {
			return fmt.Errorf("json node %s expected value was %s but got %s", arg1, arg2, result.String())
		}

		return nil
	case gjson.Number:
		match, err := regexp.MatchString(arg2, result.String())
		if err != nil {
			return err
		}

		if !match {
			return fmt.Errorf("json node %s expected value was %s but got %f", arg1, arg2, result.Float())
		}

		return nil
	case gjson.False:
		return fmt.Errorf("json node %s expected value was %s but got false", arg1, arg2)
	case gjson.True:
		return fmt.Errorf("json node %s expected value was %s but got true", arg1, arg2)
	case gjson.JSON:
		return fmt.Errorf("json node %s expected value was %s but got json object", arg1, arg2)
	case gjson.Null:
		return fmt.Errorf("json node %s expected value was %s but got null", arg1, arg2)
	default:
		return nil
	}
}

func (jc *JSONContext) theJSONNodeShouldBeEqualTo(arg1, arg2 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	switch result.Type {
	case gjson.String:
		if result.String() != arg2 {
			return fmt.Errorf("json node %s expected value was %s but got %s", arg1, arg2, result.String())
		}

		return nil
	case gjson.Number:
		const bitSize = 64
		val, err := strconv.ParseFloat(arg2, bitSize)

		if err != nil {
			return err
		}

		if result.Float() != val {
			return fmt.Errorf("json node %s expected value was %s but got %f", arg1, arg2, result.Float())
		}

		return nil
	case gjson.False:
		if arg2 != "false" {
			return fmt.Errorf("json node %s expected value was %s but got 'false'", arg1, arg2)
		}

		return nil
	case gjson.True:
		if arg2 != "true" {
			return fmt.Errorf("json node %s expected value was %s but got 'true'", arg1, arg2)
		}

		return nil
	case gjson.JSON:
		comparator := jsoncompare.NewComparator(jsoncompare.MatchStrict)

		return comparator.Compare([]byte(result.Raw), []byte(arg2))
	case gjson.Null:
		if arg2 != "null" {
			return fmt.Errorf("json node %s expected value was %s but got 'null'", arg1, arg2)
		}

		return nil
	default:
		return nil
	}
}

func (jc *JSONContext) theJSONNodeShouldBeANumber(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Type == gjson.Number {
		return nil
	}

	return fmt.Errorf("json node %s expected to be a number but was actually %s", arg1, result.Type.String())
}

func (jc *JSONContext) theJSONNodeShouldBeABoolean(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	switch result.Value().(type) {
	case bool:
		return nil
	default:
		return fmt.Errorf("json node %s expected to be a boolean but was actually %s", arg1, result.Type.String())
	}
}

func (jc *JSONContext) theJSONNodeShouldBeAnObject(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.IsObject() {
		return nil
	}

	return fmt.Errorf("json node %s expected to be an object but was actually %s", arg1, result.Type.String())
}

func (jc *JSONContext) theJSONNodeShouldBeEqualToInt(arg1 string, arg2 int) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	switch v := result.Value().(type) {
	case int, int8, int16, int32, int64:
		if v == arg2 {
			return nil
		}

		return fmt.Errorf("json node %s expected to be equal to %d but was actually equal to %s", arg1, arg2, v)
	case float32, float64:
		if v == float32(arg2) || v == float64(arg2) {
			return nil
		}

		return fmt.Errorf("json node %s expected to be equal to %d but was actually equal to %s", arg1, arg2, v)
	default:
		return fmt.Errorf("json node %s expected to be equal to %d but was actually equal to %s", arg1, arg2, v)
	}
}

func (jc *JSONContext) theJSONNodeShouldBeAnArray(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.IsArray() {
		return nil
	}

	return errors.New("the json node is not an array")
}

func (jc *JSONContext) theJSONNodeShouldBeFalse(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Type == gjson.False {
		return nil
	}

	return fmt.Errorf("json node %s expected value to be false but got %s", arg1, result.Type.String())
}

func (jc *JSONContext) theJSONNodeShouldBeTrue(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Type == gjson.True {
		return nil
	}

	return fmt.Errorf("json node %s expected value to be true but got %s", arg1, result.Type.String())
}

func (jc *JSONContext) theJSONNodeShouldContain(arg1, arg2 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	valueString, ok := result.Value().(string)
	if !ok {
		return fmt.Errorf("the JSON node %s value is not a string", arg1)
	}

	if strings.Contains(valueString, arg2) {
		return nil
	}

	return fmt.Errorf("json node %s does not contain %s", arg1, arg2)
}

func (jc *JSONContext) theJSONNodeShouldExist(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Exists() {
		return nil
	}

	return fmt.Errorf("json node %s does not exist", arg1)
}

func (jc *JSONContext) theJSONShouldHaveElements(arg1 int) error {
	result := gjson.Parse(jc.JSONObject)

	nbElements := 0

	result.ForEach(func(key, value gjson.Result) bool {
		nbElements++

		return true
	})

	if nbElements == arg1 {
		return nil
	}

	return fmt.Errorf("json expected number of elements is %d but is actually %d", arg1, nbElements)
}

func (jc *JSONContext) theJSONNodeShouldHaveElements(arg1 string, arg2 int) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	nbElements := 0

	result.ForEach(func(key, value gjson.Result) bool {
		nbElements++

		return true
	})

	if nbElements == arg2 {
		return nil
	}

	return fmt.Errorf("json object expected number of elements is %d but is actually %d", arg2, nbElements)
}

func (jc *JSONContext) theJSONNodeShouldMatch(arg1, arg2 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	regex, err := regexp.Compile(arg2)

	if err != nil {
		return err
	}

	if regex.MatchString(result.String()) {
		return nil
	}

	return fmt.Errorf("json node %s does not match regexp %s", arg1, arg2)
}

func (jc *JSONContext) theJSONNodeShouldBeNull(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Type != gjson.Null {
		return fmt.Errorf("json node %s expected value to be null but is actually %s", arg1, result.Type.String())
	}

	return nil
}

func (jc *JSONContext) theJSONNodeShouldNotBeNull(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)
	if !result.Exists() {
		return jc.NewErrorJSONNodeDoesNotExist(arg1)
	}

	if result.Type == gjson.Null {
		return fmt.Errorf("json node %s value is null", arg1)
	}

	return nil
}

func (jc *JSONContext) theJSONNodeShouldNotExist(arg1 string) error {
	result := gjson.Get(jc.JSONObject, arg1)

	if result.Exists() {
		return fmt.Errorf("json node %s exists", arg1)
	}

	return nil
}

func (jc *JSONContext) theJSONObject(arg1 *godog.DocString) error {
	jc.JSONObject = arg1.Content

	return nil
}

func (jc *JSONContext) theResponsePayloadShouldBeJSON() error {
	if jc.HTTPContext.ResponseBody == nil {
		return errors.New("the response payload is empty")
	}

	if gjson.Valid(string(jc.HTTPContext.ResponseBody)) {
		jc.JSONObject = string(jc.HTTPContext.ResponseBody)

		return nil
	}

	return errors.New("the response payload does not appear to be a JSON object")
}

func (jc *JSONContext) InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		jc.JSONObject = ""

		return ctx, nil
	})

	ctx.Step(`^the JSON should be valid$`, jc.theJSONShouldBeValid)
	ctx.Step(`^the JSON should be an array$`, jc.theJSONShouldBeAnArray)
	ctx.Step(`^the JSON should be equal to:$`, jc.theJSONShouldBeEqualTo)
	ctx.Step(`^the JSON should contain:$`, jc.theJSONShouldContain)
	ctx.Step(`^the JSON node "([^"]*)" should be an array of length (\d+)$`, jc.theJSONNodeShouldBeAnArrayOfLength)
	ctx.Step(`^the JSON should be an array of length (\d+)$`, jc.theJSONShouldBeAnArrayOfLength)
	ctx.Step(`^the JSON node "([^"]*)" should match regex "(.*)"$`, jc.theJSONNodeShouldMatchRegex)
	ctx.Step(`^the JSON node "([^"]*)" should be equal to "(.*)"$`, jc.theJSONNodeShouldBeEqualTo)
	ctx.Step(`^the JSON node "([^"]*)" should be equal to (\d+)$`, jc.theJSONNodeShouldBeEqualToInt)
	ctx.Step(`^the JSON node "([^"]*)" should be a number$`, jc.theJSONNodeShouldBeANumber)
	ctx.Step(`^the JSON node "([^"]*)" should be an object$`, jc.theJSONNodeShouldBeAnObject)
	ctx.Step(`^the JSON node "([^"]*)" should be a boolean$`, jc.theJSONNodeShouldBeABoolean)
	ctx.Step(`^the JSON node "([^"]*)" should be an array$`, jc.theJSONNodeShouldBeAnArray)
	ctx.Step(`^the JSON node "([^"]*)" should be false$`, jc.theJSONNodeShouldBeFalse)
	ctx.Step(`^the JSON node "([^"]*)" should be true$`, jc.theJSONNodeShouldBeTrue)
	ctx.Step(`^the JSON node "([^"]*)" should contain "(.*)"$`, jc.theJSONNodeShouldContain)
	ctx.Step(`^the JSON node "([^"]*)" should exist$`, jc.theJSONNodeShouldExist)
	ctx.Step(`^the JSON node "([^"]*)" should be null$`, jc.theJSONNodeShouldBeNull)
	ctx.Step(`^the JSON node "([^"]*)" should not be null$`, jc.theJSONNodeShouldNotBeNull)
	ctx.Step(`^the JSON should have (\d+) elements$`, jc.theJSONShouldHaveElements)
	ctx.Step(`^the JSON node "([^"]*)" should have (\d+) elements$`, jc.theJSONNodeShouldHaveElements)
	ctx.Step(`^the JSON node "([^"]*)" should match "([^"]*)"$`, jc.theJSONNodeShouldMatch)
	ctx.Step(`^the JSON node "([^"]*)" should not exist$`, jc.theJSONNodeShouldNotExist)
	ctx.Step(`^the JSON object:$`, jc.theJSONObject)
	ctx.Step(`^the response payload should be JSON$`, jc.theResponsePayloadShouldBeJSON)
}

func NewJSONContext(hc *HTTPContext) *JSONContext {
	return &JSONContext{HTTPContext: hc}
}
