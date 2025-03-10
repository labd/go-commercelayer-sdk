/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.6.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETRefundsRefundId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETRefundsRefundId200ResponseDataAttributes{}

// GETRefundsRefundId200ResponseDataAttributes struct for GETRefundsRefundId200ResponseDataAttributes
type GETRefundsRefundId200ResponseDataAttributes struct {
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
}

// NewGETRefundsRefundId200ResponseDataAttributes instantiates a new GETRefundsRefundId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETRefundsRefundId200ResponseDataAttributes() *GETRefundsRefundId200ResponseDataAttributes {
	this := GETRefundsRefundId200ResponseDataAttributes{}
	return &this
}

// NewGETRefundsRefundId200ResponseDataAttributesWithDefaults instantiates a new GETRefundsRefundId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETRefundsRefundId200ResponseDataAttributesWithDefaults() *GETRefundsRefundId200ResponseDataAttributes {
	this := GETRefundsRefundId200ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountCents) {
		return nil, false
	}
	return &o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasAmountCents() bool {
	if o != nil && IsNil(o.AmountCents) {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given interface{} and assigns it to the AmountCents field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountCents(v interface{}) {
	o.AmountCents = v
}

// GetAmountFloat returns the AmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AmountFloat
}

// GetAmountFloatOk returns a tuple with the AmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AmountFloat) {
		return nil, false
	}
	return &o.AmountFloat, true
}

// HasAmountFloat returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasAmountFloat() bool {
	if o != nil && IsNil(o.AmountFloat) {
		return true
	}

	return false
}

// SetAmountFloat gets a reference to the given interface{} and assigns it to the AmountFloat field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetAmountFloat(v interface{}) {
	o.AmountFloat = v
}

// GetFormattedAmount returns the FormattedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetFormattedAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedAmount
}

// GetFormattedAmountOk returns a tuple with the FormattedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedAmount) {
		return nil, false
	}
	return &o.FormattedAmount, true
}

// HasFormattedAmount returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasFormattedAmount() bool {
	if o != nil && IsNil(o.FormattedAmount) {
		return true
	}

	return false
}

// SetFormattedAmount gets a reference to the given interface{} and assigns it to the FormattedAmount field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetFormattedAmount(v interface{}) {
	o.FormattedAmount = v
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetSucceeded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetSucceededOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Succeeded) {
		return nil, false
	}
	return &o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasSucceeded() bool {
	if o != nil && IsNil(o.Succeeded) {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given interface{} and assigns it to the Succeeded field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetSucceeded(v interface{}) {
	o.Succeeded = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetMessage(v interface{}) {
	o.Message = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return &o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasErrorCode() bool {
	if o != nil && IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given interface{} and assigns it to the ErrorCode field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorCode(v interface{}) {
	o.ErrorCode = v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetErrorDetailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ErrorDetail) {
		return nil, false
	}
	return &o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasErrorDetail() bool {
	if o != nil && IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given interface{} and assigns it to the ErrorDetail field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetErrorDetail(v interface{}) {
	o.ErrorDetail = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return &o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasToken() bool {
	if o != nil && IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given interface{} and assigns it to the Token field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetToken(v interface{}) {
	o.Token = v
}

// GetGatewayTransactionId returns the GatewayTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetGatewayTransactionId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GatewayTransactionId
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetGatewayTransactionIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GatewayTransactionId) {
		return nil, false
	}
	return &o.GatewayTransactionId, true
}

// HasGatewayTransactionId returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasGatewayTransactionId() bool {
	if o != nil && IsNil(o.GatewayTransactionId) {
		return true
	}

	return false
}

// SetGatewayTransactionId gets a reference to the given interface{} and assigns it to the GatewayTransactionId field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetGatewayTransactionId(v interface{}) {
	o.GatewayTransactionId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETRefundsRefundId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETRefundsRefundId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETRefundsRefundId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETRefundsRefundId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETRefundsRefundId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETRefundsRefundId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableGETRefundsRefundId200ResponseDataAttributes struct {
	value *GETRefundsRefundId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETRefundsRefundId200ResponseDataAttributes) Get() *GETRefundsRefundId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETRefundsRefundId200ResponseDataAttributes) Set(val *GETRefundsRefundId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETRefundsRefundId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETRefundsRefundId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETRefundsRefundId200ResponseDataAttributes(val *GETRefundsRefundId200ResponseDataAttributes) *NullableGETRefundsRefundId200ResponseDataAttributes {
	return &NullableGETRefundsRefundId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETRefundsRefundId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETRefundsRefundId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
