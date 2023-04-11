# FixedPricePromotionCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleCreateDataRelationshipsMarket**](BillingInfoValidationRuleCreateDataRelationshipsMarket.md) |  | [optional] 
**PromotionRules** | Pointer to [**ExternalPromotionCreateDataRelationshipsPromotionRules**](ExternalPromotionCreateDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule**](ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**ExternalPromotionCreateDataRelationshipsSkuListPromotionRule**](ExternalPromotionCreateDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**CouponCreateDataRelationshipsPromotionRule**](CouponCreateDataRelationshipsPromotionRule.md) |  | [optional] 
**SkuList** | [**BundleCreateDataRelationshipsSkuList**](BundleCreateDataRelationshipsSkuList.md) |  | 

## Methods

### NewFixedPricePromotionCreateDataRelationships

`func NewFixedPricePromotionCreateDataRelationships(skuList BundleCreateDataRelationshipsSkuList, ) *FixedPricePromotionCreateDataRelationships`

NewFixedPricePromotionCreateDataRelationships instantiates a new FixedPricePromotionCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionCreateDataRelationshipsWithDefaults

`func NewFixedPricePromotionCreateDataRelationshipsWithDefaults() *FixedPricePromotionCreateDataRelationships`

NewFixedPricePromotionCreateDataRelationshipsWithDefaults instantiates a new FixedPricePromotionCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *FixedPricePromotionCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *FixedPricePromotionCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *FixedPricePromotionCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *FixedPricePromotionCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *FixedPricePromotionCreateDataRelationships) GetPromotionRules() ExternalPromotionCreateDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *FixedPricePromotionCreateDataRelationships) GetPromotionRulesOk() (*ExternalPromotionCreateDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *FixedPricePromotionCreateDataRelationships) SetPromotionRules(v ExternalPromotionCreateDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *FixedPricePromotionCreateDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *FixedPricePromotionCreateDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionCreateDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) GetSkuListPromotionRule() ExternalPromotionCreateDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *FixedPricePromotionCreateDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionCreateDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) SetSkuListPromotionRule(v ExternalPromotionCreateDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) GetCouponCodesPromotionRule() CouponCreateDataRelationshipsPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *FixedPricePromotionCreateDataRelationships) GetCouponCodesPromotionRuleOk() (*CouponCreateDataRelationshipsPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) SetCouponCodesPromotionRule(v CouponCreateDataRelationshipsPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *FixedPricePromotionCreateDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetSkuList

`func (o *FixedPricePromotionCreateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *FixedPricePromotionCreateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *FixedPricePromotionCreateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


