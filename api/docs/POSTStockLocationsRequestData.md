# POSTStockLocationsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStockLocationsRequestDataAttributes**](POSTStockLocationsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTMerchantsRequestDataRelationships**](POSTMerchantsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStockLocationsRequestData

`func NewPOSTStockLocationsRequestData(type_ interface{}, attributes POSTStockLocationsRequestDataAttributes, ) *POSTStockLocationsRequestData`

NewPOSTStockLocationsRequestData instantiates a new POSTStockLocationsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLocationsRequestDataWithDefaults

`func NewPOSTStockLocationsRequestDataWithDefaults() *POSTStockLocationsRequestData`

NewPOSTStockLocationsRequestDataWithDefaults instantiates a new POSTStockLocationsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTStockLocationsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStockLocationsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStockLocationsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTStockLocationsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTStockLocationsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTStockLocationsRequestData) GetAttributes() POSTStockLocationsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStockLocationsRequestData) GetAttributesOk() (*POSTStockLocationsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStockLocationsRequestData) SetAttributes(v POSTStockLocationsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTStockLocationsRequestData) GetRelationships() POSTMerchantsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStockLocationsRequestData) GetRelationshipsOk() (*POSTMerchantsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStockLocationsRequestData) SetRelationships(v POSTMerchantsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStockLocationsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


