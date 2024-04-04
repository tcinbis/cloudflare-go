// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ScriptTailService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewScriptTailService] method instead.
type ScriptTailService struct {
	Options []option.RequestOption
}

// NewScriptTailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptTailService(opts ...option.RequestOption) (r *ScriptTailService) {
	r = &ScriptTailService{}
	r.Options = opts
	return
}

// Starts a tail that receives logs and exception from a Worker.
func (r *ScriptTailService) New(ctx context.Context, scriptName string, params ScriptTailNewParams, opts ...option.RequestOption) (res *ScriptTailNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptTailNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a tail from a Worker.
func (r *ScriptTailService) Delete(ctx context.Context, scriptName string, id string, params ScriptTailDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef171, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptTailDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails/%s", params.AccountID, scriptName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get list of tails currently deployed on a Worker.
func (r *ScriptTailService) Get(ctx context.Context, scriptName string, query ScriptTailGetParams, opts ...option.RequestOption) (res *ScriptTailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptTailGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/tails", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptTailNewResponse struct {
	ID        interface{}               `json:"id"`
	ExpiresAt interface{}               `json:"expires_at"`
	URL       interface{}               `json:"url"`
	JSON      scriptTailNewResponseJSON `json:"-"`
}

// scriptTailNewResponseJSON contains the JSON metadata for the struct
// [ScriptTailNewResponse]
type scriptTailNewResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailGetResponse struct {
	ID        interface{}               `json:"id"`
	ExpiresAt interface{}               `json:"expires_at"`
	URL       interface{}               `json:"url"`
	JSON      scriptTailGetResponseJSON `json:"-"`
}

// scriptTailGetResponseJSON contains the JSON metadata for the struct
// [ScriptTailGetResponse]
type scriptTailGetResponseJSON struct {
	ID          apijson.Field
	ExpiresAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptTailNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ScriptTailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptTailNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ScriptTailNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptTailNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptTailNewResponseEnvelopeJSON    `json:"-"`
}

// scriptTailNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailNewResponseEnvelope]
type scriptTailNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailNewResponseEnvelopeSuccess bool

const (
	ScriptTailNewResponseEnvelopeSuccessTrue ScriptTailNewResponseEnvelopeSuccess = true
)

func (r ScriptTailNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptTailDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ScriptTailDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptTailDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   shared.UnnamedSchemaRef171 `json:"result,required"`
	// Whether the API call was successful
	Success ScriptTailDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptTailDeleteResponseEnvelopeJSON    `json:"-"`
}

// scriptTailDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailDeleteResponseEnvelope]
type scriptTailDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailDeleteResponseEnvelopeSuccess bool

const (
	ScriptTailDeleteResponseEnvelopeSuccessTrue ScriptTailDeleteResponseEnvelopeSuccess = true
)

func (r ScriptTailDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptTailGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptTailGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ScriptTailGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptTailGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptTailGetResponseEnvelopeJSON    `json:"-"`
}

// scriptTailGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptTailGetResponseEnvelope]
type scriptTailGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptTailGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptTailGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptTailGetResponseEnvelopeSuccess bool

const (
	ScriptTailGetResponseEnvelopeSuccessTrue ScriptTailGetResponseEnvelopeSuccess = true
)

func (r ScriptTailGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptTailGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
