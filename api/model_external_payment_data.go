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

// checks if the ExternalPaymentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPaymentData{}

// ExternalPaymentData struct for ExternalPaymentData
type ExternalPaymentData struct {
	// The resource's type
	Type          interface{}                                                   `json:"type"`
	Attributes    GETExternalPaymentsExternalPaymentId200ResponseDataAttributes `json:"attributes"`
	Relationships *ExternalPaymentDataRelationships                             `json:"relationships,omitempty"`
}

// NewExternalPaymentData instantiates a new ExternalPaymentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentData(type_ interface{}, attributes GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) *ExternalPaymentData {
	this := ExternalPaymentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExternalPaymentDataWithDefaults instantiates a new ExternalPaymentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentDataWithDefaults() *ExternalPaymentData {
	this := ExternalPaymentData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalPaymentData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalPaymentData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalPaymentData) GetAttributes() GETExternalPaymentsExternalPaymentId200ResponseDataAttributes {
	if o == nil {
		var ret GETExternalPaymentsExternalPaymentId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentData) GetAttributesOk() (*GETExternalPaymentsExternalPaymentId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalPaymentData) SetAttributes(v GETExternalPaymentsExternalPaymentId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalPaymentData) GetRelationships() ExternalPaymentDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ExternalPaymentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentData) GetRelationshipsOk() (*ExternalPaymentDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalPaymentData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPaymentDataRelationships and assigns it to the Relationships field.
func (o *ExternalPaymentData) SetRelationships(v ExternalPaymentDataRelationships) {
	o.Relationships = &v
}

func (o ExternalPaymentData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPaymentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableExternalPaymentData struct {
	value *ExternalPaymentData
	isSet bool
}

func (v NullableExternalPaymentData) Get() *ExternalPaymentData {
	return v.value
}

func (v *NullableExternalPaymentData) Set(val *ExternalPaymentData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentData(val *ExternalPaymentData) *NullableExternalPaymentData {
	return &NullableExternalPaymentData{value: val, isSet: true}
}

func (v NullableExternalPaymentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
