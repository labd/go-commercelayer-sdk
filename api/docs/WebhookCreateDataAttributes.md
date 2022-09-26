# WebhookCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name for the webhook. | [optional] 
**Topic** | **string** | The identifier of the resource/event that will trigger the webhook. | 
**CallbackUrl** | **string** | URI where the webhook subscription should send the POST request when the event occurs. | 
**IncludeResources** | Pointer to **[]string** | List of related resources that should be included in the webhook body. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewWebhookCreateDataAttributes

`func NewWebhookCreateDataAttributes(topic string, callbackUrl string, ) *WebhookCreateDataAttributes`

NewWebhookCreateDataAttributes instantiates a new WebhookCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateDataAttributesWithDefaults

`func NewWebhookCreateDataAttributesWithDefaults() *WebhookCreateDataAttributes`

NewWebhookCreateDataAttributesWithDefaults instantiates a new WebhookCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookCreateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookCreateDataAttributes) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookCreateDataAttributes) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookCreateDataAttributes) SetTopic(v string)`

SetTopic sets Topic field to given value.


### GetCallbackUrl

`func (o *WebhookCreateDataAttributes) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *WebhookCreateDataAttributes) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *WebhookCreateDataAttributes) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetIncludeResources

`func (o *WebhookCreateDataAttributes) GetIncludeResources() []string`

GetIncludeResources returns the IncludeResources field if non-nil, zero value otherwise.

### GetIncludeResourcesOk

`func (o *WebhookCreateDataAttributes) GetIncludeResourcesOk() (*[]string, bool)`

GetIncludeResourcesOk returns a tuple with the IncludeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResources

`func (o *WebhookCreateDataAttributes) SetIncludeResources(v []string)`

SetIncludeResources sets IncludeResources field to given value.

### HasIncludeResources

`func (o *WebhookCreateDataAttributes) HasIncludeResources() bool`

HasIncludeResources returns a boolean if a field has been set.

### GetReference

`func (o *WebhookCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *WebhookCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *WebhookCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *WebhookCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *WebhookCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *WebhookCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *WebhookCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *WebhookCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *WebhookCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebhookCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


