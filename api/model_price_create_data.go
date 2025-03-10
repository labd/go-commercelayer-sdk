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

// checks if the PriceCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceCreateData{}

// PriceCreateData struct for PriceCreateData
type PriceCreateData struct {
	// The resource's type
	Type          interface{}                         `json:"type"`
	Attributes    POSTPrices201ResponseDataAttributes `json:"attributes"`
	Relationships *PriceCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewPriceCreateData instantiates a new PriceCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceCreateData(type_ interface{}, attributes POSTPrices201ResponseDataAttributes) *PriceCreateData {
	this := PriceCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceCreateDataWithDefaults instantiates a new PriceCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceCreateDataWithDefaults() *PriceCreateData {
	this := PriceCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceCreateData) GetAttributes() POSTPrices201ResponseDataAttributes {
	if o == nil {
		var ret POSTPrices201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceCreateData) GetAttributesOk() (*POSTPrices201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceCreateData) SetAttributes(v POSTPrices201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceCreateData) GetRelationships() PriceCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PriceCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateData) GetRelationshipsOk() (*PriceCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceCreateDataRelationships and assigns it to the Relationships field.
func (o *PriceCreateData) SetRelationships(v PriceCreateDataRelationships) {
	o.Relationships = &v
}

func (o PriceCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceCreateData) ToMap() (map[string]interface{}, error) {
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

type NullablePriceCreateData struct {
	value *PriceCreateData
	isSet bool
}

func (v NullablePriceCreateData) Get() *PriceCreateData {
	return v.value
}

func (v *NullablePriceCreateData) Set(val *PriceCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCreateData(val *PriceCreateData) *NullablePriceCreateData {
	return &NullablePriceCreateData{value: val, isSet: true}
}

func (v NullablePriceCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
