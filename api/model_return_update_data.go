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

// checks if the ReturnUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnUpdateData{}

// ReturnUpdateData struct for ReturnUpdateData
type ReturnUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                   `json:"id"`
	Attributes    PATCHReturnsReturnId200ResponseDataAttributes `json:"attributes"`
	Relationships *ReturnUpdateDataRelationships                `json:"relationships,omitempty"`
}

// NewReturnUpdateData instantiates a new ReturnUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnUpdateData(type_ interface{}, id interface{}, attributes PATCHReturnsReturnId200ResponseDataAttributes) *ReturnUpdateData {
	this := ReturnUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewReturnUpdateDataWithDefaults instantiates a new ReturnUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnUpdateDataWithDefaults() *ReturnUpdateData {
	this := ReturnUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ReturnUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReturnUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ReturnUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReturnUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReturnUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ReturnUpdateData) GetAttributes() PATCHReturnsReturnId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHReturnsReturnId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetAttributesOk() (*PATCHReturnsReturnId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ReturnUpdateData) SetAttributes(v PATCHReturnsReturnId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReturnUpdateData) GetRelationships() ReturnUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ReturnUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnUpdateData) GetRelationshipsOk() (*ReturnUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReturnUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReturnUpdateDataRelationships and assigns it to the Relationships field.
func (o *ReturnUpdateData) SetRelationships(v ReturnUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ReturnUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableReturnUpdateData struct {
	value *ReturnUpdateData
	isSet bool
}

func (v NullableReturnUpdateData) Get() *ReturnUpdateData {
	return v.value
}

func (v *NullableReturnUpdateData) Set(val *ReturnUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnUpdateData(val *ReturnUpdateData) *NullableReturnUpdateData {
	return &NullableReturnUpdateData{value: val, isSet: true}
}

func (v NullableReturnUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
