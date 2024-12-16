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

// checks if the GETBingGeocodersBingGeocoderId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETBingGeocodersBingGeocoderId200ResponseDataAttributes{}

// GETBingGeocodersBingGeocoderId200ResponseDataAttributes struct for GETBingGeocodersBingGeocoderId200ResponseDataAttributes
type GETBingGeocodersBingGeocoderId200ResponseDataAttributes struct {
	// The geocoder's internal name.
	Name interface{} `json:"name,omitempty"`
	// The geocoder's type.
	Type interface{} `json:"type,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETBingGeocodersBingGeocoderId200ResponseDataAttributes instantiates a new GETBingGeocodersBingGeocoderId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBingGeocodersBingGeocoderId200ResponseDataAttributes() *GETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	this := GETBingGeocodersBingGeocoderId200ResponseDataAttributes{}
	return &this
}

// NewGETBingGeocodersBingGeocoderId200ResponseDataAttributesWithDefaults instantiates a new GETBingGeocodersBingGeocoderId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBingGeocodersBingGeocoderId200ResponseDataAttributesWithDefaults() *GETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	this := GETBingGeocodersBingGeocoderId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetType(v interface{}) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETBingGeocodersBingGeocoderId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETBingGeocodersBingGeocoderId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes struct {
	value *GETBingGeocodersBingGeocoderId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) Get() *GETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) Set(val *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes(val *GETBingGeocodersBingGeocoderId200ResponseDataAttributes) *NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	return &NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBingGeocodersBingGeocoderId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
