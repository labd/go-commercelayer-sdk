# BuyXPayYPromotionDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**PromotionRules** | Pointer to [**BuyXPayYPromotionDataRelationshipsPromotionRules**](BuyXPayYPromotionDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule**](BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**BuyXPayYPromotionDataRelationshipsSkuListPromotionRule**](BuyXPayYPromotionDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule**](BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule.md) |  | [optional] 
**CustomPromotionRule** | Pointer to [**BuyXPayYPromotionDataRelationshipsCustomPromotionRule**](BuyXPayYPromotionDataRelationshipsCustomPromotionRule.md) |  | [optional] 
**SkuList** | Pointer to [**BundleDataRelationshipsSkuList**](BundleDataRelationshipsSkuList.md) |  | [optional] 
**Coupons** | Pointer to [**BuyXPayYPromotionDataRelationshipsCoupons**](BuyXPayYPromotionDataRelationshipsCoupons.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Tags** | Pointer to [**AddressDataRelationshipsTags**](AddressDataRelationshipsTags.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**Skus** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 

## Methods

### NewBuyXPayYPromotionDataRelationships

`func NewBuyXPayYPromotionDataRelationships() *BuyXPayYPromotionDataRelationships`

NewBuyXPayYPromotionDataRelationships instantiates a new BuyXPayYPromotionDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyXPayYPromotionDataRelationshipsWithDefaults

`func NewBuyXPayYPromotionDataRelationshipsWithDefaults() *BuyXPayYPromotionDataRelationships`

NewBuyXPayYPromotionDataRelationshipsWithDefaults instantiates a new BuyXPayYPromotionDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BuyXPayYPromotionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BuyXPayYPromotionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BuyXPayYPromotionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *BuyXPayYPromotionDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *BuyXPayYPromotionDataRelationships) GetPromotionRules() BuyXPayYPromotionDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *BuyXPayYPromotionDataRelationships) GetPromotionRulesOk() (*BuyXPayYPromotionDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *BuyXPayYPromotionDataRelationships) SetPromotionRules(v BuyXPayYPromotionDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *BuyXPayYPromotionDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) GetOrderAmountPromotionRule() BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *BuyXPayYPromotionDataRelationships) GetOrderAmountPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) SetOrderAmountPromotionRule(v BuyXPayYPromotionDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) GetSkuListPromotionRule() BuyXPayYPromotionDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *BuyXPayYPromotionDataRelationships) GetSkuListPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) SetSkuListPromotionRule(v BuyXPayYPromotionDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) GetCouponCodesPromotionRule() BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *BuyXPayYPromotionDataRelationships) GetCouponCodesPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) SetCouponCodesPromotionRule(v BuyXPayYPromotionDataRelationshipsCouponCodesPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetCustomPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) GetCustomPromotionRule() BuyXPayYPromotionDataRelationshipsCustomPromotionRule`

GetCustomPromotionRule returns the CustomPromotionRule field if non-nil, zero value otherwise.

### GetCustomPromotionRuleOk

`func (o *BuyXPayYPromotionDataRelationships) GetCustomPromotionRuleOk() (*BuyXPayYPromotionDataRelationshipsCustomPromotionRule, bool)`

GetCustomPromotionRuleOk returns a tuple with the CustomPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) SetCustomPromotionRule(v BuyXPayYPromotionDataRelationshipsCustomPromotionRule)`

SetCustomPromotionRule sets CustomPromotionRule field to given value.

### HasCustomPromotionRule

`func (o *BuyXPayYPromotionDataRelationships) HasCustomPromotionRule() bool`

HasCustomPromotionRule returns a boolean if a field has been set.

### GetSkuList

`func (o *BuyXPayYPromotionDataRelationships) GetSkuList() BundleDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *BuyXPayYPromotionDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *BuyXPayYPromotionDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.

### HasSkuList

`func (o *BuyXPayYPromotionDataRelationships) HasSkuList() bool`

HasSkuList returns a boolean if a field has been set.

### GetCoupons

`func (o *BuyXPayYPromotionDataRelationships) GetCoupons() BuyXPayYPromotionDataRelationshipsCoupons`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *BuyXPayYPromotionDataRelationships) GetCouponsOk() (*BuyXPayYPromotionDataRelationshipsCoupons, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *BuyXPayYPromotionDataRelationships) SetCoupons(v BuyXPayYPromotionDataRelationshipsCoupons)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *BuyXPayYPromotionDataRelationships) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

### GetAttachments

`func (o *BuyXPayYPromotionDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BuyXPayYPromotionDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BuyXPayYPromotionDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BuyXPayYPromotionDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *BuyXPayYPromotionDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BuyXPayYPromotionDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BuyXPayYPromotionDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *BuyXPayYPromotionDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTags

`func (o *BuyXPayYPromotionDataRelationships) GetTags() AddressDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BuyXPayYPromotionDataRelationships) GetTagsOk() (*AddressDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BuyXPayYPromotionDataRelationships) SetTags(v AddressDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BuyXPayYPromotionDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *BuyXPayYPromotionDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BuyXPayYPromotionDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BuyXPayYPromotionDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *BuyXPayYPromotionDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetSkus

`func (o *BuyXPayYPromotionDataRelationships) GetSkus() BundleDataRelationshipsSkus`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *BuyXPayYPromotionDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *BuyXPayYPromotionDataRelationships) SetSkus(v BundleDataRelationshipsSkus)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *BuyXPayYPromotionDataRelationships) HasSkus() bool`

HasSkus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


