/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CheckoutComPaymentData struct for CheckoutComPaymentData
type CheckoutComPaymentData struct {
	// The resource's type
	Type          string                                               `json:"type"`
	Attributes    GETCheckoutComPayments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AdyenPaymentDataRelationships                       `json:"relationships,omitempty"`
}

// NewCheckoutComPaymentData instantiates a new CheckoutComPaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPaymentData(type_ string, attributes GETCheckoutComPayments200ResponseDataInnerAttributes) *CheckoutComPaymentData {
	this := CheckoutComPaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCheckoutComPaymentDataWithDefaults instantiates a new CheckoutComPaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentDataWithDefaults() *CheckoutComPaymentData {
	this := CheckoutComPaymentData{}
	return &this
}

// GetType returns the Type field value
func (o *CheckoutComPaymentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComPaymentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComPaymentData) GetAttributes() GETCheckoutComPayments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCheckoutComPayments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentData) GetAttributesOk() (*GETCheckoutComPayments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComPaymentData) SetAttributes(v GETCheckoutComPayments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComPaymentData) GetRelationships() AdyenPaymentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComPaymentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComPaymentData) SetRelationships(v AdyenPaymentDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComPaymentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComPaymentData struct {
	value *CheckoutComPaymentData
	isSet bool
}

func (v NullableCheckoutComPaymentData) Get() *CheckoutComPaymentData {
	return v.value
}

func (v *NullableCheckoutComPaymentData) Set(val *CheckoutComPaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPaymentData(val *CheckoutComPaymentData) *NullableCheckoutComPaymentData {
	return &NullableCheckoutComPaymentData{value: val, isSet: true}
}

func (v NullableCheckoutComPaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
