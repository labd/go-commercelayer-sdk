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

// KlarnaGatewayCreateData struct for KlarnaGatewayCreateData
type KlarnaGatewayCreateData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    POSTKlarnaGateways201ResponseDataAttributes `json:"attributes"`
	Relationships *KlarnaGatewayCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewKlarnaGatewayCreateData instantiates a new KlarnaGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGatewayCreateData(type_ string, attributes POSTKlarnaGateways201ResponseDataAttributes) *KlarnaGatewayCreateData {
	this := KlarnaGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewKlarnaGatewayCreateDataWithDefaults instantiates a new KlarnaGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayCreateDataWithDefaults() *KlarnaGatewayCreateData {
	this := KlarnaGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *KlarnaGatewayCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KlarnaGatewayCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *KlarnaGatewayCreateData) GetAttributes() POSTKlarnaGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTKlarnaGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayCreateData) GetAttributesOk() (*POSTKlarnaGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *KlarnaGatewayCreateData) SetAttributes(v POSTKlarnaGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *KlarnaGatewayCreateData) GetRelationships() KlarnaGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret KlarnaGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaGatewayCreateData) GetRelationshipsOk() (*KlarnaGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *KlarnaGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given KlarnaGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *KlarnaGatewayCreateData) SetRelationships(v KlarnaGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o KlarnaGatewayCreateData) MarshalJSON() ([]byte, error) {
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

type NullableKlarnaGatewayCreateData struct {
	value *KlarnaGatewayCreateData
	isSet bool
}

func (v NullableKlarnaGatewayCreateData) Get() *KlarnaGatewayCreateData {
	return v.value
}

func (v *NullableKlarnaGatewayCreateData) Set(val *KlarnaGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGatewayCreateData(val *KlarnaGatewayCreateData) *NullableKlarnaGatewayCreateData {
	return &NullableKlarnaGatewayCreateData{value: val, isSet: true}
}

func (v NullableKlarnaGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
