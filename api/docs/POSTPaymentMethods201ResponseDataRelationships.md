# POSTPaymentMethods201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets**](GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets.md) |  | [optional] 
**PaymentGateway** | [**GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway**](GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway.md) |  | 

## Methods

### NewPOSTPaymentMethods201ResponseDataRelationships

`func NewPOSTPaymentMethods201ResponseDataRelationships(paymentGateway GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway, ) *POSTPaymentMethods201ResponseDataRelationships`

NewPOSTPaymentMethods201ResponseDataRelationships instantiates a new POSTPaymentMethods201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults

`func NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults() *POSTPaymentMethods201ResponseDataRelationships`

NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults instantiates a new POSTPaymentMethods201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetPaymentGateway() GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetPaymentGatewayOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetPaymentGateway(v GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


