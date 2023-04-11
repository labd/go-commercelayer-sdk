# OrderSubscriptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes**](PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderSubscriptionUpdateDataRelationships**](OrderSubscriptionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderSubscriptionUpdateData

`func NewOrderSubscriptionUpdateData(type_ interface{}, id interface{}, attributes PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes, ) *OrderSubscriptionUpdateData`

NewOrderSubscriptionUpdateData instantiates a new OrderSubscriptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionUpdateDataWithDefaults

`func NewOrderSubscriptionUpdateDataWithDefaults() *OrderSubscriptionUpdateData`

NewOrderSubscriptionUpdateDataWithDefaults instantiates a new OrderSubscriptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderSubscriptionUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderSubscriptionUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderSubscriptionUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderSubscriptionUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderSubscriptionUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *OrderSubscriptionUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderSubscriptionUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderSubscriptionUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *OrderSubscriptionUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrderSubscriptionUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *OrderSubscriptionUpdateData) GetAttributes() PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderSubscriptionUpdateData) GetAttributesOk() (*PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderSubscriptionUpdateData) SetAttributes(v PATCHOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderSubscriptionUpdateData) GetRelationships() OrderSubscriptionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderSubscriptionUpdateData) GetRelationshipsOk() (*OrderSubscriptionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderSubscriptionUpdateData) SetRelationships(v OrderSubscriptionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderSubscriptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


