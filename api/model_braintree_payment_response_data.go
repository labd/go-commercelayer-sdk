/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BraintreePaymentResponseData struct for BraintreePaymentResponseData
type BraintreePaymentResponseData struct {
	// The resource's id
	Id *string `json:"id,omitempty"`
	// The resource's type
	Type          *string                                `json:"type,omitempty"`
	Links         *AddressResponseDataLinks              `json:"links,omitempty"`
	Attributes    *Attributes                            `json:"attributes,omitempty"`
	Relationships *AdyenPaymentResponseDataRelationships `json:"relationships,omitempty"`
}

// NewBraintreePaymentResponseData instantiates a new BraintreePaymentResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentResponseData() *BraintreePaymentResponseData {
	this := BraintreePaymentResponseData{}
	return &this
}

// NewBraintreePaymentResponseDataWithDefaults instantiates a new BraintreePaymentResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentResponseDataWithDefaults() *BraintreePaymentResponseData {
	this := BraintreePaymentResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BraintreePaymentResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BraintreePaymentResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BraintreePaymentResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BraintreePaymentResponseData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentResponseData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BraintreePaymentResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BraintreePaymentResponseData) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BraintreePaymentResponseData) GetLinks() AddressResponseDataLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentResponseData) GetLinksOk() (*AddressResponseDataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BraintreePaymentResponseData) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataLinks and assigns it to the Links field.
func (o *BraintreePaymentResponseData) SetLinks(v AddressResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BraintreePaymentResponseData) GetAttributes() Attributes {
	if o == nil || o.Attributes == nil {
		var ret Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentResponseData) GetAttributesOk() (*Attributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BraintreePaymentResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Attributes and assigns it to the Attributes field.
func (o *BraintreePaymentResponseData) SetAttributes(v Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreePaymentResponseData) GetRelationships() AdyenPaymentResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentResponseData) GetRelationshipsOk() (*AdyenPaymentResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreePaymentResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentResponseDataRelationships and assigns it to the Relationships field.
func (o *BraintreePaymentResponseData) SetRelationships(v AdyenPaymentResponseDataRelationships) {
	o.Relationships = &v
}

func (o BraintreePaymentResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableBraintreePaymentResponseData struct {
	value *BraintreePaymentResponseData
	isSet bool
}

func (v NullableBraintreePaymentResponseData) Get() *BraintreePaymentResponseData {
	return v.value
}

func (v *NullableBraintreePaymentResponseData) Set(val *BraintreePaymentResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentResponseData(val *BraintreePaymentResponseData) *NullableBraintreePaymentResponseData {
	return &NullableBraintreePaymentResponseData{value: val, isSet: true}
}

func (v NullableBraintreePaymentResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
