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

// CaptureData struct for CaptureData
type CaptureData struct {
	// The resource's type
	Type          string                                    `json:"type"`
	Attributes    GETCaptures200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *CaptureDataRelationships                 `json:"relationships,omitempty"`
}

// NewCaptureData instantiates a new CaptureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureData(type_ string, attributes GETCaptures200ResponseDataInnerAttributes) *CaptureData {
	this := CaptureData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCaptureDataWithDefaults instantiates a new CaptureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureDataWithDefaults() *CaptureData {
	this := CaptureData{}
	return &this
}

// GetType returns the Type field value
func (o *CaptureData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaptureData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CaptureData) GetAttributes() GETCaptures200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETCaptures200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaptureData) GetAttributesOk() (*GETCaptures200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CaptureData) SetAttributes(v GETCaptures200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CaptureData) GetRelationships() CaptureDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CaptureDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureData) GetRelationshipsOk() (*CaptureDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CaptureData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CaptureDataRelationships and assigns it to the Relationships field.
func (o *CaptureData) SetRelationships(v CaptureDataRelationships) {
	o.Relationships = &v
}

func (o CaptureData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableCaptureData struct {
	value *CaptureData
	isSet bool
}

func (v NullableCaptureData) Get() *CaptureData {
	return v.value
}

func (v *NullableCaptureData) Set(val *CaptureData) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureData) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureData(val *CaptureData) *NullableCaptureData {
	return &NullableCaptureData{value: val, isSet: true}
}

func (v NullableCaptureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
