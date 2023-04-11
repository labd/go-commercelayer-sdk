# POSTInStockSubscriptionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTInStockSubscriptionsRequestDataAttributes**](POSTInStockSubscriptionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTInStockSubscriptionsRequestDataRelationships**](POSTInStockSubscriptionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTInStockSubscriptionsRequestData

`func NewPOSTInStockSubscriptionsRequestData(type_ interface{}, attributes POSTInStockSubscriptionsRequestDataAttributes, ) *POSTInStockSubscriptionsRequestData`

NewPOSTInStockSubscriptionsRequestData instantiates a new POSTInStockSubscriptionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInStockSubscriptionsRequestDataWithDefaults

`func NewPOSTInStockSubscriptionsRequestDataWithDefaults() *POSTInStockSubscriptionsRequestData`

NewPOSTInStockSubscriptionsRequestDataWithDefaults instantiates a new POSTInStockSubscriptionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTInStockSubscriptionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTInStockSubscriptionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTInStockSubscriptionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTInStockSubscriptionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTInStockSubscriptionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTInStockSubscriptionsRequestData) GetAttributes() POSTInStockSubscriptionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTInStockSubscriptionsRequestData) GetAttributesOk() (*POSTInStockSubscriptionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTInStockSubscriptionsRequestData) SetAttributes(v POSTInStockSubscriptionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTInStockSubscriptionsRequestData) GetRelationships() POSTInStockSubscriptionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTInStockSubscriptionsRequestData) GetRelationshipsOk() (*POSTInStockSubscriptionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTInStockSubscriptionsRequestData) SetRelationships(v POSTInStockSubscriptionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTInStockSubscriptionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


