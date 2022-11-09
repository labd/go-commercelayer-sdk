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

// POSTLineItems201ResponseDataAttributes struct for POSTLineItems201ResponseDataAttributes
type POSTLineItems201ResponseDataAttributes struct {
	// The code of the associated SKU.
	SkuCode *string `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode *string `json:"bundle_code,omitempty"`
	// The line item quantity.
	Quantity int32 `json:"quantity"`
	// When creating a new line item, set this attribute to '1' if you want to inject the unit_amount_cents price from an external source.
	ExternalPrice *bool `json:"_external_price,omitempty"`
	// When creating a new line item, set this attribute to '1' if you want to update the line item quantity (if present) instead of creating a new line item for the same SKU.
	UpdateQuantity *bool `json:"_update_quantity,omitempty"`
	// The unit amount of the line item, in cents. Can be specified without an item, otherwise is automatically populated from the price list associated to the order's market.
	UnitAmountCents *int32 `json:"unit_amount_cents,omitempty"`
	// The name of the line item. When blank, it gets populated with the name of the associated item (if present).
	Name *string `json:"name,omitempty"`
	// The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only).
	ImageUrl *string `json:"image_url,omitempty"`
	// The type of the associate item. Can be one of 'sku', 'bundle', 'shipment', 'payment_method', 'adjustment', 'gift_card', or a valid promotion type.
	ItemType *string `json:"item_type,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTLineItems201ResponseDataAttributes instantiates a new POSTLineItems201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataAttributes(quantity int32) *POSTLineItems201ResponseDataAttributes {
	this := POSTLineItems201ResponseDataAttributes{}
	this.Quantity = quantity
	return &this
}

// NewPOSTLineItems201ResponseDataAttributesWithDefaults instantiates a new POSTLineItems201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataAttributesWithDefaults() *POSTLineItems201ResponseDataAttributes {
	this := POSTLineItems201ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *POSTLineItems201ResponseDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetBundleCode() string {
	if o == nil || o.BundleCode == nil {
		var ret string
		return ret
	}
	return *o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetBundleCodeOk() (*string, bool) {
	if o == nil || o.BundleCode == nil {
		return nil, false
	}
	return o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasBundleCode() bool {
	if o != nil && o.BundleCode != nil {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given string and assigns it to the BundleCode field.
func (o *POSTLineItems201ResponseDataAttributes) SetBundleCode(v string) {
	o.BundleCode = &v
}

// GetQuantity returns the Quantity field value
func (o *POSTLineItems201ResponseDataAttributes) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *POSTLineItems201ResponseDataAttributes) SetQuantity(v int32) {
	o.Quantity = v
}

// GetExternalPrice returns the ExternalPrice field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetExternalPrice() bool {
	if o == nil || o.ExternalPrice == nil {
		var ret bool
		return ret
	}
	return *o.ExternalPrice
}

// GetExternalPriceOk returns a tuple with the ExternalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetExternalPriceOk() (*bool, bool) {
	if o == nil || o.ExternalPrice == nil {
		return nil, false
	}
	return o.ExternalPrice, true
}

// HasExternalPrice returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasExternalPrice() bool {
	if o != nil && o.ExternalPrice != nil {
		return true
	}

	return false
}

// SetExternalPrice gets a reference to the given bool and assigns it to the ExternalPrice field.
func (o *POSTLineItems201ResponseDataAttributes) SetExternalPrice(v bool) {
	o.ExternalPrice = &v
}

// GetUpdateQuantity returns the UpdateQuantity field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetUpdateQuantity() bool {
	if o == nil || o.UpdateQuantity == nil {
		var ret bool
		return ret
	}
	return *o.UpdateQuantity
}

// GetUpdateQuantityOk returns a tuple with the UpdateQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetUpdateQuantityOk() (*bool, bool) {
	if o == nil || o.UpdateQuantity == nil {
		return nil, false
	}
	return o.UpdateQuantity, true
}

// HasUpdateQuantity returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasUpdateQuantity() bool {
	if o != nil && o.UpdateQuantity != nil {
		return true
	}

	return false
}

// SetUpdateQuantity gets a reference to the given bool and assigns it to the UpdateQuantity field.
func (o *POSTLineItems201ResponseDataAttributes) SetUpdateQuantity(v bool) {
	o.UpdateQuantity = &v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetUnitAmountCents() int32 {
	if o == nil || o.UnitAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetUnitAmountCentsOk() (*int32, bool) {
	if o == nil || o.UnitAmountCents == nil {
		return nil, false
	}
	return o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasUnitAmountCents() bool {
	if o != nil && o.UnitAmountCents != nil {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given int32 and assigns it to the UnitAmountCents field.
func (o *POSTLineItems201ResponseDataAttributes) SetUnitAmountCents(v int32) {
	o.UnitAmountCents = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *POSTLineItems201ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *POSTLineItems201ResponseDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *POSTLineItems201ResponseDataAttributes) SetItemType(v string) {
	o.ItemType = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTLineItems201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTLineItems201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTLineItems201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ExternalPrice != nil {
		toSerialize["_external_price"] = o.ExternalPrice
	}
	if o.UpdateQuantity != nil {
		toSerialize["_update_quantity"] = o.UpdateQuantity
	}
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.ItemType != nil {
		toSerialize["item_type"] = o.ItemType
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

type NullablePOSTLineItems201ResponseDataAttributes struct {
	value *POSTLineItems201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataAttributes) Get() *POSTLineItems201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataAttributes) Set(val *POSTLineItems201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataAttributes(val *POSTLineItems201ResponseDataAttributes) *NullablePOSTLineItems201ResponseDataAttributes {
	return &NullablePOSTLineItems201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}