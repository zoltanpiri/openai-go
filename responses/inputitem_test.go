// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package responses_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/zoltanpiri/openai-go"
	"github.com/zoltanpiri/openai-go/internal/testutil"
	"github.com/zoltanpiri/openai-go/option"
	"github.com/zoltanpiri/openai-go/responses"
)

func TestInputItemListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Responses.InputItems.List(
		context.TODO(),
		"response_id",
		responses.InputItemListParams{
			After:   openai.String("after"),
			Before:  openai.String("before"),
			Include: []responses.ResponseIncludable{responses.ResponseIncludableFileSearchCallResults},
			Limit:   openai.Int(0),
			Order:   responses.InputItemListParamsOrderAsc,
		},
	)
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
