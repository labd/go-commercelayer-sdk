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

// checks if the GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes{}

// GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes struct for GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes
type GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes struct {
	// The code of the associated SKU.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode interface{} `json:"bundle_code,omitempty"`
	// The subscription item quantity.
	Quantity interface{} `json:"quantity,omitempty"`
	// The unit amount of the subscription item, in cents.
	UnitAmountCents interface{} `json:"unit_amount_cents,omitempty"`
	// The unit amount of the subscription item, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce.
	UnitAmountFloat interface{} `json:"unit_amount_float,omitempty"`
	// The unit amount of the subscription item, formatted. This can be useful to display the amount with currency in you views.
	FormattedUnitAmount interface{} `json:"formatted_unit_amount,omitempty"`
	// Calculated as unit amount x quantity amount, in cents.
	TotalAmountCents interface{} `json:"total_amount_cents,omitempty"`
	// Calculated as unit amount x quantity amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce.
	TotalAmountFloat interface{} `json:"total_amount_float,omitempty"`
	// Calculated as unit amount x quantity amount, formatted. This can be useful to display the amount with currency in you views.
	FormattedTotalAmount interface{} `json:"formatted_total_amount,omitempty"`
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

// NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes instantiates a new GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes() *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes {
	this := GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes{}
	return &this
}

// NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributesWithDefaults instantiates a new GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributesWithDefaults() *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes {
	this := GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetBundleCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BundleCode) {
		return nil, false
	}
	return &o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasBundleCode() bool {
	if o != nil && IsNil(o.BundleCode) {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given interface{} and assigns it to the BundleCode field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetBundleCode(v interface{}) {
	o.BundleCode = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitAmountCents) {
		return nil, false
	}
	return &o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUnitAmountCents() bool {
	if o != nil && IsNil(o.UnitAmountCents) {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given interface{} and assigns it to the UnitAmountCents field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountCents(v interface{}) {
	o.UnitAmountCents = v
}

// GetUnitAmountFloat returns the UnitAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitAmountFloat
}

// GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUnitAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitAmountFloat) {
		return nil, false
	}
	return &o.UnitAmountFloat, true
}

// HasUnitAmountFloat returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUnitAmountFloat() bool {
	if o != nil && IsNil(o.UnitAmountFloat) {
		return true
	}

	return false
}

// SetUnitAmountFloat gets a reference to the given interface{} and assigns it to the UnitAmountFloat field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUnitAmountFloat(v interface{}) {
	o.UnitAmountFloat = v
}

// GetFormattedUnitAmount returns the FormattedUnitAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedUnitAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedUnitAmount
}

// GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedUnitAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedUnitAmount) {
		return nil, false
	}
	return &o.FormattedUnitAmount, true
}

// HasFormattedUnitAmount returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasFormattedUnitAmount() bool {
	if o != nil && IsNil(o.FormattedUnitAmount) {
		return true
	}

	return false
}

// SetFormattedUnitAmount gets a reference to the given interface{} and assigns it to the FormattedUnitAmount field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedUnitAmount(v interface{}) {
	o.FormattedUnitAmount = v
}

// GetTotalAmountCents returns the TotalAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalAmountCents
}

// GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalAmountCents) {
		return nil, false
	}
	return &o.TotalAmountCents, true
}

// HasTotalAmountCents returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasTotalAmountCents() bool {
	if o != nil && IsNil(o.TotalAmountCents) {
		return true
	}

	return false
}

// SetTotalAmountCents gets a reference to the given interface{} and assigns it to the TotalAmountCents field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountCents(v interface{}) {
	o.TotalAmountCents = v
}

// GetTotalAmountFloat returns the TotalAmountFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TotalAmountFloat
}

// GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetTotalAmountFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TotalAmountFloat) {
		return nil, false
	}
	return &o.TotalAmountFloat, true
}

// HasTotalAmountFloat returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasTotalAmountFloat() bool {
	if o != nil && IsNil(o.TotalAmountFloat) {
		return true
	}

	return false
}

// SetTotalAmountFloat gets a reference to the given interface{} and assigns it to the TotalAmountFloat field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetTotalAmountFloat(v interface{}) {
	o.TotalAmountFloat = v
}

// GetFormattedTotalAmount returns the FormattedTotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedTotalAmount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedTotalAmount
}

// GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetFormattedTotalAmountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedTotalAmount) {
		return nil, false
	}
	return &o.FormattedTotalAmount, true
}

// HasFormattedTotalAmount returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasFormattedTotalAmount() bool {
	if o != nil && IsNil(o.FormattedTotalAmount) {
		return true
	}

	return false
}

// SetFormattedTotalAmount gets a reference to the given interface{} and assigns it to the FormattedTotalAmount field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetFormattedTotalAmount(v interface{}) {
	o.FormattedTotalAmount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
	}
	if o.UnitAmountFloat != nil {
		toSerialize["unit_amount_float"] = o.UnitAmountFloat
	}
	if o.FormattedUnitAmount != nil {
		toSerialize["formatted_unit_amount"] = o.FormattedUnitAmount
	}
	if o.TotalAmountCents != nil {
		toSerialize["total_amount_cents"] = o.TotalAmountCents
	}
	if o.TotalAmountFloat != nil {
		toSerialize["total_amount_float"] = o.TotalAmountFloat
	}
	if o.FormattedTotalAmount != nil {
		toSerialize["formatted_total_amount"] = o.FormattedTotalAmount
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

type NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes struct {
	value *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) Get() *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) Set(val *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes(val *GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) *NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes {
	return &NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
