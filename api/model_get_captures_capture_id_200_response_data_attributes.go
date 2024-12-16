/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETCapturesCaptureId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETCapturesCaptureId200ResponseDataAttributes{}

// GETCapturesCaptureId200ResponseDataAttributes struct for GETCapturesCaptureId200ResponseDataAttributes
type GETCapturesCaptureId200ResponseDataAttributes struct {
	// The transaction number, auto generated.
	Number interface{} `json:"number,omitempty"`
	// The transaction's type.
	Type interface{} `json:"type,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated order.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The transaction amount, in cents.
	AmountCents interface{} `json:"amount_cents,omitempty"`
	// The transaction amount, float.
	AmountFloat interface{} `json:"amount_float,omitempty"`
	// The transaction amount, formatted.
	FormattedAmount interface{} `json:"formatted_amount,omitempty"`
	// Indicates if the transaction is successful.
	Succeeded interface{} `json:"succeeded,omitempty"`
	// The message returned by the payment gateway.
	Message interface{} `json:"message,omitempty"`
	// The error code, if any, returned by the payment gateway.
	ErrorCode interface{} `json:"error_code,omitempty"`
	// The error detail, if any, returned by the payment gateway.
	ErrorDetail interface{} `json:"error_detail,omitempty"`
	// The token identifying the transaction, returned by the payment gateway.
	Token interface{} `json:"token,omitempty"`
	// The ID identifying the transaction, returned by the payment gateway.
	GatewayTransactionId interface{} `json:"gateway_transaction_id,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The amount to be refunded, in cents.
	RefundAmountCents interface{} `json:"refund_amount_cents,omitempty"`
	// The amount to be refunded, float.
	RefundAmountFloat interface{} `json:"refund_amount_float,omitempty"`
	// The amount to be refunded, formatted.
	FormattedRefundAmount interface{} `json:"formatted_refund_amount,omitempty"`
	// The balance to be refunded, in cents.
	RefundBalanceCents interface{} `json:"refund_balance_cents,omitempty"`
	// The balance to be refunded, float.
	RefundBalanceFloat interface{} `json:"refund_balance_float,omitempty"`
	// The balance to be refunded, formatted.
	FormattedRefundBalance interface{} `json:"formatted_refund_balance,omitempty"`
}

// NewGETCapturesCaptureId200ResponseDataAttributes instantiates a new GETCapturesCaptureId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCapturesCaptureId200ResponseDataAttributes() *GETCapturesCaptureId200ResponseDataAttributes {
	this := GETCapturesCaptureId200ResponseDataAttributes{}
	return &this
}

// NewGETCapturesCaptureId200ResponseDataAttributesWithDefaults instantiates a new GETCapturesCaptureId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCapturesCaptureId200ResponseDataAttributesWithDefaults() *GETCapturesCaptureId200ResponseDataAttributes {
	this := GETCapturesCaptureId200ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return &o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasAmountCents() bool {
	if o != nil && IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given interface{} and assigns it to the AmountCents field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountFloat) {
		return nil, false
	}
	return &o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasAmountFloat() bool {
	if o != nil && IsNil(o.AmountFloat) {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given interface{} and assigns it to the AmountFloat field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetAmountFloat(v interface{}) {
	o.AmountFloat = v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedAmount) {
		return nil, false
	}
	return &o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasFormattedAmount() bool {
	if o != nil && IsNil(o.FormattedAmount) {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given interface{} and assigns it to the FormattedAmount field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetFormattedAmount(v interface{}) {
	o.FormattedAmount = v
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetSucceeded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Succeeded) {
		return nil, false
	}
	return &o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasSucceeded() bool {
	if o != nil && IsNil(o.Succeeded) {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given interface{} and assigns it to the Succeeded field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetSucceeded(v interface{}) {
	o.Succeeded = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetMessage(v interface{}) {
	o.Message = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetErrorCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetErrorCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return &o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasErrorCode() bool {
	if o != nil && IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given interface{} and assigns it to the ErrorCode field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetErrorCode(v interface{}) {
	o.ErrorCode = v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetErrorDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetErrorDetailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorDetail) {
		return nil, false
	}
	return &o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasErrorDetail() bool {
	if o != nil && IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given interface{} and assigns it to the ErrorDetail field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetErrorDetail(v interface{}) {
	o.ErrorDetail = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetGatewayTransactionId returns the GatewayTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetGatewayTransactionId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GatewayTransactionId
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetGatewayTransactionIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GatewayTransactionId) {
		return nil, false
	}
	return &o.GatewayTransactionId, true
}

// HasGatewayTransactionId returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasGatewayTransactionId() bool {
	if o != nil && IsNil(o.GatewayTransactionId) {
		return true
	}

	return false
}

// SetGatewayTransactionId gets a reference to the given interface{} and assigns it to the GatewayTransactionId field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetGatewayTransactionId(v interface{}) {
	o.GatewayTransactionId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetRefundAmountCents returns the RefundAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundAmountCents
}

// GetRefundAmountCentsOk returns a tuple with the RefundAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RefundAmountCents) {
		return nil, false
	}
	return &o.RefundAmountCents, true
}

// HasRefundAmountCents returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasRefundAmountCents() bool {
	if o != nil && IsNil(o.RefundAmountCents) {
		return true
	}

	return false
}

// SetRefundAmountCents gets a reference to the given interface{} and assigns it to the RefundAmountCents field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetRefundAmountCents(v interface{}) {
	o.RefundAmountCents = v
}

// GetRefundAmountFloat returns the RefundAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundAmountFloat
}

// GetRefundAmountFloatOk returns a tuple with the RefundAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RefundAmountFloat) {
		return nil, false
	}
	return &o.RefundAmountFloat, true
}

// HasRefundAmountFloat returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasRefundAmountFloat() bool {
	if o != nil && IsNil(o.RefundAmountFloat) {
		return true
	}

	return false
}

// SetRefundAmountFloat gets a reference to the given interface{} and assigns it to the RefundAmountFloat field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetRefundAmountFloat(v interface{}) {
	o.RefundAmountFloat = v
}

// GetFormattedRefundAmount returns the FormattedRefundAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedRefundAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedRefundAmount
}

// GetFormattedRefundAmountOk returns a tuple with the FormattedRefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedRefundAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedRefundAmount) {
		return nil, false
	}
	return &o.FormattedRefundAmount, true
}

// HasFormattedRefundAmount returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasFormattedRefundAmount() bool {
	if o != nil && IsNil(o.FormattedRefundAmount) {
		return true
	}

	return false
}

// SetFormattedRefundAmount gets a reference to the given interface{} and assigns it to the FormattedRefundAmount field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetFormattedRefundAmount(v interface{}) {
	o.FormattedRefundAmount = v
}

// GetRefundBalanceCents returns the RefundBalanceCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundBalanceCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundBalanceCents
}

// GetRefundBalanceCentsOk returns a tuple with the RefundBalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundBalanceCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RefundBalanceCents) {
		return nil, false
	}
	return &o.RefundBalanceCents, true
}

// HasRefundBalanceCents returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasRefundBalanceCents() bool {
	if o != nil && IsNil(o.RefundBalanceCents) {
		return true
	}

	return false
}

// SetRefundBalanceCents gets a reference to the given interface{} and assigns it to the RefundBalanceCents field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetRefundBalanceCents(v interface{}) {
	o.RefundBalanceCents = v
}

// GetRefundBalanceFloat returns the RefundBalanceFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundBalanceFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RefundBalanceFloat
}

// GetRefundBalanceFloatOk returns a tuple with the RefundBalanceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetRefundBalanceFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RefundBalanceFloat) {
		return nil, false
	}
	return &o.RefundBalanceFloat, true
}

// HasRefundBalanceFloat returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasRefundBalanceFloat() bool {
	if o != nil && IsNil(o.RefundBalanceFloat) {
		return true
	}

	return false
}

// SetRefundBalanceFloat gets a reference to the given interface{} and assigns it to the RefundBalanceFloat field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetRefundBalanceFloat(v interface{}) {
	o.RefundBalanceFloat = v
}

// GetFormattedRefundBalance returns the FormattedRefundBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedRefundBalance() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedRefundBalance
}

// GetFormattedRefundBalanceOk returns a tuple with the FormattedRefundBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCapturesCaptureId200ResponseDataAttributes) GetFormattedRefundBalanceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedRefundBalance) {
		return nil, false
	}
	return &o.FormattedRefundBalance, true
}

// HasFormattedRefundBalance returns a boolean if a field has been set.
func (o *GETCapturesCaptureId200ResponseDataAttributes) HasFormattedRefundBalance() bool {
	if o != nil && IsNil(o.FormattedRefundBalance) {
		return true
	}

	return false
}

// SetFormattedRefundBalance gets a reference to the given interface{} and assigns it to the FormattedRefundBalance field.
func (o *GETCapturesCaptureId200ResponseDataAttributes) SetFormattedRefundBalance(v interface{}) {
	o.FormattedRefundBalance = v
}

func (o GETCapturesCaptureId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETCapturesCaptureId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if o.AmountFloat != nil {
		toSerialize["amount_float"] = o.AmountFloat
	}
	if o.FormattedAmount != nil {
		toSerialize["formatted_amount"] = o.FormattedAmount
	}
	if o.Succeeded != nil {
		toSerialize["succeeded"] = o.Succeeded
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorDetail != nil {
		toSerialize["error_detail"] = o.ErrorDetail
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.GatewayTransactionId != nil {
		toSerialize["gateway_transaction_id"] = o.GatewayTransactionId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.RefundAmountCents != nil {
		toSerialize["refund_amount_cents"] = o.RefundAmountCents
	}
	if o.RefundAmountFloat != nil {
		toSerialize["refund_amount_float"] = o.RefundAmountFloat
	}
	if o.FormattedRefundAmount != nil {
		toSerialize["formatted_refund_amount"] = o.FormattedRefundAmount
	}
	if o.RefundBalanceCents != nil {
		toSerialize["refund_balance_cents"] = o.RefundBalanceCents
	}
	if o.RefundBalanceFloat != nil {
		toSerialize["refund_balance_float"] = o.RefundBalanceFloat
	}
	if o.FormattedRefundBalance != nil {
		toSerialize["formatted_refund_balance"] = o.FormattedRefundBalance
	}
	return toSerialize, nil
}

type NullableGETCapturesCaptureId200ResponseDataAttributes struct {
	value *GETCapturesCaptureId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETCapturesCaptureId200ResponseDataAttributes) Get() *GETCapturesCaptureId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETCapturesCaptureId200ResponseDataAttributes) Set(val *GETCapturesCaptureId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCapturesCaptureId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCapturesCaptureId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCapturesCaptureId200ResponseDataAttributes(val *GETCapturesCaptureId200ResponseDataAttributes) *NullableGETCapturesCaptureId200ResponseDataAttributes {
	return &NullableGETCapturesCaptureId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETCapturesCaptureId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCapturesCaptureId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
