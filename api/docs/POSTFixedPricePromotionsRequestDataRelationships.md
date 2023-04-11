# POSTFixedPricePromotionsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**PromotionRules** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsPromotionRules**](POSTExternalPromotionsRequestDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule**](POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule**](POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**POSTCouponsRequestDataRelationshipsPromotionRule**](POSTCouponsRequestDataRelationshipsPromotionRule.md) |  | [optional] 
**SkuList** | [**POSTBundlesRequestDataRelationshipsSkuList**](POSTBundlesRequestDataRelationshipsSkuList.md) |  | 

## Methods

### NewPOSTFixedPricePromotionsRequestDataRelationships

`func NewPOSTFixedPricePromotionsRequestDataRelationships(skuList POSTBundlesRequestDataRelationshipsSkuList, ) *POSTFixedPricePromotionsRequestDataRelationships`

NewPOSTFixedPricePromotionsRequestDataRelationships instantiates a new POSTFixedPricePromotionsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFixedPricePromotionsRequestDataRelationshipsWithDefaults

`func NewPOSTFixedPricePromotionsRequestDataRelationshipsWithDefaults() *POSTFixedPricePromotionsRequestDataRelationships`

NewPOSTFixedPricePromotionsRequestDataRelationshipsWithDefaults instantiates a new POSTFixedPricePromotionsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTFixedPricePromotionsRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetPromotionRules() POSTExternalPromotionsRequestDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetPromotionRulesOk() (*POSTExternalPromotionsRequestDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetPromotionRules(v POSTExternalPromotionsRequestDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *POSTFixedPricePromotionsRequestDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetOrderAmountPromotionRule() POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetOrderAmountPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetOrderAmountPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetSkuListPromotionRule() POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetSkuListPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetSkuListPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetCouponCodesPromotionRule() POSTCouponsRequestDataRelationshipsPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetCouponCodesPromotionRuleOk() (*POSTCouponsRequestDataRelationshipsPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetCouponCodesPromotionRule(v POSTCouponsRequestDataRelationshipsPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *POSTFixedPricePromotionsRequestDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetSkuList

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetSkuList() POSTBundlesRequestDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *POSTFixedPricePromotionsRequestDataRelationships) GetSkuListOk() (*POSTBundlesRequestDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *POSTFixedPricePromotionsRequestDataRelationships) SetSkuList(v POSTBundlesRequestDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


