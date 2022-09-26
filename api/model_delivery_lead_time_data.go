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

// DeliveryLeadTimeData struct for DeliveryLeadTimeData
type DeliveryLeadTimeData struct {
	// The resource's type
	Type          string                             `json:"type"`
	Attributes    DeliveryLeadTimeDataAttributes     `json:"attributes"`
	Relationships *DeliveryLeadTimeDataRelationships `json:"relationships,omitempty"`
}

// NewDeliveryLeadTimeData instantiates a new DeliveryLeadTimeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeData(type_ string, attributes DeliveryLeadTimeDataAttributes) *DeliveryLeadTimeData {
	this := DeliveryLeadTimeData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDeliveryLeadTimeDataWithDefaults instantiates a new DeliveryLeadTimeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeDataWithDefaults() *DeliveryLeadTimeData {
	this := DeliveryLeadTimeData{}
	return &this
}

// GetType returns the Type field value
func (o *DeliveryLeadTimeData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeliveryLeadTimeData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DeliveryLeadTimeData) GetAttributes() DeliveryLeadTimeDataAttributes {
	if o == nil {
		var ret DeliveryLeadTimeDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeData) GetAttributesOk() (*DeliveryLeadTimeDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DeliveryLeadTimeData) SetAttributes(v DeliveryLeadTimeDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DeliveryLeadTimeData) GetRelationships() DeliveryLeadTimeDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret DeliveryLeadTimeDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeData) GetRelationshipsOk() (*DeliveryLeadTimeDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DeliveryLeadTimeData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given DeliveryLeadTimeDataRelationships and assigns it to the Relationships field.
func (o *DeliveryLeadTimeData) SetRelationships(v DeliveryLeadTimeDataRelationships) {
	o.Relationships = &v
}

func (o DeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
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

type NullableDeliveryLeadTimeData struct {
	value *DeliveryLeadTimeData
	isSet bool
}

func (v NullableDeliveryLeadTimeData) Get() *DeliveryLeadTimeData {
	return v.value
}

func (v *NullableDeliveryLeadTimeData) Set(val *DeliveryLeadTimeData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeData(val *DeliveryLeadTimeData) *NullableDeliveryLeadTimeData {
	return &NullableDeliveryLeadTimeData{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
