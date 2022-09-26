# PriceTierData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GETPriceTiers200ResponseDataInnerAttributes**](GETPriceTiers200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**PriceTierDataRelationships**](PriceTierDataRelationships.md) |  | [optional] 

## Methods

### NewPriceTierData

`func NewPriceTierData(type_ string, attributes GETPriceTiers200ResponseDataInnerAttributes, ) *PriceTierData`

NewPriceTierData instantiates a new PriceTierData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierDataWithDefaults

`func NewPriceTierDataWithDefaults() *PriceTierData`

NewPriceTierDataWithDefaults instantiates a new PriceTierData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceTierData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceTierData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceTierData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PriceTierData) GetAttributes() GETPriceTiers200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceTierData) GetAttributesOk() (*GETPriceTiers200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceTierData) SetAttributes(v GETPriceTiers200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceTierData) GetRelationships() PriceTierDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceTierData) GetRelationshipsOk() (*PriceTierDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceTierData) SetRelationships(v PriceTierDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceTierData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


