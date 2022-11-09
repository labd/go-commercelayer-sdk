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

// PATCHTaxRulesTaxRuleId200Response struct for PATCHTaxRulesTaxRuleId200Response
type PATCHTaxRulesTaxRuleId200Response struct {
	Data *PATCHTaxRulesTaxRuleId200ResponseData `json:"data,omitempty"`
}

// NewPATCHTaxRulesTaxRuleId200Response instantiates a new PATCHTaxRulesTaxRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHTaxRulesTaxRuleId200Response() *PATCHTaxRulesTaxRuleId200Response {
	this := PATCHTaxRulesTaxRuleId200Response{}
	return &this
}

// NewPATCHTaxRulesTaxRuleId200ResponseWithDefaults instantiates a new PATCHTaxRulesTaxRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHTaxRulesTaxRuleId200ResponseWithDefaults() *PATCHTaxRulesTaxRuleId200Response {
	this := PATCHTaxRulesTaxRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHTaxRulesTaxRuleId200Response) GetData() PATCHTaxRulesTaxRuleId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHTaxRulesTaxRuleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHTaxRulesTaxRuleId200Response) GetDataOk() (*PATCHTaxRulesTaxRuleId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHTaxRulesTaxRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHTaxRulesTaxRuleId200ResponseData and assigns it to the Data field.
func (o *PATCHTaxRulesTaxRuleId200Response) SetData(v PATCHTaxRulesTaxRuleId200ResponseData) {
	o.Data = &v
}

func (o PATCHTaxRulesTaxRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHTaxRulesTaxRuleId200Response struct {
	value *PATCHTaxRulesTaxRuleId200Response
	isSet bool
}

func (v NullablePATCHTaxRulesTaxRuleId200Response) Get() *PATCHTaxRulesTaxRuleId200Response {
	return v.value
}

func (v *NullablePATCHTaxRulesTaxRuleId200Response) Set(val *PATCHTaxRulesTaxRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHTaxRulesTaxRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHTaxRulesTaxRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHTaxRulesTaxRuleId200Response(val *PATCHTaxRulesTaxRuleId200Response) *NullablePATCHTaxRulesTaxRuleId200Response {
	return &NullablePATCHTaxRulesTaxRuleId200Response{value: val, isSet: true}
}

func (v NullablePATCHTaxRulesTaxRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHTaxRulesTaxRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}