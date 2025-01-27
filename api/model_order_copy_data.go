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

// checks if the OrderCopyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCopyData{}

// OrderCopyData struct for OrderCopyData
type OrderCopyData struct {
	// The resource's type
	Type          interface{}                                        `json:"type"`
	Attributes    GETOrderCopiesOrderCopyId200ResponseDataAttributes `json:"attributes"`
	Relationships *OrderCopyDataRelationships                        `json:"relationships,omitempty"`
}

// NewOrderCopyData instantiates a new OrderCopyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCopyData(type_ interface{}, attributes GETOrderCopiesOrderCopyId200ResponseDataAttributes) *OrderCopyData {
	this := OrderCopyData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderCopyDataWithDefaults instantiates a new OrderCopyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCopyDataWithDefaults() *OrderCopyData {
	this := OrderCopyData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderCopyData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCopyData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderCopyData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderCopyData) GetAttributes() GETOrderCopiesOrderCopyId200ResponseDataAttributes {
	if o == nil {
		var ret GETOrderCopiesOrderCopyId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderCopyData) GetAttributesOk() (*GETOrderCopiesOrderCopyId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderCopyData) SetAttributes(v GETOrderCopiesOrderCopyId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderCopyData) GetRelationships() OrderCopyDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret OrderCopyDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCopyData) GetRelationshipsOk() (*OrderCopyDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderCopyData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderCopyDataRelationships and assigns it to the Relationships field.
func (o *OrderCopyData) SetRelationships(v OrderCopyDataRelationships) {
	o.Relationships = &v
}

func (o OrderCopyData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCopyData) ToMap() (map[string]interface{}, error) {
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

type NullableOrderCopyData struct {
	value *OrderCopyData
	isSet bool
}

func (v NullableOrderCopyData) Get() *OrderCopyData {
	return v.value
}

func (v *NullableOrderCopyData) Set(val *OrderCopyData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCopyData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCopyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCopyData(val *OrderCopyData) *NullableOrderCopyData {
	return &NullableOrderCopyData{value: val, isSet: true}
}

func (v NullableOrderCopyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCopyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
