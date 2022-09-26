# PriceVolumeTierCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**PriceVolumeTierCreateDataAttributes**](PriceVolumeTierCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceVolumeTierCreateDataRelationships**](PriceVolumeTierCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierCreateData

`func NewPriceVolumeTierCreateData(type_ string, attributes PriceVolumeTierCreateDataAttributes, ) *PriceVolumeTierCreateData`

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

`func (o *PriceVolumeTierCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceVolumeTierCreateData) GetAttributes() PriceVolumeTierCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierCreateData) GetAttributesOk() (*PriceVolumeTierCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierCreateData) SetAttributes(v PriceVolumeTierCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceVolumeTierCreateData) GetRelationships() PriceVolumeTierCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierCreateData) GetRelationshipsOk() (*PriceVolumeTierCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierCreateData) SetRelationships(v PriceVolumeTierCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


