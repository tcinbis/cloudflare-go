// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthchecks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PreviewService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPreviewService] method instead.
type PreviewService struct {
	Options []option.RequestOption
}

// NewPreviewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPreviewService(opts ...option.RequestOption) (r *PreviewService) {
	r = &PreviewService{}
	r.Options = opts
	return
}

// Create a new preview health check.
func (r *PreviewService) New(ctx context.Context, params PreviewNewParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a health check.
func (r *PreviewService) Delete(ctx context.Context, healthcheckID string, params PreviewDeleteParams, opts ...option.RequestOption) (res *PreviewDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", params.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured health check preview.
func (r *PreviewService) Get(ctx context.Context, healthcheckID string, query PreviewGetParams, opts ...option.RequestOption) (res *Healthcheck, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/healthchecks/preview/%s", query.ZoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PreviewDeleteResponse struct {
	// Identifier
	ID   string                    `json:"id"`
	JSON previewDeleteResponseJSON `json:"-"`
}

// previewDeleteResponseJSON contains the JSON metadata for the struct
// [PreviewDeleteResponse]
type previewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PreviewNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]CheckRegionItem] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HTTPConfigurationParam] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TCPConfig param.Field[TCPConfigurationParam] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r PreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PreviewNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
	// Whether the API call was successful
	Success PreviewNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewNewResponseEnvelopeJSON    `json:"-"`
}

// previewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewNewResponseEnvelope]
type previewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewNewResponseEnvelopeSuccess bool

const (
	PreviewNewResponseEnvelopeSuccessTrue PreviewNewResponseEnvelopeSuccess = true
)

func (r PreviewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PreviewDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r PreviewDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PreviewDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PreviewDeleteResponse                                     `json:"result,required"`
	// Whether the API call was successful
	Success PreviewDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewDeleteResponseEnvelopeJSON    `json:"-"`
}

// previewDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewDeleteResponseEnvelope]
type previewDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewDeleteResponseEnvelopeSuccess bool

const (
	PreviewDeleteResponseEnvelopeSuccessTrue PreviewDeleteResponseEnvelopeSuccess = true
)

func (r PreviewDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PreviewGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PreviewGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Healthcheck                                               `json:"result,required"`
	// Whether the API call was successful
	Success PreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewGetResponseEnvelopeJSON    `json:"-"`
}

// previewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelope]
type previewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewGetResponseEnvelopeSuccess bool

const (
	PreviewGetResponseEnvelopeSuccessTrue PreviewGetResponseEnvelopeSuccess = true
)

func (r PreviewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
