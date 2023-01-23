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

// BillingInfoValidationRuleCreateData struct for BillingInfoValidationRuleCreateData
type BillingInfoValidationRuleCreateData struct {
	// The resource's type
	Type          string                                            `json:"type"`
	Attributes    POSTAdyenPayments201ResponseDataAttributes        `json:"attributes"`
	Relationships *BillingInfoValidationRuleCreateDataRelationships `json:"relationships,omitempty"`
}

// NewBillingInfoValidationRuleCreateData instantiates a new BillingInfoValidationRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoValidationRuleCreateData(type_ string, attributes POSTAdyenPayments201ResponseDataAttributes) *BillingInfoValidationRuleCreateData {
	this := BillingInfoValidationRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBillingInfoValidationRuleCreateDataWithDefaults instantiates a new BillingInfoValidationRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoValidationRuleCreateDataWithDefaults() *BillingInfoValidationRuleCreateData {
	this := BillingInfoValidationRuleCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *BillingInfoValidationRuleCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BillingInfoValidationRuleCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BillingInfoValidationRuleCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BillingInfoValidationRuleCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BillingInfoValidationRuleCreateData) GetRelationships() BillingInfoValidationRuleCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret BillingInfoValidationRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfoValidationRuleCreateData) GetRelationshipsOk() (*BillingInfoValidationRuleCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BillingInfoValidationRuleCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BillingInfoValidationRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *BillingInfoValidationRuleCreateData) SetRelationships(v BillingInfoValidationRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o BillingInfoValidationRuleCreateData) MarshalJSON() ([]byte, error) {
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

type NullableBillingInfoValidationRuleCreateData struct {
	value *BillingInfoValidationRuleCreateData
	isSet bool
}

func (v NullableBillingInfoValidationRuleCreateData) Get() *BillingInfoValidationRuleCreateData {
	return v.value
}

func (v *NullableBillingInfoValidationRuleCreateData) Set(val *BillingInfoValidationRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoValidationRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoValidationRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoValidationRuleCreateData(val *BillingInfoValidationRuleCreateData) *NullableBillingInfoValidationRuleCreateData {
	return &NullableBillingInfoValidationRuleCreateData{value: val, isSet: true}
}

func (v NullableBillingInfoValidationRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoValidationRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
