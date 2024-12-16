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

// checks if the POSTStores201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStores201ResponseDataAttributes{}

// POSTStores201ResponseDataAttributes struct for POSTStores201ResponseDataAttributes
type POSTStores201ResponseDataAttributes struct {
	// The store's internal name.
	Name interface{} `json:"name"`
	// A string that you can use to identify the store (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTStores201ResponseDataAttributes instantiates a new POSTStores201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStores201ResponseDataAttributes(name interface{}) *POSTStores201ResponseDataAttributes {
	this := POSTStores201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTStores201ResponseDataAttributesWithDefaults instantiates a new POSTStores201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStores201ResponseDataAttributesWithDefaults() *POSTStores201ResponseDataAttributes {
	this := POSTStores201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTStores201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStores201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTStores201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStores201ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStores201ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *POSTStores201ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *POSTStores201ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStores201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStores201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTStores201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTStores201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStores201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStores201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTStores201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTStores201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStores201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStores201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTStores201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTStores201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTStores201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStores201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
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

type NullablePOSTStores201ResponseDataAttributes struct {
	value *POSTStores201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTStores201ResponseDataAttributes) Get() *POSTStores201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTStores201ResponseDataAttributes) Set(val *POSTStores201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStores201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStores201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStores201ResponseDataAttributes(val *POSTStores201ResponseDataAttributes) *NullablePOSTStores201ResponseDataAttributes {
	return &NullablePOSTStores201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTStores201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStores201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
