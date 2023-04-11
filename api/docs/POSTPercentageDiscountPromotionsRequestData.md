# POSTPercentageDiscountPromotionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPercentageDiscountPromotionsRequestDataAttributes**](POSTPercentageDiscountPromotionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships**](PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTPercentageDiscountPromotionsRequestData

`func NewPOSTPercentageDiscountPromotionsRequestData(type_ interface{}, attributes POSTPercentageDiscountPromotionsRequestDataAttributes, ) *POSTPercentageDiscountPromotionsRequestData`

NewPOSTPercentageDiscountPromotionsRequestData instantiates a new POSTPercentageDiscountPromotionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPercentageDiscountPromotionsRequestDataWithDefaults

`func NewPOSTPercentageDiscountPromotionsRequestDataWithDefaults() *POSTPercentageDiscountPromotionsRequestData`

NewPOSTPercentageDiscountPromotionsRequestDataWithDefaults instantiates a new POSTPercentageDiscountPromotionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTPercentageDiscountPromotionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTPercentageDiscountPromotionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTPercentageDiscountPromotionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTPercentageDiscountPromotionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTPercentageDiscountPromotionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTPercentageDiscountPromotionsRequestData) GetAttributes() POSTPercentageDiscountPromotionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTPercentageDiscountPromotionsRequestData) GetAttributesOk() (*POSTPercentageDiscountPromotionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTPercentageDiscountPromotionsRequestData) SetAttributes(v POSTPercentageDiscountPromotionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTPercentageDiscountPromotionsRequestData) GetRelationships() PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTPercentageDiscountPromotionsRequestData) GetRelationshipsOk() (*PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTPercentageDiscountPromotionsRequestData) SetRelationships(v PATCHFixedPricePromotionsFixedPricePromotionIdRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTPercentageDiscountPromotionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


