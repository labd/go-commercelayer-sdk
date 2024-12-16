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

// checks if the PATCHOrderSubscriptionsOrderSubscriptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHOrderSubscriptionsOrderSubscriptionId200Response{}

// PATCHOrderSubscriptionsOrderSubscriptionId200Response struct for PATCHOrderSubscriptionsOrderSubscriptionId200Response
type PATCHOrderSubscriptionsOrderSubscriptionId200Response struct {
	Data *PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHOrderSubscriptionsOrderSubscriptionId200Response instantiates a new PATCHOrderSubscriptionsOrderSubscriptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHOrderSubscriptionsOrderSubscriptionId200Response() *PATCHOrderSubscriptionsOrderSubscriptionId200Response {
	this := PATCHOrderSubscriptionsOrderSubscriptionId200Response{}
	return &this
}

// NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseWithDefaults instantiates a new PATCHOrderSubscriptionsOrderSubscriptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHOrderSubscriptionsOrderSubscriptionId200ResponseWithDefaults() *PATCHOrderSubscriptionsOrderSubscriptionId200Response {
	this := PATCHOrderSubscriptionsOrderSubscriptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200Response) GetData() PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200Response) GetDataOk() (*PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData and assigns it to the Data field.
func (o *PATCHOrderSubscriptionsOrderSubscriptionId200Response) SetData(v PATCHOrderSubscriptionsOrderSubscriptionId200ResponseData) {
	o.Data = &v
}

func (o PATCHOrderSubscriptionsOrderSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHOrderSubscriptionsOrderSubscriptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response struct {
	value *PATCHOrderSubscriptionsOrderSubscriptionId200Response
	isSet bool
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) Get() *PATCHOrderSubscriptionsOrderSubscriptionId200Response {
	return v.value
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) Set(val *PATCHOrderSubscriptionsOrderSubscriptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHOrderSubscriptionsOrderSubscriptionId200Response(val *PATCHOrderSubscriptionsOrderSubscriptionId200Response) *NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response {
	return &NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response{value: val, isSet: true}
}

func (v NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHOrderSubscriptionsOrderSubscriptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
