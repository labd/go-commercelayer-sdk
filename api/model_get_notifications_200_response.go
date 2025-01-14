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

// checks if the GETNotifications200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETNotifications200Response{}

// GETNotifications200Response struct for GETNotifications200Response
type GETNotifications200Response struct {
	Data interface{} `json:"data,omitempty"`
}

// NewGETNotifications200Response instantiates a new GETNotifications200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETNotifications200Response() *GETNotifications200Response {
	this := GETNotifications200Response{}
	return &this
}

// NewGETNotifications200ResponseWithDefaults instantiates a new GETNotifications200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETNotifications200ResponseWithDefaults() *GETNotifications200Response {
	this := GETNotifications200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETNotifications200Response) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETNotifications200Response) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETNotifications200Response) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *GETNotifications200Response) SetData(v interface{}) {
	o.Data = v
}

func (o GETNotifications200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETNotifications200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETNotifications200Response struct {
	value *GETNotifications200Response
	isSet bool
}

func (v NullableGETNotifications200Response) Get() *GETNotifications200Response {
	return v.value
}

func (v *NullableGETNotifications200Response) Set(val *GETNotifications200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETNotifications200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETNotifications200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETNotifications200Response(val *GETNotifications200Response) *NullableGETNotifications200Response {
	return &NullableGETNotifications200Response{value: val, isSet: true}
}

func (v NullableGETNotifications200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETNotifications200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}