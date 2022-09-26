/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ParcelUpdateData struct for ParcelUpdateData
type ParcelUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                         `json:"id"`
	Attributes    ParcelUpdateDataAttributes     `json:"attributes"`
	Relationships *ParcelUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewParcelUpdateData instantiates a new ParcelUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelUpdateData(type_ string, id string, attributes ParcelUpdateDataAttributes) *ParcelUpdateData {
	this := ParcelUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewParcelUpdateDataWithDefaults instantiates a new ParcelUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelUpdateDataWithDefaults() *ParcelUpdateData {
	this := ParcelUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *ParcelUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ParcelUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParcelUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParcelUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelUpdateData) GetAttributes() ParcelUpdateDataAttributes {
	if o == nil {
		var ret ParcelUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelUpdateData) GetAttributesOk() (*ParcelUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelUpdateData) SetAttributes(v ParcelUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelUpdateData) GetRelationships() ParcelUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ParcelUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelUpdateData) GetRelationshipsOk() (*ParcelUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ParcelUpdateDataRelationships and assigns it to the Relationships field.
func (o *ParcelUpdateData) SetRelationships(v ParcelUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ParcelUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableParcelUpdateData struct {
	value *ParcelUpdateData
	isSet bool
}

func (v NullableParcelUpdateData) Get() *ParcelUpdateData {
	return v.value
}

func (v *NullableParcelUpdateData) Set(val *ParcelUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelUpdateData(val *ParcelUpdateData) *NullableParcelUpdateData {
	return &NullableParcelUpdateData{value: val, isSet: true}
}

func (v NullableParcelUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
