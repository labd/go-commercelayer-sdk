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

// checks if the CustomerPaymentSourceCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPaymentSourceCreateData{}

// CustomerPaymentSourceCreateData struct for CustomerPaymentSourceCreateData
type CustomerPaymentSourceCreateData struct {
	// The resource's type
	Type          interface{}                                         `json:"type"`
	Attributes    POSTCustomerPaymentSources201ResponseDataAttributes `json:"attributes"`
	Relationships *CustomerPaymentSourceCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewCustomerPaymentSourceCreateData instantiates a new CustomerPaymentSourceCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceCreateData(type_ interface{}, attributes POSTCustomerPaymentSources201ResponseDataAttributes) *CustomerPaymentSourceCreateData {
	this := CustomerPaymentSourceCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerPaymentSourceCreateDataWithDefaults instantiates a new CustomerPaymentSourceCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceCreateDataWithDefaults() *CustomerPaymentSourceCreateData {
	this := CustomerPaymentSourceCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerPaymentSourceCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPaymentSourceCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerPaymentSourceCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerPaymentSourceCreateData) GetAttributes() POSTCustomerPaymentSources201ResponseDataAttributes {
	if o == nil {
		var ret POSTCustomerPaymentSources201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateData) GetAttributesOk() (*POSTCustomerPaymentSources201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerPaymentSourceCreateData) SetAttributes(v POSTCustomerPaymentSources201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerPaymentSourceCreateData) GetRelationships() CustomerPaymentSourceCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerPaymentSourceCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceCreateData) GetRelationshipsOk() (*CustomerPaymentSourceCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerPaymentSourceCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerPaymentSourceCreateDataRelationships and assigns it to the Relationships field.
func (o *CustomerPaymentSourceCreateData) SetRelationships(v CustomerPaymentSourceCreateDataRelationships) {
	o.Relationships = &v
}

func (o CustomerPaymentSourceCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPaymentSourceCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableCustomerPaymentSourceCreateData struct {
	value *CustomerPaymentSourceCreateData
	isSet bool
}

func (v NullableCustomerPaymentSourceCreateData) Get() *CustomerPaymentSourceCreateData {
	return v.value
}

func (v *NullableCustomerPaymentSourceCreateData) Set(val *CustomerPaymentSourceCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceCreateData(val *CustomerPaymentSourceCreateData) *NullableCustomerPaymentSourceCreateData {
	return &NullableCustomerPaymentSourceCreateData{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
