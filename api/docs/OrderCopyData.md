# OrderCopyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**OrderCopyDataAttributes**](OrderCopyDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderCopyDataRelationships**](OrderCopyDataRelationships.md) |  | [optional] 

## Methods

### NewOrderCopyData

`func NewOrderCopyData(type_ string, attributes OrderCopyDataAttributes, ) *OrderCopyData`

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

`func (o *OrderCopyData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderCopyData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderCopyData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrderCopyData) GetAttributes() OrderCopyDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderCopyData) GetAttributesOk() (*OrderCopyDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderCopyData) SetAttributes(v OrderCopyDataAttributes)`

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


