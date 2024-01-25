// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountIntelMiscategorizationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountIntelMiscategorizationService] method instead.
type AccountIntelMiscategorizationService struct {
	Options []option.RequestOption
}

// NewAccountIntelMiscategorizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIntelMiscategorizationService(opts ...option.RequestOption) (r *AccountIntelMiscategorizationService) {
	r = &AccountIntelMiscategorizationService{}
	r.Options = opts
	return
}

// Create Miscategorization
func (r *AccountIntelMiscategorizationService) MiscategorizationNewMiscategorization(ctx context.Context, accountIdentifier string, body AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParams, opts ...option.RequestOption) (res *AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponse struct {
	Errors   []AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseError   `json:"errors"`
	Messages []AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessage `json:"messages"`
	Result   AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseSuccess `json:"success"`
	JSON    accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseJSON    `json:"-"`
}

// accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseJSON
// contains the JSON metadata for the struct
// [AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponse]
type accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseErrorJSON `json:"-"`
}

// accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseError]
type accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessageJSON `json:"-"`
}

// accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessage]
type accountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseResultUnknown]
// or [shared.UnionString].
type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseResult interface {
	ImplementsAccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseSuccess bool

const (
	AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseSuccessTrue AccountIntelMiscategorizationMiscategorizationNewMiscategorizationResponseSuccess = true
)

type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParams struct {
	// Content category IDs to add.
	ContentAdds param.Field[interface{}] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[interface{}]                                                                           `json:"content_removes"`
	IndicatorType  param.Field[AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType] `json:"indicator_type"`
	// Provide only if indicator_type is `ipv4` or `ipv6`.
	IP param.Field[interface{}] `json:"ip"`
	// Security category IDs to add.
	SecurityAdds param.Field[interface{}] `json:"security_adds"`
	// Security category IDs to remove.
	SecurityRemoves param.Field[interface{}] `json:"security_removes"`
	// Provide only if indicator_type is `domain` or `url`. Example if indicator_type
	// is `domain`: `example.com`. Example if indicator_type is `url`:
	// `https://example.com/news/`.
	URL param.Field[string] `json:"url"`
}

func (r AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType string

const (
	AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeDomain AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "domain"
	AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeIpv4   AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "ipv4"
	AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeIpv6   AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "ipv6"
	AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorTypeURL    AccountIntelMiscategorizationMiscategorizationNewMiscategorizationParamsIndicatorType = "url"
)
