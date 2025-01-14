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

// checks if the SkuListItemCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListItemCreateDataRelationships{}

// SkuListItemCreateDataRelationships struct for SkuListItemCreateDataRelationships
type SkuListItemCreateDataRelationships struct {
	SkuList BundleCreateDataRelationshipsSkuList          `json:"sku_list"`
	Sku     InStockSubscriptionCreateDataRelationshipsSku `json:"sku"`
}

// NewSkuListItemCreateDataRelationships instantiates a new SkuListItemCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListItemCreateDataRelationships(skuList BundleCreateDataRelationshipsSkuList, sku InStockSubscriptionCreateDataRelationshipsSku) *SkuListItemCreateDataRelationships {
	this := SkuListItemCreateDataRelationships{}
	this.SkuList = skuList
	this.Sku = sku
	return &this
}

// NewSkuListItemCreateDataRelationshipsWithDefaults instantiates a new SkuListItemCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListItemCreateDataRelationshipsWithDefaults() *SkuListItemCreateDataRelationships {
	this := SkuListItemCreateDataRelationships{}
	return &this
}

// GetSkuList returns the SkuList field value
func (o *SkuListItemCreateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList {
	if o == nil {
		var ret BundleCreateDataRelationshipsSkuList
		return ret
	}

	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value
// and a boolean to check if the value has been set.
func (o *SkuListItemCreateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuList, true
}

// SetSkuList sets field value
func (o *SkuListItemCreateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList) {
	o.SkuList = v
}

// GetSku returns the Sku field value
func (o *SkuListItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *SkuListItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *SkuListItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = v
}

func (o SkuListItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListItemCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku_list"] = o.SkuList
	toSerialize["sku"] = o.Sku
	return toSerialize, nil
}

type NullableSkuListItemCreateDataRelationships struct {
	value *SkuListItemCreateDataRelationships
	isSet bool
}

func (v NullableSkuListItemCreateDataRelationships) Get() *SkuListItemCreateDataRelationships {
	return v.value
}

func (v *NullableSkuListItemCreateDataRelationships) Set(val *SkuListItemCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListItemCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListItemCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListItemCreateDataRelationships(val *SkuListItemCreateDataRelationships) *NullableSkuListItemCreateDataRelationships {
	return &NullableSkuListItemCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListItemCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
