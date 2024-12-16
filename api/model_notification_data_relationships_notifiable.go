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

// checks if the NotificationDataRelationshipsNotifiable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationDataRelationshipsNotifiable{}

// NotificationDataRelationshipsNotifiable struct for NotificationDataRelationshipsNotifiable
type NotificationDataRelationshipsNotifiable struct {
	Data *NotificationDataRelationshipsNotifiableData `json:"data,omitempty"`
}

// NewNotificationDataRelationshipsNotifiable instantiates a new NotificationDataRelationshipsNotifiable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDataRelationshipsNotifiable() *NotificationDataRelationshipsNotifiable {
	this := NotificationDataRelationshipsNotifiable{}
	return &this
}

// NewNotificationDataRelationshipsNotifiableWithDefaults instantiates a new NotificationDataRelationshipsNotifiable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDataRelationshipsNotifiableWithDefaults() *NotificationDataRelationshipsNotifiable {
	this := NotificationDataRelationshipsNotifiable{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NotificationDataRelationshipsNotifiable) GetData() NotificationDataRelationshipsNotifiableData {
	if o == nil || IsNil(o.Data) {
		var ret NotificationDataRelationshipsNotifiableData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDataRelationshipsNotifiable) GetDataOk() (*NotificationDataRelationshipsNotifiableData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NotificationDataRelationshipsNotifiable) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given NotificationDataRelationshipsNotifiableData and assigns it to the Data field.
func (o *NotificationDataRelationshipsNotifiable) SetData(v NotificationDataRelationshipsNotifiableData) {
	o.Data = &v
}

func (o NotificationDataRelationshipsNotifiable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationDataRelationshipsNotifiable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableNotificationDataRelationshipsNotifiable struct {
	value *NotificationDataRelationshipsNotifiable
	isSet bool
}

func (v NullableNotificationDataRelationshipsNotifiable) Get() *NotificationDataRelationshipsNotifiable {
	return v.value
}

func (v *NullableNotificationDataRelationshipsNotifiable) Set(val *NotificationDataRelationshipsNotifiable) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDataRelationshipsNotifiable) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDataRelationshipsNotifiable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDataRelationshipsNotifiable(val *NotificationDataRelationshipsNotifiable) *NullableNotificationDataRelationshipsNotifiable {
	return &NullableNotificationDataRelationshipsNotifiable{value: val, isSet: true}
}

func (v NullableNotificationDataRelationshipsNotifiable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDataRelationshipsNotifiable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
