/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CouponCreateData struct for CouponCreateData
type CouponCreateData struct {
	// The resource's type
	Type          string                         `json:"type"`
	Attributes    CouponCreateDataAttributes     `json:"attributes"`
	Relationships *CouponCreateDataRelationships `json:"relationships,omitempty"`
}

// NewCouponCreateData instantiates a new CouponCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCreateData(type_ string, attributes CouponCreateDataAttributes) *CouponCreateData {
	this := CouponCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCouponCreateDataWithDefaults instantiates a new CouponCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCreateDataWithDefaults() *CouponCreateData {
	this := CouponCreateData{}
	return &this
}

// GetType returns the Type field value
func (o *CouponCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CouponCreateData) GetAttributes() CouponCreateDataAttributes {
	if o == nil {
		var ret CouponCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponCreateData) GetAttributesOk() (*CouponCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponCreateData) SetAttributes(v CouponCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponCreateData) GetRelationships() CouponCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreateData) GetRelationshipsOk() (*CouponCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponCreateDataRelationships and assigns it to the Relationships field.
func (o *CouponCreateData) SetRelationships(v CouponCreateDataRelationships) {
	o.Relationships = &v
}

func (o CouponCreateData) MarshalJSON() ([]byte, error) {
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

type NullableCouponCreateData struct {
	value *CouponCreateData
	isSet bool
}

func (v NullableCouponCreateData) Get() *CouponCreateData {
	return v.value
}

func (v *NullableCouponCreateData) Set(val *CouponCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCreateData(val *CouponCreateData) *NullableCouponCreateData {
	return &NullableCouponCreateData{value: val, isSet: true}
}

func (v NullableCouponCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
