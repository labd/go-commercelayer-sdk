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

// GETWebhooks200ResponseDataInnerAttributes struct for GETWebhooks200ResponseDataInnerAttributes
type GETWebhooks200ResponseDataInnerAttributes struct {
	// Unique name for the webhook.
	Name *string `json:"name,omitempty"`
	// The identifier of the resource/event that will trigger the webhook.
	Topic *string `json:"topic,omitempty"`
	// URI where the webhook subscription should send the POST request when the event occurs.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// List of related resources that should be included in the webhook body.
	IncludeResources []string `json:"include_resources,omitempty"`
	// The circuit breaker state, by default it is 'closed'. It can become 'open' once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made.
	CircuitState *string `json:"circuit_state,omitempty"`
	// The number of consecutive failures recorded by the circuit breaker associated to this webhook, will be reset on first successful call to callback.
	CircuitFailureCount *int32 `json:"circuit_failure_count,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret *string `json:"shared_secret,omitempty"`
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

// NewGETWebhooks200ResponseDataInnerAttributes instantiates a new GETWebhooks200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWebhooks200ResponseDataInnerAttributes() *GETWebhooks200ResponseDataInnerAttributes {
	this := GETWebhooks200ResponseDataInnerAttributes{}
	return &this
}

// NewGETWebhooks200ResponseDataInnerAttributesWithDefaults instantiates a new GETWebhooks200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWebhooks200ResponseDataInnerAttributesWithDefaults() *GETWebhooks200ResponseDataInnerAttributes {
	this := GETWebhooks200ResponseDataInnerAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetTopic(v string) {
	o.Topic = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetIncludeResources returns the IncludeResources field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetIncludeResources() []string {
	if o == nil || o.IncludeResources == nil {
		var ret []string
		return ret
	}
	return o.IncludeResources
}

// GetIncludeResourcesOk returns a tuple with the IncludeResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetIncludeResourcesOk() ([]string, bool) {
	if o == nil || o.IncludeResources == nil {
		return nil, false
	}
	return o.IncludeResources, true
}

// HasIncludeResources returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasIncludeResources() bool {
	if o != nil && o.IncludeResources != nil {
		return true
	}

	return false
}

// SetIncludeResources gets a reference to the given []string and assigns it to the IncludeResources field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetIncludeResources(v []string) {
	o.IncludeResources = v
}

// GetCircuitState returns the CircuitState field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitState() string {
	if o == nil || o.CircuitState == nil {
		var ret string
		return ret
	}
	return *o.CircuitState
}

// GetCircuitStateOk returns a tuple with the CircuitState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitStateOk() (*string, bool) {
	if o == nil || o.CircuitState == nil {
		return nil, false
	}
	return o.CircuitState, true
}

// HasCircuitState returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasCircuitState() bool {
	if o != nil && o.CircuitState != nil {
		return true
	}

	return false
}

// SetCircuitState gets a reference to the given string and assigns it to the CircuitState field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitState(v string) {
	o.CircuitState = &v
}

// GetCircuitFailureCount returns the CircuitFailureCount field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitFailureCount() int32 {
	if o == nil || o.CircuitFailureCount == nil {
		var ret int32
		return ret
	}
	return *o.CircuitFailureCount
}

// GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitFailureCountOk() (*int32, bool) {
	if o == nil || o.CircuitFailureCount == nil {
		return nil, false
	}
	return o.CircuitFailureCount, true
}

// HasCircuitFailureCount returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasCircuitFailureCount() bool {
	if o != nil && o.CircuitFailureCount != nil {
		return true
	}

	return false
}

// SetCircuitFailureCount gets a reference to the given int32 and assigns it to the CircuitFailureCount field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitFailureCount(v int32) {
	o.CircuitFailureCount = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETWebhooks200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETWebhooks200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETWebhooks200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.IncludeResources != nil {
		toSerialize["include_resources"] = o.IncludeResources
	}
	if o.CircuitState != nil {
		toSerialize["circuit_state"] = o.CircuitState
	}
	if o.CircuitFailureCount != nil {
		toSerialize["circuit_failure_count"] = o.CircuitFailureCount
	}
	if o.SharedSecret != nil {
		toSerialize["shared_secret"] = o.SharedSecret
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

type NullableGETWebhooks200ResponseDataInnerAttributes struct {
	value *GETWebhooks200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETWebhooks200ResponseDataInnerAttributes) Get() *GETWebhooks200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETWebhooks200ResponseDataInnerAttributes) Set(val *GETWebhooks200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWebhooks200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWebhooks200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWebhooks200ResponseDataInnerAttributes(val *GETWebhooks200ResponseDataInnerAttributes) *NullableGETWebhooks200ResponseDataInnerAttributes {
	return &NullableGETWebhooks200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETWebhooks200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWebhooks200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
