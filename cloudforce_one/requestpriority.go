// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RequestPriorityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRequestPriorityService] method
// instead.
type RequestPriorityService struct {
	Options []option.RequestOption
}

// NewRequestPriorityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestPriorityService(opts ...option.RequestOption) (r *RequestPriorityService) {
	r = &RequestPriorityService{}
	r.Options = opts
	return
}

// Create a New Priority Requirement
func (r *RequestPriorityService) New(ctx context.Context, accountIdentifier string, body RequestPriorityNewParams, opts ...option.RequestOption) (res *Priority, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Priority Intelligence Requirement
func (r *RequestPriorityService) Update(ctx context.Context, accountIdentifier string, priorityIdentifer string, body RequestPriorityUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Priority Intelligence Report
func (r *RequestPriorityService) Delete(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *RequestPriorityDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Priority Intelligence Requirement
func (r *RequestPriorityService) Get(ctx context.Context, accountIdentifier string, priorityIdentifer string, opts ...option.RequestOption) (res *Item, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/%s", accountIdentifier, priorityIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Priority Intelligence Requirement Quota
func (r *RequestPriorityService) Quota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *Quota, err error) {
	opts = append(r.Options[:], opts...)
	var env RequestPriorityQuotaResponseEnvelope
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Label = string

type LabelParam = string

type Priority struct {
	// UUID
	ID string `json:"id,required"`
	// Priority creation time
	Created time.Time `json:"created,required" format:"date-time"`
	// List of labels
	Labels []Label `json:"labels,required"`
	// Priority
	Priority int64 `json:"priority,required"`
	// Requirement
	Requirement string `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp PriorityTlp `json:"tlp,required"`
	// Priority last updated time
	Updated time.Time    `json:"updated,required" format:"date-time"`
	JSON    priorityJSON `json:"-"`
}

// priorityJSON contains the JSON metadata for the struct [Priority]
type priorityJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Priority    apijson.Field
	Requirement apijson.Field
	Tlp         apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Priority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priorityJSON) RawJSON() string {
	return r.raw
}

// The CISA defined Traffic Light Protocol (TLP)
type PriorityTlp string

const (
	PriorityTlpClear       PriorityTlp = "clear"
	PriorityTlpAmber       PriorityTlp = "amber"
	PriorityTlpAmberStrict PriorityTlp = "amber-strict"
	PriorityTlpGreen       PriorityTlp = "green"
	PriorityTlpRed         PriorityTlp = "red"
)

func (r PriorityTlp) IsKnown() bool {
	switch r {
	case PriorityTlpClear, PriorityTlpAmber, PriorityTlpAmberStrict, PriorityTlpGreen, PriorityTlpRed:
		return true
	}
	return false
}

type PriorityEditParam struct {
	// List of labels
	Labels param.Field[[]LabelParam] `json:"labels,required"`
	// Priority
	Priority param.Field[int64] `json:"priority,required"`
	// Requirement
	Requirement param.Field[string] `json:"requirement,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[PriorityEditTlp] `json:"tlp,required"`
}

func (r PriorityEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The CISA defined Traffic Light Protocol (TLP)
type PriorityEditTlp string

const (
	PriorityEditTlpClear       PriorityEditTlp = "clear"
	PriorityEditTlpAmber       PriorityEditTlp = "amber"
	PriorityEditTlpAmberStrict PriorityEditTlp = "amber-strict"
	PriorityEditTlpGreen       PriorityEditTlp = "green"
	PriorityEditTlpRed         PriorityEditTlp = "red"
)

func (r PriorityEditTlp) IsKnown() bool {
	switch r {
	case PriorityEditTlpClear, PriorityEditTlpAmber, PriorityEditTlpAmberStrict, PriorityEditTlpGreen, PriorityEditTlpRed:
		return true
	}
	return false
}

// Union satisfied by [cloudforce_one.RequestPriorityDeleteResponseUnknown],
// [cloudforce_one.RequestPriorityDeleteResponseArray] or [shared.UnionString].
type RequestPriorityDeleteResponseUnion interface {
	ImplementsCloudforceOneRequestPriorityDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RequestPriorityDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RequestPriorityDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RequestPriorityDeleteResponseArray []interface{}

func (r RequestPriorityDeleteResponseArray) ImplementsCloudforceOneRequestPriorityDeleteResponseUnion() {
}

type RequestPriorityNewParams struct {
	PriorityEdit PriorityEditParam `json:"priority_edit,required"`
}

func (r RequestPriorityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type RequestPriorityNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Priority              `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityNewResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityNewResponseEnvelope]
type requestPriorityNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityNewResponseEnvelopeSuccess bool

const (
	RequestPriorityNewResponseEnvelopeSuccessTrue RequestPriorityNewResponseEnvelopeSuccess = true
)

func (r RequestPriorityNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityUpdateParams struct {
	PriorityEdit PriorityEditParam `json:"priority_edit,required"`
}

func (r RequestPriorityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PriorityEdit)
}

type RequestPriorityUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Item                  `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityUpdateResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityUpdateResponseEnvelope]
type requestPriorityUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityUpdateResponseEnvelopeSuccess bool

const (
	RequestPriorityUpdateResponseEnvelopeSuccessTrue RequestPriorityUpdateResponseEnvelopeSuccess = true
)

func (r RequestPriorityUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   RequestPriorityDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityDeleteResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityDeleteResponseEnvelope]
type requestPriorityDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityDeleteResponseEnvelopeSuccess bool

const (
	RequestPriorityDeleteResponseEnvelopeSuccessTrue RequestPriorityDeleteResponseEnvelopeSuccess = true
)

func (r RequestPriorityDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Item                  `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityGetResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RequestPriorityGetResponseEnvelope]
type requestPriorityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityGetResponseEnvelopeSuccess bool

const (
	RequestPriorityGetResponseEnvelopeSuccessTrue RequestPriorityGetResponseEnvelopeSuccess = true
)

func (r RequestPriorityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RequestPriorityQuotaResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Quota                 `json:"result,required"`
	// Whether the API call was successful
	Success RequestPriorityQuotaResponseEnvelopeSuccess `json:"success,required"`
	JSON    requestPriorityQuotaResponseEnvelopeJSON    `json:"-"`
}

// requestPriorityQuotaResponseEnvelopeJSON contains the JSON metadata for the
// struct [RequestPriorityQuotaResponseEnvelope]
type requestPriorityQuotaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RequestPriorityQuotaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestPriorityQuotaResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RequestPriorityQuotaResponseEnvelopeSuccess bool

const (
	RequestPriorityQuotaResponseEnvelopeSuccessTrue RequestPriorityQuotaResponseEnvelopeSuccess = true
)

func (r RequestPriorityQuotaResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RequestPriorityQuotaResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
