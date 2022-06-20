# PriceCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "prices"]
**Attributes** | [**PriceCreateDataAttributes**](PriceCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceCreateDataRelationships**](PriceCreateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceCreateData

`func NewPriceCreateData(type_ string, attributes PriceCreateDataAttributes, ) *PriceCreateData`

NewPriceCreateData instantiates a new PriceCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceCreateDataWithDefaults

`func NewPriceCreateDataWithDefaults() *PriceCreateData`

NewPriceCreateDataWithDefaults instantiates a new PriceCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceCreateData) GetAttributes() PriceCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceCreateData) GetAttributesOk() (*PriceCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceCreateData) SetAttributes(v PriceCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceCreateData) GetRelationships() PriceCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceCreateData) GetRelationshipsOk() (*PriceCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceCreateData) SetRelationships(v PriceCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


