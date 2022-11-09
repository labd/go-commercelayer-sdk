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

// GETWebhooksWebhookId200Response struct for GETWebhooksWebhookId200Response
type GETWebhooksWebhookId200Response struct {
	Data *GETWebhooks200ResponseDataInner `json:"data,omitempty"`
}

// NewGETWebhooksWebhookId200Response instantiates a new GETWebhooksWebhookId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWebhooksWebhookId200Response() *GETWebhooksWebhookId200Response {
	this := GETWebhooksWebhookId200Response{}
	return &this
}

// NewGETWebhooksWebhookId200ResponseWithDefaults instantiates a new GETWebhooksWebhookId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWebhooksWebhookId200ResponseWithDefaults() *GETWebhooksWebhookId200Response {
	this := GETWebhooksWebhookId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETWebhooksWebhookId200Response) GetData() GETWebhooks200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETWebhooks200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooksWebhookId200Response) GetDataOk() (*GETWebhooks200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETWebhooks200ResponseDataInner and assigns it to the Data field.
func (o *GETWebhooksWebhookId200Response) SetData(v GETWebhooks200ResponseDataInner) {
	o.Data = &v
}

func (o GETWebhooksWebhookId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETWebhooksWebhookId200Response struct {
	value *GETWebhooksWebhookId200Response
	isSet bool
}

func (v NullableGETWebhooksWebhookId200Response) Get() *GETWebhooksWebhookId200Response {
	return v.value
}

func (v *NullableGETWebhooksWebhookId200Response) Set(val *GETWebhooksWebhookId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWebhooksWebhookId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWebhooksWebhookId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWebhooksWebhookId200Response(val *GETWebhooksWebhookId200Response) *NullableGETWebhooksWebhookId200Response {
	return &NullableGETWebhooksWebhookId200Response{value: val, isSet: true}
}

func (v NullableGETWebhooksWebhookId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWebhooksWebhookId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
