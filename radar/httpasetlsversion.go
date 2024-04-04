// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// HTTPAseTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHTTPAseTLSVersionService] method
// instead.
type HTTPAseTLSVersionService struct {
	Options []option.RequestOption
}

// NewHTTPAseTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseTLSVersionService(opts ...option.RequestOption) (r *HTTPAseTLSVersionService) {
	r = &HTTPAseTLSVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested TLS
// protocol version. Values are a percentage out of the total traffic.
func (r *HTTPAseTLSVersionService) Get(ctx context.Context, tlsVersion HTTPAseTLSVersionGetParamsTLSVersion, query HTTPAseTLSVersionGetParams, opts ...option.RequestOption) (res *HTTPAseTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HTTPAseTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseTLSVersionGetResponse struct {
	Meta HTTPAseTLSVersionGetResponseMeta                   `json:"meta,required"`
	Top0 []UnnamedSchemaRef4124a22436f90127c7fa2c4543219752 `json:"top_0,required"`
	JSON httpAseTLSVersionGetResponseJSON                   `json:"-"`
}

// httpAseTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponse]
type httpAseTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMeta struct {
	DateRange      []UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5 `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseTLSVersionGetResponseMetaConfidenceInfo     `json:"confidenceInfo"`
	JSON           httpAseTLSVersionGetResponseMetaJSON               `json:"-"`
}

// httpAseTLSVersionGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponseMeta]
type httpAseTLSVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1 `json:"annotations"`
	Level       int64                                              `json:"level"`
	JSON        httpAseTLSVersionGetResponseMetaConfidenceInfoJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseTLSVersionGetResponseMetaConfidenceInfo]
type httpAseTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseTLSVersionGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]HTTPAseTLSVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPAseTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseTLSVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPAseTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPAseTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type HTTPAseTLSVersionGetParamsTLSVersion string

const (
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic HTTPAseTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseTLSVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3, HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsBotClass string

const (
	HTTPAseTLSVersionGetParamsBotClassLikelyAutomated HTTPAseTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseTLSVersionGetParamsBotClassLikelyHuman     HTTPAseTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseTLSVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsBotClassLikelyAutomated, HTTPAseTLSVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsDateRange string

const (
	HTTPAseTLSVersionGetParamsDateRange1d         HTTPAseTLSVersionGetParamsDateRange = "1d"
	HTTPAseTLSVersionGetParamsDateRange2d         HTTPAseTLSVersionGetParamsDateRange = "2d"
	HTTPAseTLSVersionGetParamsDateRange7d         HTTPAseTLSVersionGetParamsDateRange = "7d"
	HTTPAseTLSVersionGetParamsDateRange14d        HTTPAseTLSVersionGetParamsDateRange = "14d"
	HTTPAseTLSVersionGetParamsDateRange28d        HTTPAseTLSVersionGetParamsDateRange = "28d"
	HTTPAseTLSVersionGetParamsDateRange12w        HTTPAseTLSVersionGetParamsDateRange = "12w"
	HTTPAseTLSVersionGetParamsDateRange24w        HTTPAseTLSVersionGetParamsDateRange = "24w"
	HTTPAseTLSVersionGetParamsDateRange52w        HTTPAseTLSVersionGetParamsDateRange = "52w"
	HTTPAseTLSVersionGetParamsDateRange1dControl  HTTPAseTLSVersionGetParamsDateRange = "1dControl"
	HTTPAseTLSVersionGetParamsDateRange2dControl  HTTPAseTLSVersionGetParamsDateRange = "2dControl"
	HTTPAseTLSVersionGetParamsDateRange7dControl  HTTPAseTLSVersionGetParamsDateRange = "7dControl"
	HTTPAseTLSVersionGetParamsDateRange14dControl HTTPAseTLSVersionGetParamsDateRange = "14dControl"
	HTTPAseTLSVersionGetParamsDateRange28dControl HTTPAseTLSVersionGetParamsDateRange = "28dControl"
	HTTPAseTLSVersionGetParamsDateRange12wControl HTTPAseTLSVersionGetParamsDateRange = "12wControl"
	HTTPAseTLSVersionGetParamsDateRange24wControl HTTPAseTLSVersionGetParamsDateRange = "24wControl"
)

func (r HTTPAseTLSVersionGetParamsDateRange) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsDateRange1d, HTTPAseTLSVersionGetParamsDateRange2d, HTTPAseTLSVersionGetParamsDateRange7d, HTTPAseTLSVersionGetParamsDateRange14d, HTTPAseTLSVersionGetParamsDateRange28d, HTTPAseTLSVersionGetParamsDateRange12w, HTTPAseTLSVersionGetParamsDateRange24w, HTTPAseTLSVersionGetParamsDateRange52w, HTTPAseTLSVersionGetParamsDateRange1dControl, HTTPAseTLSVersionGetParamsDateRange2dControl, HTTPAseTLSVersionGetParamsDateRange7dControl, HTTPAseTLSVersionGetParamsDateRange14dControl, HTTPAseTLSVersionGetParamsDateRange28dControl, HTTPAseTLSVersionGetParamsDateRange12wControl, HTTPAseTLSVersionGetParamsDateRange24wControl:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsDeviceType string

const (
	HTTPAseTLSVersionGetParamsDeviceTypeDesktop HTTPAseTLSVersionGetParamsDeviceType = "DESKTOP"
	HTTPAseTLSVersionGetParamsDeviceTypeMobile  HTTPAseTLSVersionGetParamsDeviceType = "MOBILE"
	HTTPAseTLSVersionGetParamsDeviceTypeOther   HTTPAseTLSVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPAseTLSVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsDeviceTypeDesktop, HTTPAseTLSVersionGetParamsDeviceTypeMobile, HTTPAseTLSVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseTLSVersionGetParamsFormat string

const (
	HTTPAseTLSVersionGetParamsFormatJson HTTPAseTLSVersionGetParamsFormat = "JSON"
	HTTPAseTLSVersionGetParamsFormatCsv  HTTPAseTLSVersionGetParamsFormat = "CSV"
)

func (r HTTPAseTLSVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsFormatJson, HTTPAseTLSVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPProtocol string

const (
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTP  HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTP"
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseTLSVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPProtocolHTTP, HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPVersion string

const (
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv1 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv2 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv3 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseTLSVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPVersionHttPv1, HTTPAseTLSVersionGetParamsHTTPVersionHttPv2, HTTPAseTLSVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsIPVersion string

const (
	HTTPAseTLSVersionGetParamsIPVersionIPv4 HTTPAseTLSVersionGetParamsIPVersion = "IPv4"
	HTTPAseTLSVersionGetParamsIPVersionIPv6 HTTPAseTLSVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPAseTLSVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsIPVersionIPv4, HTTPAseTLSVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsOS string

const (
	HTTPAseTLSVersionGetParamsOSWindows  HTTPAseTLSVersionGetParamsOS = "WINDOWS"
	HTTPAseTLSVersionGetParamsOSMacosx   HTTPAseTLSVersionGetParamsOS = "MACOSX"
	HTTPAseTLSVersionGetParamsOSIos      HTTPAseTLSVersionGetParamsOS = "IOS"
	HTTPAseTLSVersionGetParamsOSAndroid  HTTPAseTLSVersionGetParamsOS = "ANDROID"
	HTTPAseTLSVersionGetParamsOSChromeos HTTPAseTLSVersionGetParamsOS = "CHROMEOS"
	HTTPAseTLSVersionGetParamsOSLinux    HTTPAseTLSVersionGetParamsOS = "LINUX"
	HTTPAseTLSVersionGetParamsOSSmartTv  HTTPAseTLSVersionGetParamsOS = "SMART_TV"
)

func (r HTTPAseTLSVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsOSWindows, HTTPAseTLSVersionGetParamsOSMacosx, HTTPAseTLSVersionGetParamsOSIos, HTTPAseTLSVersionGetParamsOSAndroid, HTTPAseTLSVersionGetParamsOSChromeos, HTTPAseTLSVersionGetParamsOSLinux, HTTPAseTLSVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetResponseEnvelope struct {
	Result  HTTPAseTLSVersionGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpAseTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseEnvelope]
type httpAseTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
