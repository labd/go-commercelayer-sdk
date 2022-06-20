# InStockSubscriptionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "in_stock_subscriptions"]
**Attributes** | [**InStockSubscriptionCreateDataAttributes**](InStockSubscriptionCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**InStockSubscriptionCreateDataRelationships**](InStockSubscriptionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewInStockSubscriptionCreateData

`func NewInStockSubscriptionCreateData(type_ string, attributes InStockSubscriptionCreateDataAttributes, ) *InStockSubscriptionCreateData`

NewInStockSubscriptionCreateData instantiates a new InStockSubscriptionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionCreateDataWithDefaults

`func NewInStockSubscriptionCreateDataWithDefaults() *InStockSubscriptionCreateData`

NewInStockSubscriptionCreateDataWithDefaults instantiates a new InStockSubscriptionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InStockSubscriptionCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InStockSubscriptionCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InStockSubscriptionCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InStockSubscriptionCreateData) GetAttributes() InStockSubscriptionCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InStockSubscriptionCreateData) GetAttributesOk() (*InStockSubscriptionCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InStockSubscriptionCreateData) SetAttributes(v InStockSubscriptionCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InStockSubscriptionCreateData) GetRelationships() InStockSubscriptionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InStockSubscriptionCreateData) GetRelationshipsOk() (*InStockSubscriptionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InStockSubscriptionCreateData) SetRelationships(v InStockSubscriptionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InStockSubscriptionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


