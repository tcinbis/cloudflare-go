// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_plans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]RatePlanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RatePlanGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Component struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name ComponentName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64       `json:"unit_price"`
	JSON      componentJSON `json:"-"`
}

// componentJSON contains the JSON metadata for the struct [Component]
type componentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Component) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r componentJSON) RawJSON() string {
	return r.raw
}

// The unique component.
type ComponentName string

const (
	ComponentNameZones                       ComponentName = "zones"
	ComponentNamePageRules                   ComponentName = "page_rules"
	ComponentNameDedicatedCertificates       ComponentName = "dedicated_certificates"
	ComponentNameDedicatedCertificatesCustom ComponentName = "dedicated_certificates_custom"
)

func (r ComponentName) IsKnown() bool {
	switch r {
	case ComponentNameZones, ComponentNamePageRules, ComponentNameDedicatedCertificates, ComponentNameDedicatedCertificatesCustom:
		return true
	}
	return false
}

type RatePlanGetResponse struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []Component `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency RatePlanGetResponseFrequency `json:"frequency"`
	// The plan name.
	Name string                  `json:"name"`
	JSON ratePlanGetResponseJSON `json:"-"`
}

// ratePlanGetResponseJSON contains the JSON metadata for the struct
// [RatePlanGetResponse]
type ratePlanGetResponseJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which you will be billed for this plan.
type RatePlanGetResponseFrequency string

const (
	RatePlanGetResponseFrequencyWeekly    RatePlanGetResponseFrequency = "weekly"
	RatePlanGetResponseFrequencyMonthly   RatePlanGetResponseFrequency = "monthly"
	RatePlanGetResponseFrequencyQuarterly RatePlanGetResponseFrequency = "quarterly"
	RatePlanGetResponseFrequencyYearly    RatePlanGetResponseFrequency = "yearly"
)

func (r RatePlanGetResponseFrequency) IsKnown() bool {
	switch r {
	case RatePlanGetResponseFrequencyWeekly, RatePlanGetResponseFrequencyMonthly, RatePlanGetResponseFrequencyQuarterly, RatePlanGetResponseFrequencyYearly:
		return true
	}
	return false
}

type RatePlanGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RatePlanGetResponse                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RatePlanGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RatePlanGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ratePlanGetResponseEnvelopeJSON       `json:"-"`
}

// ratePlanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RatePlanGetResponseEnvelope]
type ratePlanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RatePlanGetResponseEnvelopeSuccess bool

const (
	RatePlanGetResponseEnvelopeSuccessTrue RatePlanGetResponseEnvelopeSuccess = true
)

func (r RatePlanGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RatePlanGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RatePlanGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       ratePlanGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// ratePlanGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RatePlanGetResponseEnvelopeResultInfo]
type ratePlanGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RatePlanGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
