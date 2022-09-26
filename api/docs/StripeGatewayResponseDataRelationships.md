# StripeGatewayResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayResponseDataRelationshipsPaymentMethods**](AdyenGatewayResponseDataRelationshipsPaymentMethods.md) |  | [optional] 
**StripePayments** | Pointer to [**StripeGatewayResponseDataRelationshipsStripePayments**](StripeGatewayResponseDataRelationshipsStripePayments.md) |  | [optional] 

## Methods

### NewStripeGatewayResponseDataRelationships

`func NewStripeGatewayResponseDataRelationships() *StripeGatewayResponseDataRelationships`

NewStripeGatewayResponseDataRelationships instantiates a new StripeGatewayResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeGatewayResponseDataRelationshipsWithDefaults

`func NewStripeGatewayResponseDataRelationshipsWithDefaults() *StripeGatewayResponseDataRelationships`

NewStripeGatewayResponseDataRelationshipsWithDefaults instantiates a new StripeGatewayResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *StripeGatewayResponseDataRelationships) GetPaymentMethods() AdyenGatewayResponseDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *StripeGatewayResponseDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayResponseDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *StripeGatewayResponseDataRelationships) SetPaymentMethods(v AdyenGatewayResponseDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *StripeGatewayResponseDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetStripePayments

`func (o *StripeGatewayResponseDataRelationships) GetStripePayments() StripeGatewayResponseDataRelationshipsStripePayments`

GetStripePayments returns the StripePayments field if non-nil, zero value otherwise.

### GetStripePaymentsOk

`func (o *StripeGatewayResponseDataRelationships) GetStripePaymentsOk() (*StripeGatewayResponseDataRelationshipsStripePayments, bool)`

GetStripePaymentsOk returns a tuple with the StripePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePayments

`func (o *StripeGatewayResponseDataRelationships) SetStripePayments(v StripeGatewayResponseDataRelationshipsStripePayments)`

SetStripePayments sets StripePayments field to given value.

### HasStripePayments

`func (o *StripeGatewayResponseDataRelationships) HasStripePayments() bool`

HasStripePayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


