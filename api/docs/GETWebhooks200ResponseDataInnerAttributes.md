# GETWebhooks200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **interface{}** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **interface{}** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **[]interface{}** | List of related resources that should be included in the webhook body. | [optional] 
**CircuitState** | Pointer to **interface{}** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **interface{}** | The number of consecutive failures recorded by the circuit breaker associated to this webhook, will be reset on first successful call to callback. | [optional] 
**SharedSecret** | Pointer to **interface{}** | The shared secret used to sign the external request payload. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETWebhooks200ResponseDataInnerAttributes

`func NewGETWebhooks200ResponseDataInnerAttributes() *GETWebhooks200ResponseDataInnerAttributes`

NewGETWebhooks200ResponseDataInnerAttributes instantiates a new GETWebhooks200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETWebhooks200ResponseDataInnerAttributesWithDefaults

`func NewGETWebhooks200ResponseDataInnerAttributesWithDefaults() *GETWebhooks200ResponseDataInnerAttributes`

NewGETWebhooks200ResponseDataInnerAttributesWithDefaults instantiates a new GETWebhooks200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTopic

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetTopic(v interface{})`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCallbackUrl

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCallbackUrl() interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCallbackUrlOk() (*interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCallbackUrl(v interface{})`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetIncludeResources

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetIncludeResources() []interface{}`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetIncludeResourcesOk() (*[]interface{}, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetIncludeResources(v []interface{})`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### GetCircuitState

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitState() interface{}`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitStateOk() (*interface{}, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitState(v interface{})`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### SetCircuitStateNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitStateNil(b bool)`

 SetCircuitStateNil sets the value for CircuitState to be an explicit nil

### UnsetCircuitState
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetCircuitState()`

UnsetCircuitState ensures that no value is present for CircuitState, not even an explicit nil
### GetCircuitFailureCount

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitFailureCount() interface{}`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCircuitFailureCountOk() (*interface{}, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitFailureCount(v interface{})`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### SetCircuitFailureCountNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCircuitFailureCountNil(b bool)`

 SetCircuitFailureCountNil sets the value for CircuitFailureCount to be an explicit nil

### UnsetCircuitFailureCount
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetCircuitFailureCount()`

UnsetCircuitFailureCount ensures that no value is present for CircuitFailureCount, not even an explicit nil
### GetSharedSecret

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetSharedSecret() interface{}`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetSharedSecretOk() (*interface{}, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetSharedSecret(v interface{})`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil
### GetCreatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETWebhooks200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETWebhooks200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETWebhooks200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETWebhooks200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


