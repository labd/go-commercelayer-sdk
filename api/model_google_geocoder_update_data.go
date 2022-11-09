/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.2
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleGeocoderUpdateData struct for GoogleGeocoderUpdateData
type GoogleGeocoderUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                        `json:"id"`
	Attributes    PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                                        `json:"relationships,omitempty"`
}

// NewGoogleGeocoderUpdateData instantiates a new GoogleGeocoderUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoderUpdateData(type_ string, id string, attributes PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes) *GoogleGeocoderUpdateData {
	this := GoogleGeocoderUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewGoogleGeocoderUpdateDataWithDefaults instantiates a new GoogleGeocoderUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderUpdateDataWithDefaults() *GoogleGeocoderUpdateData {
	this := GoogleGeocoderUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *GoogleGeocoderUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GoogleGeocoderUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GoogleGeocoderUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GoogleGeocoderUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GoogleGeocoderUpdateData) GetAttributes() PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderUpdateData) GetAttributesOk() (*PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GoogleGeocoderUpdateData) SetAttributes(v PATCHGoogleGeocodersGoogleGeocoderId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GoogleGeocoderUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GoogleGeocoderUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *GoogleGeocoderUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o GoogleGeocoderUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleGeocoderUpdateData struct {
	value *GoogleGeocoderUpdateData
	isSet bool
}

func (v NullableGoogleGeocoderUpdateData) Get() *GoogleGeocoderUpdateData {
	return v.value
}

func (v *NullableGoogleGeocoderUpdateData) Set(val *GoogleGeocoderUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoderUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoderUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoderUpdateData(val *GoogleGeocoderUpdateData) *NullableGoogleGeocoderUpdateData {
	return &NullableGoogleGeocoderUpdateData{value: val, isSet: true}
}

func (v NullableGoogleGeocoderUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoderUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
