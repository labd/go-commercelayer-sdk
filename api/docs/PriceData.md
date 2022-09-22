# PriceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETPrices200ResponseDataInnerAttributes**](GETPrices200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**PriceDataRelationships**](PriceDataRelationships.md) |  | [optional] 

## Methods

### NewPriceData

`func NewPriceData(type_ string, attributes GETPrices200ResponseDataInnerAttributes, ) *PriceData`

NewPriceData instantiates a new PriceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDataWithDefaults

`func NewPriceDataWithDefaults() *PriceData`

NewPriceDataWithDefaults instantiates a new PriceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceData) GetAttributes() GETPrices200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceData) GetAttributesOk() (*GETPrices200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceData) SetAttributes(v GETPrices200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceData) GetRelationships() PriceDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceData) GetRelationshipsOk() (*PriceDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceData) SetRelationships(v PriceDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


