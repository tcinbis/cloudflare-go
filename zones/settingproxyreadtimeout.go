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

// SettingProxyReadTimeoutService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingProxyReadTimeoutService] method instead.
type SettingProxyReadTimeoutService struct {
	Options []option.RequestOption
}

// NewSettingProxyReadTimeoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingProxyReadTimeoutService(opts ...option.RequestOption) (r *SettingProxyReadTimeoutService) {
	r = &SettingProxyReadTimeoutService{}
	r.Options = opts
	return
}

// Maximum time between two read operations from origin.
func (r *SettingProxyReadTimeoutService) Edit(ctx context.Context, params SettingProxyReadTimeoutEditParams, opts ...option.RequestOption) (res *ZonesProxyReadTimeout, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
func (r *SettingProxyReadTimeoutService) Get(ctx context.Context, query SettingProxyReadTimeoutGetParams, opts ...option.RequestOption) (res *ZonesProxyReadTimeout, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingProxyReadTimeoutGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
type ZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesProxyReadTimeoutJSON `json:"-"`
}

// zonesProxyReadTimeoutJSON contains the JSON metadata for the struct
// [ZonesProxyReadTimeout]
type zonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r ZonesProxyReadTimeout) implementsZonesSettingEditResponse() {}

func (r ZonesProxyReadTimeout) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesProxyReadTimeoutID string

const (
	ZonesProxyReadTimeoutIDProxyReadTimeout ZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesProxyReadTimeoutEditable bool

const (
	ZonesProxyReadTimeoutEditableTrue  ZonesProxyReadTimeoutEditable = true
	ZonesProxyReadTimeoutEditableFalse ZonesProxyReadTimeoutEditable = false
)

// Maximum time between two read operations from origin.
type ZonesProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ZonesProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesProxyReadTimeoutParam) implementsZonesSettingEditParamsItem() {}

type SettingProxyReadTimeoutEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum time between two read operations from origin.
	Value param.Field[ZonesProxyReadTimeoutParam] `json:"value,required"`
}

func (r SettingProxyReadTimeoutEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingProxyReadTimeoutEditResponseEnvelope struct {
	Errors   []SettingProxyReadTimeoutEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingProxyReadTimeoutEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZonesProxyReadTimeout                           `json:"result"`
	JSON   settingProxyReadTimeoutEditResponseEnvelopeJSON `json:"-"`
}

// settingProxyReadTimeoutEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingProxyReadTimeoutEditResponseEnvelope]
type settingProxyReadTimeoutEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingProxyReadTimeoutEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingProxyReadTimeoutEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingProxyReadTimeoutEditResponseEnvelopeErrors]
type settingProxyReadTimeoutEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingProxyReadTimeoutEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingProxyReadTimeoutEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingProxyReadTimeoutEditResponseEnvelopeMessages]
type settingProxyReadTimeoutEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingProxyReadTimeoutGetResponseEnvelope struct {
	Errors   []SettingProxyReadTimeoutGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingProxyReadTimeoutGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZonesProxyReadTimeout                          `json:"result"`
	JSON   settingProxyReadTimeoutGetResponseEnvelopeJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingProxyReadTimeoutGetResponseEnvelope]
type settingProxyReadTimeoutGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingProxyReadTimeoutGetResponseEnvelopeErrors]
type settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingProxyReadTimeoutGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingProxyReadTimeoutGetResponseEnvelopeMessages]
type settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingProxyReadTimeoutGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingProxyReadTimeoutGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}