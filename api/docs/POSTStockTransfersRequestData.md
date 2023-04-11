# POSTStockTransfersRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockTransfersRequestDataAttributes**](POSTStockTransfersRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTStockTransfersRequestDataRelationships**](POSTStockTransfersRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStockTransfersRequestData

`func NewPOSTStockTransfersRequestData(type_ interface{}, attributes POSTStockTransfersRequestDataAttributes, ) *POSTStockTransfersRequestData`

NewPOSTStockTransfersRequestData instantiates a new POSTStockTransfersRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockTransfersRequestDataWithDefaults

`func NewPOSTStockTransfersRequestDataWithDefaults() *POSTStockTransfersRequestData`

NewPOSTStockTransfersRequestDataWithDefaults instantiates a new POSTStockTransfersRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTStockTransfersRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStockTransfersRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStockTransfersRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTStockTransfersRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTStockTransfersRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTStockTransfersRequestData) GetAttributes() POSTStockTransfersRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStockTransfersRequestData) GetAttributesOk() (*POSTStockTransfersRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStockTransfersRequestData) SetAttributes(v POSTStockTransfersRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTStockTransfersRequestData) GetRelationships() POSTStockTransfersRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStockTransfersRequestData) GetRelationshipsOk() (*POSTStockTransfersRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStockTransfersRequestData) SetRelationships(v POSTStockTransfersRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStockTransfersRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


