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

// checks if the AddressUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressUpdateData{}

// AddressUpdateData struct for AddressUpdateData
type AddressUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                      `json:"id"`
	Attributes    PATCHAddressesAddressId200ResponseDataAttributes `json:"attributes"`
	Relationships *AddressCreateDataRelationships                  `json:"relationships,omitempty"`
}

// NewAddressUpdateData instantiates a new AddressUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressUpdateData(type_ interface{}, id interface{}, attributes PATCHAddressesAddressId200ResponseDataAttributes) *AddressUpdateData {
	this := AddressUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewAddressUpdateDataWithDefaults instantiates a new AddressUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressUpdateDataWithDefaults() *AddressUpdateData {
	this := AddressUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AddressUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddressUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AddressUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddressUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *AddressUpdateData) GetAttributes() PATCHAddressesAddressId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHAddressesAddressId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetAttributesOk() (*PATCHAddressesAddressId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AddressUpdateData) SetAttributes(v PATCHAddressesAddressId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AddressUpdateData) GetRelationships() AddressCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AddressCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressUpdateData) GetRelationshipsOk() (*AddressCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AddressUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AddressCreateDataRelationships and assigns it to the Relationships field.
func (o *AddressUpdateData) SetRelationships(v AddressCreateDataRelationships) {
	o.Relationships = &v
}

func (o AddressUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableAddressUpdateData struct {
	value *AddressUpdateData
	isSet bool
}

func (v NullableAddressUpdateData) Get() *AddressUpdateData {
	return v.value
}

func (v *NullableAddressUpdateData) Set(val *AddressUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressUpdateData(val *AddressUpdateData) *NullableAddressUpdateData {
	return &NullableAddressUpdateData{value: val, isSet: true}
}

func (v NullableAddressUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
