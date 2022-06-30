# PriceVolumeTierUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "price_volume_tiers"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PriceVolumeTierUpdateDataAttributes**](PriceVolumeTierUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceVolumeTierUpdateDataRelationships**](PriceVolumeTierUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierUpdateData

`func NewPriceVolumeTierUpdateData(type_ string, id string, attributes PriceVolumeTierUpdateDataAttributes, ) *PriceVolumeTierUpdateData`

NewPriceVolumeTierUpdateData instantiates a new PriceVolumeTierUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceVolumeTierUpdateDataWithDefaults

`func NewPriceVolumeTierUpdateDataWithDefaults() *PriceVolumeTierUpdateData`

NewPriceVolumeTierUpdateDataWithDefaults instantiates a new PriceVolumeTierUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceVolumeTierUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PriceVolumeTierUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceVolumeTierUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceVolumeTierUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PriceVolumeTierUpdateData) GetAttributes() PriceVolumeTierUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierUpdateData) GetAttributesOk() (*PriceVolumeTierUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierUpdateData) SetAttributes(v PriceVolumeTierUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceVolumeTierUpdateData) GetRelationships() PriceVolumeTierUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierUpdateData) GetRelationshipsOk() (*PriceVolumeTierUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierUpdateData) SetRelationships(v PriceVolumeTierUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


