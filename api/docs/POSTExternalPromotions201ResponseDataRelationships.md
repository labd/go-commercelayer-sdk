# POSTExternalPromotions201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket**](POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket.md) |  | [optional] 
**PromotionRules** | Pointer to [**POSTExternalPromotions201ResponseDataRelationshipsPromotionRules**](POSTExternalPromotions201ResponseDataRelationshipsPromotionRules.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRule**](POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**POSTExternalPromotions201ResponseDataRelationshipsSkuListPromotionRule**](POSTExternalPromotions201ResponseDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule**](POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule.md) |  | [optional] 
**Attachments** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationshipsAttachments**](POSTAvalaraAccounts201ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents.md) |  | [optional] 

## Methods

### NewPOSTExternalPromotions201ResponseDataRelationships

`func NewPOSTExternalPromotions201ResponseDataRelationships() *POSTExternalPromotions201ResponseDataRelationships`

NewPOSTExternalPromotions201ResponseDataRelationships instantiates a new POSTExternalPromotions201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalPromotions201ResponseDataRelationshipsWithDefaults

`func NewPOSTExternalPromotions201ResponseDataRelationshipsWithDefaults() *POSTExternalPromotions201ResponseDataRelationships`

NewPOSTExternalPromotions201ResponseDataRelationshipsWithDefaults instantiates a new POSTExternalPromotions201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetMarket() POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetMarket(v POSTBillingInfoValidationRules201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPromotionRules

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetPromotionRules() POSTExternalPromotions201ResponseDataRelationshipsPromotionRules`

GetPromotionRules returns the PromotionRules field if non-nil, zero value otherwise.

### GetPromotionRulesOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetPromotionRulesOk() (*POSTExternalPromotions201ResponseDataRelationshipsPromotionRules, bool)`

GetPromotionRulesOk returns a tuple with the PromotionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionRules

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetPromotionRules(v POSTExternalPromotions201ResponseDataRelationshipsPromotionRules)`

SetPromotionRules sets PromotionRules field to given value.

### HasPromotionRules

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasPromotionRules() bool`

HasPromotionRules returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetOrderAmountPromotionRule() POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetOrderAmountPromotionRuleOk() (*POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetOrderAmountPromotionRule(v POSTExternalPromotions201ResponseDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetSkuListPromotionRule() POSTExternalPromotions201ResponseDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetSkuListPromotionRuleOk() (*POSTExternalPromotions201ResponseDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetSkuListPromotionRule(v POSTExternalPromotions201ResponseDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetCouponCodesPromotionRule() POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetCouponCodesPromotionRuleOk() (*POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetCouponCodesPromotionRule(v POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *POSTExternalPromotions201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *POSTExternalPromotions201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *POSTExternalPromotions201ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


