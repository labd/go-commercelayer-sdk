# OrderSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderSubscriptions200ResponseDataInnerAttributes**](GETOrderSubscriptions200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**OrderSubscriptionDataRelationships**](OrderSubscriptionDataRelationships.md) |  | [optional] 

## Methods

### NewOrderSubscriptionData

`func NewOrderSubscriptionData(type_ interface{}, attributes GETOrderSubscriptions200ResponseDataInnerAttributes, ) *OrderSubscriptionData`

NewOrderSubscriptionData instantiates a new OrderSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionDataWithDefaults

`func NewOrderSubscriptionDataWithDefaults() *OrderSubscriptionData`

NewOrderSubscriptionDataWithDefaults instantiates a new OrderSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderSubscriptionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderSubscriptionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderSubscriptionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderSubscriptionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderSubscriptionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderSubscriptionData) GetAttributes() GETOrderSubscriptions200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderSubscriptionData) GetAttributesOk() (*GETOrderSubscriptions200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderSubscriptionData) SetAttributes(v GETOrderSubscriptions200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderSubscriptionData) GetRelationships() OrderSubscriptionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderSubscriptionData) GetRelationshipsOk() (*OrderSubscriptionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderSubscriptionData) SetRelationships(v OrderSubscriptionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderSubscriptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


