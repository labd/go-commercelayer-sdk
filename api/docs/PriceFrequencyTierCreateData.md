# PriceFrequencyTierCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPriceFrequencyTiers201ResponseDataAttributes**](POSTPriceFrequencyTiers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceFrequencyTierCreateDataRelationships**](PriceFrequencyTierCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceFrequencyTierCreateData

`func NewPriceFrequencyTierCreateData(type_ interface{}, attributes POSTPriceFrequencyTiers201ResponseDataAttributes, ) *PriceFrequencyTierCreateData`

NewPriceFrequencyTierCreateData instantiates a new PriceFrequencyTierCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceFrequencyTierCreateDataWithDefaults

`func NewPriceFrequencyTierCreateDataWithDefaults() *PriceFrequencyTierCreateData`

NewPriceFrequencyTierCreateDataWithDefaults instantiates a new PriceFrequencyTierCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceFrequencyTierCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceFrequencyTierCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceFrequencyTierCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceFrequencyTierCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceFrequencyTierCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceFrequencyTierCreateData) GetAttributes() POSTPriceFrequencyTiers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceFrequencyTierCreateData) GetAttributesOk() (*POSTPriceFrequencyTiers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceFrequencyTierCreateData) SetAttributes(v POSTPriceFrequencyTiers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceFrequencyTierCreateData) GetRelationships() PriceFrequencyTierCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceFrequencyTierCreateData) GetRelationshipsOk() (*PriceFrequencyTierCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceFrequencyTierCreateData) SetRelationships(v PriceFrequencyTierCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceFrequencyTierCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


