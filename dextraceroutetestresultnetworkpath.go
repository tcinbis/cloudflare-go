// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DEXTracerouteTestResultNetworkPathService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDEXTracerouteTestResultNetworkPathService] method instead.
type DEXTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewDEXTracerouteTestResultNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDEXTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *DEXTracerouteTestResultNetworkPathService) {
	r = &DEXTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *DEXTracerouteTestResultNetworkPathService) List(ctx context.Context, accountID string, testResultID string, opts ...option.RequestOption) (res *DEXTracerouteTestResultNetworkPathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DEXTracerouteTestResultNetworkPathListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", accountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXTracerouteTestResultNetworkPathListResponse struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []DEXTracerouteTestResultNetworkPathListResponseHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// date time of this traceroute test
	TimeStart string `json:"time_start,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                             `json:"testName"`
	JSON     dexTracerouteTestResultNetworkPathListResponseJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseJSON contains the JSON metadata
// for the struct [DEXTracerouteTestResultNetworkPathListResponse]
type dexTracerouteTestResultNetworkPathListResponseJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	TimeStart   apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTracerouteTestResultNetworkPathListResponseHop struct {
	TTL           int64                                                      `json:"ttl,required"`
	Asn           int64                                                      `json:"asn,nullable"`
	Aso           string                                                     `json:"aso,nullable"`
	IPAddress     string                                                     `json:"ipAddress,nullable"`
	Location      DEXTracerouteTestResultNetworkPathListResponseHopsLocation `json:"location,nullable"`
	Mile          DEXTracerouteTestResultNetworkPathListResponseHopsMile     `json:"mile,nullable"`
	Name          string                                                     `json:"name,nullable"`
	PacketLossPct float64                                                    `json:"packetLossPct,nullable"`
	RTTMs         int64                                                      `json:"rttMs,nullable"`
	JSON          dexTracerouteTestResultNetworkPathListResponseHopJSON      `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseHopJSON contains the JSON metadata
// for the struct [DEXTracerouteTestResultNetworkPathListResponseHop]
type dexTracerouteTestResultNetworkPathListResponseHopJSON struct {
	TTL           apijson.Field
	Asn           apijson.Field
	Aso           apijson.Field
	IPAddress     apijson.Field
	Location      apijson.Field
	Mile          apijson.Field
	Name          apijson.Field
	PacketLossPct apijson.Field
	RTTMs         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTracerouteTestResultNetworkPathListResponseHopsLocation struct {
	City  string                                                         `json:"city,nullable"`
	State string                                                         `json:"state,nullable"`
	Zip   string                                                         `json:"zip,nullable"`
	JSON  dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestResultNetworkPathListResponseHopsLocation]
type dexTracerouteTestResultNetworkPathListResponseHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTracerouteTestResultNetworkPathListResponseHopsMile string

const (
	DEXTracerouteTestResultNetworkPathListResponseHopsMileClientToApp       DEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-app"
	DEXTracerouteTestResultNetworkPathListResponseHopsMileClientToCfEgress  DEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-egress"
	DEXTracerouteTestResultNetworkPathListResponseHopsMileClientToCfIngress DEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-cf-ingress"
	DEXTracerouteTestResultNetworkPathListResponseHopsMileClientToIsp       DEXTracerouteTestResultNetworkPathListResponseHopsMile = "client-to-isp"
)

type DEXTracerouteTestResultNetworkPathListResponseEnvelope struct {
	Errors   []DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages `json:"messages,required"`
	Result   DEXTracerouteTestResultNetworkPathListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON contains the JSON
// metadata for the struct [DEXTracerouteTestResultNetworkPathListResponseEnvelope]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages]
type dexTracerouteTestResultNetworkPathListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccessTrue DEXTracerouteTestResultNetworkPathListResponseEnvelopeSuccess = true
)
