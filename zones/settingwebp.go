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

// SettingWebpService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingWebpService] method
// instead.
type SettingWebpService struct {
	Options []option.RequestOption
}

// NewSettingWebpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingWebpService(opts ...option.RequestOption) (r *SettingWebpService) {
	r = &SettingWebpService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebpService) Edit(ctx context.Context, params SettingWebpEditParams, opts ...option.RequestOption) (res *ZonesWebp, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebpEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *SettingWebpService) Get(ctx context.Context, query SettingWebpGetParams, opts ...option.RequestOption) (res *ZonesWebp, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingWebpGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/webp", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebp struct {
	// ID of the zone setting.
	ID ZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesWebpJSON `json:"-"`
}

// zonesWebpJSON contains the JSON metadata for the struct [ZonesWebp]
type zonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWebpJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWebp) implementsZonesSettingEditResponse() {}

func (r ZonesWebp) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesWebpID string

const (
	ZonesWebpIDWebp ZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZonesWebpValue string

const (
	ZonesWebpValueOff ZonesWebpValue = "off"
	ZonesWebpValueOn  ZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesWebpEditable bool

const (
	ZonesWebpEditableTrue  ZonesWebpEditable = true
	ZonesWebpEditableFalse ZonesWebpEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebpParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebpID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesWebpValue] `json:"value,required"`
}

func (r ZonesWebpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebpParam) implementsZonesSettingEditParamsItem() {}

type SettingWebpEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingWebpEditParamsValue] `json:"value,required"`
}

func (r SettingWebpEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingWebpEditParamsValue string

const (
	SettingWebpEditParamsValueOff SettingWebpEditParamsValue = "off"
	SettingWebpEditParamsValueOn  SettingWebpEditParamsValue = "on"
)

type SettingWebpEditResponseEnvelope struct {
	Errors   []SettingWebpEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebpEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZonesWebp                           `json:"result"`
	JSON   settingWebpEditResponseEnvelopeJSON `json:"-"`
}

// settingWebpEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebpEditResponseEnvelope]
type settingWebpEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingWebpEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingWebpEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebpEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebpEditResponseEnvelopeErrors]
type settingWebpEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebpEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingWebpEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebpEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebpEditResponseEnvelopeMessages]
type settingWebpEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingWebpGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingWebpGetResponseEnvelope struct {
	Errors   []SettingWebpGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingWebpGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result ZonesWebp                          `json:"result"`
	JSON   settingWebpGetResponseEnvelopeJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingWebpGetResponseEnvelope]
type settingWebpGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingWebpGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    settingWebpGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingWebpGetResponseEnvelopeErrors]
type settingWebpGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingWebpGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingWebpGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingWebpGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingWebpGetResponseEnvelopeMessages]
type settingWebpGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingWebpGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingWebpGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}