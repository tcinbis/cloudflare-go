// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DestinationWebhookService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDestinationWebhookService] method
// instead.
type DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewDestinationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationWebhookService(opts ...option.RequestOption) (r *DestinationWebhookService) {
	r = &DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *DestinationWebhookService) New(ctx context.Context, params DestinationWebhookNewParams, opts ...option.RequestOption) (res *DestinationWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a webhook destination.
func (r *DestinationWebhookService) Update(ctx context.Context, webhookID string, params DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *DestinationWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all configured webhook destinations.
func (r *DestinationWebhookService) List(ctx context.Context, query DestinationWebhookListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AlertingWebhooks], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", query.AccountID)
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

// Gets a list of all configured webhook destinations.
func (r *DestinationWebhookService) ListAutoPaging(ctx context.Context, query DestinationWebhookListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AlertingWebhooks] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured webhook destination.
func (r *DestinationWebhookService) Delete(ctx context.Context, webhookID string, body DestinationWebhookDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef167, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single webhooks destination.
func (r *DestinationWebhookService) Get(ctx context.Context, webhookID string, query DestinationWebhookGetParams, opts ...option.RequestOption) (res *AlertingWebhooks, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingWebhooks struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AlertingWebhooksType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string               `json:"url"`
	JSON alertingWebhooksJSON `json:"-"`
}

// alertingWebhooksJSON contains the JSON metadata for the struct
// [AlertingWebhooks]
type alertingWebhooksJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingWebhooks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertingWebhooksJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type AlertingWebhooksType string

const (
	AlertingWebhooksTypeSlack   AlertingWebhooksType = "slack"
	AlertingWebhooksTypeGeneric AlertingWebhooksType = "generic"
	AlertingWebhooksTypeGchat   AlertingWebhooksType = "gchat"
)

func (r AlertingWebhooksType) IsKnown() bool {
	switch r {
	case AlertingWebhooksTypeSlack, AlertingWebhooksTypeGeneric, AlertingWebhooksTypeGchat:
		return true
	}
	return false
}

type DestinationWebhookNewResponse struct {
	// UUID
	ID   string                            `json:"id"`
	JSON destinationWebhookNewResponseJSON `json:"-"`
}

// destinationWebhookNewResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookNewResponse]
type destinationWebhookNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookUpdateResponse struct {
	// UUID
	ID   string                               `json:"id"`
	JSON destinationWebhookUpdateResponseJSON `json:"-"`
}

// destinationWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [DestinationWebhookUpdateResponse]
type destinationWebhookUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   DestinationWebhookNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookNewResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookNewResponseEnvelope]
type destinationWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookNewResponseEnvelopeSuccess bool

const (
	DestinationWebhookNewResponseEnvelopeSuccessTrue DestinationWebhookNewResponseEnvelopeSuccess = true
)

func (r DestinationWebhookNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DestinationWebhookUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   DestinationWebhookUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookUpdateResponseEnvelope]
type destinationWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookUpdateResponseEnvelopeSuccess bool

const (
	DestinationWebhookUpdateResponseEnvelopeSuccessTrue DestinationWebhookUpdateResponseEnvelopeSuccess = true
)

func (r DestinationWebhookUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   shared.UnnamedSchemaRef167 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationWebhookDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationWebhookDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationWebhookDeleteResponseEnvelopeJSON       `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookDeleteResponseEnvelope]
type destinationWebhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookDeleteResponseEnvelopeSuccess bool

const (
	DestinationWebhookDeleteResponseEnvelopeSuccessTrue DestinationWebhookDeleteResponseEnvelopeSuccess = true
)

func (r DestinationWebhookDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationWebhookDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       destinationWebhookDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationWebhookDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DestinationWebhookDeleteResponseEnvelopeResultInfo]
type destinationWebhookDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DestinationWebhookGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationWebhookGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AlertingWebhooks      `json:"result,required"`
	// Whether the API call was successful
	Success DestinationWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    destinationWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// destinationWebhookGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationWebhookGetResponseEnvelope]
type destinationWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationWebhookGetResponseEnvelopeSuccess bool

const (
	DestinationWebhookGetResponseEnvelopeSuccessTrue DestinationWebhookGetResponseEnvelopeSuccess = true
)

func (r DestinationWebhookGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationWebhookGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
