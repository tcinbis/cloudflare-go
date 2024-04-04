// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// HTTPAseService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewHTTPAseService] method instead.
type HTTPAseService struct {
	Options      []option.RequestOption
	BotClass     *HTTPAseBotClassService
	DeviceType   *HTTPAseDeviceTypeService
	HTTPProtocol *HTTPAseHTTPProtocolService
	HTTPMethod   *HTTPAseHTTPMethodService
	IPVersion    *HTTPAseIPVersionService
	OS           *HTTPAseOSService
	TLSVersion   *HTTPAseTLSVersionService
}

// NewHTTPAseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPAseService(opts ...option.RequestOption) (r *HTTPAseService) {
	r = &HTTPAseService{}
	r.Options = opts
	r.BotClass = NewHTTPAseBotClassService(opts...)
	r.DeviceType = NewHTTPAseDeviceTypeService(opts...)
	r.HTTPProtocol = NewHTTPAseHTTPProtocolService(opts...)
	r.HTTPMethod = NewHTTPAseHTTPMethodService(opts...)
	r.IPVersion = NewHTTPAseIPVersionService(opts...)
	r.OS = NewHTTPAseOSService(opts...)
	r.TLSVersion = NewHTTPAseTLSVersionService(opts...)
	return
}

// Get the top autonomous systems by HTTP traffic. Values are a percentage out of
// the total traffic.
func (r *HTTPAseService) Get(ctx context.Context, query HTTPAseGetParams, opts ...option.RequestOption) (res *HTTPAseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env HTTPAseGetResponseEnvelope
	path := "radar/http/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseGetResponse struct {
	Meta HTTPAseGetResponseMeta                             `json:"meta,required"`
	Top0 []UnnamedSchemaRef4124a22436f90127c7fa2c4543219752 `json:"top_0,required"`
	JSON httpAseGetResponseJSON                             `json:"-"`
}

// httpAseGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseGetResponse]
type httpAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetResponseMeta struct {
	DateRange      []UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5 `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseGetResponseMetaConfidenceInfo               `json:"confidenceInfo"`
	JSON           httpAseGetResponseMetaJSON                         `json:"-"`
}

// httpAseGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseMeta]
type httpAseGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1 `json:"annotations"`
	Level       int64                                              `json:"level"`
	JSON        httpAseGetResponseMetaConfidenceInfoJSON           `json:"-"`
}

// httpAseGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPAseGetResponseMetaConfidenceInfo]
type httpAseGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]HTTPAseGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPAseGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPAseGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPAseGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseGetParams]'s query parameters as `url.Values`.
func (r HTTPAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HTTPAseGetParamsBotClass string

const (
	HTTPAseGetParamsBotClassLikelyAutomated HTTPAseGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseGetParamsBotClassLikelyHuman     HTTPAseGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsBotClassLikelyAutomated, HTTPAseGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseGetParamsDateRange string

const (
	HTTPAseGetParamsDateRange1d         HTTPAseGetParamsDateRange = "1d"
	HTTPAseGetParamsDateRange2d         HTTPAseGetParamsDateRange = "2d"
	HTTPAseGetParamsDateRange7d         HTTPAseGetParamsDateRange = "7d"
	HTTPAseGetParamsDateRange14d        HTTPAseGetParamsDateRange = "14d"
	HTTPAseGetParamsDateRange28d        HTTPAseGetParamsDateRange = "28d"
	HTTPAseGetParamsDateRange12w        HTTPAseGetParamsDateRange = "12w"
	HTTPAseGetParamsDateRange24w        HTTPAseGetParamsDateRange = "24w"
	HTTPAseGetParamsDateRange52w        HTTPAseGetParamsDateRange = "52w"
	HTTPAseGetParamsDateRange1dControl  HTTPAseGetParamsDateRange = "1dControl"
	HTTPAseGetParamsDateRange2dControl  HTTPAseGetParamsDateRange = "2dControl"
	HTTPAseGetParamsDateRange7dControl  HTTPAseGetParamsDateRange = "7dControl"
	HTTPAseGetParamsDateRange14dControl HTTPAseGetParamsDateRange = "14dControl"
	HTTPAseGetParamsDateRange28dControl HTTPAseGetParamsDateRange = "28dControl"
	HTTPAseGetParamsDateRange12wControl HTTPAseGetParamsDateRange = "12wControl"
	HTTPAseGetParamsDateRange24wControl HTTPAseGetParamsDateRange = "24wControl"
)

func (r HTTPAseGetParamsDateRange) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsDateRange1d, HTTPAseGetParamsDateRange2d, HTTPAseGetParamsDateRange7d, HTTPAseGetParamsDateRange14d, HTTPAseGetParamsDateRange28d, HTTPAseGetParamsDateRange12w, HTTPAseGetParamsDateRange24w, HTTPAseGetParamsDateRange52w, HTTPAseGetParamsDateRange1dControl, HTTPAseGetParamsDateRange2dControl, HTTPAseGetParamsDateRange7dControl, HTTPAseGetParamsDateRange14dControl, HTTPAseGetParamsDateRange28dControl, HTTPAseGetParamsDateRange12wControl, HTTPAseGetParamsDateRange24wControl:
		return true
	}
	return false
}

type HTTPAseGetParamsDeviceType string

const (
	HTTPAseGetParamsDeviceTypeDesktop HTTPAseGetParamsDeviceType = "DESKTOP"
	HTTPAseGetParamsDeviceTypeMobile  HTTPAseGetParamsDeviceType = "MOBILE"
	HTTPAseGetParamsDeviceTypeOther   HTTPAseGetParamsDeviceType = "OTHER"
)

func (r HTTPAseGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsDeviceTypeDesktop, HTTPAseGetParamsDeviceTypeMobile, HTTPAseGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseGetParamsFormat string

const (
	HTTPAseGetParamsFormatJson HTTPAseGetParamsFormat = "JSON"
	HTTPAseGetParamsFormatCsv  HTTPAseGetParamsFormat = "CSV"
)

func (r HTTPAseGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsFormatJson, HTTPAseGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseGetParamsHTTPProtocol string

const (
	HTTPAseGetParamsHTTPProtocolHTTP  HTTPAseGetParamsHTTPProtocol = "HTTP"
	HTTPAseGetParamsHTTPProtocolHTTPS HTTPAseGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsHTTPProtocolHTTP, HTTPAseGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseGetParamsHTTPVersion string

const (
	HTTPAseGetParamsHTTPVersionHttPv1 HTTPAseGetParamsHTTPVersion = "HTTPv1"
	HTTPAseGetParamsHTTPVersionHttPv2 HTTPAseGetParamsHTTPVersion = "HTTPv2"
	HTTPAseGetParamsHTTPVersionHttPv3 HTTPAseGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsHTTPVersionHttPv1, HTTPAseGetParamsHTTPVersionHttPv2, HTTPAseGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseGetParamsIPVersion string

const (
	HTTPAseGetParamsIPVersionIPv4 HTTPAseGetParamsIPVersion = "IPv4"
	HTTPAseGetParamsIPVersionIPv6 HTTPAseGetParamsIPVersion = "IPv6"
)

func (r HTTPAseGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsIPVersionIPv4, HTTPAseGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseGetParamsOS string

const (
	HTTPAseGetParamsOSWindows  HTTPAseGetParamsOS = "WINDOWS"
	HTTPAseGetParamsOSMacosx   HTTPAseGetParamsOS = "MACOSX"
	HTTPAseGetParamsOSIos      HTTPAseGetParamsOS = "IOS"
	HTTPAseGetParamsOSAndroid  HTTPAseGetParamsOS = "ANDROID"
	HTTPAseGetParamsOSChromeos HTTPAseGetParamsOS = "CHROMEOS"
	HTTPAseGetParamsOSLinux    HTTPAseGetParamsOS = "LINUX"
	HTTPAseGetParamsOSSmartTv  HTTPAseGetParamsOS = "SMART_TV"
)

func (r HTTPAseGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsOSWindows, HTTPAseGetParamsOSMacosx, HTTPAseGetParamsOSIos, HTTPAseGetParamsOSAndroid, HTTPAseGetParamsOSChromeos, HTTPAseGetParamsOSLinux, HTTPAseGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseGetParamsTLSVersion string

const (
	HTTPAseGetParamsTLSVersionTlSv1_0  HTTPAseGetParamsTLSVersion = "TLSv1_0"
	HTTPAseGetParamsTLSVersionTlSv1_1  HTTPAseGetParamsTLSVersion = "TLSv1_1"
	HTTPAseGetParamsTLSVersionTlSv1_2  HTTPAseGetParamsTLSVersion = "TLSv1_2"
	HTTPAseGetParamsTLSVersionTlSv1_3  HTTPAseGetParamsTLSVersion = "TLSv1_3"
	HTTPAseGetParamsTLSVersionTlSvQuic HTTPAseGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseGetParamsTLSVersionTlSv1_0, HTTPAseGetParamsTLSVersionTlSv1_1, HTTPAseGetParamsTLSVersionTlSv1_2, HTTPAseGetParamsTLSVersionTlSv1_3, HTTPAseGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseGetResponseEnvelope struct {
	Result  HTTPAseGetResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    httpAseGetResponseEnvelopeJSON `json:"-"`
}

// httpAseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPAseGetResponseEnvelope]
type httpAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
