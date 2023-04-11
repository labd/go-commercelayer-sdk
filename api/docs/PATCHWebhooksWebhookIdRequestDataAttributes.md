# PATCHWebhooksWebhookIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **interface{}** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **interface{}** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **interface{}** | List of related resources that should be included in the webhook body. | [optional] 
**ResetCircuit** | Pointer to **interface{}** | Send this attribute if you want to reset the circuit breaker associated to this webhook to &#39;closed&#39; state and zero failures count. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHWebhooksWebhookIdRequestDataAttributes

`func NewPATCHWebhooksWebhookIdRequestDataAttributes() *PATCHWebhooksWebhookIdRequestDataAttributes`

NewPATCHWebhooksWebhookIdRequestDataAttributes instantiates a new PATCHWebhooksWebhookIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHWebhooksWebhookIdRequestDataAttributesWithDefaults

`func NewPATCHWebhooksWebhookIdRequestDataAttributesWithDefaults() *PATCHWebhooksWebhookIdRequestDataAttributes`

NewPATCHWebhooksWebhookIdRequestDataAttributesWithDefaults instantiates a new PATCHWebhooksWebhookIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTopic

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetTopic(v interface{})`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCallbackUrl

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetCallbackUrl() interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetCallbackUrlOk() (*interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetCallbackUrl(v interface{})`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetIncludeResources

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetIncludeResources() interface{}`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetIncludeResourcesOk() (*interface{}, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetIncludeResources(v interface{})`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### SetIncludeResourcesNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetIncludeResourcesNil(b bool)`

 SetIncludeResourcesNil sets the value for IncludeResources to be an explicit nil

### UnsetIncludeResources
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetIncludeResources()`

UnsetIncludeResources ensures that no value is present for IncludeResources, not even an explicit nil
### GetResetCircuit

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetResetCircuit() interface{}`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetResetCircuitOk() (*interface{}, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetResetCircuit(v interface{})`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### SetResetCircuitNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetResetCircuitNil(b bool)`

 SetResetCircuitNil sets the value for ResetCircuit to be an explicit nil

### UnsetResetCircuit
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetResetCircuit()`

UnsetResetCircuit ensures that no value is present for ResetCircuit, not even an explicit nil
### GetReference

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHWebhooksWebhookIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


