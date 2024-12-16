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

// checks if the CarrierAccountData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAccountData{}

// CarrierAccountData struct for CarrierAccountData
type CarrierAccountData struct {
	// The resource's type
	Type          interface{}                                                 `json:"type"`
	Attributes    GETCarrierAccountsCarrierAccountId200ResponseDataAttributes `json:"attributes"`
	Relationships *CarrierAccountDataRelationships                            `json:"relationships,omitempty"`
}

// NewCarrierAccountData instantiates a new CarrierAccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccountData(type_ interface{}, attributes GETCarrierAccountsCarrierAccountId200ResponseDataAttributes) *CarrierAccountData {
	this := CarrierAccountData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCarrierAccountDataWithDefaults instantiates a new CarrierAccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountDataWithDefaults() *CarrierAccountData {
	this := CarrierAccountData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CarrierAccountData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CarrierAccountData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CarrierAccountData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CarrierAccountData) GetAttributes() GETCarrierAccountsCarrierAccountId200ResponseDataAttributes {
	if o == nil {
		var ret GETCarrierAccountsCarrierAccountId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CarrierAccountData) GetAttributesOk() (*GETCarrierAccountsCarrierAccountId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CarrierAccountData) SetAttributes(v GETCarrierAccountsCarrierAccountId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CarrierAccountData) GetRelationships() CarrierAccountDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CarrierAccountDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountData) GetRelationshipsOk() (*CarrierAccountDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CarrierAccountData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CarrierAccountDataRelationships and assigns it to the Relationships field.
func (o *CarrierAccountData) SetRelationships(v CarrierAccountDataRelationships) {
	o.Relationships = &v
}

func (o CarrierAccountData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarrierAccountData) ToMap() (map[string]interface{}, error) {
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

type NullableCarrierAccountData struct {
	value *CarrierAccountData
	isSet bool
}

func (v NullableCarrierAccountData) Get() *CarrierAccountData {
	return v.value
}

func (v *NullableCarrierAccountData) Set(val *CarrierAccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccountData(val *CarrierAccountData) *NullableCarrierAccountData {
	return &NullableCarrierAccountData{value: val, isSet: true}
}

func (v NullableCarrierAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
