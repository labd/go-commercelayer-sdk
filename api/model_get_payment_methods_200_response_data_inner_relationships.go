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

// GETPaymentMethods200ResponseDataInnerRelationships struct for GETPaymentMethods200ResponseDataInnerRelationships
type GETPaymentMethods200ResponseDataInnerRelationships struct {
	Market         *GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket `json:"market,omitempty"`
	PaymentGateway *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway      `json:"payment_gateway,omitempty"`
	Attachments    *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments       `json:"attachments,omitempty"`
}

// NewGETPaymentMethods200ResponseDataInnerRelationships instantiates a new GETPaymentMethods200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPaymentMethods200ResponseDataInnerRelationships() *GETPaymentMethods200ResponseDataInnerRelationships {
	this := GETPaymentMethods200ResponseDataInnerRelationships{}
	return &this
}

// NewGETPaymentMethods200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETPaymentMethods200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPaymentMethods200ResponseDataInnerRelationshipsWithDefaults() *GETPaymentMethods200ResponseDataInnerRelationships {
	this := GETPaymentMethods200ResponseDataInnerRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetMarket() GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetMarketOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket and assigns it to the Market field.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) SetMarket(v GETBillingInfoValidationRules200ResponseDataInnerRelationshipsMarket) {
	o.Market = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetPaymentGateway() GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetPaymentGatewayOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) SetPaymentGateway(v GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETPaymentMethods200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETPaymentMethods200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PaymentGateway != nil {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETPaymentMethods200ResponseDataInnerRelationships struct {
	value *GETPaymentMethods200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETPaymentMethods200ResponseDataInnerRelationships) Get() *GETPaymentMethods200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETPaymentMethods200ResponseDataInnerRelationships) Set(val *GETPaymentMethods200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPaymentMethods200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPaymentMethods200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPaymentMethods200ResponseDataInnerRelationships(val *GETPaymentMethods200ResponseDataInnerRelationships) *NullableGETPaymentMethods200ResponseDataInnerRelationships {
	return &NullableGETPaymentMethods200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETPaymentMethods200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPaymentMethods200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
