package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/ozontech/cute"
)

func TestMetrics(t *testing.T) {
	c := cute.NewTestBuilder().
		Title("Prometheus metrics").
		Description("Test endpoints accessibility").
		Create()
	//
	c.RequestBuilder(
		cute.WithURI(Host+"/metrics"),
		cute.WithMethod(http.MethodGet),
	).ExpectStatus(
		http.StatusOK,
	).ExecuteTest(context.Background(), t)
}
