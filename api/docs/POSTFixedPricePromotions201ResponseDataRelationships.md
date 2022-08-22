# POSTFixedPricePromotions201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets**](GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets.md) |  | [optional] 
**PromotionRules** | Pointer to [**GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules**](GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule**](GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule**](GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**GETCoupons200ResponseDataInnerRelationshipsPromotionRule**](GETCoupons200ResponseDataInnerRelationshipsPromotionRule.md) |  | [optional] 
**SkuList** | [**GETBundles200ResponseDataInnerRelationshipsSkuList**](GETBundles200ResponseDataInnerRelationshipsSkuList.md) |  | 

## Methods

### NewPOSTFixedPricePromotions201ResponseDataRelationships

`func NewPOSTFixedPricePromotions201ResponseDataRelationships(skuList GETBundles200ResponseDataInnerRelationshipsSkuList, ) *POSTFixedPricePromotions201ResponseDataRelationships`

NewPOSTFixedPricePromotions201ResponseDataRelationships instantiates a new POSTFixedPricePromotions201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFixedPricePromotions201ResponseDataRelationshipsWithDefaults

`func NewPOSTFixedPricePromotions201ResponseDataRelationshipsWithDefaults() *POSTFixedPricePromotions201ResponseDataRelationships`

NewPOSTFixedPricePromotions201ResponseDataRelationshipsWithDefaults instantiates a new POSTFixedPricePromotions201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetPromotionRules() GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetPromotionRulesOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetPromotionRules(v GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetOrderAmountPromotionRule() GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetOrderAmountPromotionRuleOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetOrderAmountPromotionRule(v GETExternalPromotions200ResponseDataInnerRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetSkuListPromotionRule() GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetSkuListPromotionRuleOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetSkuListPromotionRule(v GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetCouponCodesPromotionRule() GETCoupons200ResponseDataInnerRelationshipsPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetCouponCodesPromotionRuleOk() (*GETCoupons200ResponseDataInnerRelationshipsPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetCouponCodesPromotionRule(v GETCoupons200ResponseDataInnerRelationshipsPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetSkuList

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *POSTFixedPricePromotions201ResponseDataRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


