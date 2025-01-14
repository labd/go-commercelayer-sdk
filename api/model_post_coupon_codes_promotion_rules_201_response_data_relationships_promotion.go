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

// checks if the POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion{}

// POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion struct for POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion
type POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData `json:"data,omitempty"`
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion{}
	return &this
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionWithDefaults instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionWithDefaults() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) GetData() POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) GetDataOk() (*POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData and assigns it to the Data field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) SetData(v POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotionData) {
	o.Data = &v
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion struct {
	value *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion
	isSet bool
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) Get() *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	return v.value
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) Set(val *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion(val *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	return &NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion{value: val, isSet: true}
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
