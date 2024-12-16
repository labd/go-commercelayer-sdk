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

// checks if the POSTSkuListPromotionRules201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuListPromotionRules201ResponseDataAttributes{}

// POSTSkuListPromotionRules201ResponseDataAttributes struct for POSTSkuListPromotionRules201ResponseDataAttributes
type POSTSkuListPromotionRules201ResponseDataAttributes struct {
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Indicates if the rule is activated only when all of the SKUs of the list is also part of the order.
	AllSkus interface{} `json:"all_skus,omitempty"`
	// The min quantity of SKUs of the list that must be also part of the order. If positive, overwrites the 'all_skus' option. When the SKU list is manual, its items quantities are honoured.
	MinQuantity interface{} `json:"min_quantity,omitempty"`
}

// NewPOSTSkuListPromotionRules201ResponseDataAttributes instantiates a new POSTSkuListPromotionRules201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListPromotionRules201ResponseDataAttributes() *POSTSkuListPromotionRules201ResponseDataAttributes {
	this := POSTSkuListPromotionRules201ResponseDataAttributes{}
	return &this
}

// NewPOSTSkuListPromotionRules201ResponseDataAttributesWithDefaults instantiates a new POSTSkuListPromotionRules201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListPromotionRules201ResponseDataAttributesWithDefaults() *POSTSkuListPromotionRules201ResponseDataAttributes {
	this := POSTSkuListPromotionRules201ResponseDataAttributes{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetAllSkus returns the AllSkus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetAllSkus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllSkus
}

// GetAllSkusOk returns a tuple with the AllSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetAllSkusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllSkus) {
		return nil, false
	}
	return &o.AllSkus, true
}

// HasAllSkus returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasAllSkus() bool {
	if o != nil && IsNil(o.AllSkus) {
		return true
	}

	return false
}

// SetAllSkus gets a reference to the given interface{} and assigns it to the AllSkus field.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetAllSkus(v interface{}) {
	o.AllSkus = v
}

// GetMinQuantity returns the MinQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMinQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MinQuantity
}

// GetMinQuantityOk returns a tuple with the MinQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) GetMinQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinQuantity) {
		return nil, false
	}
	return &o.MinQuantity, true
}

// HasMinQuantity returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) HasMinQuantity() bool {
	if o != nil && IsNil(o.MinQuantity) {
		return true
	}

	return false
}

// SetMinQuantity gets a reference to the given interface{} and assigns it to the MinQuantity field.
func (o *POSTSkuListPromotionRules201ResponseDataAttributes) SetMinQuantity(v interface{}) {
	o.MinQuantity = v
}

func (o POSTSkuListPromotionRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuListPromotionRules201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.AllSkus != nil {
		toSerialize["all_skus"] = o.AllSkus
	}
	if o.MinQuantity != nil {
		toSerialize["min_quantity"] = o.MinQuantity
	}
	return toSerialize, nil
}

type NullablePOSTSkuListPromotionRules201ResponseDataAttributes struct {
	value *POSTSkuListPromotionRules201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataAttributes) Get() *POSTSkuListPromotionRules201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataAttributes) Set(val *POSTSkuListPromotionRules201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListPromotionRules201ResponseDataAttributes(val *POSTSkuListPromotionRules201ResponseDataAttributes) *NullablePOSTSkuListPromotionRules201ResponseDataAttributes {
	return &NullablePOSTSkuListPromotionRules201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
