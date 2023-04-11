# POSTRecurringOrderCopiesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTRecurringOrderCopiesRequestDataAttributes**](POSTRecurringOrderCopiesRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTRecurringOrderCopiesRequestDataRelationships**](POSTRecurringOrderCopiesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTRecurringOrderCopiesRequestData

`func NewPOSTRecurringOrderCopiesRequestData(type_ interface{}, attributes POSTRecurringOrderCopiesRequestDataAttributes, ) *POSTRecurringOrderCopiesRequestData`

NewPOSTRecurringOrderCopiesRequestData instantiates a new POSTRecurringOrderCopiesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTRecurringOrderCopiesRequestDataWithDefaults

`func NewPOSTRecurringOrderCopiesRequestDataWithDefaults() *POSTRecurringOrderCopiesRequestData`

NewPOSTRecurringOrderCopiesRequestDataWithDefaults instantiates a new POSTRecurringOrderCopiesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTRecurringOrderCopiesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTRecurringOrderCopiesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTRecurringOrderCopiesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTRecurringOrderCopiesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTRecurringOrderCopiesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTRecurringOrderCopiesRequestData) GetAttributes() POSTRecurringOrderCopiesRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTRecurringOrderCopiesRequestData) GetAttributesOk() (*POSTRecurringOrderCopiesRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTRecurringOrderCopiesRequestData) SetAttributes(v POSTRecurringOrderCopiesRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTRecurringOrderCopiesRequestData) GetRelationships() POSTRecurringOrderCopiesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTRecurringOrderCopiesRequestData) GetRelationshipsOk() (*POSTRecurringOrderCopiesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTRecurringOrderCopiesRequestData) SetRelationships(v POSTRecurringOrderCopiesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTRecurringOrderCopiesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


