/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InStockSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStockSubscriptionData{}

// InStockSubscriptionData struct for InStockSubscriptionData
type InStockSubscriptionData struct {
	// The resource's type
	Type          interface{}                                                           `json:"type"`
	Attributes    GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes `json:"attributes"`
	Relationships *InStockSubscriptionDataRelationships                                 `json:"relationships,omitempty"`
}

// NewInStockSubscriptionData instantiates a new InStockSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionData(type_ interface{}, attributes GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) *InStockSubscriptionData {
	this := InStockSubscriptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInStockSubscriptionDataWithDefaults instantiates a new InStockSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionDataWithDefaults() *InStockSubscriptionData {
	this := InStockSubscriptionData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InStockSubscriptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InStockSubscriptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InStockSubscriptionData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InStockSubscriptionData) GetAttributes() GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	if o == nil {
		var ret GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionData) GetAttributesOk() (*GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InStockSubscriptionData) SetAttributes(v GETInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionData) GetRelationships() InStockSubscriptionDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InStockSubscriptionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionData) GetRelationshipsOk() (*InStockSubscriptionDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionData) SetRelationships(v InStockSubscriptionDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InStockSubscriptionData) ToMap() (map[string]interface{}, error) {
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

type NullableInStockSubscriptionData struct {
	value *InStockSubscriptionData
	isSet bool
}

func (v NullableInStockSubscriptionData) Get() *InStockSubscriptionData {
	return v.value
}

func (v *NullableInStockSubscriptionData) Set(val *InStockSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionData(val *InStockSubscriptionData) *NullableInStockSubscriptionData {
	return &NullableInStockSubscriptionData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
