// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationService] method
// instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	return
}

// Lists organizations the user is associated with.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[IamOrganization], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/organizations"
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

// Lists organizations the user is associated with.
func (r *OrganizationService) ListAutoPaging(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[IamOrganization] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Removes association to an organization.
func (r *OrganizationService) Delete(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets a specific organization the user is associated with.
func (r *OrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OrganizationGetResponseEnvelope
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IamOrganization struct {
	// Identifier
	ID string `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Access permissions for this User.
	Permissions []string `json:"permissions"`
	// List of roles that a user has within an organization.
	Roles []string `json:"roles"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status IamOrganizationStatus `json:"status"`
	JSON   iamOrganizationJSON   `json:"-"`
}

// iamOrganizationJSON contains the JSON metadata for the struct [IamOrganization]
type iamOrganizationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamOrganizationJSON) RawJSON() string {
	return r.raw
}

// Whether the user is a member of the organization or has an inivitation pending.
type IamOrganizationStatus string

const (
	IamOrganizationStatusMember  IamOrganizationStatus = "member"
	IamOrganizationStatusInvited IamOrganizationStatus = "invited"
)

func (r IamOrganizationStatus) IsKnown() bool {
	switch r {
	case IamOrganizationStatusMember, IamOrganizationStatusInvited:
		return true
	}
	return false
}

type OrganizationDeleteResponse struct {
	// Identifier
	ID   string                         `json:"id"`
	JSON organizationDeleteResponseJSON `json:"-"`
}

// organizationDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponse]
type organizationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [user.OrganizationGetResponseUnknown] or
// [shared.UnionString].
type OrganizationGetResponse interface {
	ImplementsUserOrganizationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OrganizationGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OrganizationListParams struct {
	// Direction to order organizations.
	Direction param.Field[OrganizationListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[OrganizationListParamsMatch] `query:"match"`
	// Organization name.
	Name param.Field[string] `query:"name"`
	// Field to order organizations by.
	Order param.Field[OrganizationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of organizations per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status param.Field[OrganizationListParamsStatus] `query:"status"`
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order organizations.
type OrganizationListParamsDirection string

const (
	OrganizationListParamsDirectionAsc  OrganizationListParamsDirection = "asc"
	OrganizationListParamsDirectionDesc OrganizationListParamsDirection = "desc"
)

func (r OrganizationListParamsDirection) IsKnown() bool {
	switch r {
	case OrganizationListParamsDirectionAsc, OrganizationListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any).
type OrganizationListParamsMatch string

const (
	OrganizationListParamsMatchAny OrganizationListParamsMatch = "any"
	OrganizationListParamsMatchAll OrganizationListParamsMatch = "all"
)

func (r OrganizationListParamsMatch) IsKnown() bool {
	switch r {
	case OrganizationListParamsMatchAny, OrganizationListParamsMatchAll:
		return true
	}
	return false
}

// Field to order organizations by.
type OrganizationListParamsOrder string

const (
	OrganizationListParamsOrderID     OrganizationListParamsOrder = "id"
	OrganizationListParamsOrderName   OrganizationListParamsOrder = "name"
	OrganizationListParamsOrderStatus OrganizationListParamsOrder = "status"
)

func (r OrganizationListParamsOrder) IsKnown() bool {
	switch r {
	case OrganizationListParamsOrderID, OrganizationListParamsOrderName, OrganizationListParamsOrderStatus:
		return true
	}
	return false
}

// Whether the user is a member of the organization or has an inivitation pending.
type OrganizationListParamsStatus string

const (
	OrganizationListParamsStatusMember  OrganizationListParamsStatus = "member"
	OrganizationListParamsStatusInvited OrganizationListParamsStatus = "invited"
)

func (r OrganizationListParamsStatus) IsKnown() bool {
	switch r {
	case OrganizationListParamsStatusMember, OrganizationListParamsStatusInvited:
		return true
	}
	return false
}

type OrganizationGetResponseEnvelope struct {
	Errors   []OrganizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OrganizationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OrganizationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OrganizationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    organizationGetResponseEnvelopeJSON    `json:"-"`
}

// organizationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationGetResponseEnvelope]
type organizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    organizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// organizationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationGetResponseEnvelopeErrors]
type organizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    organizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// organizationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationGetResponseEnvelopeMessages]
type organizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OrganizationGetResponseEnvelopeSuccess bool

const (
	OrganizationGetResponseEnvelopeSuccessTrue OrganizationGetResponseEnvelopeSuccess = true
)

func (r OrganizationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}