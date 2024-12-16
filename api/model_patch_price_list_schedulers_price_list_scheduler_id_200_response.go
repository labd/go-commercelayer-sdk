/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.5.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHPriceListSchedulersPriceListSchedulerId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPriceListSchedulersPriceListSchedulerId200Response{}

// PATCHPriceListSchedulersPriceListSchedulerId200Response struct for PATCHPriceListSchedulersPriceListSchedulerId200Response
type PATCHPriceListSchedulersPriceListSchedulerId200Response struct {
	Data *PATCHPriceListSchedulersPriceListSchedulerId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPriceListSchedulersPriceListSchedulerId200Response instantiates a new PATCHPriceListSchedulersPriceListSchedulerId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPriceListSchedulersPriceListSchedulerId200Response() *PATCHPriceListSchedulersPriceListSchedulerId200Response {
	this := PATCHPriceListSchedulersPriceListSchedulerId200Response{}
	return &this
}

// NewPATCHPriceListSchedulersPriceListSchedulerId200ResponseWithDefaults instantiates a new PATCHPriceListSchedulersPriceListSchedulerId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPriceListSchedulersPriceListSchedulerId200ResponseWithDefaults() *PATCHPriceListSchedulersPriceListSchedulerId200Response {
	this := PATCHPriceListSchedulersPriceListSchedulerId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPriceListSchedulersPriceListSchedulerId200Response) GetData() PATCHPriceListSchedulersPriceListSchedulerId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHPriceListSchedulersPriceListSchedulerId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPriceListSchedulersPriceListSchedulerId200Response) GetDataOk() (*PATCHPriceListSchedulersPriceListSchedulerId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPriceListSchedulersPriceListSchedulerId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPriceListSchedulersPriceListSchedulerId200ResponseData and assigns it to the Data field.
func (o *PATCHPriceListSchedulersPriceListSchedulerId200Response) SetData(v PATCHPriceListSchedulersPriceListSchedulerId200ResponseData) {
	o.Data = &v
}

func (o PATCHPriceListSchedulersPriceListSchedulerId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPriceListSchedulersPriceListSchedulerId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHPriceListSchedulersPriceListSchedulerId200Response struct {
	value *PATCHPriceListSchedulersPriceListSchedulerId200Response
	isSet bool
}

func (v NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) Get() *PATCHPriceListSchedulersPriceListSchedulerId200Response {
	return v.value
}

func (v *NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) Set(val *PATCHPriceListSchedulersPriceListSchedulerId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPriceListSchedulersPriceListSchedulerId200Response(val *PATCHPriceListSchedulersPriceListSchedulerId200Response) *NullablePATCHPriceListSchedulersPriceListSchedulerId200Response {
	return &NullablePATCHPriceListSchedulersPriceListSchedulerId200Response{value: val, isSet: true}
}

func (v NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPriceListSchedulersPriceListSchedulerId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
