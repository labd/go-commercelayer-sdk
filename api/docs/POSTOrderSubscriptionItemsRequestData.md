# POSTOrderSubscriptionItemsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTOrderSubscriptionItemsRequestDataAttributes**](POSTOrderSubscriptionItemsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTOrderSubscriptionItemsRequestDataRelationships**](POSTOrderSubscriptionItemsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTOrderSubscriptionItemsRequestData

`func NewPOSTOrderSubscriptionItemsRequestData(type_ interface{}, attributes POSTOrderSubscriptionItemsRequestDataAttributes, ) *POSTOrderSubscriptionItemsRequestData`

NewPOSTOrderSubscriptionItemsRequestData instantiates a new POSTOrderSubscriptionItemsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptionItemsRequestDataWithDefaults

`func NewPOSTOrderSubscriptionItemsRequestDataWithDefaults() *POSTOrderSubscriptionItemsRequestData`

NewPOSTOrderSubscriptionItemsRequestDataWithDefaults instantiates a new POSTOrderSubscriptionItemsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTOrderSubscriptionItemsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTOrderSubscriptionItemsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTOrderSubscriptionItemsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTOrderSubscriptionItemsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTOrderSubscriptionItemsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTOrderSubscriptionItemsRequestData) GetAttributes() POSTOrderSubscriptionItemsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTOrderSubscriptionItemsRequestData) GetAttributesOk() (*POSTOrderSubscriptionItemsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTOrderSubscriptionItemsRequestData) SetAttributes(v POSTOrderSubscriptionItemsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTOrderSubscriptionItemsRequestData) GetRelationships() POSTOrderSubscriptionItemsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTOrderSubscriptionItemsRequestData) GetRelationshipsOk() (*POSTOrderSubscriptionItemsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTOrderSubscriptionItemsRequestData) SetRelationships(v POSTOrderSubscriptionItemsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTOrderSubscriptionItemsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


