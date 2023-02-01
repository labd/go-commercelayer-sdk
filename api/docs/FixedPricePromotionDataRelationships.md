# FixedPricePromotionDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**PromotionRules** | Pointer to [**ExternalPromotionDataRelationshipsPromotionRules**](ExternalPromotionDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**ExternalPromotionDataRelationshipsOrderAmountPromotionRule**](ExternalPromotionDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**ExternalPromotionDataRelationshipsSkuListPromotionRule**](ExternalPromotionDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**CouponDataRelationshipsPromotionRule**](CouponDataRelationshipsPromotionRule.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountDataRelationshipsAttachments**](AvalaraAccountDataRelationshipsAttachments.md) |  | [optional] 
**SkuList** | Pointer to [**BundleDataRelationshipsSkuList**](BundleDataRelationshipsSkuList.md) |  | [optional] 
**Skus** | Pointer to [**BundleDataRelationshipsSkus**](BundleDataRelationshipsSkus.md) |  | [optional] 
**Events** | Pointer to [**AuthorizationDataRelationshipsEvents**](AuthorizationDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewFixedPricePromotionDataRelationships

`func NewFixedPricePromotionDataRelationships() *FixedPricePromotionDataRelationships`

NewFixedPricePromotionDataRelationships instantiates a new FixedPricePromotionDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionDataRelationshipsWithDefaults

`func NewFixedPricePromotionDataRelationshipsWithDefaults() *FixedPricePromotionDataRelationships`

NewFixedPricePromotionDataRelationshipsWithDefaults instantiates a new FixedPricePromotionDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *FixedPricePromotionDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *FixedPricePromotionDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *FixedPricePromotionDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *FixedPricePromotionDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *FixedPricePromotionDataRelationships) GetPromotionRules() ExternalPromotionDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *FixedPricePromotionDataRelationships) GetPromotionRulesOk() (*ExternalPromotionDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *FixedPricePromotionDataRelationships) SetPromotionRules(v ExternalPromotionDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *FixedPricePromotionDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *FixedPricePromotionDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *FixedPricePromotionDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *FixedPricePromotionDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *FixedPricePromotionDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *FixedPricePromotionDataRelationships) GetSkuListPromotionRule() ExternalPromotionDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *FixedPricePromotionDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *FixedPricePromotionDataRelationships) SetSkuListPromotionRule(v ExternalPromotionDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *FixedPricePromotionDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *FixedPricePromotionDataRelationships) GetCouponCodesPromotionRule() CouponDataRelationshipsPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *FixedPricePromotionDataRelationships) GetCouponCodesPromotionRuleOk() (*CouponDataRelationshipsPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *FixedPricePromotionDataRelationships) SetCouponCodesPromotionRule(v CouponDataRelationshipsPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *FixedPricePromotionDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetAttachments

`func (o *FixedPricePromotionDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *FixedPricePromotionDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *FixedPricePromotionDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *FixedPricePromotionDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSkuList

`func (o *FixedPricePromotionDataRelationships) GetSkuList() BundleDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *FixedPricePromotionDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *FixedPricePromotionDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.

### HasSkuList

`func (o *FixedPricePromotionDataRelationships) HasSkuList() bool`

HasSkuList returns a boolean if a field has been set.

### GetSkus

`func (o *FixedPricePromotionDataRelationships) GetSkus() BundleDataRelationshipsSkus`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *FixedPricePromotionDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *FixedPricePromotionDataRelationships) SetSkus(v BundleDataRelationshipsSkus)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *FixedPricePromotionDataRelationships) HasSkus() bool`

HasSkus returns a boolean if a field has been set.

### GetEvents

`func (o *FixedPricePromotionDataRelationships) GetEvents() AuthorizationDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FixedPricePromotionDataRelationships) GetEventsOk() (*AuthorizationDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FixedPricePromotionDataRelationships) SetEvents(v AuthorizationDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FixedPricePromotionDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


