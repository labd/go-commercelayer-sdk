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

// checks if the PATCHPackagesPackageId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPackagesPackageId200ResponseDataAttributes{}

// PATCHPackagesPackageId200ResponseDataAttributes struct for PATCHPackagesPackageId200ResponseDataAttributes
type PATCHPackagesPackageId200ResponseDataAttributes struct {
	// Unique name for the package.
	Name interface{} `json:"name,omitempty"`
	// The package identifying code.
	Code interface{} `json:"code,omitempty"`
	// The package length, used to automatically calculate the tax rates from the available carrier accounts.
	Length interface{} `json:"length,omitempty"`
	// The package width, used to automatically calculate the tax rates from the available carrier accounts.
	Width interface{} `json:"width,omitempty"`
	// The package height, used to automatically calculate the tax rates from the available carrier accounts.
	Height interface{} `json:"height,omitempty"`
	// The unit of length. Can be one of 'cm', or 'in'.
	UnitOfLength interface{} `json:"unit_of_length,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHPackagesPackageId200ResponseDataAttributes instantiates a new PATCHPackagesPackageId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPackagesPackageId200ResponseDataAttributes() *PATCHPackagesPackageId200ResponseDataAttributes {
	this := PATCHPackagesPackageId200ResponseDataAttributes{}
	return &this
}

// NewPATCHPackagesPackageId200ResponseDataAttributesWithDefaults instantiates a new PATCHPackagesPackageId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPackagesPackageId200ResponseDataAttributesWithDefaults() *PATCHPackagesPackageId200ResponseDataAttributes {
	this := PATCHPackagesPackageId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetLength() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetLengthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return &o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasLength() bool {
	if o != nil && IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given interface{} and assigns it to the Length field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetLength(v interface{}) {
	o.Length = v
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetWidth() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetWidthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return &o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasWidth() bool {
	if o != nil && IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given interface{} and assigns it to the Width field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetWidth(v interface{}) {
	o.Width = v
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetHeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetHeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return &o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasHeight() bool {
	if o != nil && IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given interface{} and assigns it to the Height field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetHeight(v interface{}) {
	o.Height = v
}

// GetUnitOfLength returns the UnitOfLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetUnitOfLength() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitOfLength
}

// GetUnitOfLengthOk returns a tuple with the UnitOfLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetUnitOfLengthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitOfLength) {
		return nil, false
	}
	return &o.UnitOfLength, true
}

// HasUnitOfLength returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasUnitOfLength() bool {
	if o != nil && IsNil(o.UnitOfLength) {
		return true
	}

	return false
}

// SetUnitOfLength gets a reference to the given interface{} and assigns it to the UnitOfLength field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetUnitOfLength(v interface{}) {
	o.UnitOfLength = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHPackagesPackageId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHPackagesPackageId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHPackagesPackageId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPackagesPackageId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.UnitOfLength != nil {
		toSerialize["unit_of_length"] = o.UnitOfLength
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

type NullablePATCHPackagesPackageId200ResponseDataAttributes struct {
	value *PATCHPackagesPackageId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHPackagesPackageId200ResponseDataAttributes) Get() *PATCHPackagesPackageId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHPackagesPackageId200ResponseDataAttributes) Set(val *PATCHPackagesPackageId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPackagesPackageId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPackagesPackageId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPackagesPackageId200ResponseDataAttributes(val *PATCHPackagesPackageId200ResponseDataAttributes) *NullablePATCHPackagesPackageId200ResponseDataAttributes {
	return &NullablePATCHPackagesPackageId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHPackagesPackageId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPackagesPackageId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
