// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessUserService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessUserService] method instead.
type AccessUserService struct {
	Options      []option.RequestOption
	FailedLogins *AccessUserFailedLoginService
}

// NewAccessUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessUserService(opts ...option.RequestOption) (r *AccessUserService) {
	r = &AccessUserService{}
	r.Options = opts
	r.FailedLogins = NewAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccessUserService) ZeroTrustUsersGetUsers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessUserZeroTrustUsersGetUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserZeroTrustUsersGetUsersResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserZeroTrustUsersGetUsersResponse struct {
	// UUID
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid interface{} `json:"seat_uid"`
	// The unique API identifier for the user.
	Uid       interface{}                                  `json:"uid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      accessUserZeroTrustUsersGetUsersResponseJSON `json:"-"`
}

// accessUserZeroTrustUsersGetUsersResponseJSON contains the JSON metadata for the
// struct [AccessUserZeroTrustUsersGetUsersResponse]
type accessUserZeroTrustUsersGetUsersResponseJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUid             apijson.Field
	Uid                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUserZeroTrustUsersGetUsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserZeroTrustUsersGetUsersResponseEnvelope struct {
	Errors     []AccessUserZeroTrustUsersGetUsersResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessUserZeroTrustUsersGetUsersResponseEnvelopeMessages `json:"messages"`
	Result     []AccessUserZeroTrustUsersGetUsersResponse                 `json:"result"`
	ResultInfo AccessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessUserZeroTrustUsersGetUsersResponseEnvelopeSuccess `json:"success"`
	JSON    accessUserZeroTrustUsersGetUsersResponseEnvelopeJSON    `json:"-"`
}

// accessUserZeroTrustUsersGetUsersResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessUserZeroTrustUsersGetUsersResponseEnvelope]
type accessUserZeroTrustUsersGetUsersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserZeroTrustUsersGetUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserZeroTrustUsersGetUsersResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessUserZeroTrustUsersGetUsersResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserZeroTrustUsersGetUsersResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserZeroTrustUsersGetUsersResponseEnvelopeErrors]
type accessUserZeroTrustUsersGetUsersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserZeroTrustUsersGetUsersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserZeroTrustUsersGetUsersResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accessUserZeroTrustUsersGetUsersResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserZeroTrustUsersGetUsersResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AccessUserZeroTrustUsersGetUsersResponseEnvelopeMessages]
type accessUserZeroTrustUsersGetUsersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserZeroTrustUsersGetUsersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfo struct {
	Count      interface{}                                                    `json:"count"`
	Page       interface{}                                                    `json:"page"`
	PerPage    interface{}                                                    `json:"per_page"`
	TotalCount interface{}                                                    `json:"total_count"`
	JSON       accessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AccessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfo]
type accessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserZeroTrustUsersGetUsersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserZeroTrustUsersGetUsersResponseEnvelopeSuccess bool

const (
	AccessUserZeroTrustUsersGetUsersResponseEnvelopeSuccessTrue AccessUserZeroTrustUsersGetUsersResponseEnvelopeSuccess = true
)
