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

// checks if the POSTOrderCopies201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopies201ResponseDataAttributes{}

// POSTOrderCopies201ResponseDataAttributes struct for POSTOrderCopies201ResponseDataAttributes
type POSTOrderCopies201ResponseDataAttributes struct {
	// Indicates if the target order must be placed upon copy.
	PlaceTargetOrder interface{} `json:"place_target_order,omitempty"`
	// Indicates if the payment source within the source order customer's wallet must be copied.
	ReuseWallet interface{} `json:"reuse_wallet,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Indicates if the source order must be cancelled upon copy.
	CancelSourceOrder interface{} `json:"cancel_source_order,omitempty"`
	// Indicates if promotions got applied upon copy.
	ApplyPromotions interface{} `json:"apply_promotions,omitempty"`
	// Indicates to ignore any errors during copy.
	SkipErrors interface{} `json:"skip_errors,omitempty"`
	// Indicates to ignore invalid coupon code during copy.
	IgnoreInvalidCoupon interface{} `json:"ignore_invalid_coupon,omitempty"`
}

// NewPOSTOrderCopies201ResponseDataAttributes instantiates a new POSTOrderCopies201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataAttributes() *POSTOrderCopies201ResponseDataAttributes {
	this := POSTOrderCopies201ResponseDataAttributes{}
	return &this
}

// NewPOSTOrderCopies201ResponseDataAttributesWithDefaults instantiates a new POSTOrderCopies201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataAttributesWithDefaults() *POSTOrderCopies201ResponseDataAttributes {
	this := POSTOrderCopies201ResponseDataAttributes{}
	return &this
}

// GetPlaceTargetOrder returns the PlaceTargetOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PlaceTargetOrder
}

// GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PlaceTargetOrder) {
		return nil, false
	}
	return &o.PlaceTargetOrder, true
}

// HasPlaceTargetOrder returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasPlaceTargetOrder() bool {
	if o != nil && IsNil(o.PlaceTargetOrder) {
		return true
	}

	return false
}

// SetPlaceTargetOrder gets a reference to the given interface{} and assigns it to the PlaceTargetOrder field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetPlaceTargetOrder(v interface{}) {
	o.PlaceTargetOrder = v
}

// GetReuseWallet returns the ReuseWallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWallet() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReuseWallet
}

// GetReuseWalletOk returns a tuple with the ReuseWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWalletOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReuseWallet) {
		return nil, false
	}
	return &o.ReuseWallet, true
}

// HasReuseWallet returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReuseWallet() bool {
	if o != nil && IsNil(o.ReuseWallet) {
		return true
	}

	return false
}

// SetReuseWallet gets a reference to the given interface{} and assigns it to the ReuseWallet field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReuseWallet(v interface{}) {
	o.ReuseWallet = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCancelSourceOrder returns the CancelSourceOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CancelSourceOrder
}

// GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CancelSourceOrder) {
		return nil, false
	}
	return &o.CancelSourceOrder, true
}

// HasCancelSourceOrder returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasCancelSourceOrder() bool {
	if o != nil && IsNil(o.CancelSourceOrder) {
		return true
	}

	return false
}

// SetCancelSourceOrder gets a reference to the given interface{} and assigns it to the CancelSourceOrder field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetCancelSourceOrder(v interface{}) {
	o.CancelSourceOrder = v
}

// GetApplyPromotions returns the ApplyPromotions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetApplyPromotions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ApplyPromotions
}

// GetApplyPromotionsOk returns a tuple with the ApplyPromotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetApplyPromotionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ApplyPromotions) {
		return nil, false
	}
	return &o.ApplyPromotions, true
}

// HasApplyPromotions returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasApplyPromotions() bool {
	if o != nil && IsNil(o.ApplyPromotions) {
		return true
	}

	return false
}

// SetApplyPromotions gets a reference to the given interface{} and assigns it to the ApplyPromotions field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetApplyPromotions(v interface{}) {
	o.ApplyPromotions = v
}

// GetSkipErrors returns the SkipErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetSkipErrors() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkipErrors
}

// GetSkipErrorsOk returns a tuple with the SkipErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetSkipErrorsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkipErrors) {
		return nil, false
	}
	return &o.SkipErrors, true
}

// HasSkipErrors returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasSkipErrors() bool {
	if o != nil && IsNil(o.SkipErrors) {
		return true
	}

	return false
}

// SetSkipErrors gets a reference to the given interface{} and assigns it to the SkipErrors field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetSkipErrors(v interface{}) {
	o.SkipErrors = v
}

// GetIgnoreInvalidCoupon returns the IgnoreInvalidCoupon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopies201ResponseDataAttributes) GetIgnoreInvalidCoupon() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IgnoreInvalidCoupon
}

// GetIgnoreInvalidCouponOk returns a tuple with the IgnoreInvalidCoupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopies201ResponseDataAttributes) GetIgnoreInvalidCouponOk() (*interface{}, bool) {
	if o == nil || IsNil(o.IgnoreInvalidCoupon) {
		return nil, false
	}
	return &o.IgnoreInvalidCoupon, true
}

// HasIgnoreInvalidCoupon returns a boolean if a field has been set.
func (o *POSTOrderCopies201ResponseDataAttributes) HasIgnoreInvalidCoupon() bool {
	if o != nil && IsNil(o.IgnoreInvalidCoupon) {
		return true
	}

	return false
}

// SetIgnoreInvalidCoupon gets a reference to the given interface{} and assigns it to the IgnoreInvalidCoupon field.
func (o *POSTOrderCopies201ResponseDataAttributes) SetIgnoreInvalidCoupon(v interface{}) {
	o.IgnoreInvalidCoupon = v
}

func (o POSTOrderCopies201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopies201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlaceTargetOrder != nil {
		toSerialize["place_target_order"] = o.PlaceTargetOrder
	}
	if o.ReuseWallet != nil {
		toSerialize["reuse_wallet"] = o.ReuseWallet
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
	if o.CancelSourceOrder != nil {
		toSerialize["cancel_source_order"] = o.CancelSourceOrder
	}
	if o.ApplyPromotions != nil {
		toSerialize["apply_promotions"] = o.ApplyPromotions
	}
	if o.SkipErrors != nil {
		toSerialize["skip_errors"] = o.SkipErrors
	}
	if o.IgnoreInvalidCoupon != nil {
		toSerialize["ignore_invalid_coupon"] = o.IgnoreInvalidCoupon
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopies201ResponseDataAttributes struct {
	value *POSTOrderCopies201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) Get() *POSTOrderCopies201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) Set(val *POSTOrderCopies201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataAttributes(val *POSTOrderCopies201ResponseDataAttributes) *NullablePOSTOrderCopies201ResponseDataAttributes {
	return &NullablePOSTOrderCopies201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
