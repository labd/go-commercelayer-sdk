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

// DeliveryLeadTimeCreateData struct for DeliveryLeadTimeCreateData
type DeliveryLeadTimeCreateData struct {
	// The resource's type
	Type          string                                   `json:"type"`
	Attributes    DeliveryLeadTimeCreateDataAttributes     `json:"attributes"`
	Relationships *DeliveryLeadTimeCreateDataRelationships `json:"relationships,omitempty"`
}

// NewDeliveryLeadTimeCreateData instantiates a new DeliveryLeadTimeCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeCreateData(type_ string, attributes DeliveryLeadTimeCreateDataAttributes) *DeliveryLeadTimeCreateData {
	this := DeliveryLeadTimeCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDeliveryLeadTimeCreateDataWithDefaults instantiates a new DeliveryLeadTimeCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeCreateDataWithDefaults() *DeliveryLeadTimeCreateData {
	this := DeliveryLeadTimeCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *DeliveryLeadTimeCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DeliveryLeadTimeCreateData) GetAttributes() DeliveryLeadTimeCreateDataAttributes {
	if o == nil {
		var ret DeliveryLeadTimeCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateData) GetAttributesOk() (*DeliveryLeadTimeCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DeliveryLeadTimeCreateData) SetAttributes(v DeliveryLeadTimeCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DeliveryLeadTimeCreateData) GetRelationships() DeliveryLeadTimeCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret DeliveryLeadTimeCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeCreateData) GetRelationshipsOk() (*DeliveryLeadTimeCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DeliveryLeadTimeCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DeliveryLeadTimeCreateDataRelationships and assigns it to the Relationships field.
func (o *DeliveryLeadTimeCreateData) SetRelationships(v DeliveryLeadTimeCreateDataRelationships) {
	o.Relationships = &v
}

func (o DeliveryLeadTimeCreateData) MarshalJSON() ([]byte, error) {
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

type NullableDeliveryLeadTimeCreateData struct {
	value *DeliveryLeadTimeCreateData
	isSet bool
}

func (v NullableDeliveryLeadTimeCreateData) Get() *DeliveryLeadTimeCreateData {
	return v.value
}

func (v *NullableDeliveryLeadTimeCreateData) Set(val *DeliveryLeadTimeCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeCreateData(val *DeliveryLeadTimeCreateData) *NullableDeliveryLeadTimeCreateData {
	return &NullableDeliveryLeadTimeCreateData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
