// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ScriptBindingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptBindingService] method
// instead.
type ScriptBindingService struct {
	Options []option.RequestOption
}

// NewScriptBindingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptBindingService(opts ...option.RequestOption) (r *ScriptBindingService) {
	r = &ScriptBindingService{}
	r.Options = opts
	return
}

// List the bindings for a Workers script.
func (r *ScriptBindingService) Get(ctx context.Context, query ScriptBindingGetParams, opts ...option.RequestOption) (res *[]ScriptBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptBindingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptBindingGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type  ScriptBindingGetResponseType `json:"type,required"`
	JSON  scriptBindingGetResponseJSON `json:"-"`
	union ScriptBindingGetResponseUnion
}

// scriptBindingGetResponseJSON contains the JSON metadata for the struct
// [ScriptBindingGetResponse]
type scriptBindingGetResponseJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r scriptBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ScriptBindingGetResponse) AsUnion() ScriptBindingGetResponseUnion {
	return r.union
}

// Union satisfied by [workers.KVNamespaceBinding] or
// [workers.ScriptBindingGetResponseWorkersWasmModuleBinding].
type ScriptBindingGetResponseUnion interface {
	implementsWorkersScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptBindingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptBindingGetResponseWorkersWasmModuleBinding{}),
		},
	)
}

type ScriptBindingGetResponseWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptBindingGetResponseWorkersWasmModuleBindingType `json:"type,required"`
	JSON scriptBindingGetResponseWorkersWasmModuleBindingJSON `json:"-"`
}

// scriptBindingGetResponseWorkersWasmModuleBindingJSON contains the JSON metadata
// for the struct [ScriptBindingGetResponseWorkersWasmModuleBinding]
type scriptBindingGetResponseWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptBindingGetResponseWorkersWasmModuleBinding) implementsWorkersScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type ScriptBindingGetResponseWorkersWasmModuleBindingType string

const (
	ScriptBindingGetResponseWorkersWasmModuleBindingTypeWasmModule ScriptBindingGetResponseWorkersWasmModuleBindingType = "wasm_module"
)

func (r ScriptBindingGetResponseWorkersWasmModuleBindingType) IsKnown() bool {
	switch r {
	case ScriptBindingGetResponseWorkersWasmModuleBindingTypeWasmModule:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type ScriptBindingGetResponseType string

const (
	ScriptBindingGetResponseTypeKVNamespace ScriptBindingGetResponseType = "kv_namespace"
	ScriptBindingGetResponseTypeWasmModule  ScriptBindingGetResponseType = "wasm_module"
)

func (r ScriptBindingGetResponseType) IsKnown() bool {
	switch r {
	case ScriptBindingGetResponseTypeKVNamespace, ScriptBindingGetResponseTypeWasmModule:
		return true
	}
	return false
}

type ScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptBindingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []ScriptBindingGetResponse                                `json:"result,required"`
	// Whether the API call was successful
	Success ScriptBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptBindingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptBindingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptBindingGetResponseEnvelope]
type scriptBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptBindingGetResponseEnvelopeSuccess bool

const (
	ScriptBindingGetResponseEnvelopeSuccessTrue ScriptBindingGetResponseEnvelopeSuccess = true
)

func (r ScriptBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
