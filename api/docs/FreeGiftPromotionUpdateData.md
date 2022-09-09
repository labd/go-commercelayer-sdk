# FreeGiftPromotionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "free_gift_promotions"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes**](PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionUpdateDataRelationships**](FixedPricePromotionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewFreeGiftPromotionUpdateData

`func NewFreeGiftPromotionUpdateData(type_ string, id string, attributes PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes, ) *FreeGiftPromotionUpdateData`

NewFreeGiftPromotionUpdateData instantiates a new FreeGiftPromotionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeGiftPromotionUpdateDataWithDefaults

`func NewFreeGiftPromotionUpdateDataWithDefaults() *FreeGiftPromotionUpdateData`

NewFreeGiftPromotionUpdateDataWithDefaults instantiates a new FreeGiftPromotionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FreeGiftPromotionUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeGiftPromotionUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeGiftPromotionUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FreeGiftPromotionUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreeGiftPromotionUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreeGiftPromotionUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FreeGiftPromotionUpdateData) GetAttributes() PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FreeGiftPromotionUpdateData) GetAttributesOk() (*PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FreeGiftPromotionUpdateData) SetAttributes(v PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FreeGiftPromotionUpdateData) GetRelationships() FixedPricePromotionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FreeGiftPromotionUpdateData) GetRelationshipsOk() (*FixedPricePromotionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FreeGiftPromotionUpdateData) SetRelationships(v FixedPricePromotionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FreeGiftPromotionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


