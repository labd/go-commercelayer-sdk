# PaypalGatewayResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayResponseDataRelationshipsPaymentMethods**](AdyenGatewayResponseDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaypalPayments** | Pointer to [**PaypalGatewayResponseDataRelationshipsPaypalPayments**](PaypalGatewayResponseDataRelationshipsPaypalPayments.md) |  | [optional] 

## Methods

### NewPaypalGatewayResponseDataRelationships

`func NewPaypalGatewayResponseDataRelationships() *PaypalGatewayResponseDataRelationships`

NewPaypalGatewayResponseDataRelationships instantiates a new PaypalGatewayResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalGatewayResponseDataRelationshipsWithDefaults

`func NewPaypalGatewayResponseDataRelationshipsWithDefaults() *PaypalGatewayResponseDataRelationships`

NewPaypalGatewayResponseDataRelationshipsWithDefaults instantiates a new PaypalGatewayResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *PaypalGatewayResponseDataRelationships) GetPaymentMethods() AdyenGatewayResponseDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaypalGatewayResponseDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayResponseDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaypalGatewayResponseDataRelationships) SetPaymentMethods(v AdyenGatewayResponseDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaypalGatewayResponseDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPaypalPayments

`func (o *PaypalGatewayResponseDataRelationships) GetPaypalPayments() PaypalGatewayResponseDataRelationshipsPaypalPayments`

GetPaypalPayments returns the PaypalPayments field if non-nil, zero value otherwise.

### GetPaypalPaymentsOk

`func (o *PaypalGatewayResponseDataRelationships) GetPaypalPaymentsOk() (*PaypalGatewayResponseDataRelationshipsPaypalPayments, bool)`

GetPaypalPaymentsOk returns a tuple with the PaypalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayments

`func (o *PaypalGatewayResponseDataRelationships) SetPaypalPayments(v PaypalGatewayResponseDataRelationshipsPaypalPayments)`

SetPaypalPayments sets PaypalPayments field to given value.

### HasPaypalPayments

`func (o *PaypalGatewayResponseDataRelationships) HasPaypalPayments() bool`

HasPaypalPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


