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

// checks if the ParcelCreateDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelCreateDataRelationships{}

// ParcelCreateDataRelationships struct for ParcelCreateDataRelationships
type ParcelCreateDataRelationships struct {
	Shipment ParcelCreateDataRelationshipsShipment `json:"shipment"`
	Package  ParcelCreateDataRelationshipsPackage  `json:"package"`
}

// NewParcelCreateDataRelationships instantiates a new ParcelCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelCreateDataRelationships(shipment ParcelCreateDataRelationshipsShipment, package_ ParcelCreateDataRelationshipsPackage) *ParcelCreateDataRelationships {
	this := ParcelCreateDataRelationships{}
	this.Shipment = shipment
	this.Package = package_
	return &this
}

// NewParcelCreateDataRelationshipsWithDefaults instantiates a new ParcelCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelCreateDataRelationshipsWithDefaults() *ParcelCreateDataRelationships {
	this := ParcelCreateDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value
func (o *ParcelCreateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment {
	if o == nil {
		var ret ParcelCreateDataRelationshipsShipment
		return ret
	}

	return o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shipment, true
}

// SetShipment sets field value
func (o *ParcelCreateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment) {
	o.Shipment = v
}

// GetPackage returns the Package field value
func (o *ParcelCreateDataRelationships) GetPackage() ParcelCreateDataRelationshipsPackage {
	if o == nil {
		var ret ParcelCreateDataRelationshipsPackage
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateDataRelationships) GetPackageOk() (*ParcelCreateDataRelationshipsPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *ParcelCreateDataRelationships) SetPackage(v ParcelCreateDataRelationshipsPackage) {
	o.Package = v
}

func (o ParcelCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelCreateDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipment"] = o.Shipment
	toSerialize["package"] = o.Package
	return toSerialize, nil
}

type NullableParcelCreateDataRelationships struct {
	value *ParcelCreateDataRelationships
	isSet bool
}

func (v NullableParcelCreateDataRelationships) Get() *ParcelCreateDataRelationships {
	return v.value
}

func (v *NullableParcelCreateDataRelationships) Set(val *ParcelCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelCreateDataRelationships(val *ParcelCreateDataRelationships) *NullableParcelCreateDataRelationships {
	return &NullableParcelCreateDataRelationships{value: val, isSet: true}
}

func (v NullableParcelCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
