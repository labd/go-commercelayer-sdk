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

// InventoryStockLocationDataAttributes struct for InventoryStockLocationDataAttributes
type InventoryStockLocationDataAttributes struct {
	// The stock location priority within the associated invetory model.
	Priority *int32 `json:"priority,omitempty"`
	// Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled.
	OnHold *bool `json:"on_hold,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewInventoryStockLocationDataAttributes instantiates a new InventoryStockLocationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocationDataAttributes() *InventoryStockLocationDataAttributes {
	this := InventoryStockLocationDataAttributes{}
	return &this
}

// NewInventoryStockLocationDataAttributesWithDefaults instantiates a new InventoryStockLocationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationDataAttributesWithDefaults() *InventoryStockLocationDataAttributes {
	this := InventoryStockLocationDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InventoryStockLocationDataAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetOnHold() bool {
	if o == nil || o.OnHold == nil {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetOnHoldOk() (*bool, bool) {
	if o == nil || o.OnHold == nil {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasOnHold() bool {
	if o != nil && o.OnHold != nil {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *InventoryStockLocationDataAttributes) SetOnHold(v bool) {
	o.OnHold = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InventoryStockLocationDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InventoryStockLocationDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *InventoryStockLocationDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *InventoryStockLocationDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *InventoryStockLocationDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InventoryStockLocationDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InventoryStockLocationDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *InventoryStockLocationDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o InventoryStockLocationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.OnHold != nil {
		toSerialize["on_hold"] = o.OnHold
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryStockLocationDataAttributes struct {
	value *InventoryStockLocationDataAttributes
	isSet bool
}

func (v NullableInventoryStockLocationDataAttributes) Get() *InventoryStockLocationDataAttributes {
	return v.value
}

func (v *NullableInventoryStockLocationDataAttributes) Set(val *InventoryStockLocationDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocationDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocationDataAttributes(val *InventoryStockLocationDataAttributes) *NullableInventoryStockLocationDataAttributes {
	return &NullableInventoryStockLocationDataAttributes{value: val, isSet: true}
}

func (v NullableInventoryStockLocationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
