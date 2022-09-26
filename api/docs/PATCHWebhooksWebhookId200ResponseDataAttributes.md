# PATCHWebhooksWebhookId200ResponseDataAttributes

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

### NewPATCHWebhooksWebhookId200ResponseDataAttributes

`func NewPATCHWebhooksWebhookId200ResponseDataAttributes() *PATCHWebhooksWebhookId200ResponseDataAttributes`

NewPATCHWebhooksWebhookId200ResponseDataAttributes instantiates a new PATCHWebhooksWebhookId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHWebhooksWebhookId200ResponseDataAttributesWithDefaults

`func NewPATCHWebhooksWebhookId200ResponseDataAttributesWithDefaults() *PATCHWebhooksWebhookId200ResponseDataAttributes`

NewPATCHWebhooksWebhookId200ResponseDataAttributesWithDefaults instantiates a new PATCHWebhooksWebhookId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResources() []string`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResourcesOk() (*[]string, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetIncludeResources(v []string)`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### GetResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuit() bool`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuitOk() (*bool, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetResetCircuit(v bool)`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### GetReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


