# MarketUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHMarketsMarketId200ResponseDataAttributes**](PATCHMarketsMarketId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MarketUpdateDataRelationships**](MarketUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewMarketUpdateData

`func NewMarketUpdateData(type_ interface{}, id interface{}, attributes PATCHMarketsMarketId200ResponseDataAttributes, ) *MarketUpdateData`

NewMarketUpdateData instantiates a new MarketUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketUpdateDataWithDefaults

`func NewMarketUpdateDataWithDefaults() *MarketUpdateData`

NewMarketUpdateDataWithDefaults instantiates a new MarketUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MarketUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MarketUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *MarketUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *MarketUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MarketUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *MarketUpdateData) GetAttributes() PATCHMarketsMarketId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketUpdateData) GetAttributesOk() (*PATCHMarketsMarketId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketUpdateData) SetAttributes(v PATCHMarketsMarketId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MarketUpdateData) GetRelationships() MarketUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MarketUpdateData) GetRelationshipsOk() (*MarketUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MarketUpdateData) SetRelationships(v MarketUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MarketUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


