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

// checks if the PriceListCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListCreateData{}

// PriceListCreateData struct for PriceListCreateData
type PriceListCreateData struct {
	// The resource's type
	Type          interface{}                             `json:"type"`
	Attributes    POSTPriceLists201ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                             `json:"relationships,omitempty"`
}

// NewPriceListCreateData instantiates a new PriceListCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListCreateData(type_ interface{}, attributes POSTPriceLists201ResponseDataAttributes) *PriceListCreateData {
	this := PriceListCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceListCreateDataWithDefaults instantiates a new PriceListCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListCreateDataWithDefaults() *PriceListCreateData {
	this := PriceListCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceListCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceListCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceListCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceListCreateData) GetAttributes() POSTPriceLists201ResponseDataAttributes {
	if o == nil {
		var ret POSTPriceLists201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceListCreateData) GetAttributesOk() (*POSTPriceLists201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceListCreateData) SetAttributes(v POSTPriceLists201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceListCreateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceListCreateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceListCreateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *PriceListCreateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o PriceListCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePriceListCreateData struct {
	value *PriceListCreateData
	isSet bool
}

func (v NullablePriceListCreateData) Get() *PriceListCreateData {
	return v.value
}

func (v *NullablePriceListCreateData) Set(val *PriceListCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListCreateData(val *PriceListCreateData) *NullablePriceListCreateData {
	return &NullablePriceListCreateData{value: val, isSet: true}
}

func (v NullablePriceListCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
