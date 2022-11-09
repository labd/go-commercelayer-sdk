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

// GETEventCallbacks200ResponseDataInnerAttributes struct for GETEventCallbacks200ResponseDataInnerAttributes
type GETEventCallbacks200ResponseDataInnerAttributes struct {
	// The URI of the callback, inherited by the associated webhook.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// The payload sent to the callback endpoint, including the event affected resource and the specified includes.
	Payload map[string]interface{} `json:"payload,omitempty"`
	// The HTTP response code of the callback endpoint.
	ResponseCode *string `json:"response_code,omitempty"`
	// The HTTP response message of the callback endpoint.
	ResponseMessage *string `json:"response_message,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETEventCallbacks200ResponseDataInnerAttributes instantiates a new GETEventCallbacks200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventCallbacks200ResponseDataInnerAttributes() *GETEventCallbacks200ResponseDataInnerAttributes {
	this := GETEventCallbacks200ResponseDataInnerAttributes{}
	return &this
}

// NewGETEventCallbacks200ResponseDataInnerAttributesWithDefaults instantiates a new GETEventCallbacks200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventCallbacks200ResponseDataInnerAttributesWithDefaults() *GETEventCallbacks200ResponseDataInnerAttributes {
	this := GETEventCallbacks200ResponseDataInnerAttributes{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetResponseCode() string {
	if o == nil || o.ResponseCode == nil {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetResponseCodeOk() (*string, bool) {
	if o == nil || o.ResponseCode == nil {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasResponseCode() bool {
	if o != nil && o.ResponseCode != nil {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetResponseMessage() string {
	if o == nil || o.ResponseMessage == nil {
		var ret string
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetResponseMessageOk() (*string, bool) {
	if o == nil || o.ResponseMessage == nil {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasResponseMessage() bool {
	if o != nil && o.ResponseMessage != nil {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given string and assigns it to the ResponseMessage field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetResponseMessage(v string) {
	o.ResponseMessage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETEventCallbacks200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETEventCallbacks200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.ResponseCode != nil {
		toSerialize["response_code"] = o.ResponseCode
	}
	if o.ResponseMessage != nil {
		toSerialize["response_message"] = o.ResponseMessage
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableGETEventCallbacks200ResponseDataInnerAttributes struct {
	value *GETEventCallbacks200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETEventCallbacks200ResponseDataInnerAttributes) Get() *GETEventCallbacks200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETEventCallbacks200ResponseDataInnerAttributes) Set(val *GETEventCallbacks200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventCallbacks200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventCallbacks200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventCallbacks200ResponseDataInnerAttributes(val *GETEventCallbacks200ResponseDataInnerAttributes) *NullableGETEventCallbacks200ResponseDataInnerAttributes {
	return &NullableGETEventCallbacks200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETEventCallbacks200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventCallbacks200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
