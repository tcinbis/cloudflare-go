// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCachService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneCachService] method instead.
type ZoneCachService struct {
	Options                         []option.RequestOption
	CacheReserves                   *ZoneCachCacheReserveService
	TieredCacheSmartTopologyEnables *ZoneCachTieredCacheSmartTopologyEnableService
	Variants                        *ZoneCachVariantService
}

// NewZoneCachService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCachService(opts ...option.RequestOption) (r *ZoneCachService) {
	r = &ZoneCachService{}
	r.Options = opts
	r.CacheReserves = NewZoneCachCacheReserveService(opts...)
	r.TieredCacheSmartTopologyEnables = NewZoneCachTieredCacheSmartTopologyEnableService(opts...)
	r.Variants = NewZoneCachVariantService(opts...)
	return
}
