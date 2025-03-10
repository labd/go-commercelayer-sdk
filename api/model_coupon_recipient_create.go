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

// checks if the CouponRecipientCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponRecipientCreate{}

// CouponRecipientCreate struct for CouponRecipientCreate
type CouponRecipientCreate struct {
	Data CouponRecipientCreateData `json:"data"`
}

// NewCouponRecipientCreate instantiates a new CouponRecipientCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientCreate(data CouponRecipientCreateData) *CouponRecipientCreate {
	this := CouponRecipientCreate{}
	this.Data = data
	return &this
}

// NewCouponRecipientCreateWithDefaults instantiates a new CouponRecipientCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientCreateWithDefaults() *CouponRecipientCreate {
	this := CouponRecipientCreate{}
	return &this
}

// GetData returns the Data field value
func (o *CouponRecipientCreate) GetData() CouponRecipientCreateData {
	if o == nil {
		var ret CouponRecipientCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientCreate) GetDataOk() (*CouponRecipientCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponRecipientCreate) SetData(v CouponRecipientCreateData) {
	o.Data = v
}

func (o CouponRecipientCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponRecipientCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCouponRecipientCreate struct {
	value *CouponRecipientCreate
	isSet bool
}

func (v NullableCouponRecipientCreate) Get() *CouponRecipientCreate {
	return v.value
}

func (v *NullableCouponRecipientCreate) Set(val *CouponRecipientCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientCreate(val *CouponRecipientCreate) *NullableCouponRecipientCreate {
	return &NullableCouponRecipientCreate{value: val, isSet: true}
}

func (v NullableCouponRecipientCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
