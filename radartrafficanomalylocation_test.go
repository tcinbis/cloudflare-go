// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarTrafficAnomalyLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.TrafficAnomalies.Locations.List(context.TODO(), cloudflare.RadarTrafficAnomalyLocationListParams{
		DateEnd:   cloudflare.F(time.Now()),
		DateRange: cloudflare.F(cloudflare.RadarTrafficAnomalyLocationListParamsDateRange7d),
		DateStart: cloudflare.F(time.Now()),
		Format:    cloudflare.F(cloudflare.RadarTrafficAnomalyLocationListParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
		Status:    cloudflare.F(cloudflare.RadarTrafficAnomalyLocationListParamsStatusVerified),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
