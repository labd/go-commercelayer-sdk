# GETWebhooksWebhookId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **interface{}** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **interface{}** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **interface{}** | List of related resources that should be included in the webhook body. | [optional] 
**CircuitState** | Pointer to **interface{}** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **interface{}** | The number of consecutive failures recorded by the circuit breaker associated to this webhook, will be reset on first successful call to callback. | [optional] 
**SharedSecret** | Pointer to **interface{}** | The shared secret used to sign the external request payload. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETWebhooksWebhookId200ResponseDataAttributes

`func NewGETWebhooksWebhookId200ResponseDataAttributes() *GETWebhooksWebhookId200ResponseDataAttributes`

NewGETWebhooksWebhookId200ResponseDataAttributes instantiates a new GETWebhooksWebhookId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETWebhooksWebhookId200ResponseDataAttributesWithDefaults

`func NewGETWebhooksWebhookId200ResponseDataAttributesWithDefaults() *GETWebhooksWebhookId200ResponseDataAttributes`

NewGETWebhooksWebhookId200ResponseDataAttributesWithDefaults instantiates a new GETWebhooksWebhookId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTopic

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetTopic(v interface{})`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCallbackUrl

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrl() interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrlOk() (*interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrl(v interface{})`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetIncludeResources

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetIncludeResources() interface{}`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetIncludeResourcesOk() (*interface{}, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetIncludeResources(v interface{})`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### SetIncludeResourcesNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetIncludeResourcesNil(b bool)`

 SetIncludeResourcesNil sets the value for IncludeResources to be an explicit nil

### UnsetIncludeResources
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetIncludeResources()`

UnsetIncludeResources ensures that no value is present for IncludeResources, not even an explicit nil
### GetCircuitState

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitState() interface{}`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitState(v interface{})`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### SetCircuitStateNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitStateNil(b bool)`

 SetCircuitStateNil sets the value for CircuitState to be an explicit nil

### UnsetCircuitState
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetCircuitState()`

UnsetCircuitState ensures that no value is present for CircuitState, not even an explicit nil
### GetCircuitFailureCount

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitFailureCount() interface{}`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitFailureCount(v interface{})`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### SetCircuitFailureCountNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCircuitFailureCountNil(b bool)`

 SetCircuitFailureCountNil sets the value for CircuitFailureCount to be an explicit nil

### UnsetCircuitFailureCount
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetCircuitFailureCount()`

UnsetCircuitFailureCount ensures that no value is present for CircuitFailureCount, not even an explicit nil
### GetSharedSecret

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetSharedSecret() interface{}`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetSharedSecretOk() (*interface{}, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetSharedSecret(v interface{})`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil
### GetCreatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETWebhooksWebhookId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETWebhooksWebhookId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


