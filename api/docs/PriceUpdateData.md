# PriceUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "prices"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PriceUpdateDataAttributes**](PriceUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceUpdateDataRelationships**](PriceUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceUpdateData

`func NewPriceUpdateData(type_ string, id string, attributes PriceUpdateDataAttributes, ) *PriceUpdateData`

NewPriceUpdateData instantiates a new PriceUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceUpdateDataWithDefaults

`func NewPriceUpdateDataWithDefaults() *PriceUpdateData`

NewPriceUpdateDataWithDefaults instantiates a new PriceUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PriceUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PriceUpdateData) GetAttributes() PriceUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceUpdateData) GetAttributesOk() (*PriceUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceUpdateData) SetAttributes(v PriceUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceUpdateData) GetRelationships() PriceUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceUpdateData) GetRelationshipsOk() (*PriceUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceUpdateData) SetRelationships(v PriceUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


