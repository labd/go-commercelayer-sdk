/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETDeliveryLeadTimesDeliveryLeadTimeId200Response struct for GETDeliveryLeadTimesDeliveryLeadTimeId200Response
type GETDeliveryLeadTimesDeliveryLeadTimeId200Response struct {
	Data *GETDeliveryLeadTimes200ResponseDataInner `json:"data,omitempty"`
}

// NewGETDeliveryLeadTimesDeliveryLeadTimeId200Response instantiates a new GETDeliveryLeadTimesDeliveryLeadTimeId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETDeliveryLeadTimesDeliveryLeadTimeId200Response() *GETDeliveryLeadTimesDeliveryLeadTimeId200Response {
	this := GETDeliveryLeadTimesDeliveryLeadTimeId200Response{}
	return &this
}

// NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseWithDefaults instantiates a new GETDeliveryLeadTimesDeliveryLeadTimeId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETDeliveryLeadTimesDeliveryLeadTimeId200ResponseWithDefaults() *GETDeliveryLeadTimesDeliveryLeadTimeId200Response {
	this := GETDeliveryLeadTimesDeliveryLeadTimeId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) GetData() GETDeliveryLeadTimes200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) GetDataOk() (*GETDeliveryLeadTimes200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETDeliveryLeadTimes200ResponseDataInner and assigns it to the Data field.
func (o *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) SetData(v GETDeliveryLeadTimes200ResponseDataInner) {
	o.Data = &v
}

func (o GETDeliveryLeadTimesDeliveryLeadTimeId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response struct {
	value *GETDeliveryLeadTimesDeliveryLeadTimeId200Response
	isSet bool
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) Get() *GETDeliveryLeadTimesDeliveryLeadTimeId200Response {
	return v.value
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) Set(val *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response(val *GETDeliveryLeadTimesDeliveryLeadTimeId200Response) *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response {
	return &NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response{value: val, isSet: true}
}

func (v NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETDeliveryLeadTimesDeliveryLeadTimeId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
