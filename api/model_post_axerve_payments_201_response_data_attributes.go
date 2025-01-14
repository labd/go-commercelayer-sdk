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

// checks if the POSTAxervePayments201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAxervePayments201ResponseDataAttributes{}

// POSTAxervePayments201ResponseDataAttributes struct for POSTAxervePayments201ResponseDataAttributes
type POSTAxervePayments201ResponseDataAttributes struct {
	// The URL where the payer is redirected after they approve the payment.
	ReturnUrl interface{} `json:"return_url"`
	// The IP adress of the client creating the payment.
	ClientIp interface{} `json:"client_ip,omitempty"`
	// The details of the buyer creating the payment.
	BuyerDetails interface{} `json:"buyer_details,omitempty"`
	// Requires the creation of a token to represent this payment, mandatory to use customer's wallet and order subscriptions.
	RequestToken interface{} `json:"request_token,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTAxervePayments201ResponseDataAttributes instantiates a new POSTAxervePayments201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAxervePayments201ResponseDataAttributes(returnUrl interface{}) *POSTAxervePayments201ResponseDataAttributes {
	this := POSTAxervePayments201ResponseDataAttributes{}
	this.ReturnUrl = returnUrl
	return &this
}

// NewPOSTAxervePayments201ResponseDataAttributesWithDefaults instantiates a new POSTAxervePayments201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAxervePayments201ResponseDataAttributesWithDefaults() *POSTAxervePayments201ResponseDataAttributes {
	this := POSTAxervePayments201ResponseDataAttributes{}
	return &this
}

// GetReturnUrl returns the ReturnUrl field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetReturnUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return &o.ReturnUrl, true
}

// SetReturnUrl sets field value
func (o *POSTAxervePayments201ResponseDataAttributes) SetReturnUrl(v interface{}) {
	o.ReturnUrl = v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetClientIp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetClientIpOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientIp) {
		return nil, false
	}
	return &o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasClientIp() bool {
	if o != nil && IsNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given interface{} and assigns it to the ClientIp field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetClientIp(v interface{}) {
	o.ClientIp = v
}

// GetBuyerDetails returns the BuyerDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetBuyerDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BuyerDetails
}

// GetBuyerDetailsOk returns a tuple with the BuyerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetBuyerDetailsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BuyerDetails) {
		return nil, false
	}
	return &o.BuyerDetails, true
}

// HasBuyerDetails returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasBuyerDetails() bool {
	if o != nil && IsNil(o.BuyerDetails) {
		return true
	}

	return false
}

// SetBuyerDetails gets a reference to the given interface{} and assigns it to the BuyerDetails field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetBuyerDetails(v interface{}) {
	o.BuyerDetails = v
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetRequestToken() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequestToken
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetRequestTokenOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequestToken) {
		return nil, false
	}
	return &o.RequestToken, true
}

// HasRequestToken returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasRequestToken() bool {
	if o != nil && IsNil(o.RequestToken) {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given interface{} and assigns it to the RequestToken field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetRequestToken(v interface{}) {
	o.RequestToken = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxervePayments201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxervePayments201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTAxervePayments201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTAxervePayments201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTAxervePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAxervePayments201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReturnUrl != nil {
		toSerialize["return_url"] = o.ReturnUrl
	}
	if o.ClientIp != nil {
		toSerialize["client_ip"] = o.ClientIp
	}
	if o.BuyerDetails != nil {
		toSerialize["buyer_details"] = o.BuyerDetails
	}
	if o.RequestToken != nil {
		toSerialize["request_token"] = o.RequestToken
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTAxervePayments201ResponseDataAttributes struct {
	value *POSTAxervePayments201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTAxervePayments201ResponseDataAttributes) Get() *POSTAxervePayments201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTAxervePayments201ResponseDataAttributes) Set(val *POSTAxervePayments201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAxervePayments201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAxervePayments201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAxervePayments201ResponseDataAttributes(val *POSTAxervePayments201ResponseDataAttributes) *NullablePOSTAxervePayments201ResponseDataAttributes {
	return &NullablePOSTAxervePayments201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTAxervePayments201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAxervePayments201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
