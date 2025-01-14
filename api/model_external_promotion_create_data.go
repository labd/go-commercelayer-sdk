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

// checks if the ExternalPromotionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPromotionCreateData{}

// ExternalPromotionCreateData struct for ExternalPromotionCreateData
type ExternalPromotionCreateData struct {
	// The resource's type
	Type          interface{}                                     `json:"type"`
	Attributes    POSTExternalPromotions201ResponseDataAttributes `json:"attributes"`
	Relationships *BuyXPayYPromotionUpdateDataRelationships       `json:"relationships,omitempty"`
}

// NewExternalPromotionCreateData instantiates a new ExternalPromotionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionCreateData(type_ interface{}, attributes POSTExternalPromotions201ResponseDataAttributes) *ExternalPromotionCreateData {
	this := ExternalPromotionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewExternalPromotionCreateDataWithDefaults instantiates a new ExternalPromotionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionCreateDataWithDefaults() *ExternalPromotionCreateData {
	this := ExternalPromotionCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ExternalPromotionCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPromotionCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalPromotionCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ExternalPromotionCreateData) GetAttributes() POSTExternalPromotions201ResponseDataAttributes {
	if o == nil {
		var ret POSTExternalPromotions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateData) GetAttributesOk() (*POSTExternalPromotions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ExternalPromotionCreateData) SetAttributes(v POSTExternalPromotions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ExternalPromotionCreateData) GetRelationships() BuyXPayYPromotionUpdateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BuyXPayYPromotionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionUpdateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ExternalPromotionCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BuyXPayYPromotionUpdateDataRelationships and assigns it to the Relationships field.
func (o *ExternalPromotionCreateData) SetRelationships(v BuyXPayYPromotionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o ExternalPromotionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPromotionCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableExternalPromotionCreateData struct {
	value *ExternalPromotionCreateData
	isSet bool
}

func (v NullableExternalPromotionCreateData) Get() *ExternalPromotionCreateData {
	return v.value
}

func (v *NullableExternalPromotionCreateData) Set(val *ExternalPromotionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionCreateData(val *ExternalPromotionCreateData) *NullableExternalPromotionCreateData {
	return &NullableExternalPromotionCreateData{value: val, isSet: true}
}

func (v NullableExternalPromotionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
