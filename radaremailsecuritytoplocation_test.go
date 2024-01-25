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

func TestRadarEmailSecurityTopLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.Email.Security.Top.Locations.List(context.TODO(), cloudflare.RadarEmailSecurityTopLocationListParams{
		Arc:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationListParamsArc{cloudflare.RadarEmailSecurityTopLocationListParamsArcPass, cloudflare.RadarEmailSecurityTopLocationListParamsArcNone, cloudflare.RadarEmailSecurityTopLocationListParamsArcFail}),
		ASN:       cloudflare.F([]string{"string", "string", "string"}),
		DateEnd:   cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange: cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationListParamsDateRange{cloudflare.RadarEmailSecurityTopLocationListParamsDateRange1d, cloudflare.RadarEmailSecurityTopLocationListParamsDateRange2d, cloudflare.RadarEmailSecurityTopLocationListParamsDateRange7d}),
		DateStart: cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		Dkim:      cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationListParamsDkim{cloudflare.RadarEmailSecurityTopLocationListParamsDkimPass, cloudflare.RadarEmailSecurityTopLocationListParamsDkimNone, cloudflare.RadarEmailSecurityTopLocationListParamsDkimFail}),
		Dmarc:     cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationListParamsDmarc{cloudflare.RadarEmailSecurityTopLocationListParamsDmarcPass, cloudflare.RadarEmailSecurityTopLocationListParamsDmarcNone, cloudflare.RadarEmailSecurityTopLocationListParamsDmarcFail}),
		Format:    cloudflare.F(cloudflare.RadarEmailSecurityTopLocationListParamsFormatJson),
		Limit:     cloudflare.F(int64(5)),
		Location:  cloudflare.F([]string{"string", "string", "string"}),
		Name:      cloudflare.F([]string{"string", "string", "string"}),
		Spf:       cloudflare.F([]cloudflare.RadarEmailSecurityTopLocationListParamsSpf{cloudflare.RadarEmailSecurityTopLocationListParamsSpfPass, cloudflare.RadarEmailSecurityTopLocationListParamsSpfNone, cloudflare.RadarEmailSecurityTopLocationListParamsSpfFail}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
