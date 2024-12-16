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

// checks if the POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData{}

// POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData struct for POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData
type POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData{}
	return &this
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionDataWithDefaults instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionDataWithDefaults() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData struct {
	value *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData
	isSet bool
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) Get() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData {
	return v.value
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) Set(val *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData(val *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData {
	return &NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData{value: val, isSet: true}
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
