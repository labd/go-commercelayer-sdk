/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ParcelDataRelationships struct for ParcelDataRelationships
type ParcelDataRelationships struct {
	Shipment        *OrderDataRelationshipsShipments            `json:"shipment,omitempty"`
	Package         *ParcelDataRelationshipsPackage             `json:"package,omitempty"`
	ParcelLineItems *ParcelDataRelationshipsParcelLineItems     `json:"parcel_line_items,omitempty"`
	Attachments     *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
	Events          *CustomerAddressDataRelationshipsEvents     `json:"events,omitempty"`
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
func (o *ParcelDataRelationships) GetShipment() OrderDataRelationshipsShipments {
	if o == nil || o.Shipment == nil {
		var ret OrderDataRelationshipsShipments
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given OrderDataRelationshipsShipments and assigns it to the Shipment field.
func (o *ParcelDataRelationships) SetShipment(v OrderDataRelationshipsShipments) {
	o.Shipment = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetPackage() ParcelDataRelationshipsPackage {
	if o == nil || o.Package == nil {
		var ret ParcelDataRelationshipsPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetPackageOk() (*ParcelDataRelationshipsPackage, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasPackage() bool {
	if o != nil && o.Package != nil {
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
	if o == nil || o.ParcelLineItems == nil {
		var ret ParcelDataRelationshipsParcelLineItems
		return ret
	}
	return *o.ParcelLineItems
}

// GetParcelLineItemsOk returns a tuple with the ParcelLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetParcelLineItemsOk() (*ParcelDataRelationshipsParcelLineItems, bool) {
	if o == nil || o.ParcelLineItems == nil {
		return nil, false
	}
	return o.ParcelLineItems, true
}

// HasParcelLineItems returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasParcelLineItems() bool {
	if o != nil && o.ParcelLineItems != nil {
		return true
	}

	return false
}

// SetParcelLineItems gets a reference to the given ParcelDataRelationshipsParcelLineItems and assigns it to the ParcelLineItems field.
func (o *ParcelDataRelationships) SetParcelLineItems(v ParcelDataRelationshipsParcelLineItems) {
	o.ParcelLineItems = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ParcelDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ParcelDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ParcelDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *ParcelDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o ParcelDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.ParcelLineItems != nil {
		toSerialize["parcel_line_items"] = o.ParcelLineItems
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
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
