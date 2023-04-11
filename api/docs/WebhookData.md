# WebhookData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETWebhooksWebhookId200ResponseDataAttributes**](GETWebhooksWebhookId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**WebhookDataRelationships**](WebhookDataRelationships.md) |  | [optional] 

## Methods

### NewWebhookData

`func NewWebhookData(type_ interface{}, attributes GETWebhooksWebhookId200ResponseDataAttributes, ) *WebhookData`

NewWebhookData instantiates a new WebhookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDataWithDefaults

`func NewWebhookDataWithDefaults() *WebhookData`

NewWebhookDataWithDefaults instantiates a new WebhookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *WebhookData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WebhookData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *WebhookData) GetAttributes() GETWebhooksWebhookId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WebhookData) GetAttributesOk() (*GETWebhooksWebhookId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WebhookData) SetAttributes(v GETWebhooksWebhookId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WebhookData) GetRelationships() WebhookDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WebhookData) GetRelationshipsOk() (*WebhookDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WebhookData) SetRelationships(v WebhookDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WebhookData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


