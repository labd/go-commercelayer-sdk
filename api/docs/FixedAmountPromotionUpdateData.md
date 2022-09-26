# FixedAmountPromotionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**FixedAmountPromotionUpdateDataAttributes**](FixedAmountPromotionUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionCreateDataRelationships**](ExternalPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFixedAmountPromotionUpdateData

`func NewFixedAmountPromotionUpdateData(type_ string, id string, attributes FixedAmountPromotionUpdateDataAttributes, ) *FixedAmountPromotionUpdateData`

NewFixedAmountPromotionUpdateData instantiates a new FixedAmountPromotionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAmountPromotionUpdateDataWithDefaults

`func NewFixedAmountPromotionUpdateDataWithDefaults() *FixedAmountPromotionUpdateData`

NewFixedAmountPromotionUpdateDataWithDefaults instantiates a new FixedAmountPromotionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedAmountPromotionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedAmountPromotionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedAmountPromotionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FixedAmountPromotionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedAmountPromotionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedAmountPromotionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FixedAmountPromotionUpdateData) GetAttributes() FixedAmountPromotionUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedAmountPromotionUpdateData) GetAttributesOk() (*FixedAmountPromotionUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedAmountPromotionUpdateData) SetAttributes(v FixedAmountPromotionUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedAmountPromotionUpdateData) GetRelationships() ExternalPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedAmountPromotionUpdateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedAmountPromotionUpdateData) SetRelationships(v ExternalPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedAmountPromotionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


