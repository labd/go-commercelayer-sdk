# OrderCopyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "order_copies"]
**Attributes** | [**GETOrderCopies200ResponseDataInnerAttributes**](GETOrderCopies200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETOrderCopies200ResponseDataInnerRelationships**](GETOrderCopies200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewOrderCopyData

`func NewOrderCopyData(type_ string, attributes GETOrderCopies200ResponseDataInnerAttributes, ) *OrderCopyData`

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

`func (o *OrderCopyData) GetAttributes() GETOrderCopies200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderCopyData) GetAttributesOk() (*GETOrderCopies200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderCopyData) SetAttributes(v GETOrderCopies200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderCopyData) GetRelationships() GETOrderCopies200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderCopyData) GetRelationshipsOk() (*GETOrderCopies200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderCopyData) SetRelationships(v GETOrderCopies200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderCopyData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


