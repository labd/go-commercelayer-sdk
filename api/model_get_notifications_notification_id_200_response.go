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

// checks if the GETNotificationsNotificationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETNotificationsNotificationId200Response{}

// GETNotificationsNotificationId200Response struct for GETNotificationsNotificationId200Response
type GETNotificationsNotificationId200Response struct {
	Data *GETNotificationsNotificationId200ResponseData `json:"data,omitempty"`
}

// NewGETNotificationsNotificationId200Response instantiates a new GETNotificationsNotificationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETNotificationsNotificationId200Response() *GETNotificationsNotificationId200Response {
	this := GETNotificationsNotificationId200Response{}
	return &this
}

// NewGETNotificationsNotificationId200ResponseWithDefaults instantiates a new GETNotificationsNotificationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETNotificationsNotificationId200ResponseWithDefaults() *GETNotificationsNotificationId200Response {
	this := GETNotificationsNotificationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETNotificationsNotificationId200Response) GetData() GETNotificationsNotificationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETNotificationsNotificationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETNotificationsNotificationId200Response) GetDataOk() (*GETNotificationsNotificationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETNotificationsNotificationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETNotificationsNotificationId200ResponseData and assigns it to the Data field.
func (o *GETNotificationsNotificationId200Response) SetData(v GETNotificationsNotificationId200ResponseData) {
	o.Data = &v
}

func (o GETNotificationsNotificationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETNotificationsNotificationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETNotificationsNotificationId200Response struct {
	value *GETNotificationsNotificationId200Response
	isSet bool
}

func (v NullableGETNotificationsNotificationId200Response) Get() *GETNotificationsNotificationId200Response {
	return v.value
}

func (v *NullableGETNotificationsNotificationId200Response) Set(val *GETNotificationsNotificationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETNotificationsNotificationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETNotificationsNotificationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETNotificationsNotificationId200Response(val *GETNotificationsNotificationId200Response) *NullableGETNotificationsNotificationId200Response {
	return &NullableGETNotificationsNotificationId200Response{value: val, isSet: true}
}

func (v NullableGETNotificationsNotificationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETNotificationsNotificationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
