// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// GatewayAppTypeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayAppTypeService] method
// instead.
type GatewayAppTypeService struct {
	Options []option.RequestOption
}

// NewGatewayAppTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayAppTypeService(opts ...option.RequestOption) (r *GatewayAppTypeService) {
	r = &GatewayAppTypeService{}
	r.Options = opts
	return
}

// Fetches all application and application type mappings.
func (r *GatewayAppTypeService) List(ctx context.Context, query GatewayAppTypeListParams, opts ...option.RequestOption) (res *shared.SinglePage[ZeroTrustGatewayAppTypes], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/app_types", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches all application and application type mappings.
func (r *GatewayAppTypeService) ListAutoPaging(ctx context.Context, query GatewayAppTypeListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[ZeroTrustGatewayAppTypes] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Union satisfied by
// [zero_trust.ZeroTrustGatewayAppTypesZeroTrustGatewayApplication] or
// [zero_trust.ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType].
type ZeroTrustGatewayAppTypes interface {
	implementsZeroTrustZeroTrustGatewayAppTypes()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayAppTypes)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypesZeroTrustGatewayApplication{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType{}),
		},
	)
}

type ZeroTrustGatewayAppTypesZeroTrustGatewayApplication struct {
	// The identifier for this application. There is only one application per ID.
	ID int64 `json:"id"`
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ApplicationTypeID int64     `json:"application_type_id"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The name of the application or application type.
	Name string                                                  `json:"name"`
	JSON zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON `json:"-"`
}

// zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayAppTypesZeroTrustGatewayApplication]
type zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON struct {
	ID                apijson.Field
	ApplicationTypeID apijson.Field
	CreatedAt         apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypesZeroTrustGatewayApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypesZeroTrustGatewayApplicationJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypesZeroTrustGatewayApplication) implementsZeroTrustZeroTrustGatewayAppTypes() {
}

type ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType struct {
	// The identifier for the type of this application. There can be many applications
	// with the same type. This refers to the `id` of a returned application type.
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A short summary of applications with this type.
	Description string `json:"description"`
	// The name of the application or application type.
	Name string                                                      `json:"name"`
	JSON zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON `json:"-"`
}

// zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType]
type zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayAppTypesZeroTrustGatewayApplicationTypeJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGatewayAppTypesZeroTrustGatewayApplicationType) implementsZeroTrustZeroTrustGatewayAppTypes() {
}

type GatewayAppTypeListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
