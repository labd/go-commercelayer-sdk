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

// checks if the NotificationCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationCreateData{}

// NotificationCreateData struct for NotificationCreateData
type NotificationCreateData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    POSTNotifications201ResponseDataAttributes `json:"attributes"`
	Relationships *NotificationCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewNotificationCreateData instantiates a new NotificationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationCreateData(type_ interface{}, attributes POSTNotifications201ResponseDataAttributes) *NotificationCreateData {
	this := NotificationCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewNotificationCreateDataWithDefaults instantiates a new NotificationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationCreateDataWithDefaults() *NotificationCreateData {
	this := NotificationCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *NotificationCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *NotificationCreateData) GetAttributes() POSTNotifications201ResponseDataAttributes {
	if o == nil {
		var ret POSTNotifications201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *NotificationCreateData) GetAttributesOk() (*POSTNotifications201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *NotificationCreateData) SetAttributes(v POSTNotifications201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *NotificationCreateData) GetRelationships() NotificationCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret NotificationCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCreateData) GetRelationshipsOk() (*NotificationCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *NotificationCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given NotificationCreateDataRelationships and assigns it to the Relationships field.
func (o *NotificationCreateData) SetRelationships(v NotificationCreateDataRelationships) {
	o.Relationships = &v
}

func (o NotificationCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableNotificationCreateData struct {
	value *NotificationCreateData
	isSet bool
}

func (v NullableNotificationCreateData) Get() *NotificationCreateData {
	return v.value
}

func (v *NullableNotificationCreateData) Set(val *NotificationCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationCreateData(val *NotificationCreateData) *NullableNotificationCreateData {
	return &NullableNotificationCreateData{value: val, isSet: true}
}

func (v NullableNotificationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}