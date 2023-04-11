# POSTReturnLineItemsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTReturnLineItemsRequestDataAttributes**](POSTReturnLineItemsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTReturnLineItemsRequestDataRelationships**](POSTReturnLineItemsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTReturnLineItemsRequestData

`func NewPOSTReturnLineItemsRequestData(type_ interface{}, attributes POSTReturnLineItemsRequestDataAttributes, ) *POSTReturnLineItemsRequestData`

NewPOSTReturnLineItemsRequestData instantiates a new POSTReturnLineItemsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTReturnLineItemsRequestDataWithDefaults

`func NewPOSTReturnLineItemsRequestDataWithDefaults() *POSTReturnLineItemsRequestData`

NewPOSTReturnLineItemsRequestDataWithDefaults instantiates a new POSTReturnLineItemsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTReturnLineItemsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTReturnLineItemsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTReturnLineItemsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTReturnLineItemsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTReturnLineItemsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTReturnLineItemsRequestData) GetAttributes() POSTReturnLineItemsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTReturnLineItemsRequestData) GetAttributesOk() (*POSTReturnLineItemsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTReturnLineItemsRequestData) SetAttributes(v POSTReturnLineItemsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTReturnLineItemsRequestData) GetRelationships() POSTReturnLineItemsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTReturnLineItemsRequestData) GetRelationshipsOk() (*POSTReturnLineItemsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTReturnLineItemsRequestData) SetRelationships(v POSTReturnLineItemsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTReturnLineItemsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


