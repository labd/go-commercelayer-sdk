# InStockSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "in_stock_subscriptions"]
**Attributes** | [**GETInStockSubscriptions200ResponseDataInnerAttributes**](GETInStockSubscriptions200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETInStockSubscriptions200ResponseDataInnerRelationships**](GETInStockSubscriptions200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewInStockSubscriptionData

`func NewInStockSubscriptionData(type_ string, attributes GETInStockSubscriptions200ResponseDataInnerAttributes, ) *InStockSubscriptionData`

NewInStockSubscriptionData instantiates a new InStockSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionDataWithDefaults

`func NewInStockSubscriptionDataWithDefaults() *InStockSubscriptionData`

NewInStockSubscriptionDataWithDefaults instantiates a new InStockSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InStockSubscriptionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InStockSubscriptionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InStockSubscriptionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InStockSubscriptionData) GetAttributes() GETInStockSubscriptions200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InStockSubscriptionData) GetAttributesOk() (*GETInStockSubscriptions200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InStockSubscriptionData) SetAttributes(v GETInStockSubscriptions200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InStockSubscriptionData) GetRelationships() GETInStockSubscriptions200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InStockSubscriptionData) GetRelationshipsOk() (*GETInStockSubscriptions200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InStockSubscriptionData) SetRelationships(v GETInStockSubscriptions200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InStockSubscriptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


