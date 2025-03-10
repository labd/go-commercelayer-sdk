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

// checks if the POSTSkuLists201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuLists201ResponseDataAttributes{}

// POSTSkuLists201ResponseDataAttributes struct for POSTSkuLists201ResponseDataAttributes
type POSTSkuLists201ResponseDataAttributes struct {
	// The SKU list's internal name.
	Name interface{} `json:"name"`
	// An internal description of the SKU list.
	Description interface{} `json:"description,omitempty"`
	// The URL of an image that represents the SKU list.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// Indicates if the SKU list is populated manually.
	Manual interface{} `json:"manual,omitempty"`
	// The regex that will be evaluated to match the SKU codes, max size is 5000.
	SkuCodeRegex interface{} `json:"sku_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTSkuLists201ResponseDataAttributes instantiates a new POSTSkuLists201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuLists201ResponseDataAttributes(name interface{}) *POSTSkuLists201ResponseDataAttributes {
	this := POSTSkuLists201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTSkuLists201ResponseDataAttributesWithDefaults instantiates a new POSTSkuLists201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuLists201ResponseDataAttributesWithDefaults() *POSTSkuLists201ResponseDataAttributes {
	this := POSTSkuLists201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTSkuLists201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *POSTSkuLists201ResponseDataAttributes) SetDescription(v interface{}) {
	o.Description = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *POSTSkuLists201ResponseDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetManual returns the Manual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetManual() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetManualOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return &o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasManual() bool {
	if o != nil && IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given interface{} and assigns it to the Manual field.
func (o *POSTSkuLists201ResponseDataAttributes) SetManual(v interface{}) {
	o.Manual = v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetSkuCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetSkuCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCodeRegex) {
		return nil, false
	}
	return &o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasSkuCodeRegex() bool {
	if o != nil && IsNil(o.SkuCodeRegex) {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given interface{} and assigns it to the SkuCodeRegex field.
func (o *POSTSkuLists201ResponseDataAttributes) SetSkuCodeRegex(v interface{}) {
	o.SkuCodeRegex = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTSkuLists201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTSkuLists201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSkuLists201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSkuLists201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSkuLists201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTSkuLists201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTSkuLists201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuLists201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Manual != nil {
		toSerialize["manual"] = o.Manual
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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

type NullablePOSTSkuLists201ResponseDataAttributes struct {
	value *POSTSkuLists201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSkuLists201ResponseDataAttributes) Get() *POSTSkuLists201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSkuLists201ResponseDataAttributes) Set(val *POSTSkuLists201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuLists201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuLists201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuLists201ResponseDataAttributes(val *POSTSkuLists201ResponseDataAttributes) *NullablePOSTSkuLists201ResponseDataAttributes {
	return &NullablePOSTSkuLists201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSkuLists201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuLists201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
