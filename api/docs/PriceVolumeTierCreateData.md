# PriceVolumeTierCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPriceVolumeTiers201ResponseDataAttributes**](POSTPriceVolumeTiers201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceFrequencyTierCreateDataRelationships**](PriceFrequencyTierCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierCreateData

`func NewPriceVolumeTierCreateData(type_ interface{}, attributes POSTPriceVolumeTiers201ResponseDataAttributes, ) *PriceVolumeTierCreateData`

NewPriceVolumeTierCreateData instantiates a new PriceVolumeTierCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceVolumeTierCreateDataWithDefaults

`func NewPriceVolumeTierCreateDataWithDefaults() *PriceVolumeTierCreateData`

NewPriceVolumeTierCreateDataWithDefaults instantiates a new PriceVolumeTierCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceVolumeTierCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceVolumeTierCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceVolumeTierCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PriceVolumeTierCreateData) GetAttributes() POSTPriceVolumeTiers201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierCreateData) GetAttributesOk() (*POSTPriceVolumeTiers201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierCreateData) SetAttributes(v POSTPriceVolumeTiers201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceVolumeTierCreateData) GetRelationships() PriceFrequencyTierCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierCreateData) GetRelationshipsOk() (*PriceFrequencyTierCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierCreateData) SetRelationships(v PriceFrequencyTierCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


