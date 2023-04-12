# MarketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETMarketsMarketId200ResponseDataAttributes**](GETMarketsMarketId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MarketDataRelationships**](MarketDataRelationships.md) |  | [optional] 

## Methods

### NewMarketData

`func NewMarketData(type_ interface{}, attributes GETMarketsMarketId200ResponseDataAttributes, ) *MarketData`

NewMarketData instantiates a new MarketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataWithDefaults

`func NewMarketDataWithDefaults() *MarketData`

NewMarketDataWithDefaults instantiates a new MarketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *MarketData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MarketData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *MarketData) GetAttributes() GETMarketsMarketId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketData) GetAttributesOk() (*GETMarketsMarketId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketData) SetAttributes(v GETMarketsMarketId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MarketData) GetRelationships() MarketDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MarketData) GetRelationshipsOk() (*MarketDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MarketData) SetRelationships(v MarketDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MarketData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


