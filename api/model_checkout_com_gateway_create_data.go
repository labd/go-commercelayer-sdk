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

// checks if the CheckoutComGatewayCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayCreateData{}

// CheckoutComGatewayCreateData struct for CheckoutComGatewayCreateData
type CheckoutComGatewayCreateData struct {
	// The resource's type
	Type          interface{}                                      `json:"type"`
	Attributes    POSTCheckoutComGateways201ResponseDataAttributes `json:"attributes"`
	Relationships *CheckoutComGatewayCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewCheckoutComGatewayCreateData instantiates a new CheckoutComGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayCreateData(type_ interface{}, attributes POSTCheckoutComGateways201ResponseDataAttributes) *CheckoutComGatewayCreateData {
	this := CheckoutComGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCheckoutComGatewayCreateDataWithDefaults instantiates a new CheckoutComGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayCreateDataWithDefaults() *CheckoutComGatewayCreateData {
	this := CheckoutComGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckoutComGatewayCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComGatewayCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComGatewayCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComGatewayCreateData) GetAttributes() POSTCheckoutComGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTCheckoutComGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayCreateData) GetAttributesOk() (*POSTCheckoutComGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComGatewayCreateData) SetAttributes(v POSTCheckoutComGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComGatewayCreateData) GetRelationships() CheckoutComGatewayCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CheckoutComGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayCreateData) GetRelationshipsOk() (*CheckoutComGatewayCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComGatewayCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CheckoutComGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComGatewayCreateData) SetRelationships(v CheckoutComGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComGatewayCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCheckoutComGatewayCreateData struct {
	value *CheckoutComGatewayCreateData
	isSet bool
}

func (v NullableCheckoutComGatewayCreateData) Get() *CheckoutComGatewayCreateData {
	return v.value
}

func (v *NullableCheckoutComGatewayCreateData) Set(val *CheckoutComGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayCreateData(val *CheckoutComGatewayCreateData) *NullableCheckoutComGatewayCreateData {
	return &NullableCheckoutComGatewayCreateData{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
