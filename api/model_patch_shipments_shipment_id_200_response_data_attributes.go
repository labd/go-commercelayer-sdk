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

// checks if the PATCHShipmentsShipmentId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHShipmentsShipmentId200ResponseDataAttributes{}

// PATCHShipmentsShipmentId200ResponseDataAttributes struct for PATCHShipmentsShipmentId200ResponseDataAttributes
type PATCHShipmentsShipmentId200ResponseDataAttributes struct {
	// Unique identifier for the shipment. Cannot be passed by sales channels.
	Number interface{} `json:"number,omitempty"`
	// Send this attribute if you want to mark this shipment as upcoming. Cannot be passed by sales channels.
	Upcoming interface{} `json:"_upcoming,omitempty"`
	// Send this attribute if you want to mark this shipment as cancelled (unless already shipped or delivered). Cannot be passed by sales channels.
	Cancel interface{} `json:"_cancel,omitempty"`
	// Send this attribute if you want to put this shipment on hold.
	OnHold interface{} `json:"_on_hold,omitempty"`
	// Send this attribute if you want to start picking this shipment.
	Picking interface{} `json:"_picking,omitempty"`
	// Send this attribute if you want to start packing this shipment.
	Packing interface{} `json:"_packing,omitempty"`
	// Send this attribute if you want to mark this shipment as ready to ship.
	ReadyToShip interface{} `json:"_ready_to_ship,omitempty"`
	// Send this attribute if you want to mark this shipment as shipped.
	Ship interface{} `json:"_ship,omitempty"`
	// Send this attribute if you want to mark this shipment as delivered.
	Deliver interface{} `json:"_deliver,omitempty"`
	// Send this attribute if you want to automatically reserve the stock for each of the associated stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels.
	ReserveStock interface{} `json:"_reserve_stock,omitempty"`
	// Send this attribute if you want to automatically destroy the stock reservations for each of the associated stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels.
	ReleaseStock interface{} `json:"_release_stock,omitempty"`
	// Send this attribute if you want to automatically decrement and release the stock for each of the associated stock line item. Can be done only when fulfillment is in progress. Cannot be passed by sales channels.
	DecrementStock interface{} `json:"_decrement_stock,omitempty"`
	// Send this attribute if you want get the shipping rates from the associated carrier accounts.
	GetRates interface{} `json:"_get_rates,omitempty"`
	// The selected purchase rate from the available shipping rates.
	SelectedRateId interface{} `json:"selected_rate_id,omitempty"`
	// Send this attribute if you want to purchase this shipment with the selected rate.
	Purchase interface{} `json:"_purchase,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHShipmentsShipmentId200ResponseDataAttributes instantiates a new PATCHShipmentsShipmentId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShipmentsShipmentId200ResponseDataAttributes() *PATCHShipmentsShipmentId200ResponseDataAttributes {
	this := PATCHShipmentsShipmentId200ResponseDataAttributes{}
	return &this
}

// NewPATCHShipmentsShipmentId200ResponseDataAttributesWithDefaults instantiates a new PATCHShipmentsShipmentId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShipmentsShipmentId200ResponseDataAttributesWithDefaults() *PATCHShipmentsShipmentId200ResponseDataAttributes {
	this := PATCHShipmentsShipmentId200ResponseDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return &o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasNumber() bool {
	if o != nil && IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given interface{} and assigns it to the Number field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetNumber(v interface{}) {
	o.Number = v
}

// GetUpcoming returns the Upcoming field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetUpcoming() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Upcoming
}

// GetUpcomingOk returns a tuple with the Upcoming field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetUpcomingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Upcoming) {
		return nil, false
	}
	return &o.Upcoming, true
}

// HasUpcoming returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasUpcoming() bool {
	if o != nil && IsNil(o.Upcoming) {
		return true
	}

	return false
}

// SetUpcoming gets a reference to the given interface{} and assigns it to the Upcoming field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetUpcoming(v interface{}) {
	o.Upcoming = v
}

// GetCancel returns the Cancel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetCancel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetCancelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return &o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasCancel() bool {
	if o != nil && IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given interface{} and assigns it to the Cancel field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetCancel(v interface{}) {
	o.Cancel = v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetOnHoldOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return &o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasOnHold() bool {
	if o != nil && IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given interface{} and assigns it to the OnHold field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetOnHold(v interface{}) {
	o.OnHold = v
}

// GetPicking returns the Picking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPicking() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Picking
}

// GetPickingOk returns a tuple with the Picking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPickingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Picking) {
		return nil, false
	}
	return &o.Picking, true
}

// HasPicking returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPicking() bool {
	if o != nil && IsNil(o.Picking) {
		return true
	}

	return false
}

// SetPicking gets a reference to the given interface{} and assigns it to the Picking field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPicking(v interface{}) {
	o.Picking = v
}

// GetPacking returns the Packing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPacking() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Packing
}

// GetPackingOk returns a tuple with the Packing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPackingOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Packing) {
		return nil, false
	}
	return &o.Packing, true
}

// HasPacking returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPacking() bool {
	if o != nil && IsNil(o.Packing) {
		return true
	}

	return false
}

// SetPacking gets a reference to the given interface{} and assigns it to the Packing field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPacking(v interface{}) {
	o.Packing = v
}

// GetReadyToShip returns the ReadyToShip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShip() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReadyToShip
}

// GetReadyToShipOk returns a tuple with the ReadyToShip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReadyToShipOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReadyToShip) {
		return nil, false
	}
	return &o.ReadyToShip, true
}

// HasReadyToShip returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReadyToShip() bool {
	if o != nil && IsNil(o.ReadyToShip) {
		return true
	}

	return false
}

// SetReadyToShip gets a reference to the given interface{} and assigns it to the ReadyToShip field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReadyToShip(v interface{}) {
	o.ReadyToShip = v
}

// GetShip returns the Ship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShip() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Ship
}

// GetShipOk returns a tuple with the Ship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetShipOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Ship) {
		return nil, false
	}
	return &o.Ship, true
}

// HasShip returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasShip() bool {
	if o != nil && IsNil(o.Ship) {
		return true
	}

	return false
}

// SetShip gets a reference to the given interface{} and assigns it to the Ship field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetShip(v interface{}) {
	o.Ship = v
}

// GetDeliver returns the Deliver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetDeliver() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Deliver
}

// GetDeliverOk returns a tuple with the Deliver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetDeliverOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Deliver) {
		return nil, false
	}
	return &o.Deliver, true
}

// HasDeliver returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasDeliver() bool {
	if o != nil && IsNil(o.Deliver) {
		return true
	}

	return false
}

// SetDeliver gets a reference to the given interface{} and assigns it to the Deliver field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetDeliver(v interface{}) {
	o.Deliver = v
}

// GetReserveStock returns the ReserveStock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReserveStock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReserveStock
}

// GetReserveStockOk returns a tuple with the ReserveStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReserveStockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReserveStock) {
		return nil, false
	}
	return &o.ReserveStock, true
}

// HasReserveStock returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReserveStock() bool {
	if o != nil && IsNil(o.ReserveStock) {
		return true
	}

	return false
}

// SetReserveStock gets a reference to the given interface{} and assigns it to the ReserveStock field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReserveStock(v interface{}) {
	o.ReserveStock = v
}

// GetReleaseStock returns the ReleaseStock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReleaseStock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReleaseStock
}

// GetReleaseStockOk returns a tuple with the ReleaseStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReleaseStockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReleaseStock) {
		return nil, false
	}
	return &o.ReleaseStock, true
}

// HasReleaseStock returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReleaseStock() bool {
	if o != nil && IsNil(o.ReleaseStock) {
		return true
	}

	return false
}

// SetReleaseStock gets a reference to the given interface{} and assigns it to the ReleaseStock field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReleaseStock(v interface{}) {
	o.ReleaseStock = v
}

// GetDecrementStock returns the DecrementStock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetDecrementStock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DecrementStock
}

// GetDecrementStockOk returns a tuple with the DecrementStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetDecrementStockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DecrementStock) {
		return nil, false
	}
	return &o.DecrementStock, true
}

// HasDecrementStock returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasDecrementStock() bool {
	if o != nil && IsNil(o.DecrementStock) {
		return true
	}

	return false
}

// SetDecrementStock gets a reference to the given interface{} and assigns it to the DecrementStock field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetDecrementStock(v interface{}) {
	o.DecrementStock = v
}

// GetGetRates returns the GetRates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRates() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.GetRates
}

// GetGetRatesOk returns a tuple with the GetRates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetGetRatesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.GetRates) {
		return nil, false
	}
	return &o.GetRates, true
}

// HasGetRates returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasGetRates() bool {
	if o != nil && IsNil(o.GetRates) {
		return true
	}

	return false
}

// SetGetRates gets a reference to the given interface{} and assigns it to the GetRates field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetGetRates(v interface{}) {
	o.GetRates = v
}

// GetSelectedRateId returns the SelectedRateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SelectedRateId
}

// GetSelectedRateIdOk returns a tuple with the SelectedRateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetSelectedRateIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SelectedRateId) {
		return nil, false
	}
	return &o.SelectedRateId, true
}

// HasSelectedRateId returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasSelectedRateId() bool {
	if o != nil && IsNil(o.SelectedRateId) {
		return true
	}

	return false
}

// SetSelectedRateId gets a reference to the given interface{} and assigns it to the SelectedRateId field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetSelectedRateId(v interface{}) {
	o.SelectedRateId = v
}

// GetPurchase returns the Purchase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchase() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Purchase
}

// GetPurchaseOk returns a tuple with the Purchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetPurchaseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Purchase) {
		return nil, false
	}
	return &o.Purchase, true
}

// HasPurchase returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasPurchase() bool {
	if o != nil && IsNil(o.Purchase) {
		return true
	}

	return false
}

// SetPurchase gets a reference to the given interface{} and assigns it to the Purchase field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetPurchase(v interface{}) {
	o.Purchase = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHShipmentsShipmentId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHShipmentsShipmentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHShipmentsShipmentId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Upcoming != nil {
		toSerialize["_upcoming"] = o.Upcoming
	}
	if o.Cancel != nil {
		toSerialize["_cancel"] = o.Cancel
	}
	if o.OnHold != nil {
		toSerialize["_on_hold"] = o.OnHold
	}
	if o.Picking != nil {
		toSerialize["_picking"] = o.Picking
	}
	if o.Packing != nil {
		toSerialize["_packing"] = o.Packing
	}
	if o.ReadyToShip != nil {
		toSerialize["_ready_to_ship"] = o.ReadyToShip
	}
	if o.Ship != nil {
		toSerialize["_ship"] = o.Ship
	}
	if o.Deliver != nil {
		toSerialize["_deliver"] = o.Deliver
	}
	if o.ReserveStock != nil {
		toSerialize["_reserve_stock"] = o.ReserveStock
	}
	if o.ReleaseStock != nil {
		toSerialize["_release_stock"] = o.ReleaseStock
	}
	if o.DecrementStock != nil {
		toSerialize["_decrement_stock"] = o.DecrementStock
	}
	if o.GetRates != nil {
		toSerialize["_get_rates"] = o.GetRates
	}
	if o.SelectedRateId != nil {
		toSerialize["selected_rate_id"] = o.SelectedRateId
	}
	if o.Purchase != nil {
		toSerialize["_purchase"] = o.Purchase
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

type NullablePATCHShipmentsShipmentId200ResponseDataAttributes struct {
	value *PATCHShipmentsShipmentId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHShipmentsShipmentId200ResponseDataAttributes) Get() *PATCHShipmentsShipmentId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHShipmentsShipmentId200ResponseDataAttributes) Set(val *PATCHShipmentsShipmentId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShipmentsShipmentId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShipmentsShipmentId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShipmentsShipmentId200ResponseDataAttributes(val *PATCHShipmentsShipmentId200ResponseDataAttributes) *NullablePATCHShipmentsShipmentId200ResponseDataAttributes {
	return &NullablePATCHShipmentsShipmentId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHShipmentsShipmentId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShipmentsShipmentId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
