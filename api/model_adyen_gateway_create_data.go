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

// AdyenGatewayCreateData struct for AdyenGatewayCreateData
type AdyenGatewayCreateData struct {
	// The resource's type
	Type          string                               `json:"type"`
	Attributes    AdyenGatewayCreateDataAttributes     `json:"attributes"`
	Relationships *AdyenGatewayCreateDataRelationships `json:"relationships,omitempty"`
}

// NewAdyenGatewayCreateData instantiates a new AdyenGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayCreateData(type_ string, attributes AdyenGatewayCreateDataAttributes) *AdyenGatewayCreateData {
	this := AdyenGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdyenGatewayCreateDataWithDefaults instantiates a new AdyenGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayCreateDataWithDefaults() *AdyenGatewayCreateData {
	this := AdyenGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *AdyenGatewayCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenGatewayCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenGatewayCreateData) GetAttributes() AdyenGatewayCreateDataAttributes {
	if o == nil {
		var ret AdyenGatewayCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateData) GetAttributesOk() (*AdyenGatewayCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenGatewayCreateData) SetAttributes(v AdyenGatewayCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenGatewayCreateData) GetRelationships() AdyenGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateData) GetRelationshipsOk() (*AdyenGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *AdyenGatewayCreateData) SetRelationships(v AdyenGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o AdyenGatewayCreateData) MarshalJSON() ([]byte, error) {
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

type NullableAdyenGatewayCreateData struct {
	value *AdyenGatewayCreateData
	isSet bool
}

func (v NullableAdyenGatewayCreateData) Get() *AdyenGatewayCreateData {
	return v.value
}

func (v *NullableAdyenGatewayCreateData) Set(val *AdyenGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayCreateData(val *AdyenGatewayCreateData) *NullableAdyenGatewayCreateData {
	return &NullableAdyenGatewayCreateData{value: val, isSet: true}
}

func (v NullableAdyenGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
