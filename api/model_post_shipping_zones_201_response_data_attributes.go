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

// checks if the POSTShippingZones201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingZones201ResponseDataAttributes{}

// POSTShippingZones201ResponseDataAttributes struct for POSTShippingZones201ResponseDataAttributes
type POSTShippingZones201ResponseDataAttributes struct {
	// The shipping zone's internal name.
	Name interface{} `json:"name"`
	// The regex that will be evaluated to match the shipping address country code, max size is 5000.
	CountryCodeRegex interface{} `json:"country_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address country code, max size is 5000.
	NotCountryCodeRegex interface{} `json:"not_country_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address state code, max size is 5000.
	StateCodeRegex interface{} `json:"state_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address state code, max size is 5000.
	NotStateCodeRegex interface{} `json:"not_state_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address zip code, max size is 5000.
	ZipCodeRegex interface{} `json:"zip_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping zip country code, max size is 5000.
	NotZipCodeRegex interface{} `json:"not_zip_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTShippingZones201ResponseDataAttributes instantiates a new POSTShippingZones201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingZones201ResponseDataAttributes(name interface{}) *POSTShippingZones201ResponseDataAttributes {
	this := POSTShippingZones201ResponseDataAttributes{}
	this.Name = name
	return &this
}

// NewPOSTShippingZones201ResponseDataAttributesWithDefaults instantiates a new POSTShippingZones201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingZones201ResponseDataAttributesWithDefaults() *POSTShippingZones201ResponseDataAttributes {
	this := POSTShippingZones201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTShippingZones201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCountryCodeRegex returns the CountryCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetCountryCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CountryCodeRegex
}

// GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CountryCodeRegex) {
		return nil, false
	}
	return &o.CountryCodeRegex, true
}

// HasCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasCountryCodeRegex() bool {
	if o != nil && IsNil(o.CountryCodeRegex) {
		return true
	}

	return false
}

// SetCountryCodeRegex gets a reference to the given interface{} and assigns it to the CountryCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetCountryCodeRegex(v interface{}) {
	o.CountryCodeRegex = v
}

// GetNotCountryCodeRegex returns the NotCountryCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetNotCountryCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotCountryCodeRegex
}

// GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotCountryCodeRegex) {
		return nil, false
	}
	return &o.NotCountryCodeRegex, true
}

// HasNotCountryCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotCountryCodeRegex() bool {
	if o != nil && IsNil(o.NotCountryCodeRegex) {
		return true
	}

	return false
}

// SetNotCountryCodeRegex gets a reference to the given interface{} and assigns it to the NotCountryCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotCountryCodeRegex(v interface{}) {
	o.NotCountryCodeRegex = v
}

// GetStateCodeRegex returns the StateCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetStateCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StateCodeRegex
}

// GetStateCodeRegexOk returns a tuple with the StateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StateCodeRegex) {
		return nil, false
	}
	return &o.StateCodeRegex, true
}

// HasStateCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasStateCodeRegex() bool {
	if o != nil && IsNil(o.StateCodeRegex) {
		return true
	}

	return false
}

// SetStateCodeRegex gets a reference to the given interface{} and assigns it to the StateCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetStateCodeRegex(v interface{}) {
	o.StateCodeRegex = v
}

// GetNotStateCodeRegex returns the NotStateCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetNotStateCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotStateCodeRegex
}

// GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotStateCodeRegex) {
		return nil, false
	}
	return &o.NotStateCodeRegex, true
}

// HasNotStateCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotStateCodeRegex() bool {
	if o != nil && IsNil(o.NotStateCodeRegex) {
		return true
	}

	return false
}

// SetNotStateCodeRegex gets a reference to the given interface{} and assigns it to the NotStateCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotStateCodeRegex(v interface{}) {
	o.NotStateCodeRegex = v
}

// GetZipCodeRegex returns the ZipCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetZipCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ZipCodeRegex
}

// GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ZipCodeRegex) {
		return nil, false
	}
	return &o.ZipCodeRegex, true
}

// HasZipCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasZipCodeRegex() bool {
	if o != nil && IsNil(o.ZipCodeRegex) {
		return true
	}

	return false
}

// SetZipCodeRegex gets a reference to the given interface{} and assigns it to the ZipCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetZipCodeRegex(v interface{}) {
	o.ZipCodeRegex = v
}

// GetNotZipCodeRegex returns the NotZipCodeRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetNotZipCodeRegex() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NotZipCodeRegex
}

// GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NotZipCodeRegex) {
		return nil, false
	}
	return &o.NotZipCodeRegex, true
}

// HasNotZipCodeRegex returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasNotZipCodeRegex() bool {
	if o != nil && IsNil(o.NotZipCodeRegex) {
		return true
	}

	return false
}

// SetNotZipCodeRegex gets a reference to the given interface{} and assigns it to the NotZipCodeRegex field.
func (o *POSTShippingZones201ResponseDataAttributes) SetNotZipCodeRegex(v interface{}) {
	o.NotZipCodeRegex = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTShippingZones201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTShippingZones201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingZones201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingZones201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTShippingZones201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTShippingZones201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTShippingZones201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingZones201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CountryCodeRegex != nil {
		toSerialize["country_code_regex"] = o.CountryCodeRegex
	}
	if o.NotCountryCodeRegex != nil {
		toSerialize["not_country_code_regex"] = o.NotCountryCodeRegex
	}
	if o.StateCodeRegex != nil {
		toSerialize["state_code_regex"] = o.StateCodeRegex
	}
	if o.NotStateCodeRegex != nil {
		toSerialize["not_state_code_regex"] = o.NotStateCodeRegex
	}
	if o.ZipCodeRegex != nil {
		toSerialize["zip_code_regex"] = o.ZipCodeRegex
	}
	if o.NotZipCodeRegex != nil {
		toSerialize["not_zip_code_regex"] = o.NotZipCodeRegex
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

type NullablePOSTShippingZones201ResponseDataAttributes struct {
	value *POSTShippingZones201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) Get() *POSTShippingZones201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) Set(val *POSTShippingZones201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingZones201ResponseDataAttributes(val *POSTShippingZones201ResponseDataAttributes) *NullablePOSTShippingZones201ResponseDataAttributes {
	return &NullablePOSTShippingZones201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTShippingZones201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingZones201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
