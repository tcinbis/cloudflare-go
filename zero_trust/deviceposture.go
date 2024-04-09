// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DevicePostureService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePostureService] method
// instead.
type DevicePostureService struct {
	Options      []option.RequestOption
	Integrations *DevicePostureIntegrationService
}

// NewDevicePostureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicePostureService(opts ...option.RequestOption) (r *DevicePostureService) {
	r = &DevicePostureService{}
	r.Options = opts
	r.Integrations = NewDevicePostureIntegrationService(opts...)
	return
}

// Creates a new device posture rule.
func (r *DevicePostureService) New(ctx context.Context, params DevicePostureNewParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a device posture rule.
func (r *DevicePostureService) Update(ctx context.Context, ruleID string, params DevicePostureUpdateParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) List(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DevicePostureRule], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/posture", query.AccountID)
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

// Fetches device posture rules for a Zero Trust account.
func (r *DevicePostureService) ListAutoPaging(ctx context.Context, query DevicePostureListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DevicePostureRule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device posture rule.
func (r *DevicePostureService) Delete(ctx context.Context, ruleID string, params DevicePostureDeleteParams, opts ...option.RequestOption) (res *DevicePostureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single device posture rule.
func (r *DevicePostureService) Get(ctx context.Context, ruleID string, query DevicePostureGetParams, opts ...option.RequestOption) (res *DevicePostureRule, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePostureGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/posture/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The value to be checked against.
type DeviceInput struct {
	// Whether or not file exists
	Exists bool `json:"exists"`
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system"`
	// File path.
	Path string `json:"path"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string `json:"thumbprint"`
	// List ID.
	ID string `json:"id"`
	// Domain
	Domain string `json:"domain"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string `json:"os_version_extra"`
	// Version of OS
	Version string `json:"version"`
	// Enabled
	Enabled    bool        `json:"enabled"`
	CheckDisks interface{} `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll bool `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn string `json:"cn"`
	// Compliance Status
	ComplianceStatus DeviceInputComplianceStatus `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID string `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DeviceInputState `json:"state"`
	// Version Operator
	VersionOperator DeviceInputVersionOperator `json:"versionOperator"`
	// Count Operator
	CountOperator DeviceInputCountOperator `json:"countOperator"`
	// The Number of Issues.
	IssueCount string `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DeviceInputRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DeviceInputScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64 `json:"total_score"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DeviceInputNetworkStatus `json:"network_status"`
	JSON          deviceInputJSON          `json:"-"`
	union         DeviceInputUnion
}

// deviceInputJSON contains the JSON metadata for the struct [DeviceInput]
type deviceInputJSON struct {
	Exists           apijson.Field
	OperatingSystem  apijson.Field
	Path             apijson.Field
	Sha256           apijson.Field
	Thumbprint       apijson.Field
	ID               apijson.Field
	Domain           apijson.Field
	Operator         apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	Version          apijson.Field
	Enabled          apijson.Field
	CheckDisks       apijson.Field
	RequireAll       apijson.Field
	CertificateID    apijson.Field
	Cn               apijson.Field
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	LastSeen         apijson.Field
	OS               apijson.Field
	Overall          apijson.Field
	SensorConfig     apijson.Field
	State            apijson.Field
	VersionOperator  apijson.Field
	CountOperator    apijson.Field
	IssueCount       apijson.Field
	EidLastSeen      apijson.Field
	RiskLevel        apijson.Field
	ScoreOperator    apijson.Field
	TotalScore       apijson.Field
	ActiveThreats    apijson.Field
	Infected         apijson.Field
	IsActive         apijson.Field
	NetworkStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r deviceInputJSON) RawJSON() string {
	return r.raw
}

func (r *DeviceInput) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DeviceInput) AsUnion() DeviceInputUnion {
	return r.union
}

// The value to be checked against.
//
// Union satisfied by [zero_trust.DeviceInputTeamsDevicesFileInputRequest],
// [zero_trust.DeviceInputTeamsDevicesUniqueClientIDInputRequest],
// [zero_trust.DeviceInputTeamsDevicesDomainJoinedInputRequest],
// [zero_trust.DeviceInputTeamsDevicesOSVersionInputRequest],
// [zero_trust.DeviceInputTeamsDevicesFirewallInputRequest],
// [zero_trust.DeviceInputTeamsDevicesSentineloneInputRequest],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequest],
// [zero_trust.DeviceInputTeamsDevicesDiskEncryptionInputRequest],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequest],
// [zero_trust.DeviceInputTeamsDevicesClientCertificateInputRequest],
// [zero_trust.DeviceInputTeamsDevicesWorkspaceOneInputRequest],
// [zero_trust.DeviceInputTeamsDevicesCrowdstrikeInputRequest],
// [zero_trust.DeviceInputTeamsDevicesIntuneInputRequest],
// [zero_trust.DeviceInputTeamsDevicesKolideInputRequest],
// [zero_trust.DeviceInputTeamsDevicesTaniumInputRequest] or
// [zero_trust.DeviceInputTeamsDevicesSentineloneS2sInputRequest].
type DeviceInputUnion interface {
	implementsZeroTrustDeviceInput()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceInputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesFileInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesUniqueClientIDInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesDomainJoinedInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesOSVersionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesFirewallInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesSentineloneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesCarbonblackInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesDiskEncryptionInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesApplicationInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesClientCertificateInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesWorkspaceOneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesCrowdstrikeInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesIntuneInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesKolideInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesTaniumInputRequest{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeviceInputTeamsDevicesSentineloneS2sInputRequest{}),
		},
	)
}

type DeviceInputTeamsDevicesFileInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// Whether or not file exists
	Exists bool `json:"exists"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                      `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesFileInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesFileInputRequestJSON contains the JSON metadata for the
// struct [DeviceInputTeamsDevicesFileInputRequest]
type deviceInputTeamsDevicesFileInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Exists          apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesFileInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesFileInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesFileInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesUniqueClientIDInputRequest struct {
	// List ID.
	ID string `json:"id,required"`
	// Operating System
	OperatingSystem DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            deviceInputTeamsDevicesUniqueClientIDInputRequestJSON            `json:"-"`
}

// deviceInputTeamsDevicesUniqueClientIDInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesUniqueClientIDInputRequest]
type deviceInputTeamsDevicesUniqueClientIDInputRequestJSON struct {
	ID              apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesUniqueClientIDInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesUniqueClientIDInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesUniqueClientIDInputRequest) implementsZeroTrustDeviceInput() {}

// Operating System
type DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid  DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "android"
	DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos      DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "ios"
	DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem = "chromeos"
)

func (r DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemAndroid, DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemIos, DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystemChromeos:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesDomainJoinedInputRequest struct {
	// Operating System
	OperatingSystem DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystem `json:"operating_system,required"`
	// Domain
	Domain string                                              `json:"domain"`
	JSON   deviceInputTeamsDevicesDomainJoinedInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesDomainJoinedInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesDomainJoinedInputRequest]
type deviceInputTeamsDevicesDomainJoinedInputRequestJSON struct {
	OperatingSystem apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesDomainJoinedInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesDomainJoinedInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesDomainJoinedInputRequest) implementsZeroTrustDeviceInput() {}

// Operating System
type DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystem = "windows"
)

func (r DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesOSVersionInputRequest struct {
	// Operating System
	OperatingSystem DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystem `json:"operating_system,required"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator,required"`
	// Version of OS
	Version string `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName string `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision string `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra string                                           `json:"os_version_extra"`
	JSON           deviceInputTeamsDevicesOSVersionInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesOSVersionInputRequestJSON contains the JSON metadata for
// the struct [DeviceInputTeamsDevicesOSVersionInputRequest]
type deviceInputTeamsDevicesOSVersionInputRequestJSON struct {
	OperatingSystem  apijson.Field
	Operator         apijson.Field
	Version          apijson.Field
	OSDistroName     apijson.Field
	OSDistroRevision apijson.Field
	OSVersionExtra   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesOSVersionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesOSVersionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesOSVersionInputRequest) implementsZeroTrustDeviceInput() {}

// Operating System
type DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystem = "windows"
)

func (r DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystemWindows:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesFirewallInputRequest struct {
	// Enabled
	Enabled bool `json:"enabled,required"`
	// Operating System
	OperatingSystem DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem `json:"operating_system,required"`
	JSON            deviceInputTeamsDevicesFirewallInputRequestJSON            `json:"-"`
}

// deviceInputTeamsDevicesFirewallInputRequestJSON contains the JSON metadata for
// the struct [DeviceInputTeamsDevicesFirewallInputRequest]
type deviceInputTeamsDevicesFirewallInputRequestJSON struct {
	Enabled         apijson.Field
	OperatingSystem apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesFirewallInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesFirewallInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesFirewallInputRequest) implementsZeroTrustDeviceInput() {}

// Operating System
type DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem string

const (
	DeviceInputTeamsDevicesFirewallInputRequestOperatingSystemWindows DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem = "windows"
	DeviceInputTeamsDevicesFirewallInputRequestOperatingSystemMac     DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem = "mac"
)

func (r DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesFirewallInputRequestOperatingSystemWindows, DeviceInputTeamsDevicesFirewallInputRequestOperatingSystemMac:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesSentineloneInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesSentineloneInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesSentineloneInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesSentineloneInputRequest]
type deviceInputTeamsDevicesSentineloneInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesSentineloneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesSentineloneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesSentineloneInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesCarbonblackInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// File path.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesCarbonblackInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesCarbonblackInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesCarbonblackInputRequest]
type deviceInputTeamsDevicesCarbonblackInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesCarbonblackInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesCarbonblackInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesDiskEncryptionInputRequest struct {
	// List of volume names to be checked for encryption.
	CheckDisks []string `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll bool                                                  `json:"requireAll"`
	JSON       deviceInputTeamsDevicesDiskEncryptionInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesDiskEncryptionInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesDiskEncryptionInputRequest]
type deviceInputTeamsDevicesDiskEncryptionInputRequestJSON struct {
	CheckDisks  apijson.Field
	RequireAll  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesDiskEncryptionInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesDiskEncryptionInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesDiskEncryptionInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesApplicationInputRequest struct {
	// Operating system
	OperatingSystem UnnamedSchemaRef41885dd46b9e0294254c49305a273681 `json:"operating_system,required"`
	// Path for the application.
	Path string `json:"path,required"`
	// SHA-256.
	Sha256 string `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint string                                             `json:"thumbprint"`
	JSON       deviceInputTeamsDevicesApplicationInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesApplicationInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesApplicationInputRequest]
type deviceInputTeamsDevicesApplicationInputRequestJSON struct {
	OperatingSystem apijson.Field
	Path            apijson.Field
	Sha256          apijson.Field
	Thumbprint      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesApplicationInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesApplicationInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesApplicationInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesClientCertificateInputRequest struct {
	// UUID of Cloudflare managed certificate.
	CertificateID string `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn   string                                                   `json:"cn,required"`
	JSON deviceInputTeamsDevicesClientCertificateInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesClientCertificateInputRequestJSON contains the JSON
// metadata for the struct [DeviceInputTeamsDevicesClientCertificateInputRequest]
type deviceInputTeamsDevicesClientCertificateInputRequestJSON struct {
	CertificateID apijson.Field
	Cn            apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesClientCertificateInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesClientCertificateInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesClientCertificateInputRequest) implementsZeroTrustDeviceInput() {}

type DeviceInputTeamsDevicesWorkspaceOneInputRequest struct {
	// Compliance Status
	ComplianceStatus DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                              `json:"connection_id,required"`
	JSON         deviceInputTeamsDevicesWorkspaceOneInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesWorkspaceOneInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesWorkspaceOneInputRequest]
type deviceInputTeamsDevicesWorkspaceOneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesWorkspaceOneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesWorkspaceOneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesWorkspaceOneInputRequest) implementsZeroTrustDeviceInput() {}

// Compliance Status
type DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus string

const (
	DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant    DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "compliant"
	DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "noncompliant"
	DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown      DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus = "unknown"
)

func (r DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusCompliant, DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusNoncompliant, DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatusUnknown:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesCrowdstrikeInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen string `json:"last_seen"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 `json:"operator"`
	// Os Version
	OS string `json:"os"`
	// overall
	Overall string `json:"overall"`
	// SensorConfig
	SensorConfig string `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State DeviceInputTeamsDevicesCrowdstrikeInputRequestState `json:"state"`
	// Version
	Version string `json:"version"`
	// Version Operator
	VersionOperator DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator `json:"versionOperator"`
	JSON            deviceInputTeamsDevicesCrowdstrikeInputRequestJSON            `json:"-"`
}

// deviceInputTeamsDevicesCrowdstrikeInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesCrowdstrikeInputRequest]
type deviceInputTeamsDevicesCrowdstrikeInputRequestJSON struct {
	ConnectionID    apijson.Field
	LastSeen        apijson.Field
	Operator        apijson.Field
	OS              apijson.Field
	Overall         apijson.Field
	SensorConfig    apijson.Field
	State           apijson.Field
	Version         apijson.Field
	VersionOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesCrowdstrikeInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesCrowdstrikeInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesCrowdstrikeInputRequest) implementsZeroTrustDeviceInput() {}

// For more details on state, please refer to the Crowdstrike documentation.
type DeviceInputTeamsDevicesCrowdstrikeInputRequestState string

const (
	DeviceInputTeamsDevicesCrowdstrikeInputRequestStateOnline  DeviceInputTeamsDevicesCrowdstrikeInputRequestState = "online"
	DeviceInputTeamsDevicesCrowdstrikeInputRequestStateOffline DeviceInputTeamsDevicesCrowdstrikeInputRequestState = "offline"
	DeviceInputTeamsDevicesCrowdstrikeInputRequestStateUnknown DeviceInputTeamsDevicesCrowdstrikeInputRequestState = "unknown"
)

func (r DeviceInputTeamsDevicesCrowdstrikeInputRequestState) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesCrowdstrikeInputRequestStateOnline, DeviceInputTeamsDevicesCrowdstrikeInputRequestStateOffline, DeviceInputTeamsDevicesCrowdstrikeInputRequestStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator string

const (
	DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess            DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<"
	DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals    DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "<="
	DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater         DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">"
	DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = ">="
	DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals          DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator = "=="
)

func (r DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLess, DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorLessOrEquals, DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreater, DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorGreaterOrEquals, DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperatorEquals:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesIntuneInputRequest struct {
	// Compliance Status
	ComplianceStatus DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID string                                        `json:"connection_id,required"`
	JSON         deviceInputTeamsDevicesIntuneInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesIntuneInputRequestJSON contains the JSON metadata for the
// struct [DeviceInputTeamsDevicesIntuneInputRequest]
type deviceInputTeamsDevicesIntuneInputRequestJSON struct {
	ComplianceStatus apijson.Field
	ConnectionID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesIntuneInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesIntuneInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesIntuneInputRequest) implementsZeroTrustDeviceInput() {}

// Compliance Status
type DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus string

const (
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant     DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "compliant"
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant  DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "noncompliant"
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown       DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "unknown"
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "notapplicable"
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "ingraceperiod"
	DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusError         DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus = "error"
)

func (r DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusCompliant, DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusNoncompliant, DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusUnknown, DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusNotapplicable, DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusIngraceperiod, DeviceInputTeamsDevicesIntuneInputRequestComplianceStatusError:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesKolideInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// Count Operator
	CountOperator DeviceInputTeamsDevicesKolideInputRequestCountOperator `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount string                                        `json:"issue_count,required"`
	JSON       deviceInputTeamsDevicesKolideInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesKolideInputRequestJSON contains the JSON metadata for the
// struct [DeviceInputTeamsDevicesKolideInputRequest]
type deviceInputTeamsDevicesKolideInputRequestJSON struct {
	ConnectionID  apijson.Field
	CountOperator apijson.Field
	IssueCount    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesKolideInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesKolideInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesKolideInputRequest) implementsZeroTrustDeviceInput() {}

// Count Operator
type DeviceInputTeamsDevicesKolideInputRequestCountOperator string

const (
	DeviceInputTeamsDevicesKolideInputRequestCountOperatorLess            DeviceInputTeamsDevicesKolideInputRequestCountOperator = "<"
	DeviceInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals    DeviceInputTeamsDevicesKolideInputRequestCountOperator = "<="
	DeviceInputTeamsDevicesKolideInputRequestCountOperatorGreater         DeviceInputTeamsDevicesKolideInputRequestCountOperator = ">"
	DeviceInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals DeviceInputTeamsDevicesKolideInputRequestCountOperator = ">="
	DeviceInputTeamsDevicesKolideInputRequestCountOperatorEquals          DeviceInputTeamsDevicesKolideInputRequestCountOperator = "=="
)

func (r DeviceInputTeamsDevicesKolideInputRequestCountOperator) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesKolideInputRequestCountOperatorLess, DeviceInputTeamsDevicesKolideInputRequestCountOperatorLessOrEquals, DeviceInputTeamsDevicesKolideInputRequestCountOperatorGreater, DeviceInputTeamsDevicesKolideInputRequestCountOperatorGreaterOrEquals, DeviceInputTeamsDevicesKolideInputRequestCountOperatorEquals:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesTaniumInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen string `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator DeviceInputTeamsDevicesTaniumInputRequestOperator `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel DeviceInputTeamsDevicesTaniumInputRequestRiskLevel `json:"risk_level"`
	// Score Operator
	ScoreOperator DeviceInputTeamsDevicesTaniumInputRequestScoreOperator `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore float64                                       `json:"total_score"`
	JSON       deviceInputTeamsDevicesTaniumInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesTaniumInputRequestJSON contains the JSON metadata for the
// struct [DeviceInputTeamsDevicesTaniumInputRequest]
type deviceInputTeamsDevicesTaniumInputRequestJSON struct {
	ConnectionID  apijson.Field
	EidLastSeen   apijson.Field
	Operator      apijson.Field
	RiskLevel     apijson.Field
	ScoreOperator apijson.Field
	TotalScore    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesTaniumInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesTaniumInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesTaniumInputRequest) implementsZeroTrustDeviceInput() {}

// Operator to evaluate risk_level or eid_last_seen.
type DeviceInputTeamsDevicesTaniumInputRequestOperator string

const (
	DeviceInputTeamsDevicesTaniumInputRequestOperatorLess            DeviceInputTeamsDevicesTaniumInputRequestOperator = "<"
	DeviceInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals    DeviceInputTeamsDevicesTaniumInputRequestOperator = "<="
	DeviceInputTeamsDevicesTaniumInputRequestOperatorGreater         DeviceInputTeamsDevicesTaniumInputRequestOperator = ">"
	DeviceInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals DeviceInputTeamsDevicesTaniumInputRequestOperator = ">="
	DeviceInputTeamsDevicesTaniumInputRequestOperatorEquals          DeviceInputTeamsDevicesTaniumInputRequestOperator = "=="
)

func (r DeviceInputTeamsDevicesTaniumInputRequestOperator) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesTaniumInputRequestOperatorLess, DeviceInputTeamsDevicesTaniumInputRequestOperatorLessOrEquals, DeviceInputTeamsDevicesTaniumInputRequestOperatorGreater, DeviceInputTeamsDevicesTaniumInputRequestOperatorGreaterOrEquals, DeviceInputTeamsDevicesTaniumInputRequestOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DeviceInputTeamsDevicesTaniumInputRequestRiskLevel string

const (
	DeviceInputTeamsDevicesTaniumInputRequestRiskLevelLow      DeviceInputTeamsDevicesTaniumInputRequestRiskLevel = "low"
	DeviceInputTeamsDevicesTaniumInputRequestRiskLevelMedium   DeviceInputTeamsDevicesTaniumInputRequestRiskLevel = "medium"
	DeviceInputTeamsDevicesTaniumInputRequestRiskLevelHigh     DeviceInputTeamsDevicesTaniumInputRequestRiskLevel = "high"
	DeviceInputTeamsDevicesTaniumInputRequestRiskLevelCritical DeviceInputTeamsDevicesTaniumInputRequestRiskLevel = "critical"
)

func (r DeviceInputTeamsDevicesTaniumInputRequestRiskLevel) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesTaniumInputRequestRiskLevelLow, DeviceInputTeamsDevicesTaniumInputRequestRiskLevelMedium, DeviceInputTeamsDevicesTaniumInputRequestRiskLevelHigh, DeviceInputTeamsDevicesTaniumInputRequestRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DeviceInputTeamsDevicesTaniumInputRequestScoreOperator string

const (
	DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorLess            DeviceInputTeamsDevicesTaniumInputRequestScoreOperator = "<"
	DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals    DeviceInputTeamsDevicesTaniumInputRequestScoreOperator = "<="
	DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorGreater         DeviceInputTeamsDevicesTaniumInputRequestScoreOperator = ">"
	DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals DeviceInputTeamsDevicesTaniumInputRequestScoreOperator = ">="
	DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorEquals          DeviceInputTeamsDevicesTaniumInputRequestScoreOperator = "=="
)

func (r DeviceInputTeamsDevicesTaniumInputRequestScoreOperator) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorLess, DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorLessOrEquals, DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorGreater, DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorGreaterOrEquals, DeviceInputTeamsDevicesTaniumInputRequestScoreOperatorEquals:
		return true
	}
	return false
}

type DeviceInputTeamsDevicesSentineloneS2sInputRequest struct {
	// Posture Integration ID.
	ConnectionID string `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats float64 `json:"active_threats"`
	// Whether device is infected.
	Infected bool `json:"infected"`
	// Whether device is active.
	IsActive bool `json:"is_active"`
	// Network status of device.
	NetworkStatus DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus `json:"network_status"`
	// operator
	Operator UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930      `json:"operator"`
	JSON     deviceInputTeamsDevicesSentineloneS2sInputRequestJSON `json:"-"`
}

// deviceInputTeamsDevicesSentineloneS2sInputRequestJSON contains the JSON metadata
// for the struct [DeviceInputTeamsDevicesSentineloneS2sInputRequest]
type deviceInputTeamsDevicesSentineloneS2sInputRequestJSON struct {
	ConnectionID  apijson.Field
	ActiveThreats apijson.Field
	Infected      apijson.Field
	IsActive      apijson.Field
	NetworkStatus apijson.Field
	Operator      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceInputTeamsDevicesSentineloneS2sInputRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceInputTeamsDevicesSentineloneS2sInputRequestJSON) RawJSON() string {
	return r.raw
}

func (r DeviceInputTeamsDevicesSentineloneS2sInputRequest) implementsZeroTrustDeviceInput() {}

// Network status of device.
type DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus string

const (
	DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected     DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connected"
	DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected  DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnected"
	DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "disconnecting"
	DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting    DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus = "connecting"
)

func (r DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus) IsKnown() bool {
	switch r {
	case DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnected, DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnected, DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusDisconnecting, DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatusConnecting:
		return true
	}
	return false
}

// Compliance Status
type DeviceInputComplianceStatus string

const (
	DeviceInputComplianceStatusCompliant     DeviceInputComplianceStatus = "compliant"
	DeviceInputComplianceStatusNoncompliant  DeviceInputComplianceStatus = "noncompliant"
	DeviceInputComplianceStatusUnknown       DeviceInputComplianceStatus = "unknown"
	DeviceInputComplianceStatusNotapplicable DeviceInputComplianceStatus = "notapplicable"
	DeviceInputComplianceStatusIngraceperiod DeviceInputComplianceStatus = "ingraceperiod"
	DeviceInputComplianceStatusError         DeviceInputComplianceStatus = "error"
)

func (r DeviceInputComplianceStatus) IsKnown() bool {
	switch r {
	case DeviceInputComplianceStatusCompliant, DeviceInputComplianceStatusNoncompliant, DeviceInputComplianceStatusUnknown, DeviceInputComplianceStatusNotapplicable, DeviceInputComplianceStatusIngraceperiod, DeviceInputComplianceStatusError:
		return true
	}
	return false
}

// For more details on state, please refer to the Crowdstrike documentation.
type DeviceInputState string

const (
	DeviceInputStateOnline  DeviceInputState = "online"
	DeviceInputStateOffline DeviceInputState = "offline"
	DeviceInputStateUnknown DeviceInputState = "unknown"
)

func (r DeviceInputState) IsKnown() bool {
	switch r {
	case DeviceInputStateOnline, DeviceInputStateOffline, DeviceInputStateUnknown:
		return true
	}
	return false
}

// Version Operator
type DeviceInputVersionOperator string

const (
	DeviceInputVersionOperatorLess            DeviceInputVersionOperator = "<"
	DeviceInputVersionOperatorLessOrEquals    DeviceInputVersionOperator = "<="
	DeviceInputVersionOperatorGreater         DeviceInputVersionOperator = ">"
	DeviceInputVersionOperatorGreaterOrEquals DeviceInputVersionOperator = ">="
	DeviceInputVersionOperatorEquals          DeviceInputVersionOperator = "=="
)

func (r DeviceInputVersionOperator) IsKnown() bool {
	switch r {
	case DeviceInputVersionOperatorLess, DeviceInputVersionOperatorLessOrEquals, DeviceInputVersionOperatorGreater, DeviceInputVersionOperatorGreaterOrEquals, DeviceInputVersionOperatorEquals:
		return true
	}
	return false
}

// Count Operator
type DeviceInputCountOperator string

const (
	DeviceInputCountOperatorLess            DeviceInputCountOperator = "<"
	DeviceInputCountOperatorLessOrEquals    DeviceInputCountOperator = "<="
	DeviceInputCountOperatorGreater         DeviceInputCountOperator = ">"
	DeviceInputCountOperatorGreaterOrEquals DeviceInputCountOperator = ">="
	DeviceInputCountOperatorEquals          DeviceInputCountOperator = "=="
)

func (r DeviceInputCountOperator) IsKnown() bool {
	switch r {
	case DeviceInputCountOperatorLess, DeviceInputCountOperatorLessOrEquals, DeviceInputCountOperatorGreater, DeviceInputCountOperatorGreaterOrEquals, DeviceInputCountOperatorEquals:
		return true
	}
	return false
}

// For more details on risk level, refer to the Tanium documentation.
type DeviceInputRiskLevel string

const (
	DeviceInputRiskLevelLow      DeviceInputRiskLevel = "low"
	DeviceInputRiskLevelMedium   DeviceInputRiskLevel = "medium"
	DeviceInputRiskLevelHigh     DeviceInputRiskLevel = "high"
	DeviceInputRiskLevelCritical DeviceInputRiskLevel = "critical"
)

func (r DeviceInputRiskLevel) IsKnown() bool {
	switch r {
	case DeviceInputRiskLevelLow, DeviceInputRiskLevelMedium, DeviceInputRiskLevelHigh, DeviceInputRiskLevelCritical:
		return true
	}
	return false
}

// Score Operator
type DeviceInputScoreOperator string

const (
	DeviceInputScoreOperatorLess            DeviceInputScoreOperator = "<"
	DeviceInputScoreOperatorLessOrEquals    DeviceInputScoreOperator = "<="
	DeviceInputScoreOperatorGreater         DeviceInputScoreOperator = ">"
	DeviceInputScoreOperatorGreaterOrEquals DeviceInputScoreOperator = ">="
	DeviceInputScoreOperatorEquals          DeviceInputScoreOperator = "=="
)

func (r DeviceInputScoreOperator) IsKnown() bool {
	switch r {
	case DeviceInputScoreOperatorLess, DeviceInputScoreOperatorLessOrEquals, DeviceInputScoreOperatorGreater, DeviceInputScoreOperatorGreaterOrEquals, DeviceInputScoreOperatorEquals:
		return true
	}
	return false
}

// Network status of device.
type DeviceInputNetworkStatus string

const (
	DeviceInputNetworkStatusConnected     DeviceInputNetworkStatus = "connected"
	DeviceInputNetworkStatusDisconnected  DeviceInputNetworkStatus = "disconnected"
	DeviceInputNetworkStatusDisconnecting DeviceInputNetworkStatus = "disconnecting"
	DeviceInputNetworkStatusConnecting    DeviceInputNetworkStatus = "connecting"
)

func (r DeviceInputNetworkStatus) IsKnown() bool {
	switch r {
	case DeviceInputNetworkStatusConnected, DeviceInputNetworkStatusDisconnected, DeviceInputNetworkStatusDisconnecting, DeviceInputNetworkStatusConnecting:
		return true
	}
	return false
}

// The value to be checked against.
type DeviceInputParam struct {
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system"`
	// File path.
	Path param.Field[string] `json:"path"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
	// List ID.
	ID param.Field[string] `json:"id"`
	// Domain
	Domain param.Field[string] `json:"domain"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
	// Version of OS
	Version param.Field[string] `json:"version"`
	// Enabled
	Enabled    param.Field[bool]        `json:"enabled"`
	CheckDisks param.Field[interface{}] `json:"checkDisks,required"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn"`
	// Compliance Status
	ComplianceStatus param.Field[DeviceInputComplianceStatus] `json:"compliance_status"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DeviceInputState] `json:"state"`
	// Version Operator
	VersionOperator param.Field[DeviceInputVersionOperator] `json:"versionOperator"`
	// Count Operator
	CountOperator param.Field[DeviceInputCountOperator] `json:"countOperator"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DeviceInputRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DeviceInputScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DeviceInputNetworkStatus] `json:"network_status"`
}

func (r DeviceInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputParam) implementsZeroTrustDeviceInputUnionParam() {}

// The value to be checked against.
//
// Satisfied by [zero_trust.DeviceInputTeamsDevicesFileInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesUniqueClientIDInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesDomainJoinedInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesOSVersionInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesFirewallInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesSentineloneInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesCarbonblackInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesDiskEncryptionInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesApplicationInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesClientCertificateInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesWorkspaceOneInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesCrowdstrikeInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesIntuneInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesKolideInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesTaniumInputRequestParam],
// [zero_trust.DeviceInputTeamsDevicesSentineloneS2sInputRequestParam],
// [DeviceInputParam].
type DeviceInputUnionParam interface {
	implementsZeroTrustDeviceInputUnionParam()
}

type DeviceInputTeamsDevicesFileInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// Whether or not file exists
	Exists param.Field[bool] `json:"exists"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesFileInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesFileInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesUniqueClientIDInputRequestParam struct {
	// List ID.
	ID param.Field[string] `json:"id,required"`
	// Operating System
	OperatingSystem param.Field[DeviceInputTeamsDevicesUniqueClientIDInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DeviceInputTeamsDevicesUniqueClientIDInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesUniqueClientIDInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesDomainJoinedInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[DeviceInputTeamsDevicesDomainJoinedInputRequestOperatingSystem] `json:"operating_system,required"`
	// Domain
	Domain param.Field[string] `json:"domain"`
}

func (r DeviceInputTeamsDevicesDomainJoinedInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesDomainJoinedInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesOSVersionInputRequestParam struct {
	// Operating System
	OperatingSystem param.Field[DeviceInputTeamsDevicesOSVersionInputRequestOperatingSystem] `json:"operating_system,required"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator,required"`
	// Version of OS
	Version param.Field[string] `json:"version,required"`
	// Operating System Distribution Name (linux only)
	OSDistroName param.Field[string] `json:"os_distro_name"`
	// Version of OS Distribution (linux only)
	OSDistroRevision param.Field[string] `json:"os_distro_revision"`
	// Additional version data. For Mac or iOS, the Product Verison Extra. For Linux,
	// the kernel release version. (Mac, iOS, and Linux only)
	OSVersionExtra param.Field[string] `json:"os_version_extra"`
}

func (r DeviceInputTeamsDevicesOSVersionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesOSVersionInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesFirewallInputRequestParam struct {
	// Enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// Operating System
	OperatingSystem param.Field[DeviceInputTeamsDevicesFirewallInputRequestOperatingSystem] `json:"operating_system,required"`
}

func (r DeviceInputTeamsDevicesFirewallInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesFirewallInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesSentineloneInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesSentineloneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesSentineloneInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesCarbonblackInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// File path.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesCarbonblackInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesDiskEncryptionInputRequestParam struct {
	// List of volume names to be checked for encryption.
	CheckDisks param.Field[[]string] `json:"checkDisks"`
	// Whether to check all disks for encryption.
	RequireAll param.Field[bool] `json:"requireAll"`
}

func (r DeviceInputTeamsDevicesDiskEncryptionInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesDiskEncryptionInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesApplicationInputRequestParam struct {
	// Operating system
	OperatingSystem param.Field[UnnamedSchemaRef41885dd46b9e0294254c49305a273681] `json:"operating_system,required"`
	// Path for the application.
	Path param.Field[string] `json:"path,required"`
	// SHA-256.
	Sha256 param.Field[string] `json:"sha256"`
	// Signing certificate thumbprint.
	Thumbprint param.Field[string] `json:"thumbprint"`
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesApplicationInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesClientCertificateInputRequestParam struct {
	// UUID of Cloudflare managed certificate.
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// Common Name that is protected by the certificate
	Cn param.Field[string] `json:"cn,required"`
}

func (r DeviceInputTeamsDevicesClientCertificateInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesClientCertificateInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesWorkspaceOneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[DeviceInputTeamsDevicesWorkspaceOneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DeviceInputTeamsDevicesWorkspaceOneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesWorkspaceOneInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesCrowdstrikeInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on last seen, please refer to the Crowdstrike documentation.
	LastSeen param.Field[string] `json:"last_seen"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
	// Os Version
	OS param.Field[string] `json:"os"`
	// overall
	Overall param.Field[string] `json:"overall"`
	// SensorConfig
	SensorConfig param.Field[string] `json:"sensor_config"`
	// For more details on state, please refer to the Crowdstrike documentation.
	State param.Field[DeviceInputTeamsDevicesCrowdstrikeInputRequestState] `json:"state"`
	// Version
	Version param.Field[string] `json:"version"`
	// Version Operator
	VersionOperator param.Field[DeviceInputTeamsDevicesCrowdstrikeInputRequestVersionOperator] `json:"versionOperator"`
}

func (r DeviceInputTeamsDevicesCrowdstrikeInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesCrowdstrikeInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceInputTeamsDevicesIntuneInputRequestParam struct {
	// Compliance Status
	ComplianceStatus param.Field[DeviceInputTeamsDevicesIntuneInputRequestComplianceStatus] `json:"compliance_status,required"`
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r DeviceInputTeamsDevicesIntuneInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesIntuneInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesKolideInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Count Operator
	CountOperator param.Field[DeviceInputTeamsDevicesKolideInputRequestCountOperator] `json:"countOperator,required"`
	// The Number of Issues.
	IssueCount param.Field[string] `json:"issue_count,required"`
}

func (r DeviceInputTeamsDevicesKolideInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesKolideInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesTaniumInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// For more details on eid last seen, refer to the Tanium documentation.
	EidLastSeen param.Field[string] `json:"eid_last_seen"`
	// Operator to evaluate risk_level or eid_last_seen.
	Operator param.Field[DeviceInputTeamsDevicesTaniumInputRequestOperator] `json:"operator"`
	// For more details on risk level, refer to the Tanium documentation.
	RiskLevel param.Field[DeviceInputTeamsDevicesTaniumInputRequestRiskLevel] `json:"risk_level"`
	// Score Operator
	ScoreOperator param.Field[DeviceInputTeamsDevicesTaniumInputRequestScoreOperator] `json:"scoreOperator"`
	// For more details on total score, refer to the Tanium documentation.
	TotalScore param.Field[float64] `json:"total_score"`
}

func (r DeviceInputTeamsDevicesTaniumInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesTaniumInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {}

type DeviceInputTeamsDevicesSentineloneS2sInputRequestParam struct {
	// Posture Integration ID.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The Number of active threats.
	ActiveThreats param.Field[float64] `json:"active_threats"`
	// Whether device is infected.
	Infected param.Field[bool] `json:"infected"`
	// Whether device is active.
	IsActive param.Field[bool] `json:"is_active"`
	// Network status of device.
	NetworkStatus param.Field[DeviceInputTeamsDevicesSentineloneS2sInputRequestNetworkStatus] `json:"network_status"`
	// operator
	Operator param.Field[UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930] `json:"operator"`
}

func (r DeviceInputTeamsDevicesSentineloneS2sInputRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DeviceInputTeamsDevicesSentineloneS2sInputRequestParam) implementsZeroTrustDeviceInputUnionParam() {
}

type DeviceMatch struct {
	Platform DeviceMatchPlatform `json:"platform"`
	JSON     deviceMatchJSON     `json:"-"`
}

// deviceMatchJSON contains the JSON metadata for the struct [DeviceMatch]
type deviceMatchJSON struct {
	Platform    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceMatchJSON) RawJSON() string {
	return r.raw
}

type DeviceMatchPlatform string

const (
	DeviceMatchPlatformWindows DeviceMatchPlatform = "windows"
	DeviceMatchPlatformMac     DeviceMatchPlatform = "mac"
	DeviceMatchPlatformLinux   DeviceMatchPlatform = "linux"
	DeviceMatchPlatformAndroid DeviceMatchPlatform = "android"
	DeviceMatchPlatformIos     DeviceMatchPlatform = "ios"
)

func (r DeviceMatchPlatform) IsKnown() bool {
	switch r {
	case DeviceMatchPlatformWindows, DeviceMatchPlatformMac, DeviceMatchPlatformLinux, DeviceMatchPlatformAndroid, DeviceMatchPlatformIos:
		return true
	}
	return false
}

type DeviceMatchParam struct {
	Platform param.Field[DeviceMatchPlatform] `json:"platform"`
}

func (r DeviceMatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePostureRule struct {
	// API UUID.
	ID string `json:"id"`
	// The description of the device posture rule.
	Description string `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration string `json:"expiration"`
	// The value to be checked against.
	Input DeviceInput `json:"input"`
	// The conditions that the client must match to run the rule.
	Match []DeviceMatch `json:"match"`
	// The name of the device posture rule.
	Name string `json:"name"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule string `json:"schedule"`
	// The type of device posture rule.
	Type DevicePostureRuleType `json:"type"`
	JSON devicePostureRuleJSON `json:"-"`
}

// devicePostureRuleJSON contains the JSON metadata for the struct
// [DevicePostureRule]
type devicePostureRuleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expiration  apijson.Field
	Input       apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Schedule    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureRuleJSON) RawJSON() string {
	return r.raw
}

// The type of device posture rule.
type DevicePostureRuleType string

const (
	DevicePostureRuleTypeFile              DevicePostureRuleType = "file"
	DevicePostureRuleTypeApplication       DevicePostureRuleType = "application"
	DevicePostureRuleTypeTanium            DevicePostureRuleType = "tanium"
	DevicePostureRuleTypeGateway           DevicePostureRuleType = "gateway"
	DevicePostureRuleTypeWARP              DevicePostureRuleType = "warp"
	DevicePostureRuleTypeDiskEncryption    DevicePostureRuleType = "disk_encryption"
	DevicePostureRuleTypeSentinelone       DevicePostureRuleType = "sentinelone"
	DevicePostureRuleTypeCarbonblack       DevicePostureRuleType = "carbonblack"
	DevicePostureRuleTypeFirewall          DevicePostureRuleType = "firewall"
	DevicePostureRuleTypeOSVersion         DevicePostureRuleType = "os_version"
	DevicePostureRuleTypeDomainJoined      DevicePostureRuleType = "domain_joined"
	DevicePostureRuleTypeClientCertificate DevicePostureRuleType = "client_certificate"
	DevicePostureRuleTypeUniqueClientID    DevicePostureRuleType = "unique_client_id"
	DevicePostureRuleTypeKolide            DevicePostureRuleType = "kolide"
	DevicePostureRuleTypeTaniumS2s         DevicePostureRuleType = "tanium_s2s"
	DevicePostureRuleTypeCrowdstrikeS2s    DevicePostureRuleType = "crowdstrike_s2s"
	DevicePostureRuleTypeIntune            DevicePostureRuleType = "intune"
	DevicePostureRuleTypeWorkspaceOne      DevicePostureRuleType = "workspace_one"
	DevicePostureRuleTypeSentineloneS2s    DevicePostureRuleType = "sentinelone_s2s"
)

func (r DevicePostureRuleType) IsKnown() bool {
	switch r {
	case DevicePostureRuleTypeFile, DevicePostureRuleTypeApplication, DevicePostureRuleTypeTanium, DevicePostureRuleTypeGateway, DevicePostureRuleTypeWARP, DevicePostureRuleTypeDiskEncryption, DevicePostureRuleTypeSentinelone, DevicePostureRuleTypeCarbonblack, DevicePostureRuleTypeFirewall, DevicePostureRuleTypeOSVersion, DevicePostureRuleTypeDomainJoined, DevicePostureRuleTypeClientCertificate, DevicePostureRuleTypeUniqueClientID, DevicePostureRuleTypeKolide, DevicePostureRuleTypeTaniumS2s, DevicePostureRuleTypeCrowdstrikeS2s, DevicePostureRuleTypeIntune, DevicePostureRuleTypeWorkspaceOne, DevicePostureRuleTypeSentineloneS2s:
		return true
	}
	return false
}

// operator
type UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 string

const (
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less            UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals    UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "<="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater         UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">"
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = ">="
	UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals          UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930 = "=="
)

func (r UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Less, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930LessOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Greater, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930GreaterOrEquals, UnnamedSchemaRef34ef0ad73a63c3f76ed170adca181930Equals:
		return true
	}
	return false
}

// Operating system
type UnnamedSchemaRef41885dd46b9e0294254c49305a273681 string

const (
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "windows"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux   UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "linux"
	UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac     UnnamedSchemaRef41885dd46b9e0294254c49305a273681 = "mac"
)

func (r UnnamedSchemaRef41885dd46b9e0294254c49305a273681) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef41885dd46b9e0294254c49305a273681Windows, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Linux, UnnamedSchemaRef41885dd46b9e0294254c49305a273681Mac:
		return true
	}
	return false
}

type DevicePostureDeleteResponse struct {
	// API UUID.
	ID   string                          `json:"id"`
	JSON devicePostureDeleteResponseJSON `json:"-"`
}

// devicePostureDeleteResponseJSON contains the JSON metadata for the struct
// [DevicePostureDeleteResponse]
type devicePostureDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DevicePostureNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureNewParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureNewParamsType string

const (
	DevicePostureNewParamsTypeFile              DevicePostureNewParamsType = "file"
	DevicePostureNewParamsTypeApplication       DevicePostureNewParamsType = "application"
	DevicePostureNewParamsTypeTanium            DevicePostureNewParamsType = "tanium"
	DevicePostureNewParamsTypeGateway           DevicePostureNewParamsType = "gateway"
	DevicePostureNewParamsTypeWARP              DevicePostureNewParamsType = "warp"
	DevicePostureNewParamsTypeDiskEncryption    DevicePostureNewParamsType = "disk_encryption"
	DevicePostureNewParamsTypeSentinelone       DevicePostureNewParamsType = "sentinelone"
	DevicePostureNewParamsTypeCarbonblack       DevicePostureNewParamsType = "carbonblack"
	DevicePostureNewParamsTypeFirewall          DevicePostureNewParamsType = "firewall"
	DevicePostureNewParamsTypeOSVersion         DevicePostureNewParamsType = "os_version"
	DevicePostureNewParamsTypeDomainJoined      DevicePostureNewParamsType = "domain_joined"
	DevicePostureNewParamsTypeClientCertificate DevicePostureNewParamsType = "client_certificate"
	DevicePostureNewParamsTypeUniqueClientID    DevicePostureNewParamsType = "unique_client_id"
	DevicePostureNewParamsTypeKolide            DevicePostureNewParamsType = "kolide"
	DevicePostureNewParamsTypeTaniumS2s         DevicePostureNewParamsType = "tanium_s2s"
	DevicePostureNewParamsTypeCrowdstrikeS2s    DevicePostureNewParamsType = "crowdstrike_s2s"
	DevicePostureNewParamsTypeIntune            DevicePostureNewParamsType = "intune"
	DevicePostureNewParamsTypeWorkspaceOne      DevicePostureNewParamsType = "workspace_one"
	DevicePostureNewParamsTypeSentineloneS2s    DevicePostureNewParamsType = "sentinelone_s2s"
)

func (r DevicePostureNewParamsType) IsKnown() bool {
	switch r {
	case DevicePostureNewParamsTypeFile, DevicePostureNewParamsTypeApplication, DevicePostureNewParamsTypeTanium, DevicePostureNewParamsTypeGateway, DevicePostureNewParamsTypeWARP, DevicePostureNewParamsTypeDiskEncryption, DevicePostureNewParamsTypeSentinelone, DevicePostureNewParamsTypeCarbonblack, DevicePostureNewParamsTypeFirewall, DevicePostureNewParamsTypeOSVersion, DevicePostureNewParamsTypeDomainJoined, DevicePostureNewParamsTypeClientCertificate, DevicePostureNewParamsTypeUniqueClientID, DevicePostureNewParamsTypeKolide, DevicePostureNewParamsTypeTaniumS2s, DevicePostureNewParamsTypeCrowdstrikeS2s, DevicePostureNewParamsTypeIntune, DevicePostureNewParamsTypeWorkspaceOne, DevicePostureNewParamsTypeSentineloneS2s:
		return true
	}
	return false
}

type DevicePostureNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureNewResponseEnvelopeJSON    `json:"-"`
}

// devicePostureNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureNewResponseEnvelope]
type devicePostureNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureNewResponseEnvelopeSuccess bool

const (
	DevicePostureNewResponseEnvelopeSuccessTrue DevicePostureNewResponseEnvelopeSuccess = true
)

func (r DevicePostureNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the device posture rule.
	Name param.Field[string] `json:"name,required"`
	// The type of device posture rule.
	Type param.Field[DevicePostureUpdateParamsType] `json:"type,required"`
	// The description of the device posture rule.
	Description param.Field[string] `json:"description"`
	// Sets the expiration time for a posture check result. If empty, the result
	// remains valid until it is overwritten by new data from the WARP client.
	Expiration param.Field[string] `json:"expiration"`
	// The value to be checked against.
	Input param.Field[DeviceInputUnionParam] `json:"input"`
	// The conditions that the client must match to run the rule.
	Match param.Field[[]DeviceMatchParam] `json:"match"`
	// Polling frequency for the WARP client posture check. Default: `5m` (poll every
	// five minutes). Minimum: `1m`.
	Schedule param.Field[string] `json:"schedule"`
}

func (r DevicePostureUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device posture rule.
type DevicePostureUpdateParamsType string

const (
	DevicePostureUpdateParamsTypeFile              DevicePostureUpdateParamsType = "file"
	DevicePostureUpdateParamsTypeApplication       DevicePostureUpdateParamsType = "application"
	DevicePostureUpdateParamsTypeTanium            DevicePostureUpdateParamsType = "tanium"
	DevicePostureUpdateParamsTypeGateway           DevicePostureUpdateParamsType = "gateway"
	DevicePostureUpdateParamsTypeWARP              DevicePostureUpdateParamsType = "warp"
	DevicePostureUpdateParamsTypeDiskEncryption    DevicePostureUpdateParamsType = "disk_encryption"
	DevicePostureUpdateParamsTypeSentinelone       DevicePostureUpdateParamsType = "sentinelone"
	DevicePostureUpdateParamsTypeCarbonblack       DevicePostureUpdateParamsType = "carbonblack"
	DevicePostureUpdateParamsTypeFirewall          DevicePostureUpdateParamsType = "firewall"
	DevicePostureUpdateParamsTypeOSVersion         DevicePostureUpdateParamsType = "os_version"
	DevicePostureUpdateParamsTypeDomainJoined      DevicePostureUpdateParamsType = "domain_joined"
	DevicePostureUpdateParamsTypeClientCertificate DevicePostureUpdateParamsType = "client_certificate"
	DevicePostureUpdateParamsTypeUniqueClientID    DevicePostureUpdateParamsType = "unique_client_id"
	DevicePostureUpdateParamsTypeKolide            DevicePostureUpdateParamsType = "kolide"
	DevicePostureUpdateParamsTypeTaniumS2s         DevicePostureUpdateParamsType = "tanium_s2s"
	DevicePostureUpdateParamsTypeCrowdstrikeS2s    DevicePostureUpdateParamsType = "crowdstrike_s2s"
	DevicePostureUpdateParamsTypeIntune            DevicePostureUpdateParamsType = "intune"
	DevicePostureUpdateParamsTypeWorkspaceOne      DevicePostureUpdateParamsType = "workspace_one"
	DevicePostureUpdateParamsTypeSentineloneS2s    DevicePostureUpdateParamsType = "sentinelone_s2s"
)

func (r DevicePostureUpdateParamsType) IsKnown() bool {
	switch r {
	case DevicePostureUpdateParamsTypeFile, DevicePostureUpdateParamsTypeApplication, DevicePostureUpdateParamsTypeTanium, DevicePostureUpdateParamsTypeGateway, DevicePostureUpdateParamsTypeWARP, DevicePostureUpdateParamsTypeDiskEncryption, DevicePostureUpdateParamsTypeSentinelone, DevicePostureUpdateParamsTypeCarbonblack, DevicePostureUpdateParamsTypeFirewall, DevicePostureUpdateParamsTypeOSVersion, DevicePostureUpdateParamsTypeDomainJoined, DevicePostureUpdateParamsTypeClientCertificate, DevicePostureUpdateParamsTypeUniqueClientID, DevicePostureUpdateParamsTypeKolide, DevicePostureUpdateParamsTypeTaniumS2s, DevicePostureUpdateParamsTypeCrowdstrikeS2s, DevicePostureUpdateParamsTypeIntune, DevicePostureUpdateParamsTypeWorkspaceOne, DevicePostureUpdateParamsTypeSentineloneS2s:
		return true
	}
	return false
}

type DevicePostureUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePostureUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureUpdateResponseEnvelope]
type devicePostureUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureUpdateResponseEnvelopeSuccess bool

const (
	DevicePostureUpdateResponseEnvelopeSuccessTrue DevicePostureUpdateResponseEnvelopeSuccess = true
)

func (r DevicePostureUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureDeleteParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r DevicePostureDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePostureDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DevicePostureDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureDeleteResponseEnvelopeJSON    `json:"-"`
}

// devicePostureDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePostureDeleteResponseEnvelope]
type devicePostureDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureDeleteResponseEnvelopeSuccess bool

const (
	DevicePostureDeleteResponseEnvelopeSuccessTrue DevicePostureDeleteResponseEnvelopeSuccess = true
)

func (r DevicePostureDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePostureGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePostureGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DevicePostureRule     `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePostureGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePostureGetResponseEnvelopeJSON    `json:"-"`
}

// devicePostureGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DevicePostureGetResponseEnvelope]
type devicePostureGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePostureGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePostureGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePostureGetResponseEnvelopeSuccess bool

const (
	DevicePostureGetResponseEnvelopeSuccessTrue DevicePostureGetResponseEnvelopeSuccess = true
)

func (r DevicePostureGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePostureGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
