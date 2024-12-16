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

// checks if the AxervePaymentCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxervePaymentCreateData{}

// AxervePaymentCreateData struct for AxervePaymentCreateData
type AxervePaymentCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTAxervePayments201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships        `json:"relationships,omitempty"`
}

// NewAxervePaymentCreateData instantiates a new AxervePaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxervePaymentCreateData(type_ interface{}, attributes POSTAxervePayments201ResponseDataAttributes) *AxervePaymentCreateData {
	this := AxervePaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAxervePaymentCreateDataWithDefaults instantiates a new AxervePaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxervePaymentCreateDataWithDefaults() *AxervePaymentCreateData {
	this := AxervePaymentCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AxervePaymentCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxervePaymentCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AxervePaymentCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AxervePaymentCreateData) GetAttributes() POSTAxervePayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAxervePayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AxervePaymentCreateData) GetAttributesOk() (*POSTAxervePayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AxervePaymentCreateData) SetAttributes(v POSTAxervePayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AxervePaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AxervePaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AxervePaymentCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *AxervePaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o AxervePaymentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxervePaymentCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableAxervePaymentCreateData struct {
	value *AxervePaymentCreateData
	isSet bool
}

func (v NullableAxervePaymentCreateData) Get() *AxervePaymentCreateData {
	return v.value
}

func (v *NullableAxervePaymentCreateData) Set(val *AxervePaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAxervePaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAxervePaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxervePaymentCreateData(val *AxervePaymentCreateData) *NullableAxervePaymentCreateData {
	return &NullableAxervePaymentCreateData{value: val, isSet: true}
}

func (v NullableAxervePaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxervePaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
