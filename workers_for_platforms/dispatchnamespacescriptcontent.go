// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers"
)

// DispatchNamespaceScriptContentService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptContentService] method instead.
type DispatchNamespaceScriptContentService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptContentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentService) {
	r = &DispatchNamespaceScriptContentService{}
	r.Options = opts
	return
}

// Put script content for a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptContentService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptContentUpdateParams, opts ...option.RequestOption) (res *workers.WorkersScript, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptContentUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch script content from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptContentService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptContentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/content", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DispatchNamespaceScriptContentUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A module comprising a Worker script, often a javascript file. Multiple modules
	// may be provided as separate named parts, but at least one module must be
	// present. This should be referenced either in the metadata as `main_module`
	// (esm)/`body_part` (service worker) or as a header `CF-WORKER-MAIN-MODULE-PART`
	// (esm) /`CF-WORKER-BODY-PART` (service worker) by part name.
	AnyPartName param.Field[[]io.Reader] `json:"<any part name>" format:"binary"`
	// JSON encoded metadata about the uploaded parts and Worker configuration.
	Metadata               param.Field[DispatchNamespaceScriptContentUpdateParamsMetadata] `json:"metadata"`
	CfWorkerBodyPart       param.Field[string]                                             `header:"CF-WORKER-BODY-PART"`
	CfWorkerMainModulePart param.Field[string]                                             `header:"CF-WORKER-MAIN-MODULE-PART"`
}

func (r DispatchNamespaceScriptContentUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// JSON encoded metadata about the uploaded parts and Worker configuration.
type DispatchNamespaceScriptContentUpdateParamsMetadata struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r DispatchNamespaceScriptContentUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptContentUpdateResponseEnvelope struct {
	Errors   []DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   workers.WorkersScript                                          `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptContentUpdateResponseEnvelope]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages]
type dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptContentUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptContentGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}