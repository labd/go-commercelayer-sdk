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

// checks if the POSTPackages201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPackages201ResponseDataRelationships{}

// POSTPackages201ResponseDataRelationships struct for POSTPackages201ResponseDataRelationships
type POSTPackages201ResponseDataRelationships struct {
	StockLocation *POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation          `json:"stock_location,omitempty"`
	Parcels       *POSTPackages201ResponseDataRelationshipsParcels                         `json:"parcels,omitempty"`
	Attachments   *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Versions      *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
}

// NewPOSTPackages201ResponseDataRelationships instantiates a new POSTPackages201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPackages201ResponseDataRelationships() *POSTPackages201ResponseDataRelationships {
	this := POSTPackages201ResponseDataRelationships{}
	return &this
}

// NewPOSTPackages201ResponseDataRelationshipsWithDefaults instantiates a new POSTPackages201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPackages201ResponseDataRelationshipsWithDefaults() *POSTPackages201ResponseDataRelationships {
	this := POSTPackages201ResponseDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *POSTPackages201ResponseDataRelationships) GetStockLocation() POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPackages201ResponseDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *POSTPackages201ResponseDataRelationships) SetStockLocation(v POSTDeliveryLeadTimes201ResponseDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *POSTPackages201ResponseDataRelationships) GetParcels() POSTPackages201ResponseDataRelationshipsParcels {
	if o == nil || IsNil(o.Parcels) {
		var ret POSTPackages201ResponseDataRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPackages201ResponseDataRelationships) GetParcelsOk() (*POSTPackages201ResponseDataRelationshipsParcels, bool) {
	if o == nil || IsNil(o.Parcels) {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataRelationships) HasParcels() bool {
	if o != nil && !IsNil(o.Parcels) {
		return true
	}

	return false
}

// SetParcels gets a reference to the given POSTPackages201ResponseDataRelationshipsParcels and assigns it to the Parcels field.
func (o *POSTPackages201ResponseDataRelationships) SetParcels(v POSTPackages201ResponseDataRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTPackages201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPackages201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTPackages201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTPackages201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTPackages201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTPackages201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTPackages201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

func (o POSTPackages201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPackages201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.Parcels) {
		toSerialize["parcels"] = o.Parcels
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullablePOSTPackages201ResponseDataRelationships struct {
	value *POSTPackages201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTPackages201ResponseDataRelationships) Get() *POSTPackages201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTPackages201ResponseDataRelationships) Set(val *POSTPackages201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPackages201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPackages201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPackages201ResponseDataRelationships(val *POSTPackages201ResponseDataRelationships) *NullablePOSTPackages201ResponseDataRelationships {
	return &NullablePOSTPackages201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTPackages201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPackages201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
