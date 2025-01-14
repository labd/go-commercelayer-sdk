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

// checks if the ParcelDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelDataRelationships{}

// ParcelDataRelationships struct for ParcelDataRelationships
type ParcelDataRelationships struct {
	Shipment        *LineItemDataRelationshipsShipment         `json:"shipment,omitempty"`
	Package         *ParcelDataRelationshipsPackage            `json:"package,omitempty"`
	ParcelLineItems *ParcelDataRelationshipsParcelLineItems    `json:"parcel_line_items,omitempty"`
	Attachments     *AuthorizationDataRelationshipsAttachments `json:"attachments,omitempty"`
	Events          *AddressDataRelationshipsEvents            `json:"events,omitempty"`
	Versions        *AddressDataRelationshipsVersions          `json:"versions,omitempty"`
}

// NewParcelDataRelationships instantiates a new ParcelDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelDataRelationships() *ParcelDataRelationships {
	this := ParcelDataRelationships{}
	return &this
}

// NewParcelDataRelationshipsWithDefaults instantiates a new ParcelDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelDataRelationshipsWithDefaults() *ParcelDataRelationships {
	this := ParcelDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetShipment() LineItemDataRelationshipsShipment {
	if o == nil || IsNil(o.Shipment) {
		var ret LineItemDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetShipmentOk() (*LineItemDataRelationshipsShipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given LineItemDataRelationshipsShipment and assigns it to the Shipment field.
func (o *ParcelDataRelationships) SetShipment(v LineItemDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetPackage() ParcelDataRelationshipsPackage {
	if o == nil || IsNil(o.Package) {
		var ret ParcelDataRelationshipsPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetPackageOk() (*ParcelDataRelationshipsPackage, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given ParcelDataRelationshipsPackage and assigns it to the Package field.
func (o *ParcelDataRelationships) SetPackage(v ParcelDataRelationshipsPackage) {
	o.Package = &v
}

// GetParcelLineItems returns the ParcelLineItems field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetParcelLineItems() ParcelDataRelationshipsParcelLineItems {
	if o == nil || IsNil(o.ParcelLineItems) {
		var ret ParcelDataRelationshipsParcelLineItems
		return ret
	}
	return *o.ParcelLineItems
}

// GetParcelLineItemsOk returns a tuple with the ParcelLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetParcelLineItemsOk() (*ParcelDataRelationshipsParcelLineItems, bool) {
	if o == nil || IsNil(o.ParcelLineItems) {
		return nil, false
	}
	return o.ParcelLineItems, true
}

// HasParcelLineItems returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasParcelLineItems() bool {
	if o != nil && !IsNil(o.ParcelLineItems) {
		return true
	}

	return false
}

// SetParcelLineItems gets a reference to the given ParcelDataRelationshipsParcelLineItems and assigns it to the ParcelLineItems field.
func (o *ParcelDataRelationships) SetParcelLineItems(v ParcelDataRelationshipsParcelLineItems) {
	o.ParcelLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AuthorizationDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AuthorizationDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ParcelDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetEvents() AddressDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret AddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given AddressDataRelationshipsEvents and assigns it to the Events field.
func (o *ParcelDataRelationships) SetEvents(v AddressDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetVersions() AddressDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret AddressDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given AddressDataRelationshipsVersions and assigns it to the Versions field.
func (o *ParcelDataRelationships) SetVersions(v AddressDataRelationshipsVersions) {
	o.Versions = &v
}

func (o ParcelDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.ParcelLineItems) {
		toSerialize["parcel_line_items"] = o.ParcelLineItems
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableParcelDataRelationships struct {
	value *ParcelDataRelationships
	isSet bool
}

func (v NullableParcelDataRelationships) Get() *ParcelDataRelationships {
	return v.value
}

func (v *NullableParcelDataRelationships) Set(val *ParcelDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelDataRelationships(val *ParcelDataRelationships) *NullableParcelDataRelationships {
	return &NullableParcelDataRelationships{value: val, isSet: true}
}

func (v NullableParcelDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
