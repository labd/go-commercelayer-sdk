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

// checks if the PATCHCustomerSubscriptionsCustomerSubscriptionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerSubscriptionsCustomerSubscriptionId200Response{}

// PATCHCustomerSubscriptionsCustomerSubscriptionId200Response struct for PATCHCustomerSubscriptionsCustomerSubscriptionId200Response
type PATCHCustomerSubscriptionsCustomerSubscriptionId200Response struct {
	Data *PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHCustomerSubscriptionsCustomerSubscriptionId200Response instantiates a new PATCHCustomerSubscriptionsCustomerSubscriptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerSubscriptionsCustomerSubscriptionId200Response() *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response {
	this := PATCHCustomerSubscriptionsCustomerSubscriptionId200Response{}
	return &this
}

// NewPATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseWithDefaults instantiates a new PATCHCustomerSubscriptionsCustomerSubscriptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseWithDefaults() *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response {
	this := PATCHCustomerSubscriptionsCustomerSubscriptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) GetData() PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) GetDataOk() (*PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData and assigns it to the Data field.
func (o *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) SetData(v PATCHCustomerSubscriptionsCustomerSubscriptionId200ResponseData) {
	o.Data = &v
}

func (o PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response struct {
	value *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response
	isSet bool
}

func (v NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) Get() *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response {
	return v.value
}

func (v *NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) Set(val *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response(val *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response) *NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response {
	return &NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response{value: val, isSet: true}
}

func (v NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerSubscriptionsCustomerSubscriptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
