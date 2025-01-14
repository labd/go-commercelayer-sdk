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

// checks if the NotificationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationCreate{}

// NotificationCreate struct for NotificationCreate
type NotificationCreate struct {
	Data NotificationCreateData `json:"data"`
}

// NewNotificationCreate instantiates a new NotificationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationCreate(data NotificationCreateData) *NotificationCreate {
	this := NotificationCreate{}
	this.Data = data
	return &this
}

// NewNotificationCreateWithDefaults instantiates a new NotificationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationCreateWithDefaults() *NotificationCreate {
	this := NotificationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *NotificationCreate) GetData() NotificationCreateData {
	if o == nil {
		var ret NotificationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NotificationCreate) GetDataOk() (*NotificationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NotificationCreate) SetData(v NotificationCreateData) {
	o.Data = v
}

func (o NotificationCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableNotificationCreate struct {
	value *NotificationCreate
	isSet bool
}

func (v NullableNotificationCreate) Get() *NotificationCreate {
	return v.value
}

func (v *NullableNotificationCreate) Set(val *NotificationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationCreate(val *NotificationCreate) *NullableNotificationCreate {
	return &NullableNotificationCreate{value: val, isSet: true}
}

func (v NullableNotificationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}