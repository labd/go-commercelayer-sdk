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

// checks if the ReturnLineItemCreateDataRelationshipsReturn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLineItemCreateDataRelationshipsReturn{}

// ReturnLineItemCreateDataRelationshipsReturn struct for ReturnLineItemCreateDataRelationshipsReturn
type ReturnLineItemCreateDataRelationshipsReturn struct {
	Data CaptureDataRelationshipsReturnData `json:"data"`
}

// NewReturnLineItemCreateDataRelationshipsReturn instantiates a new ReturnLineItemCreateDataRelationshipsReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemCreateDataRelationshipsReturn(data CaptureDataRelationshipsReturnData) *ReturnLineItemCreateDataRelationshipsReturn {
	this := ReturnLineItemCreateDataRelationshipsReturn{}
	this.Data = data
	return &this
}

// NewReturnLineItemCreateDataRelationshipsReturnWithDefaults instantiates a new ReturnLineItemCreateDataRelationshipsReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemCreateDataRelationshipsReturnWithDefaults() *ReturnLineItemCreateDataRelationshipsReturn {
	this := ReturnLineItemCreateDataRelationshipsReturn{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnLineItemCreateDataRelationshipsReturn) GetData() CaptureDataRelationshipsReturnData {
	if o == nil {
		var ret CaptureDataRelationshipsReturnData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreateDataRelationshipsReturn) GetDataOk() (*CaptureDataRelationshipsReturnData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnLineItemCreateDataRelationshipsReturn) SetData(v CaptureDataRelationshipsReturnData) {
	o.Data = v
}

func (o ReturnLineItemCreateDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLineItemCreateDataRelationshipsReturn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableReturnLineItemCreateDataRelationshipsReturn struct {
	value *ReturnLineItemCreateDataRelationshipsReturn
	isSet bool
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) Get() *ReturnLineItemCreateDataRelationshipsReturn {
	return v.value
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) Set(val *ReturnLineItemCreateDataRelationshipsReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemCreateDataRelationshipsReturn(val *ReturnLineItemCreateDataRelationshipsReturn) *NullableReturnLineItemCreateDataRelationshipsReturn {
	return &NullableReturnLineItemCreateDataRelationshipsReturn{value: val, isSet: true}
}

func (v NullableReturnLineItemCreateDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemCreateDataRelationshipsReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
