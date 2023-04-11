# PriceFrequencyTierData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETPriceFrequencyTiers200ResponseDataInnerAttributes**](GETPriceFrequencyTiers200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**PriceFrequencyTierDataRelationships**](PriceFrequencyTierDataRelationships.md) |  | [optional] 

## Methods

### NewPriceFrequencyTierData

`func NewPriceFrequencyTierData(type_ interface{}, attributes GETPriceFrequencyTiers200ResponseDataInnerAttributes, ) *PriceFrequencyTierData`

NewPriceFrequencyTierData instantiates a new PriceFrequencyTierData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceFrequencyTierDataWithDefaults

`func NewPriceFrequencyTierDataWithDefaults() *PriceFrequencyTierData`

NewPriceFrequencyTierDataWithDefaults instantiates a new PriceFrequencyTierData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceFrequencyTierData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceFrequencyTierData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceFrequencyTierData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceFrequencyTierData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceFrequencyTierData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceFrequencyTierData) GetAttributes() GETPriceFrequencyTiers200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceFrequencyTierData) GetAttributesOk() (*GETPriceFrequencyTiers200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceFrequencyTierData) SetAttributes(v GETPriceFrequencyTiers200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceFrequencyTierData) GetRelationships() PriceFrequencyTierDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceFrequencyTierData) GetRelationshipsOk() (*PriceFrequencyTierDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceFrequencyTierData) SetRelationships(v PriceFrequencyTierDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceFrequencyTierData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


