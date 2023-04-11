# POSTOrderSubscriptionsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**SourceOrder** | [**POSTAdyenPaymentsRequestDataRelationshipsOrder**](POSTAdyenPaymentsRequestDataRelationshipsOrder.md) |  | 

## Methods

### NewPOSTOrderSubscriptionsRequestDataRelationships

`func NewPOSTOrderSubscriptionsRequestDataRelationships(sourceOrder POSTAdyenPaymentsRequestDataRelationshipsOrder, ) *POSTOrderSubscriptionsRequestDataRelationships`

NewPOSTOrderSubscriptionsRequestDataRelationships instantiates a new POSTOrderSubscriptionsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptionsRequestDataRelationshipsWithDefaults

`func NewPOSTOrderSubscriptionsRequestDataRelationshipsWithDefaults() *POSTOrderSubscriptionsRequestDataRelationships`

NewPOSTOrderSubscriptionsRequestDataRelationshipsWithDefaults instantiates a new POSTOrderSubscriptionsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrderSubscriptionsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrderSubscriptionsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrderSubscriptionsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrderSubscriptionsRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSourceOrder

`func (o *POSTOrderSubscriptionsRequestDataRelationships) GetSourceOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *POSTOrderSubscriptionsRequestDataRelationships) GetSourceOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *POSTOrderSubscriptionsRequestDataRelationships) SetSourceOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


