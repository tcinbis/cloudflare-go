// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// GatewayLoggingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayLoggingService] method
// instead.
type GatewayLoggingService struct {
	Options []option.RequestOption
}

// NewGatewayLoggingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayLoggingService(opts ...option.RequestOption) (r *GatewayLoggingService) {
	r = &GatewayLoggingService{}
	r.Options = opts
	return
}

// Updates logging settings for the current Zero Trust account.
func (r *GatewayLoggingService) Update(ctx context.Context, params GatewayLoggingUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayGatewayAccountLoggingSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/logging", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current logging settings for Zero Trust account.
func (r *GatewayLoggingService) Get(ctx context.Context, query GatewayLoggingGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayGatewayAccountLoggingSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/logging", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayGatewayAccountLoggingSettings struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType shared.UnnamedSchemaRef28                         `json:"settings_by_rule_type"`
	JSON               zeroTrustGatewayGatewayAccountLoggingSettingsJSON `json:"-"`
}

// zeroTrustGatewayGatewayAccountLoggingSettingsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayGatewayAccountLoggingSettings]
type zeroTrustGatewayGatewayAccountLoggingSettingsJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGatewayGatewayAccountLoggingSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayGatewayAccountLoggingSettingsJSON) RawJSON() string {
	return r.raw
}

type GatewayLoggingUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[shared.UnnamedSchemaRef28Param] `json:"settings_by_rule_type"`
}

func (r GatewayLoggingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLoggingUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172                  `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172                  `json:"messages,required"`
	Result   ZeroTrustGatewayGatewayAccountLoggingSettings `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLoggingUpdateResponseEnvelope]
type gatewayLoggingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLoggingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLoggingUpdateResponseEnvelopeSuccess bool

const (
	GatewayLoggingUpdateResponseEnvelopeSuccessTrue GatewayLoggingUpdateResponseEnvelopeSuccess = true
)

func (r GatewayLoggingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLoggingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayLoggingGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayLoggingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172                  `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172                  `json:"messages,required"`
	Result   ZeroTrustGatewayGatewayAccountLoggingSettings `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLoggingGetResponseEnvelope]
type gatewayLoggingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLoggingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayLoggingGetResponseEnvelopeSuccess bool

const (
	GatewayLoggingGetResponseEnvelopeSuccessTrue GatewayLoggingGetResponseEnvelopeSuccess = true
)

func (r GatewayLoggingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayLoggingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
