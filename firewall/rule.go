// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
func (r *RuleService) New(ctx context.Context, zoneIdentifier string, body RuleNewParams, opts ...option.RequestOption) (res *[]Rule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing firewall rule.
func (r *RuleService) Update(ctx context.Context, zoneIdentifier string, id string, body RuleUpdateParams, opts ...option.RequestOption) (res *Rule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *RuleService) List(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Rule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
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

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *RuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Rule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing firewall rule.
func (r *RuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body RuleDeleteParams, opts ...option.RequestOption) (res *Rule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of an existing firewall rule.
func (r *RuleService) Edit(ctx context.Context, zoneIdentifier string, id string, body RuleEditParams, opts ...option.RequestOption) (res *[]Rule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a firewall rule.
func (r *RuleService) Get(ctx context.Context, zoneIdentifier string, params RuleGetParams, opts ...option.RequestOption) (res *Rule, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A list of products to bypass for a request when using the `bypass` action.
type ProductsItem string

const (
	ProductsItemZoneLockdown  ProductsItem = "zoneLockdown"
	ProductsItemUABlock       ProductsItem = "uaBlock"
	ProductsItemBic           ProductsItem = "bic"
	ProductsItemHot           ProductsItem = "hot"
	ProductsItemSecurityLevel ProductsItem = "securityLevel"
	ProductsItemRateLimit     ProductsItem = "rateLimit"
	ProductsItemWAF           ProductsItem = "waf"
)

func (r ProductsItem) IsKnown() bool {
	switch r {
	case ProductsItemZoneLockdown, ProductsItemUABlock, ProductsItemBic, ProductsItemHot, ProductsItemSecurityLevel, ProductsItemRateLimit, ProductsItemWAF:
		return true
	}
	return false
}

type Rule struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string     `json:"description"`
	Filter      RuleFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64        `json:"priority"`
	Products []ProductsItem `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string   `json:"ref"`
	JSON ruleJSON `json:"-"`
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleAction string

const (
	RuleActionBlock            RuleAction = "block"
	RuleActionChallenge        RuleAction = "challenge"
	RuleActionJsChallenge      RuleAction = "js_challenge"
	RuleActionManagedChallenge RuleAction = "managed_challenge"
	RuleActionAllow            RuleAction = "allow"
	RuleActionLog              RuleAction = "log"
	RuleActionBypass           RuleAction = "bypass"
)

func (r RuleAction) IsKnown() bool {
	switch r {
	case RuleActionBlock, RuleActionChallenge, RuleActionJsChallenge, RuleActionManagedChallenge, RuleActionAllow, RuleActionLog, RuleActionBypass:
		return true
	}
	return false
}

type RuleFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool           `json:"deleted"`
	JSON    ruleFilterJSON `json:"-"`
	union   RuleFilterUnion
}

// ruleFilterJSON contains the JSON metadata for the struct [RuleFilter]
type ruleFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleFilter) AsUnion() RuleFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleFilterFirewallFilter] or
// [firewall.DeletedFilter].
type RuleFilterUnion interface {
	implementsFirewallRuleFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeletedFilter{}),
		},
	)
}

type RuleFilterFirewallFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                       `json:"ref"`
	JSON ruleFilterFirewallFilterJSON `json:"-"`
}

// ruleFilterFirewallFilterJSON contains the JSON metadata for the struct
// [RuleFilterFirewallFilter]
type ruleFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleFilterFirewallFilter) implementsFirewallRuleFilter() {}

type RuleParam struct {
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action param.Field[RuleAction] `json:"action"`
	// An informative summary of the firewall rule.
	Description param.Field[string]               `json:"description"`
	Filter      param.Field[RuleFilterUnionParam] `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority param.Field[float64]        `json:"priority"`
	Products param.Field[[]ProductsItem] `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref param.Field[string] `json:"ref"`
}

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleFilterParam struct {
	// An informative summary of the filter.
	Description param.Field[string] `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref param.Field[string] `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted param.Field[bool] `json:"deleted"`
}

func (r RuleFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFilterParam) implementsFirewallRuleFilterUnionParam() {}

// Satisfied by [firewall.RuleFilterFirewallFilterParam],
// [firewall.DeletedFilterParam], [RuleFilterParam].
type RuleFilterUnionParam interface {
	implementsFirewallRuleFilterUnionParam()
}

type RuleFilterFirewallFilterParam struct {
	// An informative summary of the filter.
	Description param.Field[string] `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref param.Field[string] `json:"ref"`
}

func (r RuleFilterFirewallFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleFilterFirewallFilterParam) implementsFirewallRuleFilterUnionParam() {}

type DeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool              `json:"deleted,required"`
	JSON    deletedFilterJSON `json:"-"`
}

// deletedFilterJSON contains the JSON metadata for the struct [DeletedFilter]
type deletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r DeletedFilter) implementsFirewallRuleFilter() {}

type RuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []Rule                                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleNewResponseEnvelopeJSON       `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       ruleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeResultInfo]
type ruleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Rule                                                      `json:"result,required"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r RuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Rule                                                      `json:"result,required"`
	// Whether the API call was successful
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []Rule                                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleEditResponseEnvelopeJSON       `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       ruleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeResultInfo]
type ruleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleGetParams struct {
	// The unique identifier of the firewall rule.
	PathID param.Field[string] `path:"id,required"`
	// The unique identifier of the firewall rule.
	QueryID param.Field[string] `query:"id"`
}

// URLQuery serializes [RuleGetParams]'s query parameters as `url.Values`.
func (r RuleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Rule                                                      `json:"result,required"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
