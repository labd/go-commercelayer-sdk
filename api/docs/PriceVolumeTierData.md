# PriceVolumeTierData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**PriceTierDataAttributes**](PriceTierDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceTierDataRelationships**](PriceTierDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierData

`func NewPriceVolumeTierData(type_ string, attributes PriceTierDataAttributes, ) *PriceVolumeTierData`

NewPriceVolumeTierData instantiates a new PriceVolumeTierData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceVolumeTierDataWithDefaults

`func NewPriceVolumeTierDataWithDefaults() *PriceVolumeTierData`

NewPriceVolumeTierDataWithDefaults instantiates a new PriceVolumeTierData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceVolumeTierData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceVolumeTierData) GetAttributes() PriceTierDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierData) GetAttributesOk() (*PriceTierDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierData) SetAttributes(v PriceTierDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceVolumeTierData) GetRelationships() PriceTierDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierData) GetRelationshipsOk() (*PriceTierDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierData) SetRelationships(v PriceTierDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


