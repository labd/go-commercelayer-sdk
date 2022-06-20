# PaymentMethodCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**PaymentGateway** | [**AdyenPaymentDataRelationshipsPaymentGateway**](AdyenPaymentDataRelationshipsPaymentGateway.md) |  | 

## Methods

### NewPaymentMethodCreateDataRelationships

`func NewPaymentMethodCreateDataRelationships(paymentGateway AdyenPaymentDataRelationshipsPaymentGateway, ) *PaymentMethodCreateDataRelationships`

NewPaymentMethodCreateDataRelationships instantiates a new PaymentMethodCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCreateDataRelationshipsWithDefaults

`func NewPaymentMethodCreateDataRelationshipsWithDefaults() *PaymentMethodCreateDataRelationships`

NewPaymentMethodCreateDataRelationshipsWithDefaults instantiates a new PaymentMethodCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *PaymentMethodCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PaymentMethodCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PaymentMethodCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *PaymentMethodCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *PaymentMethodCreateDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *PaymentMethodCreateDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *PaymentMethodCreateDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


