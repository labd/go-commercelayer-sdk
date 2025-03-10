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

// checks if the PATCHNotificationsNotificationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHNotificationsNotificationId200Response{}

// PATCHNotificationsNotificationId200Response struct for PATCHNotificationsNotificationId200Response
type PATCHNotificationsNotificationId200Response struct {
	Data *PATCHNotificationsNotificationId200ResponseData `json:"data,omitempty"`
}

// NewPATCHNotificationsNotificationId200Response instantiates a new PATCHNotificationsNotificationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHNotificationsNotificationId200Response() *PATCHNotificationsNotificationId200Response {
	this := PATCHNotificationsNotificationId200Response{}
	return &this
}

// NewPATCHNotificationsNotificationId200ResponseWithDefaults instantiates a new PATCHNotificationsNotificationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHNotificationsNotificationId200ResponseWithDefaults() *PATCHNotificationsNotificationId200Response {
	this := PATCHNotificationsNotificationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHNotificationsNotificationId200Response) GetData() PATCHNotificationsNotificationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHNotificationsNotificationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHNotificationsNotificationId200Response) GetDataOk() (*PATCHNotificationsNotificationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHNotificationsNotificationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHNotificationsNotificationId200ResponseData and assigns it to the Data field.
func (o *PATCHNotificationsNotificationId200Response) SetData(v PATCHNotificationsNotificationId200ResponseData) {
	o.Data = &v
}

func (o PATCHNotificationsNotificationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHNotificationsNotificationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHNotificationsNotificationId200Response struct {
	value *PATCHNotificationsNotificationId200Response
	isSet bool
}

func (v NullablePATCHNotificationsNotificationId200Response) Get() *PATCHNotificationsNotificationId200Response {
	return v.value
}

func (v *NullablePATCHNotificationsNotificationId200Response) Set(val *PATCHNotificationsNotificationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHNotificationsNotificationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHNotificationsNotificationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHNotificationsNotificationId200Response(val *PATCHNotificationsNotificationId200Response) *NullablePATCHNotificationsNotificationId200Response {
	return &NullablePATCHNotificationsNotificationId200Response{value: val, isSet: true}
}

func (v NullablePATCHNotificationsNotificationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHNotificationsNotificationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
