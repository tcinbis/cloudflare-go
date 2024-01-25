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

func TestRadarEmailSecuritySpfTimeseryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.SpfTimeseries.List(context.TODO(), cloudflare.RadarEmailSecuritySpfTimeseryListParams{
		AggInterval: cloudflare.F(cloudflare.RadarEmailSecuritySpfTimeseryListParamsAggInterval1h),
		Arc:         cloudflare.F([]cloudflare.RadarEmailSecuritySpfTimeseryListParamsArc{cloudflare.RadarEmailSecuritySpfTimeseryListParamsArcPass, cloudflare.RadarEmailSecuritySpfTimeseryListParamsArcNone, cloudflare.RadarEmailSecuritySpfTimeseryListParamsArcFail}),
		ASN:         cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:     cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:   cloudflare.F([]cloudflare.RadarEmailSecuritySpfTimeseryListParamsDateRange{cloudflare.RadarEmailSecuritySpfTimeseryListParamsDateRange1d, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDateRange2d, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDateRange7d}),
		DateStart:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:        cloudflare.F([]cloudflare.RadarEmailSecuritySpfTimeseryListParamsDkim{cloudflare.RadarEmailSecuritySpfTimeseryListParamsDkimPass, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDkimNone, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDkimFail}),
		Dmarc:       cloudflare.F([]cloudflare.RadarEmailSecuritySpfTimeseryListParamsDmarc{cloudflare.RadarEmailSecuritySpfTimeseryListParamsDmarcPass, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDmarcNone, cloudflare.RadarEmailSecuritySpfTimeseryListParamsDmarcFail}),
		Format:      cloudflare.F(cloudflare.RadarEmailSecuritySpfTimeseryListParamsFormatJson),
		Location:    cloudflare.F([]string{"string", "string", "string"}),
		Name:        cloudflare.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
