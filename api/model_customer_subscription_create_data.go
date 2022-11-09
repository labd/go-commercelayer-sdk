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

// CustomerSubscriptionCreateData struct for CustomerSubscriptionCreateData
type CustomerSubscriptionCreateData struct {
	// The resource's type
	Type          string                                             `json:"type"`
	Attributes    POSTCustomerSubscriptions201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                             `json:"relationships,omitempty"`
}

// NewCustomerSubscriptionCreateData instantiates a new CustomerSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscriptionCreateData(type_ string, attributes POSTCustomerSubscriptions201ResponseDataAttributes) *CustomerSubscriptionCreateData {
	this := CustomerSubscriptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCustomerSubscriptionCreateDataWithDefaults instantiates a new CustomerSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionCreateDataWithDefaults() *CustomerSubscriptionCreateData {
	this := CustomerSubscriptionCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerSubscriptionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerSubscriptionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerSubscriptionCreateData) GetAttributes() POSTCustomerSubscriptions201ResponseDataAttributes {
	if o == nil {
		var ret POSTCustomerSubscriptions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionCreateData) GetAttributesOk() (*POSTCustomerSubscriptions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerSubscriptionCreateData) SetAttributes(v POSTCustomerSubscriptions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerSubscriptionCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSubscriptionCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerSubscriptionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CustomerSubscriptionCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CustomerSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSubscriptionCreateData struct {
	value *CustomerSubscriptionCreateData
	isSet bool
}

func (v NullableCustomerSubscriptionCreateData) Get() *CustomerSubscriptionCreateData {
	return v.value
}

func (v *NullableCustomerSubscriptionCreateData) Set(val *CustomerSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscriptionCreateData(val *CustomerSubscriptionCreateData) *NullableCustomerSubscriptionCreateData {
	return &NullableCustomerSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableCustomerSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
