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

// checks if the PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes{}

// PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes struct for PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes
type PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes struct {
	// The customer group's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to identify the customer group (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes instantiates a new PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes() *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes {
	this := PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes{}
	return &this
}

// NewPATCHCustomerGroupsCustomerGroupId200ResponseDataAttributesWithDefaults instantiates a new PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerGroupsCustomerGroupId200ResponseDataAttributesWithDefaults() *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes {
	this := PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes struct {
	value *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) Get() *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) Set(val *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes(val *PATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) *NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes {
	return &NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
