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

func TestZoneWaitingRoomGet(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneWaitingRoomUpdateParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.ZoneWaitingRoomUpdateParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.ZoneWaitingRoomUpdateParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestZoneWaitingRoomDelete(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneWaitingRoomPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.Patch(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneWaitingRoomPatchParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.ZoneWaitingRoomPatchParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.ZoneWaitingRoomPatchParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestZoneWaitingRoomWaitingRoomNewWaitingRoomWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.WaitingRoomNewWaitingRoom(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParams{
			Host:              cloudflare.F("shop.example.com"),
			Name:              cloudflare.F("production_webinar"),
			NewUsersPerMinute: cloudflare.F(int64(200)),
			TotalActiveUsers:  cloudflare.F(int64(200)),
			AdditionalRoutes: cloudflare.F([]cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsAdditionalRoute{{
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}, {
				Host: cloudflare.F("shop2.example.com"),
				Path: cloudflare.F("/shop2/checkout"),
			}}),
			CookieAttributes: cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributes{
				Samesite: cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSamesiteAuto),
				Secure:   cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsCookieAttributesSecureAuto),
			}),
			CookieSuffix:            cloudflare.F("abcd"),
			CustomPageHTML:          cloudflare.F("{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Queue all enabled {{/waitTimeKnown}}"),
			DefaultTemplateLanguage: cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsDefaultTemplateLanguageEsEs),
			Description:             cloudflare.F("Production - DO NOT MODIFY"),
			DisableSessionRenewal:   cloudflare.F(false),
			JsonResponseEnabled:     cloudflare.F(false),
			Path:                    cloudflare.F("/shop/checkout"),
			QueueAll:                cloudflare.F(true),
			QueueingMethod:          cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingMethodFifo),
			QueueingStatusCode:      cloudflare.F(cloudflare.ZoneWaitingRoomWaitingRoomNewWaitingRoomParamsQueueingStatusCode202),
			SessionDuration:         cloudflare.F(int64(1)),
			Suspended:               cloudflare.F(true),
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

func TestZoneWaitingRoomWaitingRoomListWaitingRooms(t *testing.T) {
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
	_, err := client.Zones.WaitingRooms.WaitingRoomListWaitingRooms(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
