/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SubscriptionModelUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionModelUpdateData{}

// SubscriptionModelUpdateData struct for SubscriptionModelUpdateData
type SubscriptionModelUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                         `json:"id"`
	Attributes    PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                                         `json:"relationships,omitempty"`
}

// NewSubscriptionModelUpdateData instantiates a new SubscriptionModelUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionModelUpdateData(type_ interface{}, id interface{}, attributes PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) *SubscriptionModelUpdateData {
	this := SubscriptionModelUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewSubscriptionModelUpdateDataWithDefaults instantiates a new SubscriptionModelUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionModelUpdateDataWithDefaults() *SubscriptionModelUpdateData {
	this := SubscriptionModelUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SubscriptionModelUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionModelUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionModelUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SubscriptionModelUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionModelUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionModelUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionModelUpdateData) GetAttributes() PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionModelUpdateData) GetAttributesOk() (*PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionModelUpdateData) SetAttributes(v PATCHSubscriptionModelsSubscriptionModelId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionModelUpdateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionModelUpdateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionModelUpdateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *SubscriptionModelUpdateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o SubscriptionModelUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionModelUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionModelUpdateData struct {
	value *SubscriptionModelUpdateData
	isSet bool
}

func (v NullableSubscriptionModelUpdateData) Get() *SubscriptionModelUpdateData {
	return v.value
}

func (v *NullableSubscriptionModelUpdateData) Set(val *SubscriptionModelUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionModelUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionModelUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionModelUpdateData(val *SubscriptionModelUpdateData) *NullableSubscriptionModelUpdateData {
	return &NullableSubscriptionModelUpdateData{value: val, isSet: true}
}

func (v NullableSubscriptionModelUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionModelUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}