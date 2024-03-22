// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDNSService] method instead.
type DNSService struct {
	Options []option.RequestOption
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	return
}

// Show the DNS records needed to configure your Email Routing zone.
func (r *DNSService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]EmailDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/dns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of records needed to enable an Email Routing zone.
type EmailDNSRecord struct {
	// DNS record content.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name"`
	// Required for MX, SRV and URI records. Unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
	// for 'automatic'.
	TTL EmailDNSRecordTTL `json:"ttl"`
	// DNS record type.
	Type EmailDNSRecordType `json:"type"`
	JSON emailDNSRecordJSON `json:"-"`
}

// emailDNSRecordJSON contains the JSON metadata for the struct [EmailDNSRecord]
type emailDNSRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailDNSRecordJSON) RawJSON() string {
	return r.raw
}

// Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1
// for 'automatic'.
//
// Union satisfied by [shared.UnionFloat] or
// [email_routing.EmailDNSRecordTTLNumber].
type EmailDNSRecordTTL interface {
	ImplementsEmailRoutingEmailDNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmailDNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(EmailDNSRecordTTLNumber(0)),
		},
	)
}

type EmailDNSRecordTTLNumber float64

const (
	EmailDNSRecordTTLNumber1 EmailDNSRecordTTLNumber = 1
)

func (r EmailDNSRecordTTLNumber) IsKnown() bool {
	switch r {
	case EmailDNSRecordTTLNumber1:
		return true
	}
	return false
}

// DNS record type.
type EmailDNSRecordType string

const (
	EmailDNSRecordTypeA      EmailDNSRecordType = "A"
	EmailDNSRecordTypeAAAA   EmailDNSRecordType = "AAAA"
	EmailDNSRecordTypeCNAME  EmailDNSRecordType = "CNAME"
	EmailDNSRecordTypeHTTPS  EmailDNSRecordType = "HTTPS"
	EmailDNSRecordTypeTXT    EmailDNSRecordType = "TXT"
	EmailDNSRecordTypeSRV    EmailDNSRecordType = "SRV"
	EmailDNSRecordTypeLOC    EmailDNSRecordType = "LOC"
	EmailDNSRecordTypeMX     EmailDNSRecordType = "MX"
	EmailDNSRecordTypeNS     EmailDNSRecordType = "NS"
	EmailDNSRecordTypeCERT   EmailDNSRecordType = "CERT"
	EmailDNSRecordTypeDNSKEY EmailDNSRecordType = "DNSKEY"
	EmailDNSRecordTypeDS     EmailDNSRecordType = "DS"
	EmailDNSRecordTypeNAPTR  EmailDNSRecordType = "NAPTR"
	EmailDNSRecordTypeSMIMEA EmailDNSRecordType = "SMIMEA"
	EmailDNSRecordTypeSSHFP  EmailDNSRecordType = "SSHFP"
	EmailDNSRecordTypeSVCB   EmailDNSRecordType = "SVCB"
	EmailDNSRecordTypeTLSA   EmailDNSRecordType = "TLSA"
	EmailDNSRecordTypeURI    EmailDNSRecordType = "URI"
)

func (r EmailDNSRecordType) IsKnown() bool {
	switch r {
	case EmailDNSRecordTypeA, EmailDNSRecordTypeAAAA, EmailDNSRecordTypeCNAME, EmailDNSRecordTypeHTTPS, EmailDNSRecordTypeTXT, EmailDNSRecordTypeSRV, EmailDNSRecordTypeLOC, EmailDNSRecordTypeMX, EmailDNSRecordTypeNS, EmailDNSRecordTypeCERT, EmailDNSRecordTypeDNSKEY, EmailDNSRecordTypeDS, EmailDNSRecordTypeNAPTR, EmailDNSRecordTypeSMIMEA, EmailDNSRecordTypeSSHFP, EmailDNSRecordTypeSVCB, EmailDNSRecordTypeTLSA, EmailDNSRecordTypeURI:
		return true
	}
	return false
}

type DNSGetResponseEnvelope struct {
	Errors   []DNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailDNSRecord                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DNSGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DNSGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dnsGetResponseEnvelopeJSON       `json:"-"`
}

// dnsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelope]
type dnsGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dnsGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeErrors]
type dnsGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    dnsGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeMessages]
type dnsGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSGetResponseEnvelopeSuccess bool

const (
	DNSGetResponseEnvelopeSuccessTrue DNSGetResponseEnvelopeSuccess = true
)

func (r DNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       dnsGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// dnsGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeResultInfo]
type dnsGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}