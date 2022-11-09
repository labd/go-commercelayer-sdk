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

// PATCHParcelLineItemsParcelLineItemId200Response struct for PATCHParcelLineItemsParcelLineItemId200Response
type PATCHParcelLineItemsParcelLineItemId200Response struct {
	Data *PATCHParcelLineItemsParcelLineItemId200ResponseData `json:"data,omitempty"`
}

// NewPATCHParcelLineItemsParcelLineItemId200Response instantiates a new PATCHParcelLineItemsParcelLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHParcelLineItemsParcelLineItemId200Response() *PATCHParcelLineItemsParcelLineItemId200Response {
	this := PATCHParcelLineItemsParcelLineItemId200Response{}
	return &this
}

// NewPATCHParcelLineItemsParcelLineItemId200ResponseWithDefaults instantiates a new PATCHParcelLineItemsParcelLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHParcelLineItemsParcelLineItemId200ResponseWithDefaults() *PATCHParcelLineItemsParcelLineItemId200Response {
	this := PATCHParcelLineItemsParcelLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHParcelLineItemsParcelLineItemId200Response) GetData() PATCHParcelLineItemsParcelLineItemId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHParcelLineItemsParcelLineItemId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHParcelLineItemsParcelLineItemId200Response) GetDataOk() (*PATCHParcelLineItemsParcelLineItemId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHParcelLineItemsParcelLineItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHParcelLineItemsParcelLineItemId200ResponseData and assigns it to the Data field.
func (o *PATCHParcelLineItemsParcelLineItemId200Response) SetData(v PATCHParcelLineItemsParcelLineItemId200ResponseData) {
	o.Data = &v
}

func (o PATCHParcelLineItemsParcelLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHParcelLineItemsParcelLineItemId200Response struct {
	value *PATCHParcelLineItemsParcelLineItemId200Response
	isSet bool
}

func (v NullablePATCHParcelLineItemsParcelLineItemId200Response) Get() *PATCHParcelLineItemsParcelLineItemId200Response {
	return v.value
}

func (v *NullablePATCHParcelLineItemsParcelLineItemId200Response) Set(val *PATCHParcelLineItemsParcelLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHParcelLineItemsParcelLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHParcelLineItemsParcelLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHParcelLineItemsParcelLineItemId200Response(val *PATCHParcelLineItemsParcelLineItemId200Response) *NullablePATCHParcelLineItemsParcelLineItemId200Response {
	return &NullablePATCHParcelLineItemsParcelLineItemId200Response{value: val, isSet: true}
}

func (v NullablePATCHParcelLineItemsParcelLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHParcelLineItemsParcelLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
