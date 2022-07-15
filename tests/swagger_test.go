package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/ozontech/cute"
)

func TestSwagger(t *testing.T) {
	c := cute.NewTestBuilder().
		Title("Swagger").
		Description("Test endpoints accessibility").
		Create()
	//
	c.RequestBuilder(
		cute.WithURI(Host+"/swagger"),
		cute.WithMethod(http.MethodGet),
	).ExpectStatus(
		http.StatusOK,
	).ExecuteTest(context.Background(), t)
}
