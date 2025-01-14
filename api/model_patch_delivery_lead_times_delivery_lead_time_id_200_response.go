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

// checks if the PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response{}

// PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response struct for PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response
type PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response struct {
	Data *PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData `json:"data,omitempty"`
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeId200Response instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeId200Response() *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response{}
	return &this
}

// NewPATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseWithDefaults instantiates a new PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseWithDefaults() *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response {
	this := PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) GetData() PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) GetDataOk() (*PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData and assigns it to the Data field.
func (o *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) SetData(v PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseData) {
	o.Data = &v
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response struct {
	value *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response
	isSet bool
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) Get() *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response {
	return v.value
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) Set(val *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response(val *PATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response {
	return &NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response{value: val, isSet: true}
}

func (v NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHDeliveryLeadTimesDeliveryLeadTimeId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
