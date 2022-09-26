# ExternalPromotionResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleResponseDataRelationshipsMarket**](BillingInfoValidationRuleResponseDataRelationshipsMarket.md) |  | [optional] 
**PromotionRules** | Pointer to [**ExternalPromotionResponseDataRelationshipsPromotionRules**](ExternalPromotionResponseDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule**](ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**ExternalPromotionResponseDataRelationshipsSkuListPromotionRule**](ExternalPromotionResponseDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule**](ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewExternalPromotionResponseDataRelationships

`func NewExternalPromotionResponseDataRelationships() *ExternalPromotionResponseDataRelationships`

NewExternalPromotionResponseDataRelationships instantiates a new ExternalPromotionResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPromotionResponseDataRelationshipsWithDefaults

`func NewExternalPromotionResponseDataRelationshipsWithDefaults() *ExternalPromotionResponseDataRelationships`

NewExternalPromotionResponseDataRelationshipsWithDefaults instantiates a new ExternalPromotionResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *ExternalPromotionResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *ExternalPromotionResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *ExternalPromotionResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *ExternalPromotionResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *ExternalPromotionResponseDataRelationships) GetPromotionRules() ExternalPromotionResponseDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *ExternalPromotionResponseDataRelationships) GetPromotionRulesOk() (*ExternalPromotionResponseDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *ExternalPromotionResponseDataRelationships) SetPromotionRules(v ExternalPromotionResponseDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *ExternalPromotionResponseDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) GetOrderAmountPromotionRule() ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *ExternalPromotionResponseDataRelationships) GetOrderAmountPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) SetOrderAmountPromotionRule(v ExternalPromotionResponseDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) GetSkuListPromotionRule() ExternalPromotionResponseDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *ExternalPromotionResponseDataRelationships) GetSkuListPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) SetSkuListPromotionRule(v ExternalPromotionResponseDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) GetCouponCodesPromotionRule() ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *ExternalPromotionResponseDataRelationships) GetCouponCodesPromotionRuleOk() (*ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) SetCouponCodesPromotionRule(v ExternalPromotionResponseDataRelationshipsCouponCodesPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *ExternalPromotionResponseDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetAttachments

`func (o *ExternalPromotionResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ExternalPromotionResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ExternalPromotionResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ExternalPromotionResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


