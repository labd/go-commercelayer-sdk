# OrderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**OrderDataAttributes**](OrderDataAttributes.md) |  | 
**Relationships** | Pointer to [**OrderDataRelationships**](OrderDataRelationships.md) |  | [optional] 

## Methods

### NewOrderData

`func NewOrderData(type_ string, attributes OrderDataAttributes, ) *OrderData`

NewOrderData instantiates a new OrderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDataWithDefaults

`func NewOrderDataWithDefaults() *OrderData`

NewOrderDataWithDefaults instantiates a new OrderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *OrderData) GetAttributes() OrderDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderData) GetAttributesOk() (*OrderDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderData) SetAttributes(v OrderDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderData) GetRelationships() OrderDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderData) GetRelationshipsOk() (*OrderDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderData) SetRelationships(v OrderDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


