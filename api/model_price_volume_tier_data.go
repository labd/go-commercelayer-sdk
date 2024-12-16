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

// checks if the PriceVolumeTierData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceVolumeTierData{}

// PriceVolumeTierData struct for PriceVolumeTierData
type PriceVolumeTierData struct {
	// The resource's type
	Type          interface{}                                                   `json:"type"`
	Attributes    GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes `json:"attributes"`
	Relationships *PriceFrequencyTierDataRelationships                          `json:"relationships,omitempty"`
}

// NewPriceVolumeTierData instantiates a new PriceVolumeTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierData(type_ interface{}, attributes GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) *PriceVolumeTierData {
	this := PriceVolumeTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceVolumeTierDataWithDefaults instantiates a new PriceVolumeTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierDataWithDefaults() *PriceVolumeTierData {
	this := PriceVolumeTierData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceVolumeTierData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceVolumeTierData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceVolumeTierData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceVolumeTierData) GetAttributes() GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes {
	if o == nil {
		var ret GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierData) GetAttributesOk() (*GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceVolumeTierData) SetAttributes(v GETPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceVolumeTierData) GetRelationships() PriceFrequencyTierDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PriceFrequencyTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierData) GetRelationshipsOk() (*PriceFrequencyTierDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceVolumeTierData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceFrequencyTierDataRelationships and assigns it to the Relationships field.
func (o *PriceVolumeTierData) SetRelationships(v PriceFrequencyTierDataRelationships) {
	o.Relationships = &v
}

func (o PriceVolumeTierData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceVolumeTierData) ToMap() (map[string]interface{}, error) {
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

type NullablePriceVolumeTierData struct {
	value *PriceVolumeTierData
	isSet bool
}

func (v NullablePriceVolumeTierData) Get() *PriceVolumeTierData {
	return v.value
}

func (v *NullablePriceVolumeTierData) Set(val *PriceVolumeTierData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierData(val *PriceVolumeTierData) *NullablePriceVolumeTierData {
	return &NullablePriceVolumeTierData{value: val, isSet: true}
}

func (v NullablePriceVolumeTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
