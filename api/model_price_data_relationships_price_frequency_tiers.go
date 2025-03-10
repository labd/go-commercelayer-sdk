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

// checks if the PriceDataRelationshipsPriceFrequencyTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceDataRelationshipsPriceFrequencyTiers{}

// PriceDataRelationshipsPriceFrequencyTiers struct for PriceDataRelationshipsPriceFrequencyTiers
type PriceDataRelationshipsPriceFrequencyTiers struct {
	Data *PriceDataRelationshipsPriceFrequencyTiersData `json:"data,omitempty"`
}

// NewPriceDataRelationshipsPriceFrequencyTiers instantiates a new PriceDataRelationshipsPriceFrequencyTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDataRelationshipsPriceFrequencyTiers() *PriceDataRelationshipsPriceFrequencyTiers {
	this := PriceDataRelationshipsPriceFrequencyTiers{}
	return &this
}

// NewPriceDataRelationshipsPriceFrequencyTiersWithDefaults instantiates a new PriceDataRelationshipsPriceFrequencyTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDataRelationshipsPriceFrequencyTiersWithDefaults() *PriceDataRelationshipsPriceFrequencyTiers {
	this := PriceDataRelationshipsPriceFrequencyTiers{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PriceDataRelationshipsPriceFrequencyTiers) GetData() PriceDataRelationshipsPriceFrequencyTiersData {
	if o == nil || IsNil(o.Data) {
		var ret PriceDataRelationshipsPriceFrequencyTiersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceDataRelationshipsPriceFrequencyTiers) GetDataOk() (*PriceDataRelationshipsPriceFrequencyTiersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PriceDataRelationshipsPriceFrequencyTiers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PriceDataRelationshipsPriceFrequencyTiersData and assigns it to the Data field.
func (o *PriceDataRelationshipsPriceFrequencyTiers) SetData(v PriceDataRelationshipsPriceFrequencyTiersData) {
	o.Data = &v
}

func (o PriceDataRelationshipsPriceFrequencyTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceDataRelationshipsPriceFrequencyTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePriceDataRelationshipsPriceFrequencyTiers struct {
	value *PriceDataRelationshipsPriceFrequencyTiers
	isSet bool
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiers) Get() *PriceDataRelationshipsPriceFrequencyTiers {
	return v.value
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiers) Set(val *PriceDataRelationshipsPriceFrequencyTiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDataRelationshipsPriceFrequencyTiers(val *PriceDataRelationshipsPriceFrequencyTiers) *NullablePriceDataRelationshipsPriceFrequencyTiers {
	return &NullablePriceDataRelationshipsPriceFrequencyTiers{value: val, isSet: true}
}

func (v NullablePriceDataRelationshipsPriceFrequencyTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDataRelationshipsPriceFrequencyTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
