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

// checks if the PATCHPackagesPackageId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPackagesPackageId200Response{}

// PATCHPackagesPackageId200Response struct for PATCHPackagesPackageId200Response
type PATCHPackagesPackageId200Response struct {
	Data *PATCHPackagesPackageId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPackagesPackageId200Response instantiates a new PATCHPackagesPackageId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPackagesPackageId200Response() *PATCHPackagesPackageId200Response {
	this := PATCHPackagesPackageId200Response{}
	return &this
}

// NewPATCHPackagesPackageId200ResponseWithDefaults instantiates a new PATCHPackagesPackageId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPackagesPackageId200ResponseWithDefaults() *PATCHPackagesPackageId200Response {
	this := PATCHPackagesPackageId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPackagesPackageId200Response) GetData() PATCHPackagesPackageId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHPackagesPackageId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPackagesPackageId200Response) GetDataOk() (*PATCHPackagesPackageId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPackagesPackageId200ResponseData and assigns it to the Data field.
func (o *PATCHPackagesPackageId200Response) SetData(v PATCHPackagesPackageId200ResponseData) {
	o.Data = &v
}

func (o PATCHPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPackagesPackageId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHPackagesPackageId200Response struct {
	value *PATCHPackagesPackageId200Response
	isSet bool
}

func (v NullablePATCHPackagesPackageId200Response) Get() *PATCHPackagesPackageId200Response {
	return v.value
}

func (v *NullablePATCHPackagesPackageId200Response) Set(val *PATCHPackagesPackageId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPackagesPackageId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPackagesPackageId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPackagesPackageId200Response(val *PATCHPackagesPackageId200Response) *NullablePATCHPackagesPackageId200Response {
	return &NullablePATCHPackagesPackageId200Response{value: val, isSet: true}
}

func (v NullablePATCHPackagesPackageId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPackagesPackageId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
