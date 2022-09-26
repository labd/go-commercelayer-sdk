/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PackageResponseDataRelationships struct for PackageResponseDataRelationships
type PackageResponseDataRelationships struct {
	StockLocation *DeliveryLeadTimeResponseDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Parcels       *PackageResponseDataRelationshipsParcels                `json:"parcels,omitempty"`
	Attachments   *AvalaraAccountResponseDataRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewPackageResponseDataRelationships instantiates a new PackageResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageResponseDataRelationships() *PackageResponseDataRelationships {
	this := PackageResponseDataRelationships{}
	return &this
}

// NewPackageResponseDataRelationshipsWithDefaults instantiates a new PackageResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageResponseDataRelationshipsWithDefaults() *PackageResponseDataRelationships {
	this := PackageResponseDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *PackageResponseDataRelationships) GetStockLocation() DeliveryLeadTimeResponseDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageResponseDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeResponseDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *PackageResponseDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *PackageResponseDataRelationships) SetStockLocation(v DeliveryLeadTimeResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *PackageResponseDataRelationships) GetParcels() PackageResponseDataRelationshipsParcels {
	if o == nil || o.Parcels == nil {
		var ret PackageResponseDataRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageResponseDataRelationships) GetParcelsOk() (*PackageResponseDataRelationshipsParcels, bool) {
	if o == nil || o.Parcels == nil {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *PackageResponseDataRelationships) HasParcels() bool {
	if o != nil && o.Parcels != nil {
		return true
	}

	return false
}

// SetParcels gets a reference to the given PackageResponseDataRelationshipsParcels and assigns it to the Parcels field.
func (o *PackageResponseDataRelationships) SetParcels(v PackageResponseDataRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PackageResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PackageResponseDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PackageResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o PackageResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Parcels != nil {
		toSerialize["parcels"] = o.Parcels
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullablePackageResponseDataRelationships struct {
	value *PackageResponseDataRelationships
	isSet bool
}

func (v NullablePackageResponseDataRelationships) Get() *PackageResponseDataRelationships {
	return v.value
}

func (v *NullablePackageResponseDataRelationships) Set(val *PackageResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageResponseDataRelationships(val *PackageResponseDataRelationships) *NullablePackageResponseDataRelationships {
	return &NullablePackageResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePackageResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
