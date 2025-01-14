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

// checks if the ManualGatewayUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManualGatewayUpdateData{}

// ManualGatewayUpdateData struct for ManualGatewayUpdateData
type ManualGatewayUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                 `json:"id"`
	Attributes    PATCHManualGatewaysManualGatewayId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                                 `json:"relationships,omitempty"`
}

// NewManualGatewayUpdateData instantiates a new ManualGatewayUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayUpdateData(type_ interface{}, id interface{}, attributes PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) *ManualGatewayUpdateData {
	this := ManualGatewayUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewManualGatewayUpdateDataWithDefaults instantiates a new ManualGatewayUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayUpdateDataWithDefaults() *ManualGatewayUpdateData {
	this := ManualGatewayUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualGatewayUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualGatewayUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualGatewayUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualGatewayUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualGatewayUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ManualGatewayUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ManualGatewayUpdateData) GetAttributes() PATCHManualGatewaysManualGatewayId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHManualGatewaysManualGatewayId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayUpdateData) GetAttributesOk() (*PATCHManualGatewaysManualGatewayId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualGatewayUpdateData) SetAttributes(v PATCHManualGatewaysManualGatewayId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManualGatewayUpdateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualGatewayUpdateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualGatewayUpdateData) HasRelationships() bool {
	if o != nil && IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *ManualGatewayUpdateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o ManualGatewayUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManualGatewayUpdateData) ToMap() (map[string]interface{}, error) {
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

type NullableManualGatewayUpdateData struct {
	value *ManualGatewayUpdateData
	isSet bool
}

func (v NullableManualGatewayUpdateData) Get() *ManualGatewayUpdateData {
	return v.value
}

func (v *NullableManualGatewayUpdateData) Set(val *ManualGatewayUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayUpdateData(val *ManualGatewayUpdateData) *NullableManualGatewayUpdateData {
	return &NullableManualGatewayUpdateData{value: val, isSet: true}
}

func (v NullableManualGatewayUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
