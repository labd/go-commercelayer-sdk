/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTSkuListItems201ResponseDataAttributes struct for POSTSkuListItems201ResponseDataAttributes
type POSTSkuListItems201ResponseDataAttributes struct {
	// The SKU list item's position.
	Position *int32 `json:"position,omitempty"`
	// The code of the associated SKU.
	SkuCode *string `json:"sku_code,omitempty"`
	// The SKU quantity for this SKU list item.
	Quantity *int32 `json:"quantity,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPOSTSkuListItems201ResponseDataAttributes instantiates a new POSTSkuListItems201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListItems201ResponseDataAttributes() *POSTSkuListItems201ResponseDataAttributes {
	this := POSTSkuListItems201ResponseDataAttributes{}
	return &this
}

// NewPOSTSkuListItems201ResponseDataAttributesWithDefaults instantiates a new POSTSkuListItems201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListItems201ResponseDataAttributesWithDefaults() *POSTSkuListItems201ResponseDataAttributes {
	this := POSTSkuListItems201ResponseDataAttributes{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetPosition(v int32) {
	o.Position = &v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *POSTSkuListItems201ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o POSTSkuListItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
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

type NullablePOSTSkuListItems201ResponseDataAttributes struct {
	value *POSTSkuListItems201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSkuListItems201ResponseDataAttributes) Get() *POSTSkuListItems201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSkuListItems201ResponseDataAttributes) Set(val *POSTSkuListItems201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListItems201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListItems201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListItems201ResponseDataAttributes(val *POSTSkuListItems201ResponseDataAttributes) *NullablePOSTSkuListItems201ResponseDataAttributes {
	return &NullablePOSTSkuListItems201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSkuListItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListItems201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
