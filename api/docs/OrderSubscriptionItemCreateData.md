# OrderSubscriptionItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderSubscriptionItems201ResponseDataAttributes**](POSTOrderSubscriptionItems201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderSubscriptionItemCreateDataRelationships**](OrderSubscriptionItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewOrderSubscriptionItemCreateData

`func NewOrderSubscriptionItemCreateData(type_ interface{}, attributes POSTOrderSubscriptionItems201ResponseDataAttributes, ) *OrderSubscriptionItemCreateData`

NewOrderSubscriptionItemCreateData instantiates a new OrderSubscriptionItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionItemCreateDataWithDefaults

`func NewOrderSubscriptionItemCreateDataWithDefaults() *OrderSubscriptionItemCreateData`

NewOrderSubscriptionItemCreateDataWithDefaults instantiates a new OrderSubscriptionItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderSubscriptionItemCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderSubscriptionItemCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderSubscriptionItemCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderSubscriptionItemCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderSubscriptionItemCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderSubscriptionItemCreateData) GetAttributes() POSTOrderSubscriptionItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderSubscriptionItemCreateData) GetAttributesOk() (*POSTOrderSubscriptionItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderSubscriptionItemCreateData) SetAttributes(v POSTOrderSubscriptionItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderSubscriptionItemCreateData) GetRelationships() OrderSubscriptionItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderSubscriptionItemCreateData) GetRelationshipsOk() (*OrderSubscriptionItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderSubscriptionItemCreateData) SetRelationships(v OrderSubscriptionItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderSubscriptionItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


