# POSTPaymentMethodsRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**PaymentGateway** | [**POSTPaymentMethodsRequestDataRelationshipsPaymentGateway**](POSTPaymentMethodsRequestDataRelationshipsPaymentGateway.md) |  | 

## Methods

### NewPOSTPaymentMethodsRequestDataRelationships

`func NewPOSTPaymentMethodsRequestDataRelationships(paymentGateway POSTPaymentMethodsRequestDataRelationshipsPaymentGateway, ) *POSTPaymentMethodsRequestDataRelationships`

NewPOSTPaymentMethodsRequestDataRelationships instantiates a new POSTPaymentMethodsRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethodsRequestDataRelationshipsWithDefaults

`func NewPOSTPaymentMethodsRequestDataRelationshipsWithDefaults() *POSTPaymentMethodsRequestDataRelationships`

NewPOSTPaymentMethodsRequestDataRelationshipsWithDefaults instantiates a new POSTPaymentMethodsRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTPaymentMethodsRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTPaymentMethodsRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTPaymentMethodsRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTPaymentMethodsRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *POSTPaymentMethodsRequestDataRelationships) GetPaymentGateway() POSTPaymentMethodsRequestDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *POSTPaymentMethodsRequestDataRelationships) GetPaymentGatewayOk() (*POSTPaymentMethodsRequestDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *POSTPaymentMethodsRequestDataRelationships) SetPaymentGateway(v POSTPaymentMethodsRequestDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


