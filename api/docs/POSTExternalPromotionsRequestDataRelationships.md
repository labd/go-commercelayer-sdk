# POSTExternalPromotionsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**PromotionRules** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsPromotionRules**](POSTExternalPromotionsRequestDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule**](POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule**](POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**POSTCouponsRequestDataRelationshipsPromotionRule**](POSTCouponsRequestDataRelationshipsPromotionRule.md) |  | [optional] 

## Methods

### NewPOSTExternalPromotionsRequestDataRelationships

`func NewPOSTExternalPromotionsRequestDataRelationships() *POSTExternalPromotionsRequestDataRelationships`

NewPOSTExternalPromotionsRequestDataRelationships instantiates a new POSTExternalPromotionsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalPromotionsRequestDataRelationshipsWithDefaults

`func NewPOSTExternalPromotionsRequestDataRelationshipsWithDefaults() *POSTExternalPromotionsRequestDataRelationships`

NewPOSTExternalPromotionsRequestDataRelationshipsWithDefaults instantiates a new POSTExternalPromotionsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTExternalPromotionsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTExternalPromotionsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTExternalPromotionsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTExternalPromotionsRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *POSTExternalPromotionsRequestDataRelationships) GetPromotionRules() POSTExternalPromotionsRequestDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *POSTExternalPromotionsRequestDataRelationships) GetPromotionRulesOk() (*POSTExternalPromotionsRequestDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *POSTExternalPromotionsRequestDataRelationships) SetPromotionRules(v POSTExternalPromotionsRequestDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *POSTExternalPromotionsRequestDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) GetOrderAmountPromotionRule() POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *POSTExternalPromotionsRequestDataRelationships) GetOrderAmountPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) SetOrderAmountPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) GetSkuListPromotionRule() POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *POSTExternalPromotionsRequestDataRelationships) GetSkuListPromotionRuleOk() (*POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) SetSkuListPromotionRule(v POSTExternalPromotionsRequestDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) GetCouponCodesPromotionRule() POSTCouponsRequestDataRelationshipsPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *POSTExternalPromotionsRequestDataRelationships) GetCouponCodesPromotionRuleOk() (*POSTCouponsRequestDataRelationshipsPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) SetCouponCodesPromotionRule(v POSTCouponsRequestDataRelationshipsPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *POSTExternalPromotionsRequestDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


