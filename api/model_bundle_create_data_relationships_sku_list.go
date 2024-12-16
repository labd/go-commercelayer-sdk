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

// checks if the BundleCreateDataRelationshipsSkuList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleCreateDataRelationshipsSkuList{}

// BundleCreateDataRelationshipsSkuList struct for BundleCreateDataRelationshipsSkuList
type BundleCreateDataRelationshipsSkuList struct {
	Data BundleDataRelationshipsSkuListData `json:"data"`
}

// NewBundleCreateDataRelationshipsSkuList instantiates a new BundleCreateDataRelationshipsSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreateDataRelationshipsSkuList(data BundleDataRelationshipsSkuListData) *BundleCreateDataRelationshipsSkuList {
	this := BundleCreateDataRelationshipsSkuList{}
	this.Data = data
	return &this
}

// NewBundleCreateDataRelationshipsSkuListWithDefaults instantiates a new BundleCreateDataRelationshipsSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateDataRelationshipsSkuListWithDefaults() *BundleCreateDataRelationshipsSkuList {
	this := BundleCreateDataRelationshipsSkuList{}
	return &this
}

// GetData returns the Data field value
func (o *BundleCreateDataRelationshipsSkuList) GetData() BundleDataRelationshipsSkuListData {
	if o == nil {
		var ret BundleDataRelationshipsSkuListData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationshipsSkuList) GetDataOk() (*BundleDataRelationshipsSkuListData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleCreateDataRelationshipsSkuList) SetData(v BundleDataRelationshipsSkuListData) {
	o.Data = v
}

func (o BundleCreateDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleCreateDataRelationshipsSkuList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBundleCreateDataRelationshipsSkuList struct {
	value *BundleCreateDataRelationshipsSkuList
	isSet bool
}

func (v NullableBundleCreateDataRelationshipsSkuList) Get() *BundleCreateDataRelationshipsSkuList {
	return v.value
}

func (v *NullableBundleCreateDataRelationshipsSkuList) Set(val *BundleCreateDataRelationshipsSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreateDataRelationshipsSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreateDataRelationshipsSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreateDataRelationshipsSkuList(val *BundleCreateDataRelationshipsSkuList) *NullableBundleCreateDataRelationshipsSkuList {
	return &NullableBundleCreateDataRelationshipsSkuList{value: val, isSet: true}
}

func (v NullableBundleCreateDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreateDataRelationshipsSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
