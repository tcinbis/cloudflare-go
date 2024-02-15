// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorService] method instead.
type UserLoadBalancerMonitorService struct {
	Options    []option.RequestOption
	Previews   *UserLoadBalancerMonitorPreviewService
	References *UserLoadBalancerMonitorReferenceService
}

// NewUserLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorService) {
	r = &UserLoadBalancerMonitorService{}
	r.Options = opts
	r.Previews = NewUserLoadBalancerMonitorPreviewService(opts...)
	r.References = NewUserLoadBalancerMonitorReferenceService(opts...)
	return
}

// Create a configured monitor.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsNewMonitor(ctx context.Context, body UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams, opts ...option.RequestOption) (res *UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for a user.
func (r *UserLoadBalancerMonitorService) LoadBalancerMonitorsListMonitors(ctx context.Context, opts ...option.RequestOption) (res *[]UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponse struct {
	ID string `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64     `json:"consecutive_up"`
	CreatedOn     time.Time `json:"created_on" format:"date-time"`
	// Object description.
	Description string `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes string `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval int64 `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method     string    `json:"method"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path string `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone string `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType `json:"type"`
	JSON userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponse]
type userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	CreatedOn       apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	ModifiedOn      apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType string

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeHTTP     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "http"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeHTTPS    UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "https"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeTcp      UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "tcp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeUdpIcmp  UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "udp_icmp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeIcmpPing UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "icmp_ping"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseTypeSmtp     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseType = "smtp"
)

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse struct {
	ID string `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64     `json:"consecutive_up"`
	CreatedOn     time.Time `json:"created_on" format:"date-time"`
	// Object description.
	Description string `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes string `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval int64 `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method     string    `json:"method"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path string `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone string `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType `json:"type"`
	JSON userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	CreatedOn       apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	ModifiedOn      apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType string

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeHTTP     UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "http"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeHTTPS    UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "https"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeTcp      UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "tcp"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeUdpIcmp  UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "udp_icmp"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeIcmpPing UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "icmp_ping"
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseTypeSmtp     UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseType = "smtp"
)

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType string

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeHTTP     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "http"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeHTTPS    UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "https"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeTcp      UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "tcp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeUdpIcmp  UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "udp_icmp"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeIcmpPing UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "icmp_ping"
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsTypeSmtp     UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorParamsType = "smtp"
)

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelope]
type userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors]
type userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages]
type userLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccessTrue UserLoadBalancerMonitorLoadBalancerMonitorsNewMonitorResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelope]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccessTrue UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo]
type userLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorLoadBalancerMonitorsListMonitorsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
