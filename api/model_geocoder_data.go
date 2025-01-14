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

// checks if the GeocoderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeocoderData{}

// GeocoderData struct for GeocoderData
type GeocoderData struct {
	// The resource's type
	Type          interface{}                                             `json:"type"`
	Attributes    GETBingGeocodersBingGeocoderId200ResponseDataAttributes `json:"attributes"`
	Relationships *BingGeocoderDataRelationships                          `json:"relationships,omitempty"`
}

// NewGeocoderData instantiates a new GeocoderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeocoderData(type_ interface{}, attributes GETBingGeocodersBingGeocoderId200ResponseDataAttributes) *GeocoderData {
	this := GeocoderData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGeocoderDataWithDefaults instantiates a new GeocoderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeocoderDataWithDefaults() *GeocoderData {
	this := GeocoderData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GeocoderData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeocoderData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeocoderData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GeocoderData) GetAttributes() GETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	if o == nil {
		var ret GETBingGeocodersBingGeocoderId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GeocoderData) GetAttributesOk() (*GETBingGeocodersBingGeocoderId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GeocoderData) SetAttributes(v GETBingGeocodersBingGeocoderId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GeocoderData) GetRelationships() BingGeocoderDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BingGeocoderDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeocoderData) GetRelationshipsOk() (*BingGeocoderDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GeocoderData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BingGeocoderDataRelationships and assigns it to the Relationships field.
func (o *GeocoderData) SetRelationships(v BingGeocoderDataRelationships) {
	o.Relationships = &v
}

func (o GeocoderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeocoderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGeocoderData struct {
	value *GeocoderData
	isSet bool
}

func (v NullableGeocoderData) Get() *GeocoderData {
	return v.value
}

func (v *NullableGeocoderData) Set(val *GeocoderData) {
	v.value = val
	v.isSet = true
}

func (v NullableGeocoderData) IsSet() bool {
	return v.isSet
}

func (v *NullableGeocoderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeocoderData(val *GeocoderData) *NullableGeocoderData {
	return &NullableGeocoderData{value: val, isSet: true}
}

func (v NullableGeocoderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeocoderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
