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

// checks if the WebhookUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookUpdate{}

// WebhookUpdate struct for WebhookUpdate
type WebhookUpdate struct {
	Data WebhookUpdateData `json:"data"`
}

// NewWebhookUpdate instantiates a new WebhookUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUpdate(data WebhookUpdateData) *WebhookUpdate {
	this := WebhookUpdate{}
	this.Data = data
	return &this
}

// NewWebhookUpdateWithDefaults instantiates a new WebhookUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUpdateWithDefaults() *WebhookUpdate {
	this := WebhookUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *WebhookUpdate) GetData() WebhookUpdateData {
	if o == nil {
		var ret WebhookUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdate) GetDataOk() (*WebhookUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *WebhookUpdate) SetData(v WebhookUpdateData) {
	o.Data = v
}

func (o WebhookUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableWebhookUpdate struct {
	value *WebhookUpdate
	isSet bool
}

func (v NullableWebhookUpdate) Get() *WebhookUpdate {
	return v.value
}

func (v *NullableWebhookUpdate) Set(val *WebhookUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUpdate(val *WebhookUpdate) *NullableWebhookUpdate {
	return &NullableWebhookUpdate{value: val, isSet: true}
}

func (v NullableWebhookUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
