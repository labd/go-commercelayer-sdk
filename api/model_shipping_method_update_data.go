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

// ShippingMethodUpdateData struct for ShippingMethodUpdateData
type ShippingMethodUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                        `json:"id"`
	Attributes    PATCHShippingMethodsShippingMethodId200ResponseDataAttributes `json:"attributes"`
	Relationships *ShippingMethodCreateDataRelationships                        `json:"relationships,omitempty"`
}

// NewShippingMethodUpdateData instantiates a new ShippingMethodUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodUpdateData(type_ string, id string, attributes PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) *ShippingMethodUpdateData {
	this := ShippingMethodUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewShippingMethodUpdateDataWithDefaults instantiates a new ShippingMethodUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodUpdateDataWithDefaults() *ShippingMethodUpdateData {
	this := ShippingMethodUpdateData{}
	return &this
}

// GetType returns the Type field value
func (o *ShippingMethodUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShippingMethodUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ShippingMethodUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShippingMethodUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ShippingMethodUpdateData) GetAttributes() PATCHShippingMethodsShippingMethodId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHShippingMethodsShippingMethodId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetAttributesOk() (*PATCHShippingMethodsShippingMethodId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShippingMethodUpdateData) SetAttributes(v PATCHShippingMethodsShippingMethodId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShippingMethodUpdateData) GetRelationships() ShippingMethodCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShippingMethodCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateData) GetRelationshipsOk() (*ShippingMethodCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShippingMethodUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShippingMethodCreateDataRelationships and assigns it to the Relationships field.
func (o *ShippingMethodUpdateData) SetRelationships(v ShippingMethodCreateDataRelationships) {
	o.Relationships = &v
}

func (o ShippingMethodUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableShippingMethodUpdateData struct {
	value *ShippingMethodUpdateData
	isSet bool
}

func (v NullableShippingMethodUpdateData) Get() *ShippingMethodUpdateData {
	return v.value
}

func (v *NullableShippingMethodUpdateData) Set(val *ShippingMethodUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodUpdateData(val *ShippingMethodUpdateData) *NullableShippingMethodUpdateData {
	return &NullableShippingMethodUpdateData{value: val, isSet: true}
}

func (v NullableShippingMethodUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
