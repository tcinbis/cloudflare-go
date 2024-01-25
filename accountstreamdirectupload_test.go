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

func TestAccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Streams.DirectUploads.StreamVideosUploadVideosViaDirectUploadURLs(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParams{
			MaxDurationSeconds: cloudflare.F(int64(1)),
			AllowedOrigins:     cloudflare.F([]string{"example.com"}),
			Creator:            cloudflare.F("creator-id_abcde12345"),
			Expiry:             cloudflare.F(time.Now()),
			Meta: cloudflare.F[any](map[string]interface{}{
				"name": "video12345.mp4",
			}),
			RequireSignedURLs:     cloudflare.F(true),
			ScheduledDeletion:     cloudflare.F(time.Now()),
			ThumbnailTimestampPct: cloudflare.F(0.529241),
			Watermark: cloudflare.F(cloudflare.AccountStreamDirectUploadStreamVideosUploadVideosViaDirectUploadURLsParamsWatermark{
				Uid: cloudflare.F("ea95132c15732412d22c1476fa83f27a"),
			}),
			UploadCreator: cloudflare.F("creator-id_abcde12345"),
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
