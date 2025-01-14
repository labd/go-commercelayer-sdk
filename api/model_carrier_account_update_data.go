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

// checks if the CarrierAccountUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CarrierAccountUpdateData{}

// CarrierAccountUpdateData struct for CarrierAccountUpdateData
type CarrierAccountUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                   `json:"id"`
	Attributes    PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes `json:"attributes"`
	Relationships *CarrierAccountCreateDataRelationships                        `json:"relationships,omitempty"`
}

// NewCarrierAccountUpdateData instantiates a new CarrierAccountUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierAccountUpdateData(type_ interface{}, id interface{}, attributes PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes) *CarrierAccountUpdateData {
	this := CarrierAccountUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCarrierAccountUpdateDataWithDefaults instantiates a new CarrierAccountUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierAccountUpdateDataWithDefaults() *CarrierAccountUpdateData {
	this := CarrierAccountUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CarrierAccountUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CarrierAccountUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CarrierAccountUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CarrierAccountUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CarrierAccountUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CarrierAccountUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CarrierAccountUpdateData) GetAttributes() PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CarrierAccountUpdateData) GetAttributesOk() (*PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CarrierAccountUpdateData) SetAttributes(v PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CarrierAccountUpdateData) GetRelationships() CarrierAccountCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CarrierAccountCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierAccountUpdateData) GetRelationshipsOk() (*CarrierAccountCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CarrierAccountUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CarrierAccountCreateDataRelationships and assigns it to the Relationships field.
func (o *CarrierAccountUpdateData) SetRelationships(v CarrierAccountCreateDataRelationships) {
	o.Relationships = &v
}

func (o CarrierAccountUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarrierAccountUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableCarrierAccountUpdateData struct {
	value *CarrierAccountUpdateData
	isSet bool
}

func (v NullableCarrierAccountUpdateData) Get() *CarrierAccountUpdateData {
	return v.value
}

func (v *NullableCarrierAccountUpdateData) Set(val *CarrierAccountUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierAccountUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierAccountUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierAccountUpdateData(val *CarrierAccountUpdateData) *NullableCarrierAccountUpdateData {
	return &NullableCarrierAccountUpdateData{value: val, isSet: true}
}

func (v NullableCarrierAccountUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierAccountUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
