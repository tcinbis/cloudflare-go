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

func TestRadarHTTPTimeseriesGroupByOListWithOptionalParams(t *testing.T) {
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
	_, err := client.Radar.HTTP.TimeseriesGroups.ByOs.List(context.TODO(), cloudflare.RadarHTTPTimeseriesGroupByOListParams{
		AggInterval:  cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByOListParamsAggInterval1h),
		ASN:          cloudflare.F([]string{"string", "string", "string"}),
		BotClass:     cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsBotClass{cloudflare.RadarHTTPTimeseriesGroupByOListParamsBotClassLikelyAutomated, cloudflare.RadarHTTPTimeseriesGroupByOListParamsBotClassLikelyHuman}),
		DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DateRange:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsDateRange{cloudflare.RadarHTTPTimeseriesGroupByOListParamsDateRange1d, cloudflare.RadarHTTPTimeseriesGroupByOListParamsDateRange2d, cloudflare.RadarHTTPTimeseriesGroupByOListParamsDateRange7d}),
		DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
		DeviceType:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsDeviceType{cloudflare.RadarHTTPTimeseriesGroupByOListParamsDeviceTypeDesktop, cloudflare.RadarHTTPTimeseriesGroupByOListParamsDeviceTypeMobile, cloudflare.RadarHTTPTimeseriesGroupByOListParamsDeviceTypeOther}),
		Format:       cloudflare.F(cloudflare.RadarHTTPTimeseriesGroupByOListParamsFormatJson),
		HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPProtocol{cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPProtocolHTTP, cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPProtocolHTTPs}),
		HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPVersion{cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv1, cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv2, cloudflare.RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv3}),
		IPVersion:    cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsIPVersion{cloudflare.RadarHTTPTimeseriesGroupByOListParamsIPVersionIPv4, cloudflare.RadarHTTPTimeseriesGroupByOListParamsIPVersionIPv6}),
		Location:     cloudflare.F([]string{"string", "string", "string"}),
		Name:         cloudflare.F([]string{"string", "string", "string"}),
		TlsVersion:   cloudflare.F([]cloudflare.RadarHTTPTimeseriesGroupByOListParamsTlsVersion{cloudflare.RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_0, cloudflare.RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_1, cloudflare.RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_2}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
