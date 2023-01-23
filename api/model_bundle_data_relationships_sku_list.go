/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BundleDataRelationshipsSkuList struct for BundleDataRelationshipsSkuList
type BundleDataRelationshipsSkuList struct {
	Data *BundleDataRelationshipsSkuListData `json:"data,omitempty"`
}

// NewBundleDataRelationshipsSkuList instantiates a new BundleDataRelationshipsSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDataRelationshipsSkuList() *BundleDataRelationshipsSkuList {
	this := BundleDataRelationshipsSkuList{}
	return &this
}

// NewBundleDataRelationshipsSkuListWithDefaults instantiates a new BundleDataRelationshipsSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataRelationshipsSkuListWithDefaults() *BundleDataRelationshipsSkuList {
	this := BundleDataRelationshipsSkuList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BundleDataRelationshipsSkuList) GetData() BundleDataRelationshipsSkuListData {
	if o == nil || o.Data == nil {
		var ret BundleDataRelationshipsSkuListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDataRelationshipsSkuList) GetDataOk() (*BundleDataRelationshipsSkuListData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BundleDataRelationshipsSkuList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BundleDataRelationshipsSkuListData and assigns it to the Data field.
func (o *BundleDataRelationshipsSkuList) SetData(v BundleDataRelationshipsSkuListData) {
	o.Data = &v
}

func (o BundleDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundleDataRelationshipsSkuList struct {
	value *BundleDataRelationshipsSkuList
	isSet bool
}

func (v NullableBundleDataRelationshipsSkuList) Get() *BundleDataRelationshipsSkuList {
	return v.value
}

func (v *NullableBundleDataRelationshipsSkuList) Set(val *BundleDataRelationshipsSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDataRelationshipsSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDataRelationshipsSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDataRelationshipsSkuList(val *BundleDataRelationshipsSkuList) *NullableBundleDataRelationshipsSkuList {
	return &NullableBundleDataRelationshipsSkuList{value: val, isSet: true}
}

func (v NullableBundleDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDataRelationshipsSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
