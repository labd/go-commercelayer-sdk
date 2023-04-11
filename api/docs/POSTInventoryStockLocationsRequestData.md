# POSTInventoryStockLocationsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTInventoryStockLocationsRequestDataAttributes**](POSTInventoryStockLocationsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTInventoryReturnLocationsRequestDataRelationships**](POSTInventoryReturnLocationsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTInventoryStockLocationsRequestData

`func NewPOSTInventoryStockLocationsRequestData(type_ interface{}, attributes POSTInventoryStockLocationsRequestDataAttributes, ) *POSTInventoryStockLocationsRequestData`

NewPOSTInventoryStockLocationsRequestData instantiates a new POSTInventoryStockLocationsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInventoryStockLocationsRequestDataWithDefaults

`func NewPOSTInventoryStockLocationsRequestDataWithDefaults() *POSTInventoryStockLocationsRequestData`

NewPOSTInventoryStockLocationsRequestDataWithDefaults instantiates a new POSTInventoryStockLocationsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTInventoryStockLocationsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTInventoryStockLocationsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTInventoryStockLocationsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTInventoryStockLocationsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTInventoryStockLocationsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTInventoryStockLocationsRequestData) GetAttributes() POSTInventoryStockLocationsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTInventoryStockLocationsRequestData) GetAttributesOk() (*POSTInventoryStockLocationsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTInventoryStockLocationsRequestData) SetAttributes(v POSTInventoryStockLocationsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTInventoryStockLocationsRequestData) GetRelationships() POSTInventoryReturnLocationsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTInventoryStockLocationsRequestData) GetRelationshipsOk() (*POSTInventoryReturnLocationsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTInventoryStockLocationsRequestData) SetRelationships(v POSTInventoryReturnLocationsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTInventoryStockLocationsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


