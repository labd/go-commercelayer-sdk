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

// checks if the PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes{}

// PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes struct for PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes
type PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes struct {
	// The geocoder's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// The Bing Virtualearth key.
	Key interface{} `json:"key,omitempty"`
}

// NewPATCHBingGeocodersBingGeocoderId200ResponseDataAttributes instantiates a new PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBingGeocodersBingGeocoderId200ResponseDataAttributes() *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes {
	this := PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes{}
	return &this
}

// NewPATCHBingGeocodersBingGeocoderId200ResponseDataAttributesWithDefaults instantiates a new PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBingGeocodersBingGeocoderId200ResponseDataAttributesWithDefaults() *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes {
	this := PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) GetKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) HasKey() bool {
	if o != nil && IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) SetKey(v interface{}) {
	o.Key = v
}

func (o PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes struct {
	value *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) Get() *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) Set(val *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes(val *PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) *NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes {
	return &NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBingGeocodersBingGeocoderId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
