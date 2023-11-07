// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingSslRecommenderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSslRecommenderService] method instead.
type ZoneSettingSslRecommenderService struct {
	Options []option.RequestOption
}

// NewZoneSettingSslRecommenderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSslRecommenderService(opts ...option.RequestOption) (r *ZoneSettingSslRecommenderService) {
	r = &ZoneSettingSslRecommenderService{}
	r.Options = opts
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZoneSettingSslRecommenderService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSslRecommenderUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSslRecommenderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZoneSettingSslRecommenderService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSslRecommenderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSslRecommenderUpdateResponse struct {
	Errors   []ZoneSettingSslRecommenderUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingSslRecommenderUpdateResponseMessage `json:"messages"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSslRecommenderUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSslRecommenderUpdateResponseJSON
}

// zoneSettingSslRecommenderUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSslRecommenderUpdateResponse]
type zoneSettingSslRecommenderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslRecommenderUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSslRecommenderUpdateResponseErrorJSON
}

// zoneSettingSslRecommenderUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSslRecommenderUpdateResponseError]
type zoneSettingSslRecommenderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslRecommenderUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSslRecommenderUpdateResponseMessageJSON
}

// zoneSettingSslRecommenderUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingSslRecommenderUpdateResponseMessage]
type zoneSettingSslRecommenderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSslRecommenderUpdateResponseResult struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSslRecommenderUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSslRecommenderUpdateResponseResultEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSslRecommenderUpdateResponseResultJSON
}

// zoneSettingSslRecommenderUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSslRecommenderUpdateResponseResult]
type zoneSettingSslRecommenderUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSslRecommenderUpdateResponseResultID string

const (
	ZoneSettingSslRecommenderUpdateResponseResultIDSslRecommender ZoneSettingSslRecommenderUpdateResponseResultID = "ssl_recommender"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSslRecommenderUpdateResponseResultEditable bool

const (
	ZoneSettingSslRecommenderUpdateResponseResultEditableTrue  ZoneSettingSslRecommenderUpdateResponseResultEditable = true
	ZoneSettingSslRecommenderUpdateResponseResultEditableFalse ZoneSettingSslRecommenderUpdateResponseResultEditable = false
)

type ZoneSettingSslRecommenderListResponse struct {
	Errors   []ZoneSettingSslRecommenderListResponseError   `json:"errors"`
	Messages []ZoneSettingSslRecommenderListResponseMessage `json:"messages"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result ZoneSettingSslRecommenderListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingSslRecommenderListResponseJSON
}

// zoneSettingSslRecommenderListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSslRecommenderListResponse]
type zoneSettingSslRecommenderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslRecommenderListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSslRecommenderListResponseErrorJSON
}

// zoneSettingSslRecommenderListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSslRecommenderListResponseError]
type zoneSettingSslRecommenderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSslRecommenderListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingSslRecommenderListResponseMessageJSON
}

// zoneSettingSslRecommenderListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingSslRecommenderListResponseMessage]
type zoneSettingSslRecommenderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSslRecommenderListResponseResult struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingSslRecommenderListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSslRecommenderListResponseResultEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSslRecommenderListResponseResultJSON
}

// zoneSettingSslRecommenderListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSslRecommenderListResponseResult]
type zoneSettingSslRecommenderListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSslRecommenderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSslRecommenderListResponseResultID string

const (
	ZoneSettingSslRecommenderListResponseResultIDSslRecommender ZoneSettingSslRecommenderListResponseResultID = "ssl_recommender"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSslRecommenderListResponseResultEditable bool

const (
	ZoneSettingSslRecommenderListResponseResultEditableTrue  ZoneSettingSslRecommenderListResponseResultEditable = true
	ZoneSettingSslRecommenderListResponseResultEditableFalse ZoneSettingSslRecommenderListResponseResultEditable = false
)

type ZoneSettingSslRecommenderUpdateParams struct {
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value param.Field[ZoneSettingSslRecommenderUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSslRecommenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingSslRecommenderUpdateParamsValue struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingSslRecommenderUpdateParamsValueID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingSslRecommenderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingSslRecommenderUpdateParamsValueID string

const (
	ZoneSettingSslRecommenderUpdateParamsValueIDSslRecommender ZoneSettingSslRecommenderUpdateParamsValueID = "ssl_recommender"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSslRecommenderUpdateParamsValueEditable bool

const (
	ZoneSettingSslRecommenderUpdateParamsValueEditableTrue  ZoneSettingSslRecommenderUpdateParamsValueEditable = true
	ZoneSettingSslRecommenderUpdateParamsValueEditableFalse ZoneSettingSslRecommenderUpdateParamsValueEditable = false
)