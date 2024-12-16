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

// checks if the PriceFrequencyTierUpdateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceFrequencyTierUpdateDataRelationships{}

// PriceFrequencyTierUpdateDataRelationships struct for PriceFrequencyTierUpdateDataRelationships
type PriceFrequencyTierUpdateDataRelationships struct {
	Price *PriceFrequencyTierCreateDataRelationshipsPrice `json:"price,omitempty"`
}

// NewPriceFrequencyTierUpdateDataRelationships instantiates a new PriceFrequencyTierUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFrequencyTierUpdateDataRelationships() *PriceFrequencyTierUpdateDataRelationships {
	this := PriceFrequencyTierUpdateDataRelationships{}
	return &this
}

// NewPriceFrequencyTierUpdateDataRelationshipsWithDefaults instantiates a new PriceFrequencyTierUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFrequencyTierUpdateDataRelationshipsWithDefaults() *PriceFrequencyTierUpdateDataRelationships {
	this := PriceFrequencyTierUpdateDataRelationships{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceFrequencyTierUpdateDataRelationships) GetPrice() PriceFrequencyTierCreateDataRelationshipsPrice {
	if o == nil || IsNil(o.Price) {
		var ret PriceFrequencyTierCreateDataRelationshipsPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierUpdateDataRelationships) GetPriceOk() (*PriceFrequencyTierCreateDataRelationshipsPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceFrequencyTierUpdateDataRelationships) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PriceFrequencyTierCreateDataRelationshipsPrice and assigns it to the Price field.
func (o *PriceFrequencyTierUpdateDataRelationships) SetPrice(v PriceFrequencyTierCreateDataRelationshipsPrice) {
	o.Price = &v
}

func (o PriceFrequencyTierUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceFrequencyTierUpdateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullablePriceFrequencyTierUpdateDataRelationships struct {
	value *PriceFrequencyTierUpdateDataRelationships
	isSet bool
}

func (v NullablePriceFrequencyTierUpdateDataRelationships) Get() *PriceFrequencyTierUpdateDataRelationships {
	return v.value
}

func (v *NullablePriceFrequencyTierUpdateDataRelationships) Set(val *PriceFrequencyTierUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFrequencyTierUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFrequencyTierUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFrequencyTierUpdateDataRelationships(val *PriceFrequencyTierUpdateDataRelationships) *NullablePriceFrequencyTierUpdateDataRelationships {
	return &NullablePriceFrequencyTierUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceFrequencyTierUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFrequencyTierUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
