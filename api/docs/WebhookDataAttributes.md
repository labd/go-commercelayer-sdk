# WebhookDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **string** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **string** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **[]string** | List of related resources that should be included in the webhook body. | [optional] 
**CircuitState** | Pointer to **string** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **int32** | The number of consecutive failures recorded by the circuit breaker associated to this webhook, will be reset on first successful call to callback. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewWebhookDataAttributes

`func NewWebhookDataAttributes() *WebhookDataAttributes`

NewWebhookDataAttributes instantiates a new WebhookDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDataAttributesWithDefaults

`func NewWebhookDataAttributesWithDefaults() *WebhookDataAttributes`

NewWebhookDataAttributesWithDefaults instantiates a new WebhookDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookDataAttributes) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDataAttributes) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDataAttributes) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *WebhookDataAttributes) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *WebhookDataAttributes) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *WebhookDataAttributes) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *WebhookDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetIncludeResources

`func (o *WebhookDataAttributes) GetIncludeResources() []string`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *WebhookDataAttributes) GetIncludeResourcesOk() (*[]string, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *WebhookDataAttributes) SetIncludeResources(v []string)`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *WebhookDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### GetCircuitState

`func (o *WebhookDataAttributes) GetCircuitState() string`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *WebhookDataAttributes) GetCircuitStateOk() (*string, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *WebhookDataAttributes) SetCircuitState(v string)`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *WebhookDataAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### GetCircuitFailureCount

`func (o *WebhookDataAttributes) GetCircuitFailureCount() int32`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *WebhookDataAttributes) GetCircuitFailureCountOk() (*int32, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *WebhookDataAttributes) SetCircuitFailureCount(v int32)`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *WebhookDataAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### GetId

`func (o *WebhookDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *WebhookDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WebhookDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WebhookDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *WebhookDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *WebhookDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *WebhookDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *WebhookDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *WebhookDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *WebhookDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebhookDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


