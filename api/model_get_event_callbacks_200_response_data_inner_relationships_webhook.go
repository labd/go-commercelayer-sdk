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

// GETEventCallbacks200ResponseDataInnerRelationshipsWebhook struct for GETEventCallbacks200ResponseDataInnerRelationshipsWebhook
type GETEventCallbacks200ResponseDataInnerRelationshipsWebhook struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks    `json:"links,omitempty"`
	Data  *GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData `json:"data,omitempty"`
}

// NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhook instantiates a new GETEventCallbacks200ResponseDataInnerRelationshipsWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhook() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook {
	this := GETEventCallbacks200ResponseDataInnerRelationshipsWebhook{}
	return &this
}

// NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookWithDefaults instantiates a new GETEventCallbacks200ResponseDataInnerRelationshipsWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacks200ResponseDataInnerRelationshipsWebhookWithDefaults() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook {
	this := GETEventCallbacks200ResponseDataInnerRelationshipsWebhook{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) GetData() GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData {
	if o == nil || o.Data == nil {
		var ret GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) GetDataOk() (*GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData and assigns it to the Data field.
func (o *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) SetData(v GETEventCallbacks200ResponseDataInnerRelationshipsWebhookData) {
	o.Data = &v
}

func (o GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook struct {
	value *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook
	isSet bool
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) Get() *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook {
	return v.value
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) Set(val *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook(val *GETEventCallbacks200ResponseDataInnerRelationshipsWebhook) *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook {
	return &NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook{value: val, isSet: true}
}

func (v NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacks200ResponseDataInnerRelationshipsWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
