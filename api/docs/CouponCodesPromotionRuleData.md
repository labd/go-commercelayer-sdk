# CouponCodesPromotionRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "coupon_codes_promotion_rules"]
**Attributes** | [**GETBillingInfoValidationRules200ResponseDataInnerAttributes**](GETBillingInfoValidationRules200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETCouponCodesPromotionRules200ResponseDataInnerRelationships**](GETCouponCodesPromotionRules200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewCouponCodesPromotionRuleData

`func NewCouponCodesPromotionRuleData(type_ string, attributes GETBillingInfoValidationRules200ResponseDataInnerAttributes, ) *CouponCodesPromotionRuleData`

NewCouponCodesPromotionRuleData instantiates a new CouponCodesPromotionRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodesPromotionRuleDataWithDefaults

`func NewCouponCodesPromotionRuleDataWithDefaults() *CouponCodesPromotionRuleData`

NewCouponCodesPromotionRuleDataWithDefaults instantiates a new CouponCodesPromotionRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponCodesPromotionRuleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCodesPromotionRuleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCodesPromotionRuleData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CouponCodesPromotionRuleData) GetAttributes() GETBillingInfoValidationRules200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCodesPromotionRuleData) GetAttributesOk() (*GETBillingInfoValidationRules200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCodesPromotionRuleData) SetAttributes(v GETBillingInfoValidationRules200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponCodesPromotionRuleData) GetRelationships() GETCouponCodesPromotionRules200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponCodesPromotionRuleData) GetRelationshipsOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponCodesPromotionRuleData) SetRelationships(v GETCouponCodesPromotionRules200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponCodesPromotionRuleData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


