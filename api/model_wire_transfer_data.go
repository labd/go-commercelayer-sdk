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

// checks if the WireTransferData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WireTransferData{}

// WireTransferData struct for WireTransferData
type WireTransferData struct {
	// The resource's type
	Type          interface{}                                             `json:"type"`
	Attributes    GETWireTransfersWireTransferId200ResponseDataAttributes `json:"attributes"`
	Relationships *WireTransferDataRelationships                          `json:"relationships,omitempty"`
}

// NewWireTransferData instantiates a new WireTransferData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWireTransferData(type_ interface{}, attributes GETWireTransfersWireTransferId200ResponseDataAttributes) *WireTransferData {
	this := WireTransferData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewWireTransferDataWithDefaults instantiates a new WireTransferData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWireTransferDataWithDefaults() *WireTransferData {
	this := WireTransferData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WireTransferData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WireTransferData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WireTransferData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *WireTransferData) GetAttributes() GETWireTransfersWireTransferId200ResponseDataAttributes {
	if o == nil {
		var ret GETWireTransfersWireTransferId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *WireTransferData) GetAttributesOk() (*GETWireTransfersWireTransferId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *WireTransferData) SetAttributes(v GETWireTransfersWireTransferId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WireTransferData) GetRelationships() WireTransferDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret WireTransferDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WireTransferData) GetRelationshipsOk() (*WireTransferDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WireTransferData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given WireTransferDataRelationships and assigns it to the Relationships field.
func (o *WireTransferData) SetRelationships(v WireTransferDataRelationships) {
	o.Relationships = &v
}

func (o WireTransferData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WireTransferData) ToMap() (map[string]interface{}, error) {
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

type NullableWireTransferData struct {
	value *WireTransferData
	isSet bool
}

func (v NullableWireTransferData) Get() *WireTransferData {
	return v.value
}

func (v *NullableWireTransferData) Set(val *WireTransferData) {
	v.value = val
	v.isSet = true
}

func (v NullableWireTransferData) IsSet() bool {
	return v.isSet
}

func (v *NullableWireTransferData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWireTransferData(val *WireTransferData) *NullableWireTransferData {
	return &NullableWireTransferData{value: val, isSet: true}
}

func (v NullableWireTransferData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWireTransferData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
