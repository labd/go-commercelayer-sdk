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

// checks if the SkuOptionUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuOptionUpdateData{}

// SkuOptionUpdateData struct for SkuOptionUpdateData
type SkuOptionUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                         `json:"id"`
	Attributes    PATCHSkuOptionsSkuOptionId200ResponseDataAttributes `json:"attributes"`
	Relationships *SkuOptionCreateDataRelationships                   `json:"relationships,omitempty"`
}

// NewSkuOptionUpdateData instantiates a new SkuOptionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionUpdateData(type_ interface{}, id interface{}, attributes PATCHSkuOptionsSkuOptionId200ResponseDataAttributes) *SkuOptionUpdateData {
	this := SkuOptionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSkuOptionUpdateDataWithDefaults instantiates a new SkuOptionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionUpdateDataWithDefaults() *SkuOptionUpdateData {
	this := SkuOptionUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuOptionUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuOptionUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuOptionUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SkuOptionUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SkuOptionUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SkuOptionUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SkuOptionUpdateData) GetAttributes() PATCHSkuOptionsSkuOptionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHSkuOptionsSkuOptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetAttributesOk() (*PATCHSkuOptionsSkuOptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuOptionUpdateData) SetAttributes(v PATCHSkuOptionsSkuOptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuOptionUpdateData) GetRelationships() SkuOptionCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SkuOptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdateData) GetRelationshipsOk() (*SkuOptionCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuOptionUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuOptionCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuOptionUpdateData) SetRelationships(v SkuOptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuOptionUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuOptionUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableSkuOptionUpdateData struct {
	value *SkuOptionUpdateData
	isSet bool
}

func (v NullableSkuOptionUpdateData) Get() *SkuOptionUpdateData {
	return v.value
}

func (v *NullableSkuOptionUpdateData) Set(val *SkuOptionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionUpdateData(val *SkuOptionUpdateData) *NullableSkuOptionUpdateData {
	return &NullableSkuOptionUpdateData{value: val, isSet: true}
}

func (v NullableSkuOptionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
