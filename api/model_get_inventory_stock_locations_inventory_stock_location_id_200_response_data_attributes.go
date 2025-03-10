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

// checks if the GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{}

// GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct for GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes
type GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct {
	// The stock location priority within the associated invetory model.
	Priority interface{} `json:"priority,omitempty"`
	// Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled.
	OnHold interface{} `json:"on_hold,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes instantiates a new GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes() *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	this := GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{}
	return &this
}

// NewGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributesWithDefaults instantiates a new GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributesWithDefaults() *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	this := GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasPriority() bool {
	if o != nil && IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetOnHold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetOnHoldOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return &o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasOnHold() bool {
	if o != nil && IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given interface{} and assigns it to the OnHold field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetOnHold(v interface{}) {
	o.OnHold = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.OnHold != nil {
		toSerialize["on_hold"] = o.OnHold
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
	return toSerialize, nil
}

type NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct {
	value *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Get() *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Set(val *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes(val *GETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) *NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	return &NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
