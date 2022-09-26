# ExternalPaymentResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentResponseDataRelationshipsOrder**](AdyenPaymentResponseDataRelationshipsOrder.md) |  | [optional] 
**PaymentGateway** | Pointer to [**AdyenPaymentResponseDataRelationshipsPaymentGateway**](AdyenPaymentResponseDataRelationshipsPaymentGateway.md) |  | [optional] 
**Wallet** | Pointer to [**ExternalPaymentResponseDataRelationshipsWallet**](ExternalPaymentResponseDataRelationshipsWallet.md) |  | [optional] 

## Methods

### NewExternalPaymentResponseDataRelationships

`func NewExternalPaymentResponseDataRelationships() *ExternalPaymentResponseDataRelationships`

NewExternalPaymentResponseDataRelationships instantiates a new ExternalPaymentResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentResponseDataRelationshipsWithDefaults

`func NewExternalPaymentResponseDataRelationshipsWithDefaults() *ExternalPaymentResponseDataRelationships`

NewExternalPaymentResponseDataRelationshipsWithDefaults instantiates a new ExternalPaymentResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ExternalPaymentResponseDataRelationships) GetOrder() AdyenPaymentResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ExternalPaymentResponseDataRelationships) GetOrderOk() (*AdyenPaymentResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ExternalPaymentResponseDataRelationships) SetOrder(v AdyenPaymentResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ExternalPaymentResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *ExternalPaymentResponseDataRelationships) GetPaymentGateway() AdyenPaymentResponseDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *ExternalPaymentResponseDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentResponseDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *ExternalPaymentResponseDataRelationships) SetPaymentGateway(v AdyenPaymentResponseDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *ExternalPaymentResponseDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetWallet

`func (o *ExternalPaymentResponseDataRelationships) GetWallet() ExternalPaymentResponseDataRelationshipsWallet`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ExternalPaymentResponseDataRelationships) GetWalletOk() (*ExternalPaymentResponseDataRelationshipsWallet, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ExternalPaymentResponseDataRelationships) SetWallet(v ExternalPaymentResponseDataRelationshipsWallet)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ExternalPaymentResponseDataRelationships) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


