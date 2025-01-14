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

// checks if the OrderDataRelationshipsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataRelationshipsLinks{}

// OrderDataRelationshipsLinks struct for OrderDataRelationshipsLinks
type OrderDataRelationshipsLinks struct {
	Data *OrderDataRelationshipsLinksData `json:"data,omitempty"`
}

// NewOrderDataRelationshipsLinks instantiates a new OrderDataRelationshipsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsLinks() *OrderDataRelationshipsLinks {
	this := OrderDataRelationshipsLinks{}
	return &this
}

// NewOrderDataRelationshipsLinksWithDefaults instantiates a new OrderDataRelationshipsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsLinksWithDefaults() *OrderDataRelationshipsLinks {
	this := OrderDataRelationshipsLinks{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderDataRelationshipsLinks) GetData() OrderDataRelationshipsLinksData {
	if o == nil || IsNil(o.Data) {
		var ret OrderDataRelationshipsLinksData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsLinks) GetDataOk() (*OrderDataRelationshipsLinksData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderDataRelationshipsLinks) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given OrderDataRelationshipsLinksData and assigns it to the Data field.
func (o *OrderDataRelationshipsLinks) SetData(v OrderDataRelationshipsLinksData) {
	o.Data = &v
}

func (o OrderDataRelationshipsLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataRelationshipsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderDataRelationshipsLinks struct {
	value *OrderDataRelationshipsLinks
	isSet bool
}

func (v NullableOrderDataRelationshipsLinks) Get() *OrderDataRelationshipsLinks {
	return v.value
}

func (v *NullableOrderDataRelationshipsLinks) Set(val *OrderDataRelationshipsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsLinks(val *OrderDataRelationshipsLinks) *NullableOrderDataRelationshipsLinks {
	return &NullableOrderDataRelationshipsLinks{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
