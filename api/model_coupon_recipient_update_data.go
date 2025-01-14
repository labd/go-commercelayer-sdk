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

// checks if the CouponRecipientUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponRecipientUpdateData{}

// CouponRecipientUpdateData struct for CouponRecipientUpdateData
type CouponRecipientUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                     `json:"id"`
	Attributes    PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes `json:"attributes"`
	Relationships *CouponRecipientCreateDataRelationships                         `json:"relationships,omitempty"`
}

// NewCouponRecipientUpdateData instantiates a new CouponRecipientUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) *CouponRecipientUpdateData {
	this := CouponRecipientUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCouponRecipientUpdateDataWithDefaults instantiates a new CouponRecipientUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientUpdateDataWithDefaults() *CouponRecipientUpdateData {
	this := CouponRecipientUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CouponRecipientUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponRecipientUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponRecipientUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CouponRecipientUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponRecipientUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponRecipientUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CouponRecipientUpdateData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponRecipientUpdateData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponRecipientUpdateData) GetRelationships() CouponRecipientCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CouponRecipientCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponRecipientUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponRecipientUpdateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponRecipientCreateDataRelationships and assigns it to the Relationships field.
func (o *CouponRecipientUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships) {
	o.Relationships = &v
}

func (o CouponRecipientUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponRecipientUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableCouponRecipientUpdateData struct {
	value *CouponRecipientUpdateData
	isSet bool
}

func (v NullableCouponRecipientUpdateData) Get() *CouponRecipientUpdateData {
	return v.value
}

func (v *NullableCouponRecipientUpdateData) Set(val *CouponRecipientUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientUpdateData(val *CouponRecipientUpdateData) *NullableCouponRecipientUpdateData {
	return &NullableCouponRecipientUpdateData{value: val, isSet: true}
}

func (v NullableCouponRecipientUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
