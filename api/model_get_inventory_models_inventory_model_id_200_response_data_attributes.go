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

// checks if the GETInventoryModelsInventoryModelId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETInventoryModelsInventoryModelId200ResponseDataAttributes{}

// GETInventoryModelsInventoryModelId200ResponseDataAttributes struct for GETInventoryModelsInventoryModelId200ResponseDataAttributes
type GETInventoryModelsInventoryModelId200ResponseDataAttributes struct {
	// The inventory model's internal name.
	Name interface{} `json:"name,omitempty"`
	// The inventory model's shipping strategy: one between 'no_split' (default), 'split_shipments', 'ship_from_primary' and 'ship_from_first_available_or_primary'.
	Strategy interface{} `json:"strategy,omitempty"`
	// The maximum number of stock locations used for inventory computation.
	StockLocationsCutoff interface{} `json:"stock_locations_cutoff,omitempty"`
	// The duration in seconds of the generated stock reservations.
	StockReservationCutoff interface{} `json:"stock_reservation_cutoff,omitempty"`
	// Indicates if the the stock transfers must be put on hold automatically with the associated shipment.
	PutStockTransfersOnHold interface{} `json:"put_stock_transfers_on_hold,omitempty"`
	// Indicates if the the stock will be decremented manually after the order approval.
	ManualStockDecrement interface{} `json:"manual_stock_decrement,omitempty"`
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

// NewGETInventoryModelsInventoryModelId200ResponseDataAttributes instantiates a new GETInventoryModelsInventoryModelId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModelsInventoryModelId200ResponseDataAttributes() *GETInventoryModelsInventoryModelId200ResponseDataAttributes {
	this := GETInventoryModelsInventoryModelId200ResponseDataAttributes{}
	return &this
}

// NewGETInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults instantiates a new GETInventoryModelsInventoryModelId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModelsInventoryModelId200ResponseDataAttributesWithDefaults() *GETInventoryModelsInventoryModelId200ResponseDataAttributes {
	this := GETInventoryModelsInventoryModelId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategy() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStrategyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return &o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasStrategy() bool {
	if o != nil && IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given interface{} and assigns it to the Strategy field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStrategy(v interface{}) {
	o.Strategy = v
}

// GetStockLocationsCutoff returns the StockLocationsCutoff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoff() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StockLocationsCutoff
}

// GetStockLocationsCutoffOk returns a tuple with the StockLocationsCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockLocationsCutoffOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StockLocationsCutoff) {
		return nil, false
	}
	return &o.StockLocationsCutoff, true
}

// HasStockLocationsCutoff returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockLocationsCutoff() bool {
	if o != nil && IsNil(o.StockLocationsCutoff) {
		return true
	}

	return false
}

// SetStockLocationsCutoff gets a reference to the given interface{} and assigns it to the StockLocationsCutoff field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockLocationsCutoff(v interface{}) {
	o.StockLocationsCutoff = v
}

// GetStockReservationCutoff returns the StockReservationCutoff field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockReservationCutoff() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StockReservationCutoff
}

// GetStockReservationCutoffOk returns a tuple with the StockReservationCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetStockReservationCutoffOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StockReservationCutoff) {
		return nil, false
	}
	return &o.StockReservationCutoff, true
}

// HasStockReservationCutoff returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasStockReservationCutoff() bool {
	if o != nil && IsNil(o.StockReservationCutoff) {
		return true
	}

	return false
}

// SetStockReservationCutoff gets a reference to the given interface{} and assigns it to the StockReservationCutoff field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetStockReservationCutoff(v interface{}) {
	o.StockReservationCutoff = v
}

// GetPutStockTransfersOnHold returns the PutStockTransfersOnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetPutStockTransfersOnHold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PutStockTransfersOnHold
}

// GetPutStockTransfersOnHoldOk returns a tuple with the PutStockTransfersOnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetPutStockTransfersOnHoldOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PutStockTransfersOnHold) {
		return nil, false
	}
	return &o.PutStockTransfersOnHold, true
}

// HasPutStockTransfersOnHold returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasPutStockTransfersOnHold() bool {
	if o != nil && IsNil(o.PutStockTransfersOnHold) {
		return true
	}

	return false
}

// SetPutStockTransfersOnHold gets a reference to the given interface{} and assigns it to the PutStockTransfersOnHold field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetPutStockTransfersOnHold(v interface{}) {
	o.PutStockTransfersOnHold = v
}

// GetManualStockDecrement returns the ManualStockDecrement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetManualStockDecrement() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ManualStockDecrement
}

// GetManualStockDecrementOk returns a tuple with the ManualStockDecrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetManualStockDecrementOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ManualStockDecrement) {
		return nil, false
	}
	return &o.ManualStockDecrement, true
}

// HasManualStockDecrement returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasManualStockDecrement() bool {
	if o != nil && IsNil(o.ManualStockDecrement) {
		return true
	}

	return false
}

// SetManualStockDecrement gets a reference to the given interface{} and assigns it to the ManualStockDecrement field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetManualStockDecrement(v interface{}) {
	o.ManualStockDecrement = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETInventoryModelsInventoryModelId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETInventoryModelsInventoryModelId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETInventoryModelsInventoryModelId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	if o.StockLocationsCutoff != nil {
		toSerialize["stock_locations_cutoff"] = o.StockLocationsCutoff
	}
	if o.StockReservationCutoff != nil {
		toSerialize["stock_reservation_cutoff"] = o.StockReservationCutoff
	}
	if o.PutStockTransfersOnHold != nil {
		toSerialize["put_stock_transfers_on_hold"] = o.PutStockTransfersOnHold
	}
	if o.ManualStockDecrement != nil {
		toSerialize["manual_stock_decrement"] = o.ManualStockDecrement
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

type NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes struct {
	value *GETInventoryModelsInventoryModelId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) Get() *GETInventoryModelsInventoryModelId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) Set(val *GETInventoryModelsInventoryModelId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModelsInventoryModelId200ResponseDataAttributes(val *GETInventoryModelsInventoryModelId200ResponseDataAttributes) *NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes {
	return &NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModelsInventoryModelId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
