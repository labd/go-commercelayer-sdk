# ExternalPaymentDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**PaymentGateway** | Pointer to [**AdyenPaymentDataRelationshipsPaymentGateway**](AdyenPaymentDataRelationshipsPaymentGateway.md) |  | [optional] 
**Wallet** | Pointer to [**CustomerDataRelationshipsCustomerPaymentSources**](CustomerDataRelationshipsCustomerPaymentSources.md) |  | [optional] 

## Methods

### NewExternalPaymentDataRelationships

`func NewExternalPaymentDataRelationships() *ExternalPaymentDataRelationships`

NewExternalPaymentDataRelationships instantiates a new ExternalPaymentDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentDataRelationshipsWithDefaults

`func NewExternalPaymentDataRelationshipsWithDefaults() *ExternalPaymentDataRelationships`

NewExternalPaymentDataRelationshipsWithDefaults instantiates a new ExternalPaymentDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *ExternalPaymentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ExternalPaymentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ExternalPaymentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ExternalPaymentDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *ExternalPaymentDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *ExternalPaymentDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *ExternalPaymentDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *ExternalPaymentDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetWallet

`func (o *ExternalPaymentDataRelationships) GetWallet() CustomerDataRelationshipsCustomerPaymentSources`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ExternalPaymentDataRelationships) GetWalletOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ExternalPaymentDataRelationships) SetWallet(v CustomerDataRelationshipsCustomerPaymentSources)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ExternalPaymentDataRelationships) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


