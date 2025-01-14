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

// checks if the ParcelLineItemCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelLineItemCreateData{}

// ParcelLineItemCreateData struct for ParcelLineItemCreateData
type ParcelLineItemCreateData struct {
	// The resource's type
	Type          interface{}                                  `json:"type"`
	Attributes    POSTParcelLineItems201ResponseDataAttributes `json:"attributes"`
	Relationships *ParcelLineItemCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewParcelLineItemCreateData instantiates a new ParcelLineItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemCreateData(type_ interface{}, attributes POSTParcelLineItems201ResponseDataAttributes) *ParcelLineItemCreateData {
	this := ParcelLineItemCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewParcelLineItemCreateDataWithDefaults instantiates a new ParcelLineItemCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemCreateDataWithDefaults() *ParcelLineItemCreateData {
	this := ParcelLineItemCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ParcelLineItemCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParcelLineItemCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelLineItemCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelLineItemCreateData) GetAttributes() POSTParcelLineItems201ResponseDataAttributes {
	if o == nil {
		var ret POSTParcelLineItems201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateData) GetAttributesOk() (*POSTParcelLineItems201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelLineItemCreateData) SetAttributes(v POSTParcelLineItems201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelLineItemCreateData) GetRelationships() ParcelLineItemCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ParcelLineItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateData) GetRelationshipsOk() (*ParcelLineItemCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelLineItemCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ParcelLineItemCreateDataRelationships and assigns it to the Relationships field.
func (o *ParcelLineItemCreateData) SetRelationships(v ParcelLineItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o ParcelLineItemCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelLineItemCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableParcelLineItemCreateData struct {
	value *ParcelLineItemCreateData
	isSet bool
}

func (v NullableParcelLineItemCreateData) Get() *ParcelLineItemCreateData {
	return v.value
}

func (v *NullableParcelLineItemCreateData) Set(val *ParcelLineItemCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemCreateData(val *ParcelLineItemCreateData) *NullableParcelLineItemCreateData {
	return &NullableParcelLineItemCreateData{value: val, isSet: true}
}

func (v NullableParcelLineItemCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
