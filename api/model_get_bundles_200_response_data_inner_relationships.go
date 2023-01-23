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

// GETBundles200ResponseDataInnerRelationships struct for GETBundles200ResponseDataInnerRelationships
type GETBundles200ResponseDataInnerRelationships struct {
	Market      *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	SkuList     *GETBundles200ResponseDataInnerRelationshipsSkuList                   `json:"sku_list,omitempty"`
	Skus        *GETBundles200ResponseDataInnerRelationshipsSkus                      `json:"skus,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments       `json:"attachments,omitempty"`
}

// NewGETBundles200ResponseDataInnerRelationships instantiates a new GETBundles200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundles200ResponseDataInnerRelationships() *GETBundles200ResponseDataInnerRelationships {
	this := GETBundles200ResponseDataInnerRelationships{}
	return &this
}

// NewGETBundles200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETBundles200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundles200ResponseDataInnerRelationshipsWithDefaults() *GETBundles200ResponseDataInnerRelationships {
	this := GETBundles200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETBundles200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkuList and assigns it to the SkuList field.
func (o *GETBundles200ResponseDataInnerRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationships) GetSkus() GETBundles200ResponseDataInnerRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationships) GetSkusOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkus and assigns it to the Skus field.
func (o *GETBundles200ResponseDataInnerRelationships) SetSkus(v GETBundles200ResponseDataInnerRelationshipsSkus) {
	o.Skus = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETBundles200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETBundles200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETBundles200ResponseDataInnerRelationships struct {
	value *GETBundles200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETBundles200ResponseDataInnerRelationships) Get() *GETBundles200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETBundles200ResponseDataInnerRelationships) Set(val *GETBundles200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundles200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundles200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundles200ResponseDataInnerRelationships(val *GETBundles200ResponseDataInnerRelationships) *NullableGETBundles200ResponseDataInnerRelationships {
	return &NullableGETBundles200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETBundles200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundles200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
