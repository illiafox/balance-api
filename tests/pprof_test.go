package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/ozontech/cute"
)

func TestPprof(t *testing.T) {
	c := cute.NewTestBuilder().
		Title("Pprof").
		Description("Test endpoints accessibility").
		Create()
	//
	c.RequestBuilder(
		cute.WithURI(Host+"/debug/pprof/"),
		cute.WithMethod(http.MethodGet),
	).ExpectStatus(
		http.StatusOK,
	).ExecuteTest(context.Background(), t)
}
