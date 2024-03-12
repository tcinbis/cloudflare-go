// File generated from our OpenAPI spec by Stainless.

package custom_certificates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PrioritizeService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPrioritizeService] method instead.
type PrioritizeService struct {
	Options []option.RequestOption
}

// NewPrioritizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrioritizeService(opts ...option.RequestOption) (r *PrioritizeService) {
	r = &PrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *PrioritizeService) Update(ctx context.Context, params PrioritizeUpdateParams, opts ...option.RequestOption) (res *[]TLSCertificatesAndHostnamesCustomCertificate, err error) {
	opts = append(r.Options[:], opts...)
	var env PrioritizeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrioritizeUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Array of ordered certificates.
	Certificates param.Field[[]PrioritizeUpdateParamsCertificate] `json:"certificates,required"`
}

func (r PrioritizeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrioritizeUpdateParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r PrioritizeUpdateParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrioritizeUpdateResponseEnvelope struct {
	Errors   []PrioritizeUpdateResponseEnvelopeErrors       `json:"errors,required"`
	Messages []PrioritizeUpdateResponseEnvelopeMessages     `json:"messages,required"`
	Result   []TLSCertificatesAndHostnamesCustomCertificate `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PrioritizeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PrioritizeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       prioritizeUpdateResponseEnvelopeJSON       `json:"-"`
}

// prioritizeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrioritizeUpdateResponseEnvelope]
type prioritizeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    prioritizeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseEnvelopeErrors]
type prioritizeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrioritizeUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    prioritizeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PrioritizeUpdateResponseEnvelopeMessages]
type prioritizeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PrioritizeUpdateResponseEnvelopeSuccess bool

const (
	PrioritizeUpdateResponseEnvelopeSuccessTrue PrioritizeUpdateResponseEnvelopeSuccess = true
)

type PrioritizeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       prioritizeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// prioritizeUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PrioritizeUpdateResponseEnvelopeResultInfo]
type prioritizeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrioritizeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prioritizeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}