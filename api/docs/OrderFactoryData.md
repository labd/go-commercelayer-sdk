# OrderFactoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETOrderFactories200ResponseDataInnerAttributes**](GETOrderFactories200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**OrderFactoryDataRelationships**](OrderFactoryDataRelationships.md) |  | [optional] 

## Methods

### NewOrderFactoryData

`func NewOrderFactoryData(type_ interface{}, attributes GETOrderFactories200ResponseDataInnerAttributes, ) *OrderFactoryData`

NewOrderFactoryData instantiates a new OrderFactoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFactoryDataWithDefaults

`func NewOrderFactoryDataWithDefaults() *OrderFactoryData`

NewOrderFactoryDataWithDefaults instantiates a new OrderFactoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OrderFactoryData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderFactoryData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderFactoryData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OrderFactoryData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OrderFactoryData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *OrderFactoryData) GetAttributes() GETOrderFactories200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrderFactoryData) GetAttributesOk() (*GETOrderFactories200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrderFactoryData) SetAttributes(v GETOrderFactories200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *OrderFactoryData) GetRelationships() OrderFactoryDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *OrderFactoryData) GetRelationshipsOk() (*OrderFactoryDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *OrderFactoryData) SetRelationships(v OrderFactoryDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *OrderFactoryData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


