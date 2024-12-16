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

// checks if the GETWebhooksWebhookId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETWebhooksWebhookId200ResponseDataAttributes{}

// GETWebhooksWebhookId200ResponseDataAttributes struct for GETWebhooksWebhookId200ResponseDataAttributes
type GETWebhooksWebhookId200ResponseDataAttributes struct {
	// Unique name for the webhook.
	Name interface{} `json:"name,omitempty"`
	// The identifier of the resource/event that will trigger the webhook.
	Topic interface{} `json:"topic,omitempty"`
	// URI where the webhook subscription should send the POST request when the event occurs.
	CallbackUrl interface{} `json:"callback_url,omitempty"`
	// List of related resources that should be included in the webhook body.
	IncludeResources interface{} `json:"include_resources,omitempty"`
	// Time at which this resource was disabled.
	DisabledAt interface{} `json:"disabled_at,omitempty"`
	// The circuit breaker state, by default it is 'closed'. It can become 'open' once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made.
	CircuitState interface{} `json:"circuit_state,omitempty"`
	// The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback.
	CircuitFailureCount interface{} `json:"circuit_failure_count,omitempty"`
	// The shared secret used to sign the external request payload.
	SharedSecret interface{} `json:"shared_secret,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETWebhooksWebhookId200ResponseDataAttributes instantiates a new GETWebhooksWebhookId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETWebhooksWebhookId200ResponseDataAttributes() *GETWebhooksWebhookId200ResponseDataAttributes {
	this := GETWebhooksWebhookId200ResponseDataAttributes{}
	return &this
}

// NewGETWebhooksWebhookId200ResponseDataAttributesWithDefaults instantiates a new GETWebhooksWebhookId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETWebhooksWebhookId200ResponseDataAttributesWithDefaults() *GETWebhooksWebhookId200ResponseDataAttributes {
	this := GETWebhooksWebhookId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetTopic() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetTopicOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return &o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasTopic() bool {
	if o != nil && IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given interface{} and assigns it to the Topic field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetTopic(v interface{}) {
	o.Topic = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return &o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCallbackUrl() bool {
	if o != nil && IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given interface{} and assigns it to the CallbackUrl field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrl(v interface{}) {
	o.CallbackUrl = v
}

// GetIncludeResources returns the IncludeResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetIncludeResources() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.IncludeResources
}

// GetIncludeResourcesOk returns a tuple with the IncludeResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetIncludeResourcesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.IncludeResources) {
		return nil, false
	}
	return &o.IncludeResources, true
}

// HasIncludeResources returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasIncludeResources() bool {
	if o != nil && IsNil(o.IncludeResources) {
		return true
	}

	return false
}

// SetIncludeResources gets a reference to the given interface{} and assigns it to the IncludeResources field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetIncludeResources(v interface{}) {
	o.IncludeResources = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetCircuitState returns the CircuitState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitState
}

// GetCircuitStateOk returns a tuple with the CircuitState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitState) {
		return nil, false
	}
	return &o.CircuitState, true
}

// HasCircuitState returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCircuitState() bool {
	if o != nil && IsNil(o.CircuitState) {
		return true
	}

	return false
}

// SetCircuitState gets a reference to the given interface{} and assigns it to the CircuitState field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitState(v interface{}) {
	o.CircuitState = v
}

// GetCircuitFailureCount returns the CircuitFailureCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitFailureCount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CircuitFailureCount
}

// GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CircuitFailureCount) {
		return nil, false
	}
	return &o.CircuitFailureCount, true
}

// HasCircuitFailureCount returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCircuitFailureCount() bool {
	if o != nil && IsNil(o.CircuitFailureCount) {
		return true
	}

	return false
}

// SetCircuitFailureCount gets a reference to the given interface{} and assigns it to the CircuitFailureCount field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitFailureCount(v interface{}) {
	o.CircuitFailureCount = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetSharedSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return &o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasSharedSecret() bool {
	if o != nil && IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given interface{} and assigns it to the SharedSecret field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetSharedSecret(v interface{}) {
	o.SharedSecret = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETWebhooksWebhookId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETWebhooksWebhookId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
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
	return toSerialize, nil
}

type NullableGETWebhooksWebhookId200ResponseDataAttributes struct {
	value *GETWebhooksWebhookId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETWebhooksWebhookId200ResponseDataAttributes) Get() *GETWebhooksWebhookId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETWebhooksWebhookId200ResponseDataAttributes) Set(val *GETWebhooksWebhookId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETWebhooksWebhookId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETWebhooksWebhookId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETWebhooksWebhookId200ResponseDataAttributes(val *GETWebhooksWebhookId200ResponseDataAttributes) *NullableGETWebhooksWebhookId200ResponseDataAttributes {
	return &NullableGETWebhooksWebhookId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETWebhooksWebhookId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETWebhooksWebhookId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
