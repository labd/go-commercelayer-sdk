# InStockSubscriptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "in_stock_subscriptions"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes**](PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataRelationships**](PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewInStockSubscriptionUpdateData

`func NewInStockSubscriptionUpdateData(type_ string, id string, attributes PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes, ) *InStockSubscriptionUpdateData`

NewInStockSubscriptionUpdateData instantiates a new InStockSubscriptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionUpdateDataWithDefaults

`func NewInStockSubscriptionUpdateDataWithDefaults() *InStockSubscriptionUpdateData`

NewInStockSubscriptionUpdateDataWithDefaults instantiates a new InStockSubscriptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InStockSubscriptionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InStockSubscriptionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InStockSubscriptionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InStockSubscriptionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InStockSubscriptionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InStockSubscriptionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *InStockSubscriptionUpdateData) GetAttributes() PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InStockSubscriptionUpdateData) GetAttributesOk() (*PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InStockSubscriptionUpdateData) SetAttributes(v PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InStockSubscriptionUpdateData) GetRelationships() PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InStockSubscriptionUpdateData) GetRelationshipsOk() (*PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InStockSubscriptionUpdateData) SetRelationships(v PATCHInStockSubscriptionsInStockSubscriptionId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InStockSubscriptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


