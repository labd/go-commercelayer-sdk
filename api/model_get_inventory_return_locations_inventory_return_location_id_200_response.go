/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETInventoryReturnLocationsInventoryReturnLocationId200Response struct for GETInventoryReturnLocationsInventoryReturnLocationId200Response
type GETInventoryReturnLocationsInventoryReturnLocationId200Response struct {
	Data *GETInventoryReturnLocations200ResponseDataInner `json:"data,omitempty"`
}

// NewGETInventoryReturnLocationsInventoryReturnLocationId200Response instantiates a new GETInventoryReturnLocationsInventoryReturnLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryReturnLocationsInventoryReturnLocationId200Response() *GETInventoryReturnLocationsInventoryReturnLocationId200Response {
	this := GETInventoryReturnLocationsInventoryReturnLocationId200Response{}
	return &this
}

// NewGETInventoryReturnLocationsInventoryReturnLocationId200ResponseWithDefaults instantiates a new GETInventoryReturnLocationsInventoryReturnLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryReturnLocationsInventoryReturnLocationId200ResponseWithDefaults() *GETInventoryReturnLocationsInventoryReturnLocationId200Response {
	this := GETInventoryReturnLocationsInventoryReturnLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryReturnLocationsInventoryReturnLocationId200Response) GetData() GETInventoryReturnLocations200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETInventoryReturnLocations200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocationsInventoryReturnLocationId200Response) GetDataOk() (*GETInventoryReturnLocations200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryReturnLocationsInventoryReturnLocationId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInventoryReturnLocations200ResponseDataInner and assigns it to the Data field.
func (o *GETInventoryReturnLocationsInventoryReturnLocationId200Response) SetData(v GETInventoryReturnLocations200ResponseDataInner) {
	o.Data = &v
}

func (o GETInventoryReturnLocationsInventoryReturnLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response struct {
	value *GETInventoryReturnLocationsInventoryReturnLocationId200Response
	isSet bool
}

func (v NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) Get() *GETInventoryReturnLocationsInventoryReturnLocationId200Response {
	return v.value
}

func (v *NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) Set(val *GETInventoryReturnLocationsInventoryReturnLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryReturnLocationsInventoryReturnLocationId200Response(val *GETInventoryReturnLocationsInventoryReturnLocationId200Response) *NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response {
	return &NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response{value: val, isSet: true}
}

func (v NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryReturnLocationsInventoryReturnLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
