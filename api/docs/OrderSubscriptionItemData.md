# OrderSubscriptionItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes**](GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderSubscriptionItemDataRelationships**](OrderSubscriptionItemDataRelationships.md) |  | [optional] 

## Methods

### NewOrderSubscriptionItemData

`func NewOrderSubscriptionItemData(type_ interface{}, attributes GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes, ) *OrderSubscriptionItemData`

NewOrderSubscriptionItemData instantiates a new OrderSubscriptionItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionItemDataWithDefaults

`func NewOrderSubscriptionItemDataWithDefaults() *OrderSubscriptionItemData`

NewOrderSubscriptionItemDataWithDefaults instantiates a new OrderSubscriptionItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderSubscriptionItemData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderSubscriptionItemData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderSubscriptionItemData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderSubscriptionItemData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderSubscriptionItemData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderSubscriptionItemData) GetAttributes() GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderSubscriptionItemData) GetAttributesOk() (*GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderSubscriptionItemData) SetAttributes(v GETOrderSubscriptionItemsOrderSubscriptionItemId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderSubscriptionItemData) GetRelationships() OrderSubscriptionItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderSubscriptionItemData) GetRelationshipsOk() (*OrderSubscriptionItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderSubscriptionItemData) SetRelationships(v OrderSubscriptionItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderSubscriptionItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


