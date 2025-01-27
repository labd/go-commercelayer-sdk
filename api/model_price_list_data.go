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

// checks if the PriceListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceListData{}

// PriceListData struct for PriceListData
type PriceListData struct {
	// The resource's type
	Type          interface{}                                       `json:"type"`
	Attributes    GETPriceListsPriceListId200ResponseDataAttributes `json:"attributes"`
	Relationships *PriceListDataRelationships                       `json:"relationships,omitempty"`
}

// NewPriceListData instantiates a new PriceListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListData(type_ interface{}, attributes GETPriceListsPriceListId200ResponseDataAttributes) *PriceListData {
	this := PriceListData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceListDataWithDefaults instantiates a new PriceListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListDataWithDefaults() *PriceListData {
	this := PriceListData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceListData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceListData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceListData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceListData) GetAttributes() GETPriceListsPriceListId200ResponseDataAttributes {
	if o == nil {
		var ret GETPriceListsPriceListId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceListData) GetAttributesOk() (*GETPriceListsPriceListId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceListData) SetAttributes(v GETPriceListsPriceListId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceListData) GetRelationships() PriceListDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PriceListDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListData) GetRelationshipsOk() (*PriceListDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceListData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceListDataRelationships and assigns it to the Relationships field.
func (o *PriceListData) SetRelationships(v PriceListDataRelationships) {
	o.Relationships = &v
}

func (o PriceListData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceListData) ToMap() (map[string]interface{}, error) {
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

type NullablePriceListData struct {
	value *PriceListData
	isSet bool
}

func (v NullablePriceListData) Get() *PriceListData {
	return v.value
}

func (v *NullablePriceListData) Set(val *PriceListData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListData(val *PriceListData) *NullablePriceListData {
	return &NullablePriceListData{value: val, isSet: true}
}

func (v NullablePriceListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
