# OrderCopyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderCopiesOrderCopyId200ResponseDataAttributes**](GETOrderCopiesOrderCopyId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCopyDataRelationships**](OrderCopyDataRelationships.md) |  | [optional] 

## Methods

### NewOrderCopyData

`func NewOrderCopyData(type_ interface{}, attributes GETOrderCopiesOrderCopyId200ResponseDataAttributes, ) *OrderCopyData`

NewOrderCopyData instantiates a new OrderCopyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCopyDataWithDefaults

`func NewOrderCopyDataWithDefaults() *OrderCopyData`

NewOrderCopyDataWithDefaults instantiates a new OrderCopyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderCopyData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCopyData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCopyData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderCopyData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderCopyData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderCopyData) GetAttributes() GETOrderCopiesOrderCopyId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderCopyData) GetAttributesOk() (*GETOrderCopiesOrderCopyId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderCopyData) SetAttributes(v GETOrderCopiesOrderCopyId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderCopyData) GetRelationships() OrderCopyDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderCopyData) GetRelationshipsOk() (*OrderCopyDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderCopyData) SetRelationships(v OrderCopyDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderCopyData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


