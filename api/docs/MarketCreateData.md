# MarketCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "markets"]
**Attributes** | [**POSTMarkets201ResponseDataAttributes**](POSTMarkets201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MarketCreateDataRelationships**](MarketCreateDataRelationships.md) |  | [optional] 

## Methods

### NewMarketCreateData

`func NewMarketCreateData(type_ string, attributes POSTMarkets201ResponseDataAttributes, ) *MarketCreateData`

NewMarketCreateData instantiates a new MarketCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketCreateDataWithDefaults

`func NewMarketCreateDataWithDefaults() *MarketCreateData`

NewMarketCreateDataWithDefaults instantiates a new MarketCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MarketCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *MarketCreateData) GetAttributes() POSTMarkets201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MarketCreateData) GetAttributesOk() (*POSTMarkets201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MarketCreateData) SetAttributes(v POSTMarkets201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *MarketCreateData) GetRelationships() MarketCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MarketCreateData) GetRelationshipsOk() (*MarketCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MarketCreateData) SetRelationships(v MarketCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MarketCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


