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

// checks if the PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes{}

// PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes struct for PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes
type PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes struct {
	// The code of the associated SKU, replace the relationship.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The threshold at which to trigger the back in stock notification.
	StockThreshold interface{} `json:"stock_threshold,omitempty"`
	// Send this attribute if you want to activate an inactive subscription.
	Activate interface{} `json:"_activate,omitempty"`
	// Send this attribute if you want to dactivate an active subscription.
	Deactivate interface{} `json:"_deactivate,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes instantiates a new PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes() *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	this := PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes{}
	return &this
}

// NewPATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributesWithDefaults instantiates a new PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributesWithDefaults() *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	this := PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetStockThreshold returns the StockThreshold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetStockThreshold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StockThreshold
}

// GetStockThresholdOk returns a tuple with the StockThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetStockThresholdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StockThreshold) {
		return nil, false
	}
	return &o.StockThreshold, true
}

// HasStockThreshold returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasStockThreshold() bool {
	if o != nil && IsNil(o.StockThreshold) {
		return true
	}

	return false
}

// SetStockThreshold gets a reference to the given interface{} and assigns it to the StockThreshold field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetStockThreshold(v interface{}) {
	o.StockThreshold = v
}

// GetActivate returns the Activate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetActivate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetActivateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return &o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasActivate() bool {
	if o != nil && IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given interface{} and assigns it to the Activate field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetActivate(v interface{}) {
	o.Activate = v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetDeactivate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetDeactivateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return &o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasDeactivate() bool {
	if o != nil && IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given interface{} and assigns it to the Deactivate field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetDeactivate(v interface{}) {
	o.Deactivate = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.StockThreshold != nil {
		toSerialize["stock_threshold"] = o.StockThreshold
	}
	if o.Activate != nil {
		toSerialize["_activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["_deactivate"] = o.Deactivate
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

type NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes struct {
	value *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) Get() *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) Set(val *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes(val *PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) *NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes {
	return &NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
