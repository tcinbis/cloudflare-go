// File generated from our OpenAPI spec by Stainless.

package logs

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ControlService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewControlService] method instead.
type ControlService struct {
	Options   []option.RequestOption
	Retention *ControlRetentionService
	Cmb       *ControlCmbService
}

// NewControlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewControlService(opts ...option.RequestOption) (r *ControlService) {
	r = &ControlService{}
	r.Options = opts
	r.Retention = NewControlRetentionService(opts...)
	r.Cmb = NewControlCmbService(opts...)
	return
}