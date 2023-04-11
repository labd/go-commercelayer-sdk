# FreeGiftPromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETFreeGiftPromotions200ResponseDataInnerAttributes**](GETFreeGiftPromotions200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionDataRelationships**](FixedPricePromotionDataRelationships.md) |  | [optional] 

## Methods

### NewFreeGiftPromotionData

`func NewFreeGiftPromotionData(type_ interface{}, attributes GETFreeGiftPromotions200ResponseDataInnerAttributes, ) *FreeGiftPromotionData`

NewFreeGiftPromotionData instantiates a new FreeGiftPromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeGiftPromotionDataWithDefaults

`func NewFreeGiftPromotionDataWithDefaults() *FreeGiftPromotionData`

NewFreeGiftPromotionDataWithDefaults instantiates a new FreeGiftPromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FreeGiftPromotionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeGiftPromotionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeGiftPromotionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FreeGiftPromotionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FreeGiftPromotionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FreeGiftPromotionData) GetAttributes() GETFreeGiftPromotions200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FreeGiftPromotionData) GetAttributesOk() (*GETFreeGiftPromotions200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FreeGiftPromotionData) SetAttributes(v GETFreeGiftPromotions200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FreeGiftPromotionData) GetRelationships() FixedPricePromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FreeGiftPromotionData) GetRelationshipsOk() (*FixedPricePromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FreeGiftPromotionData) SetRelationships(v FixedPricePromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FreeGiftPromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


