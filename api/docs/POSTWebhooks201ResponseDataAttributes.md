# POSTWebhooks201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | Unique name for the webhook. | [optional] 
**Topic** | **interface{}** | The identifier of the resource/event that will trigger the webhook. | 
**CallbackUrl** | **interface{}** | URI where the webhook subscription should send the POST request when the event occurs. | 
**IncludeResources** | Pointer to **interface{}** | List of related resources that should be included in the webhook body. | [optional] 
**Disable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as disabled. | [optional] 
**Enable** | Pointer to **interface{}** | Send this attribute if you want to mark this resource as enabled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTWebhooks201ResponseDataAttributes

`func NewPOSTWebhooks201ResponseDataAttributes(topic interface{}, callbackUrl interface{}, ) *POSTWebhooks201ResponseDataAttributes`

NewPOSTWebhooks201ResponseDataAttributes instantiates a new POSTWebhooks201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTWebhooks201ResponseDataAttributesWithDefaults

`func NewPOSTWebhooks201ResponseDataAttributesWithDefaults() *POSTWebhooks201ResponseDataAttributes`

NewPOSTWebhooks201ResponseDataAttributesWithDefaults instantiates a new POSTWebhooks201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTWebhooks201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTWebhooks201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *POSTWebhooks201ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTopic

`func (o *POSTWebhooks201ResponseDataAttributes) GetTopic() interface{}`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetTopicOk() (*interface{}, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *POSTWebhooks201ResponseDataAttributes) SetTopic(v interface{})`

SetTopic sets Topic field to given value.


### SetTopicNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetCallbackUrl

`func (o *POSTWebhooks201ResponseDataAttributes) GetCallbackUrl() interface{}`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetCallbackUrlOk() (*interface{}, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *POSTWebhooks201ResponseDataAttributes) SetCallbackUrl(v interface{})`

SetCallbackUrl sets CallbackUrl field to given value.


### SetCallbackUrlNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetIncludeResources

`func (o *POSTWebhooks201ResponseDataAttributes) GetIncludeResources() interface{}`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetIncludeResourcesOk() (*interface{}, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *POSTWebhooks201ResponseDataAttributes) SetIncludeResources(v interface{})`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *POSTWebhooks201ResponseDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### SetIncludeResourcesNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetIncludeResourcesNil(b bool)`

 SetIncludeResourcesNil sets the value for IncludeResources to be an explicit nil

### UnsetIncludeResources
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetIncludeResources()`

UnsetIncludeResources ensures that no value is present for IncludeResources, not even an explicit nil
### GetDisable

`func (o *POSTWebhooks201ResponseDataAttributes) GetDisable() interface{}`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetDisableOk() (*interface{}, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *POSTWebhooks201ResponseDataAttributes) SetDisable(v interface{})`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *POSTWebhooks201ResponseDataAttributes) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEnable

`func (o *POSTWebhooks201ResponseDataAttributes) GetEnable() interface{}`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetEnableOk() (*interface{}, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *POSTWebhooks201ResponseDataAttributes) SetEnable(v interface{})`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *POSTWebhooks201ResponseDataAttributes) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetReference

`func (o *POSTWebhooks201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTWebhooks201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTWebhooks201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTWebhooks201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTWebhooks201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTWebhooks201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTWebhooks201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTWebhooks201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTWebhooks201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTWebhooks201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTWebhooks201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTWebhooks201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


