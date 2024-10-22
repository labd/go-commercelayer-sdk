# PATCHWebhooksWebhookId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the webhook. | [optional] 
**Topic** | Pointer to **interface{}** | The identifier of the resource/event that will trigger the webhook. | [optional] 
**CallbackUrl** | Pointer to **interface{}** | URI where the webhook subscription should send the POST request when the event occurs. | [optional] 
**IncludeResources** | Pointer to **interface{}** | List of related resources that should be included in the webhook body. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**ResetCircuit** | Pointer to **interface{}** | Send this attribute if you want to reset the circuit breaker associated to this resource to &#39;closed&#39; state and zero failures count. Cannot be passed by sales channels. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetTopic(v interface{})`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrl() interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrlOk() (*interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrl(v interface{})`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResources() interface{}`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResourcesOk() (*interface{}, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetIncludeResources(v interface{})`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### SetIncludeResourcesNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetIncludeResourcesNil(b bool)`

 SetIncludeResourcesNil sets the value for IncludeResources to be an explicit nil

### UnsetIncludeResources
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetIncludeResources()`

UnsetIncludeResources ensures that no value is present for IncludeResources, not even an explicit nil
### GetDisable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuit() interface{}`

GetResetCircuit returns the ResetCircuit field if non-nil, zero value otherwise.

### GetResetCircuitOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuitOk() (*interface{}, bool)`

GetResetCircuitOk returns a tuple with the ResetCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetResetCircuit(v interface{})`

SetResetCircuit sets ResetCircuit field to given value.

### HasResetCircuit

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasResetCircuit() bool`

HasResetCircuit returns a boolean if a field has been set.

### SetResetCircuitNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetResetCircuitNil(b bool)`

 SetResetCircuitNil sets the value for ResetCircuit to be an explicit nil

### UnsetResetCircuit
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetResetCircuit()`

UnsetResetCircuit ensures that no value is present for ResetCircuit, not even an explicit nil
### GetReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


