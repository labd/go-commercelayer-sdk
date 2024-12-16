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

// checks if the CarrierAccountCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAccountCreateData{}

// CarrierAccountCreateData struct for CarrierAccountCreateData
type CarrierAccountCreateData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTCarrierAccounts201ResponseDataAttributes `json:"attributes"`
	Relationships *CarrierAccountCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewCarrierAccountCreateData instantiates a new CarrierAccountCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccountCreateData(type_ interface{}, attributes POSTCarrierAccounts201ResponseDataAttributes) *CarrierAccountCreateData {
	this := CarrierAccountCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCarrierAccountCreateDataWithDefaults instantiates a new CarrierAccountCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountCreateDataWithDefaults() *CarrierAccountCreateData {
	this := CarrierAccountCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CarrierAccountCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CarrierAccountCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CarrierAccountCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CarrierAccountCreateData) GetAttributes() POSTCarrierAccounts201ResponseDataAttributes {
	if o == nil {
		var ret POSTCarrierAccounts201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CarrierAccountCreateData) GetAttributesOk() (*POSTCarrierAccounts201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CarrierAccountCreateData) SetAttributes(v POSTCarrierAccounts201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CarrierAccountCreateData) GetRelationships() CarrierAccountCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CarrierAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountCreateData) GetRelationshipsOk() (*CarrierAccountCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CarrierAccountCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CarrierAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *CarrierAccountCreateData) SetRelationships(v CarrierAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o CarrierAccountCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarrierAccountCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableCarrierAccountCreateData struct {
	value *CarrierAccountCreateData
	isSet bool
}

func (v NullableCarrierAccountCreateData) Get() *CarrierAccountCreateData {
	return v.value
}

func (v *NullableCarrierAccountCreateData) Set(val *CarrierAccountCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccountCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccountCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccountCreateData(val *CarrierAccountCreateData) *NullableCarrierAccountCreateData {
	return &NullableCarrierAccountCreateData{value: val, isSet: true}
}

func (v NullableCarrierAccountCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierAccountCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
