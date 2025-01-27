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

// checks if the GETStockLineItemsStockLineItemId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETStockLineItemsStockLineItemId200ResponseDataAttributes{}

// GETStockLineItemsStockLineItemId200ResponseDataAttributes struct for GETStockLineItemsStockLineItemId200ResponseDataAttributes
type GETStockLineItemsStockLineItemId200ResponseDataAttributes struct {
	// The code of the associated SKU.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode interface{} `json:"bundle_code,omitempty"`
	// The line item quantity.
	Quantity interface{} `json:"quantity,omitempty"`
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

// NewGETStockLineItemsStockLineItemId200ResponseDataAttributes instantiates a new GETStockLineItemsStockLineItemId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLineItemsStockLineItemId200ResponseDataAttributes() *GETStockLineItemsStockLineItemId200ResponseDataAttributes {
	this := GETStockLineItemsStockLineItemId200ResponseDataAttributes{}
	return &this
}

// NewGETStockLineItemsStockLineItemId200ResponseDataAttributesWithDefaults instantiates a new GETStockLineItemsStockLineItemId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLineItemsStockLineItemId200ResponseDataAttributesWithDefaults() *GETStockLineItemsStockLineItemId200ResponseDataAttributes {
	this := GETStockLineItemsStockLineItemId200ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetBundleCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BundleCode) {
		return nil, false
	}
	return &o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasBundleCode() bool {
	if o != nil && IsNil(o.BundleCode) {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given interface{} and assigns it to the BundleCode field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetBundleCode(v interface{}) {
	o.BundleCode = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETStockLineItemsStockLineItemId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETStockLineItemsStockLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETStockLineItemsStockLineItemId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
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

type NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes struct {
	value *GETStockLineItemsStockLineItemId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) Get() *GETStockLineItemsStockLineItemId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) Set(val *GETStockLineItemsStockLineItemId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLineItemsStockLineItemId200ResponseDataAttributes(val *GETStockLineItemsStockLineItemId200ResponseDataAttributes) *NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes {
	return &NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLineItemsStockLineItemId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
