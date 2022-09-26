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

// ReturnLineItemResponseDataRelationshipsReturn struct for ReturnLineItemResponseDataRelationshipsReturn
type ReturnLineItemResponseDataRelationshipsReturn struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *ReturnLineItemResponseDataRelationshipsReturnData `json:"data,omitempty"`
}

// NewReturnLineItemResponseDataRelationshipsReturn instantiates a new ReturnLineItemResponseDataRelationshipsReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemResponseDataRelationshipsReturn() *ReturnLineItemResponseDataRelationshipsReturn {
	this := ReturnLineItemResponseDataRelationshipsReturn{}
	return &this
}

// NewReturnLineItemResponseDataRelationshipsReturnWithDefaults instantiates a new ReturnLineItemResponseDataRelationshipsReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemResponseDataRelationshipsReturnWithDefaults() *ReturnLineItemResponseDataRelationshipsReturn {
	this := ReturnLineItemResponseDataRelationshipsReturn{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReturnLineItemResponseDataRelationshipsReturn) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemResponseDataRelationshipsReturn) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReturnLineItemResponseDataRelationshipsReturn) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *ReturnLineItemResponseDataRelationshipsReturn) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReturnLineItemResponseDataRelationshipsReturn) GetData() ReturnLineItemResponseDataRelationshipsReturnData {
	if o == nil || o.Data == nil {
		var ret ReturnLineItemResponseDataRelationshipsReturnData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLineItemResponseDataRelationshipsReturn) GetDataOk() (*ReturnLineItemResponseDataRelationshipsReturnData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReturnLineItemResponseDataRelationshipsReturn) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ReturnLineItemResponseDataRelationshipsReturnData and assigns it to the Data field.
func (o *ReturnLineItemResponseDataRelationshipsReturn) SetData(v ReturnLineItemResponseDataRelationshipsReturnData) {
	o.Data = &v
}

func (o ReturnLineItemResponseDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnLineItemResponseDataRelationshipsReturn struct {
	value *ReturnLineItemResponseDataRelationshipsReturn
	isSet bool
}

func (v NullableReturnLineItemResponseDataRelationshipsReturn) Get() *ReturnLineItemResponseDataRelationshipsReturn {
	return v.value
}

func (v *NullableReturnLineItemResponseDataRelationshipsReturn) Set(val *ReturnLineItemResponseDataRelationshipsReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemResponseDataRelationshipsReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemResponseDataRelationshipsReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemResponseDataRelationshipsReturn(val *ReturnLineItemResponseDataRelationshipsReturn) *NullableReturnLineItemResponseDataRelationshipsReturn {
	return &NullableReturnLineItemResponseDataRelationshipsReturn{value: val, isSet: true}
}

func (v NullableReturnLineItemResponseDataRelationshipsReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemResponseDataRelationshipsReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
