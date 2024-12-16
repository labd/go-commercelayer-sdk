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

// checks if the GoogleGeocoderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleGeocoderData{}

// GoogleGeocoderData struct for GoogleGeocoderData
type GoogleGeocoderData struct {
	// The resource's type
	Type          interface{}                                             `json:"type"`
	Attributes    GETBingGeocodersBingGeocoderId200ResponseDataAttributes `json:"attributes"`
	Relationships *BingGeocoderDataRelationships                          `json:"relationships,omitempty"`
}

// NewGoogleGeocoderData instantiates a new GoogleGeocoderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoderData(type_ interface{}, attributes GETBingGeocodersBingGeocoderId200ResponseDataAttributes) *GoogleGeocoderData {
	this := GoogleGeocoderData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGoogleGeocoderDataWithDefaults instantiates a new GoogleGeocoderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderDataWithDefaults() *GoogleGeocoderData {
	this := GoogleGeocoderData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GoogleGeocoderData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoogleGeocoderData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GoogleGeocoderData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GoogleGeocoderData) GetAttributes() GETBingGeocodersBingGeocoderId200ResponseDataAttributes {
	if o == nil {
		var ret GETBingGeocodersBingGeocoderId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderData) GetAttributesOk() (*GETBingGeocodersBingGeocoderId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GoogleGeocoderData) SetAttributes(v GETBingGeocodersBingGeocoderId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GoogleGeocoderData) GetRelationships() BingGeocoderDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BingGeocoderDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderData) GetRelationshipsOk() (*BingGeocoderDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GoogleGeocoderData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BingGeocoderDataRelationships and assigns it to the Relationships field.
func (o *GoogleGeocoderData) SetRelationships(v BingGeocoderDataRelationships) {
	o.Relationships = &v
}

func (o GoogleGeocoderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleGeocoderData) ToMap() (map[string]interface{}, error) {
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

type NullableGoogleGeocoderData struct {
	value *GoogleGeocoderData
	isSet bool
}

func (v NullableGoogleGeocoderData) Get() *GoogleGeocoderData {
	return v.value
}

func (v *NullableGoogleGeocoderData) Set(val *GoogleGeocoderData) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoderData) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoderData(val *GoogleGeocoderData) *NullableGoogleGeocoderData {
	return &NullableGoogleGeocoderData{value: val, isSet: true}
}

func (v NullableGoogleGeocoderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
