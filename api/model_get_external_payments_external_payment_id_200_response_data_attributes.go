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

// checks if the GETExternalPaymentsExternalPaymentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETExternalPaymentsExternalPaymentId200ResponseDataAttributes{}

// GETExternalPaymentsExternalPaymentId200ResponseDataAttributes struct for GETExternalPaymentsExternalPaymentId200ResponseDataAttributes
type GETExternalPaymentsExternalPaymentId200ResponseDataAttributes struct {
	// The payment source token, as generated by the external gateway SDK. Credit Card numbers are rejected.
	PaymentSourceToken interface{} `json:"payment_source_token,omitempty"`
	// External payment options.
	Options interface{} `json:"options,omitempty"`
	// Information about the payment instrument used in the transaction.
	PaymentInstrument interface{} `json:"payment_instrument,omitempty"`
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

// NewGETExternalPaymentsExternalPaymentId200ResponseDataAttributes instantiates a new GETExternalPaymentsExternalPaymentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPaymentsExternalPaymentId200ResponseDataAttributes() *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	this := GETExternalPaymentsExternalPaymentId200ResponseDataAttributes{}
	return &this
}

// NewGETExternalPaymentsExternalPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETExternalPaymentsExternalPaymentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPaymentsExternalPaymentId200ResponseDataAttributesWithDefaults() *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	this := GETExternalPaymentsExternalPaymentId200ResponseDataAttributes{}
	return &this
}

// GetPaymentSourceToken returns the PaymentSourceToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetPaymentSourceToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentSourceToken
}

// GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetPaymentSourceTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentSourceToken) {
		return nil, false
	}
	return &o.PaymentSourceToken, true
}

// HasPaymentSourceToken returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasPaymentSourceToken() bool {
	if o != nil && IsNil(o.PaymentSourceToken) {
		return true
	}

	return false
}

// SetPaymentSourceToken gets a reference to the given interface{} and assigns it to the PaymentSourceToken field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetPaymentSourceToken(v interface{}) {
	o.PaymentSourceToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetOptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasOptions() bool {
	if o != nil && IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given interface{} and assigns it to the Options field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetOptions(v interface{}) {
	o.Options = v
}

// GetPaymentInstrument returns the PaymentInstrument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PaymentInstrument) {
		return nil, false
	}
	return &o.PaymentInstrument, true
}

// HasPaymentInstrument returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool {
	if o != nil && IsNil(o.PaymentInstrument) {
		return true
	}

	return false
}

// SetPaymentInstrument gets a reference to the given interface{} and assigns it to the PaymentInstrument field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{}) {
	o.PaymentInstrument = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentSourceToken != nil {
		toSerialize["payment_source_token"] = o.PaymentSourceToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PaymentInstrument != nil {
		toSerialize["payment_instrument"] = o.PaymentInstrument
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

type NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes struct {
	value *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) Get() *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) Set(val *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes(val *GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) *NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	return &NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPaymentsExternalPaymentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
