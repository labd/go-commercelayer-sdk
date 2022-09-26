# PercentageDiscountPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**POSTPercentageDiscountPromotions201ResponseDataAttributes**](POSTPercentageDiscountPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionUpdateDataRelationships**](FixedPricePromotionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPercentageDiscountPromotionCreateData

`func NewPercentageDiscountPromotionCreateData(type_ string, attributes POSTPercentageDiscountPromotions201ResponseDataAttributes, ) *PercentageDiscountPromotionCreateData`

NewPercentageDiscountPromotionCreateData instantiates a new PercentageDiscountPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPercentageDiscountPromotionCreateDataWithDefaults

`func NewPercentageDiscountPromotionCreateDataWithDefaults() *PercentageDiscountPromotionCreateData`

NewPercentageDiscountPromotionCreateDataWithDefaults instantiates a new PercentageDiscountPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PercentageDiscountPromotionCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PercentageDiscountPromotionCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PercentageDiscountPromotionCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *PercentageDiscountPromotionCreateData) GetAttributes() POSTPercentageDiscountPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PercentageDiscountPromotionCreateData) GetAttributesOk() (*POSTPercentageDiscountPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PercentageDiscountPromotionCreateData) SetAttributes(v POSTPercentageDiscountPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PercentageDiscountPromotionCreateData) GetRelationships() FixedPricePromotionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PercentageDiscountPromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PercentageDiscountPromotionCreateData) SetRelationships(v FixedPricePromotionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PercentageDiscountPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


