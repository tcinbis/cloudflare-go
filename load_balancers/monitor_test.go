// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestMonitorNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.New(context.TODO(), load_balancers.MonitorNewParams{
		AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ExpectedCodes:   cloudflare.F("2xx"),
		AllowInsecure:   cloudflare.F(true),
		ConsecutiveDown: cloudflare.F(int64(0)),
		ConsecutiveUp:   cloudflare.F(int64(0)),
		Description:     cloudflare.F("Login page monitor"),
		ExpectedBody:    cloudflare.F("alive"),
		FollowRedirects: cloudflare.F(true),
		Header: cloudflare.F[any](map[string]interface{}{
			"Host": map[string]interface{}{
				"0": "example.com",
			},
			"X-App-ID": map[string]interface{}{
				"0": "abc123",
			},
		}),
		Interval:  cloudflare.F(int64(0)),
		Method:    cloudflare.F("GET"),
		Path:      cloudflare.F("/health"),
		Port:      cloudflare.F(int64(0)),
		ProbeZone: cloudflare.F("example.com"),
		Retries:   cloudflare.F(int64(0)),
		Timeout:   cloudflare.F(int64(0)),
		Type:      cloudflare.F(load_balancers.MonitorNewParamsTypeHTTPS),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.Update(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorUpdateParams{
			AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ExpectedCodes:   cloudflare.F("2xx"),
			AllowInsecure:   cloudflare.F(true),
			ConsecutiveDown: cloudflare.F(int64(0)),
			ConsecutiveUp:   cloudflare.F(int64(0)),
			Description:     cloudflare.F("Login page monitor"),
			ExpectedBody:    cloudflare.F("alive"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F[any](map[string]interface{}{
				"Host": map[string]interface{}{
					"0": "example.com",
				},
				"X-App-ID": map[string]interface{}{
					"0": "abc123",
				},
			}),
			Interval:  cloudflare.F(int64(0)),
			Method:    cloudflare.F("GET"),
			Path:      cloudflare.F("/health"),
			Port:      cloudflare.F(int64(0)),
			ProbeZone: cloudflare.F("example.com"),
			Retries:   cloudflare.F(int64(0)),
			Timeout:   cloudflare.F(int64(0)),
			Type:      cloudflare.F(load_balancers.MonitorUpdateParamsTypeHTTPS),
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

func TestMonitorList(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.List(context.TODO(), load_balancers.MonitorListParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMonitorDelete(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.Delete(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorDeleteParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Body:      cloudflare.F[any](map[string]interface{}{}),
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

func TestMonitorEditWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.Edit(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorEditParams{
			AccountID:       cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			ExpectedCodes:   cloudflare.F("2xx"),
			AllowInsecure:   cloudflare.F(true),
			ConsecutiveDown: cloudflare.F(int64(0)),
			ConsecutiveUp:   cloudflare.F(int64(0)),
			Description:     cloudflare.F("Login page monitor"),
			ExpectedBody:    cloudflare.F("alive"),
			FollowRedirects: cloudflare.F(true),
			Header: cloudflare.F[any](map[string]interface{}{
				"Host": map[string]interface{}{
					"0": "example.com",
				},
				"X-App-ID": map[string]interface{}{
					"0": "abc123",
				},
			}),
			Interval:  cloudflare.F(int64(0)),
			Method:    cloudflare.F("GET"),
			Path:      cloudflare.F("/health"),
			Port:      cloudflare.F(int64(0)),
			ProbeZone: cloudflare.F("example.com"),
			Retries:   cloudflare.F(int64(0)),
			Timeout:   cloudflare.F(int64(0)),
			Type:      cloudflare.F(load_balancers.MonitorEditParamsTypeHTTPS),
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

func TestMonitorGet(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Monitors.Get(
		context.TODO(),
		"f1aba936b94213e5b8dca0c0dbf1f9cc",
		load_balancers.MonitorGetParams{
			AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
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
