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

// checks if the POSTStockLocations201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockLocations201ResponseDataAttributes{}

// POSTStockLocations201ResponseDataAttributes struct for POSTStockLocations201ResponseDataAttributes
type POSTStockLocations201ResponseDataAttributes struct {
	// The stock location's internal name.
	Name interface{} `json:"name"`
	// A string that you can use to identify the stock location (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// The shipping label format for this stock location. Can be one of 'PDF', 'ZPL', 'EPL2', or 'PNG'.
	LabelFormat interface{} `json:"label_format,omitempty"`
	// Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments.
	SuppressEtd interface{} `json:"suppress_etd,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTStockLocations201ResponseDataAttributes instantiates a new POSTStockLocations201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockLocations201ResponseDataAttributes(name interface{}) *POSTStockLocations201ResponseDataAttributes {
	this := POSTStockLocations201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTStockLocations201ResponseDataAttributesWithDefaults instantiates a new POSTStockLocations201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockLocations201ResponseDataAttributesWithDefaults() *POSTStockLocations201ResponseDataAttributes {
	this := POSTStockLocations201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTStockLocations201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *POSTStockLocations201ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetLabelFormat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetLabelFormatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LabelFormat) {
		return nil, false
	}
	return &o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasLabelFormat() bool {
	if o != nil && IsNil(o.LabelFormat) {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given interface{} and assigns it to the LabelFormat field.
func (o *POSTStockLocations201ResponseDataAttributes) SetLabelFormat(v interface{}) {
	o.LabelFormat = v
}

// GetSuppressEtd returns the SuppressEtd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetSuppressEtd() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SuppressEtd
}

// GetSuppressEtdOk returns a tuple with the SuppressEtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetSuppressEtdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SuppressEtd) {
		return nil, false
	}
	return &o.SuppressEtd, true
}

// HasSuppressEtd returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasSuppressEtd() bool {
	if o != nil && IsNil(o.SuppressEtd) {
		return true
	}

	return false
}

// SetSuppressEtd gets a reference to the given interface{} and assigns it to the SuppressEtd field.
func (o *POSTStockLocations201ResponseDataAttributes) SetSuppressEtd(v interface{}) {
	o.SuppressEtd = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTStockLocations201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTStockLocations201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockLocations201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockLocations201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTStockLocations201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTStockLocations201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTStockLocations201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockLocations201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.LabelFormat != nil {
		toSerialize["label_format"] = o.LabelFormat
	}
	if o.SuppressEtd != nil {
		toSerialize["suppress_etd"] = o.SuppressEtd
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

type NullablePOSTStockLocations201ResponseDataAttributes struct {
	value *POSTStockLocations201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTStockLocations201ResponseDataAttributes) Get() *POSTStockLocations201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTStockLocations201ResponseDataAttributes) Set(val *POSTStockLocations201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockLocations201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockLocations201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockLocations201ResponseDataAttributes(val *POSTStockLocations201ResponseDataAttributes) *NullablePOSTStockLocations201ResponseDataAttributes {
	return &NullablePOSTStockLocations201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTStockLocations201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockLocations201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
