/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddressData struct for AddressData
type AddressData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    GETAddresses200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *AddressDataRelationships                  `json:"relationships,omitempty"`
}

// NewAddressData instantiates a new AddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressData(type_ string, attributes GETAddresses200ResponseDataInnerAttributes) *AddressData {
	this := AddressData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAddressDataWithDefaults instantiates a new AddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataWithDefaults() *AddressData {
	this := AddressData{}
	return &this
}

// GetType returns the Type field value
func (o *AddressData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AddressData) GetAttributes() GETAddresses200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETAddresses200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetAttributesOk() (*GETAddresses200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddressData) SetAttributes(v GETAddresses200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AddressData) GetRelationships() AddressDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AddressDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressData) GetRelationshipsOk() (*AddressDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AddressData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AddressDataRelationships and assigns it to the Relationships field.
func (o *AddressData) SetRelationships(v AddressDataRelationships) {
	o.Relationships = &v
}

func (o AddressData) MarshalJSON() ([]byte, error) {
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

type NullableAddressData struct {
	value *AddressData
	isSet bool
}

func (v NullableAddressData) Get() *AddressData {
	return v.value
}

func (v *NullableAddressData) Set(val *AddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressData(val *AddressData) *NullableAddressData {
	return &NullableAddressData{value: val, isSet: true}
}

func (v NullableAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
