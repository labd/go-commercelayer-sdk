/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCarrierAccounts200ResponseDataInnerRelationships struct for GETCarrierAccounts200ResponseDataInnerRelationships
type GETCarrierAccounts200ResponseDataInnerRelationships struct {
	Market      *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments       `json:"attachments,omitempty"`
}

// NewGETCarrierAccounts200ResponseDataInnerRelationships instantiates a new GETCarrierAccounts200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCarrierAccounts200ResponseDataInnerRelationships() *GETCarrierAccounts200ResponseDataInnerRelationships {
	this := GETCarrierAccounts200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCarrierAccounts200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCarrierAccounts200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCarrierAccounts200ResponseDataInnerRelationshipsWithDefaults() *GETCarrierAccounts200ResponseDataInnerRelationships {
	this := GETCarrierAccounts200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETCarrierAccounts200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETCarrierAccounts200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETCarrierAccounts200ResponseDataInnerRelationships struct {
	value *GETCarrierAccounts200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCarrierAccounts200ResponseDataInnerRelationships) Get() *GETCarrierAccounts200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCarrierAccounts200ResponseDataInnerRelationships) Set(val *GETCarrierAccounts200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCarrierAccounts200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCarrierAccounts200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCarrierAccounts200ResponseDataInnerRelationships(val *GETCarrierAccounts200ResponseDataInnerRelationships) *NullableGETCarrierAccounts200ResponseDataInnerRelationships {
	return &NullableGETCarrierAccounts200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCarrierAccounts200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCarrierAccounts200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
