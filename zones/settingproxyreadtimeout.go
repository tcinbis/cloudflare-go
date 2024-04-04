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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingProxyReadTimeoutService) Edit(ctx context.Context, params SettingProxyReadTimeoutEditParams, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeout, err error) {
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
func (r *SettingProxyReadTimeoutService) Get(ctx context.Context, query SettingProxyReadTimeoutGetParams, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeout, err error) {
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
type ZoneSettingProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutJSON contains the JSON metadata for the struct
// [ZoneSettingProxyReadTimeout]
type zoneSettingProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingProxyReadTimeout) implementsZonesSettingEditResponse() {}

func (r ZoneSettingProxyReadTimeout) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutID string

const (
	ZoneSettingProxyReadTimeoutIDProxyReadTimeout ZoneSettingProxyReadTimeoutID = "proxy_read_timeout"
)

func (r ZoneSettingProxyReadTimeoutID) IsKnown() bool {
	switch r {
	case ZoneSettingProxyReadTimeoutIDProxyReadTimeout:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutEditable bool

const (
	ZoneSettingProxyReadTimeoutEditableTrue  ZoneSettingProxyReadTimeoutEditable = true
	ZoneSettingProxyReadTimeoutEditableFalse ZoneSettingProxyReadTimeoutEditable = false
)

func (r ZoneSettingProxyReadTimeoutEditable) IsKnown() bool {
	switch r {
	case ZoneSettingProxyReadTimeoutEditableTrue, ZoneSettingProxyReadTimeoutEditableFalse:
		return true
	}
	return false
}

// Maximum time between two read operations from origin.
type ZoneSettingProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingProxyReadTimeoutParam) implementsZonesSettingEditParamsItem() {}

type SettingProxyReadTimeoutEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum time between two read operations from origin.
	Value param.Field[ZoneSettingProxyReadTimeoutParam] `json:"value,required"`
}

func (r SettingProxyReadTimeoutEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingProxyReadTimeoutEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZoneSettingProxyReadTimeout                     `json:"result"`
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

type SettingProxyReadTimeoutGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingProxyReadTimeoutGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZoneSettingProxyReadTimeout                    `json:"result"`
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
