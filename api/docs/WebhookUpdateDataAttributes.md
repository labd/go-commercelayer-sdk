# WebhookUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **string** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **string** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **[]string** | List of related resources that should be included in the webhook body. | [optional] 
**ResetCircuit** | Pointer to **bool** | Send this attribute if you want to reset the circuit breaker associated to this webhook to &#39;closed&#39; state and zero failures count. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewWebhookUpdateDataAttributes

`func NewWebhookUpdateDataAttributes() *WebhookUpdateDataAttributes`

NewWebhookUpdateDataAttributes instantiates a new WebhookUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateDataAttributesWithDefaults

`func NewWebhookUpdateDataAttributesWithDefaults() *WebhookUpdateDataAttributes`

NewWebhookUpdateDataAttributesWithDefaults instantiates a new WebhookUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookUpdateDataAttributes) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookUpdateDataAttributes) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookUpdateDataAttributes) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookUpdateDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *WebhookUpdateDataAttributes) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *WebhookUpdateDataAttributes) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *WebhookUpdateDataAttributes) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *WebhookUpdateDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetIncludeResources

`func (o *WebhookUpdateDataAttributes) GetIncludeResources() []string`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *WebhookUpdateDataAttributes) GetIncludeResourcesOk() (*[]string, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *WebhookUpdateDataAttributes) SetIncludeResources(v []string)`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *WebhookUpdateDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### GetResetCircuit

`func (o *WebhookUpdateDataAttributes) GetResetCircuit() bool`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *WebhookUpdateDataAttributes) GetResetCircuitOk() (*bool, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *WebhookUpdateDataAttributes) SetResetCircuit(v bool)`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *WebhookUpdateDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### GetReference

`func (o *WebhookUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WebhookUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WebhookUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *WebhookUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *WebhookUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *WebhookUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *WebhookUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *WebhookUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *WebhookUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebhookUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


