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

// checks if the SatispayGatewayDataRelationshipsSatispayPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayGatewayDataRelationshipsSatispayPayments{}

// SatispayGatewayDataRelationshipsSatispayPayments struct for SatispayGatewayDataRelationshipsSatispayPayments
type SatispayGatewayDataRelationshipsSatispayPayments struct {
	Data *SatispayGatewayDataRelationshipsSatispayPaymentsData `json:"data,omitempty"`
}

// NewSatispayGatewayDataRelationshipsSatispayPayments instantiates a new SatispayGatewayDataRelationshipsSatispayPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayDataRelationshipsSatispayPayments() *SatispayGatewayDataRelationshipsSatispayPayments {
	this := SatispayGatewayDataRelationshipsSatispayPayments{}
	return &this
}

// NewSatispayGatewayDataRelationshipsSatispayPaymentsWithDefaults instantiates a new SatispayGatewayDataRelationshipsSatispayPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayDataRelationshipsSatispayPaymentsWithDefaults() *SatispayGatewayDataRelationshipsSatispayPayments {
	this := SatispayGatewayDataRelationshipsSatispayPayments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SatispayGatewayDataRelationshipsSatispayPayments) GetData() SatispayGatewayDataRelationshipsSatispayPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret SatispayGatewayDataRelationshipsSatispayPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayDataRelationshipsSatispayPayments) GetDataOk() (*SatispayGatewayDataRelationshipsSatispayPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SatispayGatewayDataRelationshipsSatispayPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SatispayGatewayDataRelationshipsSatispayPaymentsData and assigns it to the Data field.
func (o *SatispayGatewayDataRelationshipsSatispayPayments) SetData(v SatispayGatewayDataRelationshipsSatispayPaymentsData) {
	o.Data = &v
}

func (o SatispayGatewayDataRelationshipsSatispayPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayGatewayDataRelationshipsSatispayPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSatispayGatewayDataRelationshipsSatispayPayments struct {
	value *SatispayGatewayDataRelationshipsSatispayPayments
	isSet bool
}

func (v NullableSatispayGatewayDataRelationshipsSatispayPayments) Get() *SatispayGatewayDataRelationshipsSatispayPayments {
	return v.value
}

func (v *NullableSatispayGatewayDataRelationshipsSatispayPayments) Set(val *SatispayGatewayDataRelationshipsSatispayPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayDataRelationshipsSatispayPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayDataRelationshipsSatispayPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayDataRelationshipsSatispayPayments(val *SatispayGatewayDataRelationshipsSatispayPayments) *NullableSatispayGatewayDataRelationshipsSatispayPayments {
	return &NullableSatispayGatewayDataRelationshipsSatispayPayments{value: val, isSet: true}
}

func (v NullableSatispayGatewayDataRelationshipsSatispayPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayDataRelationshipsSatispayPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
