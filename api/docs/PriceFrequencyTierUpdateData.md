# PriceFrequencyTierUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes**](PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceFrequencyTierUpdateDataRelationships**](PriceFrequencyTierUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceFrequencyTierUpdateData

`func NewPriceFrequencyTierUpdateData(type_ interface{}, id interface{}, attributes PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes, ) *PriceFrequencyTierUpdateData`

NewPriceFrequencyTierUpdateData instantiates a new PriceFrequencyTierUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceFrequencyTierUpdateDataWithDefaults

`func NewPriceFrequencyTierUpdateDataWithDefaults() *PriceFrequencyTierUpdateData`

NewPriceFrequencyTierUpdateDataWithDefaults instantiates a new PriceFrequencyTierUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceFrequencyTierUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceFrequencyTierUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceFrequencyTierUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceFrequencyTierUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceFrequencyTierUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PriceFrequencyTierUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceFrequencyTierUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceFrequencyTierUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PriceFrequencyTierUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PriceFrequencyTierUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PriceFrequencyTierUpdateData) GetAttributes() PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceFrequencyTierUpdateData) GetAttributesOk() (*PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceFrequencyTierUpdateData) SetAttributes(v PATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceFrequencyTierUpdateData) GetRelationships() PriceFrequencyTierUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceFrequencyTierUpdateData) GetRelationshipsOk() (*PriceFrequencyTierUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceFrequencyTierUpdateData) SetRelationships(v PriceFrequencyTierUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceFrequencyTierUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


