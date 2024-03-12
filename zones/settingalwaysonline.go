// File generated from our OpenAPI spec by Stainless.

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

// SettingAlwaysOnlineService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAlwaysOnlineService]
// method instead.
type SettingAlwaysOnlineService struct {
	Options []option.RequestOption
}

// NewSettingAlwaysOnlineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAlwaysOnlineService(opts ...option.RequestOption) (r *SettingAlwaysOnlineService) {
	r = &SettingAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *SettingAlwaysOnlineService) Edit(ctx context.Context, params SettingAlwaysOnlineEditParams, opts ...option.RequestOption) (res *ZonesAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *SettingAlwaysOnlineService) Get(ctx context.Context, query SettingAlwaysOnlineGetParams, opts ...option.RequestOption) (res *ZonesAlwaysOnline, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAlwaysOnlineGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_online", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesAlwaysOnlineJSON `json:"-"`
}

// zonesAlwaysOnlineJSON contains the JSON metadata for the struct
// [ZonesAlwaysOnline]
type zonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r ZonesAlwaysOnline) implementsZonesSettingEditResponse() {}

func (r ZonesAlwaysOnline) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesAlwaysOnlineID string

const (
	ZonesAlwaysOnlineIDAlwaysOnline ZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZonesAlwaysOnlineValue string

const (
	ZonesAlwaysOnlineValueOn  ZonesAlwaysOnlineValue = "on"
	ZonesAlwaysOnlineValueOff ZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesAlwaysOnlineEditable bool

const (
	ZonesAlwaysOnlineEditableTrue  ZonesAlwaysOnlineEditable = true
	ZonesAlwaysOnlineEditableFalse ZonesAlwaysOnlineEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnlineParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAlwaysOnlineValue] `json:"value,required"`
}

func (r ZonesAlwaysOnlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAlwaysOnlineParam) implementsZonesSettingEditParamsItem() {}

type SettingAlwaysOnlineEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingAlwaysOnlineEditParamsValue] `json:"value,required"`
}

func (r SettingAlwaysOnlineEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingAlwaysOnlineEditParamsValue string

const (
	SettingAlwaysOnlineEditParamsValueOn  SettingAlwaysOnlineEditParamsValue = "on"
	SettingAlwaysOnlineEditParamsValueOff SettingAlwaysOnlineEditParamsValue = "off"
)

type SettingAlwaysOnlineEditResponseEnvelope struct {
	Errors   []SettingAlwaysOnlineEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysOnlineEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZonesAlwaysOnline                           `json:"result"`
	JSON   settingAlwaysOnlineEditResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineEditResponseEnvelope]
type settingAlwaysOnlineEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingAlwaysOnlineEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAlwaysOnlineEditResponseEnvelopeErrors]
type settingAlwaysOnlineEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingAlwaysOnlineEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysOnlineEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysOnlineEditResponseEnvelopeMessages]
type settingAlwaysOnlineEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAlwaysOnlineGetResponseEnvelope struct {
	Errors   []SettingAlwaysOnlineGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAlwaysOnlineGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result ZonesAlwaysOnline                          `json:"result"`
	JSON   settingAlwaysOnlineGetResponseEnvelopeJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAlwaysOnlineGetResponseEnvelope]
type settingAlwaysOnlineGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingAlwaysOnlineGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAlwaysOnlineGetResponseEnvelopeErrors]
type settingAlwaysOnlineGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAlwaysOnlineGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingAlwaysOnlineGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAlwaysOnlineGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAlwaysOnlineGetResponseEnvelopeMessages]
type settingAlwaysOnlineGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAlwaysOnlineGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAlwaysOnlineGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}