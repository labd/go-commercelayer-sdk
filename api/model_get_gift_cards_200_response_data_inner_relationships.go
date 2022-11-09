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

// GETGiftCards200ResponseDataInnerRelationships struct for GETGiftCards200ResponseDataInnerRelationships
type GETGiftCards200ResponseDataInnerRelationships struct {
	Market            *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	GiftCardRecipient *GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient       `json:"gift_card_recipient,omitempty"`
	Attachments       *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments       `json:"attachments,omitempty"`
	Events            *GETCustomerAddresses200ResponseDataInnerRelationshipsEvents          `json:"events,omitempty"`
}

// NewGETGiftCards200ResponseDataInnerRelationships instantiates a new GETGiftCards200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGiftCards200ResponseDataInnerRelationships() *GETGiftCards200ResponseDataInnerRelationships {
	this := GETGiftCards200ResponseDataInnerRelationships{}
	return &this
}

// NewGETGiftCards200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETGiftCards200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGiftCards200ResponseDataInnerRelationshipsWithDefaults() *GETGiftCards200ResponseDataInnerRelationships {
	this := GETGiftCards200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETGiftCards200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetGiftCardRecipient() GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient {
	if o == nil || o.GiftCardRecipient == nil {
		var ret GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetGiftCardRecipientOk() (*GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient, bool) {
	if o == nil || o.GiftCardRecipient == nil {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) HasGiftCardRecipient() bool {
	if o != nil && o.GiftCardRecipient != nil {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *GETGiftCards200ResponseDataInnerRelationships) SetGiftCardRecipient(v GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETGiftCards200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetEvents() GETCustomerAddresses200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) GetEventsOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETGiftCards200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETGiftCards200ResponseDataInnerRelationships) SetEvents(v GETCustomerAddresses200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETGiftCards200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.GiftCardRecipient != nil {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETGiftCards200ResponseDataInnerRelationships struct {
	value *GETGiftCards200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETGiftCards200ResponseDataInnerRelationships) Get() *GETGiftCards200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETGiftCards200ResponseDataInnerRelationships) Set(val *GETGiftCards200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGiftCards200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGiftCards200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGiftCards200ResponseDataInnerRelationships(val *GETGiftCards200ResponseDataInnerRelationships) *NullableGETGiftCards200ResponseDataInnerRelationships {
	return &NullableGETGiftCards200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETGiftCards200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGiftCards200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
