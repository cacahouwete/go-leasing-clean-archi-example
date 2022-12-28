package contexts

import (
	"context"
	"log"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v16"
)

type Context struct {
	AppContext  *AppContext
	HTTPContext *HTTPContext
	JSONContext *JSONContext
}

func (c *Context) processScenarioTags(ctx context.Context, tags []*messages.PickleTag) (context.Context, error) {
	var err error

	for _, tag := range tags {
		switch tag.Name {
		case "@db:clean":
			err = c.AppContext.DbClean()
		case "@restart":
			err = c.AppContext.Restart()
		case "@reset":
			err = c.AppContext.Reset()
		}

		if err != nil {
			return nil, err
		}
	}

	return ctx, nil
}

func (c *Context) InitializeTestSuite(*godog.TestSuiteContext) {
	if errP := c.AppContext.Reset(); errP != nil {
		log.Fatal(errP.Error())
	}
}

func (c *Context) InitializeScenario(ctx *godog.ScenarioContext) {
	c.AppContext.InitializeScenario(ctx)
	c.HTTPContext.InitializeScenario(ctx)
	c.JSONContext.InitializeScenario(ctx)

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// before each scenario we need to process tags to reset, cleandb, restart, ... on demand
		return c.processScenarioTags(ctx, sc.Tags)
	})
}

func New() *Context {
	app := NewAPPContext()
	hc := NewHTTPContext()

	return &Context{
		AppContext:  app,
		HTTPContext: hc,
		JSONContext: NewJSONContext(hc),
	}
}
