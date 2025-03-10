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

// checks if the CheckoutComGatewayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutComGatewayData{}

// CheckoutComGatewayData struct for CheckoutComGatewayData
type CheckoutComGatewayData struct {
	// The resource's type
	Type          interface{}                                                         `json:"type"`
	Attributes    GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships *CheckoutComGatewayDataRelationships                                `json:"relationships,omitempty"`
}

// NewCheckoutComGatewayData instantiates a new CheckoutComGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComGatewayData(type_ interface{}, attributes GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) *CheckoutComGatewayData {
	this := CheckoutComGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCheckoutComGatewayDataWithDefaults instantiates a new CheckoutComGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComGatewayDataWithDefaults() *CheckoutComGatewayData {
	this := CheckoutComGatewayData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CheckoutComGatewayData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutComGatewayData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutComGatewayData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CheckoutComGatewayData) GetAttributes() GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayData) GetAttributesOk() (*GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CheckoutComGatewayData) SetAttributes(v GETCheckoutComGatewaysCheckoutComGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CheckoutComGatewayData) GetRelationships() CheckoutComGatewayDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CheckoutComGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutComGatewayData) GetRelationshipsOk() (*CheckoutComGatewayDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CheckoutComGatewayData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CheckoutComGatewayDataRelationships and assigns it to the Relationships field.
func (o *CheckoutComGatewayData) SetRelationships(v CheckoutComGatewayDataRelationships) {
	o.Relationships = &v
}

func (o CheckoutComGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutComGatewayData) ToMap() (map[string]interface{}, error) {
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

type NullableCheckoutComGatewayData struct {
	value *CheckoutComGatewayData
	isSet bool
}

func (v NullableCheckoutComGatewayData) Get() *CheckoutComGatewayData {
	return v.value
}

func (v *NullableCheckoutComGatewayData) Set(val *CheckoutComGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComGatewayData(val *CheckoutComGatewayData) *NullableCheckoutComGatewayData {
	return &NullableCheckoutComGatewayData{value: val, isSet: true}
}

func (v NullableCheckoutComGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
