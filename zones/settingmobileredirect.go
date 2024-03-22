// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingMobileRedirectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMobileRedirectService]
// method instead.
type SettingMobileRedirectService struct {
	Options []option.RequestOption
}

// NewSettingMobileRedirectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingMobileRedirectService(opts ...option.RequestOption) (r *SettingMobileRedirectService) {
	r = &SettingMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *SettingMobileRedirectService) Edit(ctx context.Context, params SettingMobileRedirectEditParams, opts ...option.RequestOption) (res *ZonesMobileRedirect, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *SettingMobileRedirectService) Get(ctx context.Context, query SettingMobileRedirectGetParams, opts ...option.RequestOption) (res *ZonesMobileRedirect, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMobileRedirectGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesMobileRedirectJSON `json:"-"`
}

// zonesMobileRedirectJSON contains the JSON metadata for the struct
// [ZonesMobileRedirect]
type zonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMobileRedirect) implementsZonesSettingEditResponse() {}

func (r ZonesMobileRedirect) implementsZonesSettingGetResponse() {}

// Identifier of the zone setting.
type ZonesMobileRedirectID string

const (
	ZonesMobileRedirectIDMobileRedirect ZonesMobileRedirectID = "mobile_redirect"
)

func (r ZonesMobileRedirectID) IsKnown() bool {
	switch r {
	case ZonesMobileRedirectIDMobileRedirect:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI bool                         `json:"strip_uri"`
	JSON     zonesMobileRedirectValueJSON `json:"-"`
}

// zonesMobileRedirectValueJSON contains the JSON metadata for the struct
// [ZonesMobileRedirectValue]
type zonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMobileRedirectValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not mobile redirect is enabled.
type ZonesMobileRedirectValueStatus string

const (
	ZonesMobileRedirectValueStatusOn  ZonesMobileRedirectValueStatus = "on"
	ZonesMobileRedirectValueStatusOff ZonesMobileRedirectValueStatus = "off"
)

func (r ZonesMobileRedirectValueStatus) IsKnown() bool {
	switch r {
	case ZonesMobileRedirectValueStatusOn, ZonesMobileRedirectValueStatusOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesMobileRedirectEditable bool

const (
	ZonesMobileRedirectEditableTrue  ZonesMobileRedirectEditable = true
	ZonesMobileRedirectEditableFalse ZonesMobileRedirectEditable = false
)

func (r ZonesMobileRedirectEditable) IsKnown() bool {
	switch r {
	case ZonesMobileRedirectEditableTrue, ZonesMobileRedirectEditableFalse:
		return true
	}
	return false
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZonesMobileRedirectParam struct {
	// Identifier of the zone setting.
	ID param.Field[ZonesMobileRedirectID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesMobileRedirectValueParam] `json:"value,required"`
}

func (r ZonesMobileRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMobileRedirectParam) implementsZonesSettingEditParamsItem() {}

// Current value of the zone setting.
type ZonesMobileRedirectValueParam struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZonesMobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI param.Field[bool] `json:"strip_uri"`
}

func (r ZonesMobileRedirectValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingMobileRedirectEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMobileRedirectEditParamsValue] `json:"value,required"`
}

func (r SettingMobileRedirectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMobileRedirectEditParamsValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[SettingMobileRedirectEditParamsValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI param.Field[bool] `json:"strip_uri"`
}

func (r SettingMobileRedirectEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type SettingMobileRedirectEditParamsValueStatus string

const (
	SettingMobileRedirectEditParamsValueStatusOn  SettingMobileRedirectEditParamsValueStatus = "on"
	SettingMobileRedirectEditParamsValueStatusOff SettingMobileRedirectEditParamsValueStatus = "off"
)

func (r SettingMobileRedirectEditParamsValueStatus) IsKnown() bool {
	switch r {
	case SettingMobileRedirectEditParamsValueStatusOn, SettingMobileRedirectEditParamsValueStatusOff:
		return true
	}
	return false
}

type SettingMobileRedirectEditResponseEnvelope struct {
	Errors   []SettingMobileRedirectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMobileRedirectEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZonesMobileRedirect                           `json:"result"`
	JSON   settingMobileRedirectEditResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectEditResponseEnvelope]
type settingMobileRedirectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectEditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingMobileRedirectEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMobileRedirectEditResponseEnvelopeErrors]
type settingMobileRedirectEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectEditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingMobileRedirectEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMobileRedirectEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMobileRedirectEditResponseEnvelopeMessages]
type settingMobileRedirectEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMobileRedirectGetResponseEnvelope struct {
	Errors   []SettingMobileRedirectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMobileRedirectGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result ZonesMobileRedirect                          `json:"result"`
	JSON   settingMobileRedirectGetResponseEnvelopeJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMobileRedirectGetResponseEnvelope]
type settingMobileRedirectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingMobileRedirectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingMobileRedirectGetResponseEnvelopeErrors]
type settingMobileRedirectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMobileRedirectGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingMobileRedirectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMobileRedirectGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingMobileRedirectGetResponseEnvelopeMessages]
type settingMobileRedirectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMobileRedirectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMobileRedirectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}