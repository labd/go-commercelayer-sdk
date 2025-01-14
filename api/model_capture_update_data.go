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

// checks if the CaptureUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureUpdateData{}

// CaptureUpdateData struct for CaptureUpdateData
type CaptureUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                     `json:"id"`
	Attributes    PATCHCapturesCaptureId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                     `json:"relationships,omitempty"`
}

// NewCaptureUpdateData instantiates a new CaptureUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureUpdateData(type_ interface{}, id interface{}, attributes PATCHCapturesCaptureId200ResponseDataAttributes) *CaptureUpdateData {
	this := CaptureUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCaptureUpdateDataWithDefaults instantiates a new CaptureUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureUpdateDataWithDefaults() *CaptureUpdateData {
	this := CaptureUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CaptureUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CaptureUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CaptureUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptureUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CaptureUpdateData) GetAttributes() PATCHCapturesCaptureId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCapturesCaptureId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaptureUpdateData) GetAttributesOk() (*PATCHCapturesCaptureId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CaptureUpdateData) SetAttributes(v PATCHCapturesCaptureId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CaptureUpdateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CaptureUpdateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CaptureUpdateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *CaptureUpdateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o CaptureUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCaptureUpdateData struct {
	value *CaptureUpdateData
	isSet bool
}

func (v NullableCaptureUpdateData) Get() *CaptureUpdateData {
	return v.value
}

func (v *NullableCaptureUpdateData) Set(val *CaptureUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureUpdateData(val *CaptureUpdateData) *NullableCaptureUpdateData {
	return &NullableCaptureUpdateData{value: val, isSet: true}
}

func (v NullableCaptureUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
