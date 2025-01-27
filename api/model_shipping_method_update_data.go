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

// checks if the ShippingMethodUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodUpdateData{}

// ShippingMethodUpdateData struct for ShippingMethodUpdateData
type ShippingMethodUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                   `json:"id"`
	Attributes    PATCHShippingMethodsShippingMethodId200ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingMethodCreateDataRelationships                        `json:"relationships,omitempty"`
}

// NewShippingMethodUpdateData instantiates a new ShippingMethodUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodUpdateData(type_ interface{}, id interface{}, attributes PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) *ShippingMethodUpdateData {
	this := ShippingMethodUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewShippingMethodUpdateDataWithDefaults instantiates a new ShippingMethodUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodUpdateDataWithDefaults() *ShippingMethodUpdateData {
	this := ShippingMethodUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingMethodUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingMethodUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShippingMethodUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShippingMethodUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingMethodUpdateData) GetAttributes() PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHShippingMethodsShippingMethodId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetAttributesOk() (*PATCHShippingMethodsShippingMethodId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingMethodUpdateData) SetAttributes(v PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingMethodUpdateData) GetRelationships() ShippingMethodCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ShippingMethodCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetRelationshipsOk() (*ShippingMethodCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingMethodUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodCreateDataRelationships and assigns it to the Relationships field.
func (o *ShippingMethodUpdateData) SetRelationships(v ShippingMethodCreateDataRelationships) {
	o.Relationships = &v
}

func (o ShippingMethodUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableShippingMethodUpdateData struct {
	value *ShippingMethodUpdateData
	isSet bool
}

func (v NullableShippingMethodUpdateData) Get() *ShippingMethodUpdateData {
	return v.value
}

func (v *NullableShippingMethodUpdateData) Set(val *ShippingMethodUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodUpdateData(val *ShippingMethodUpdateData) *NullableShippingMethodUpdateData {
	return &NullableShippingMethodUpdateData{value: val, isSet: true}
}

func (v NullableShippingMethodUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
