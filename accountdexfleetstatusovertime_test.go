// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAccountDexFleetStatusOverTimeListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithAPIToken("my-cloudflare-api-token"),
		option.WithUserServiceKey("my-cloudflare-user-service-key"),
	)
	err := client.Accounts.Dex.FleetStatusOverTime.List(
		context.TODO(),
		"01a7362d577a6c3019a474fd6f485823",
		cloudflare.AccountDexFleetStatusOverTimeListParams{
			TimeEnd:   cloudflare.F("2023-10-11T00:00:00Z"),
			TimeStart: cloudflare.F("2023-10-11T00:00:00Z"),
			Colo:      cloudflare.F("SJC"),
			DeviceID:  cloudflare.F("cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
