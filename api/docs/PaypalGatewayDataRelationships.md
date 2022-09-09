# PaypalGatewayDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaypalPayments** | Pointer to [**PaypalGatewayDataRelationshipsPaypalPayments**](PaypalGatewayDataRelationshipsPaypalPayments.md) |  | [optional] 

## Methods

### NewPaypalGatewayDataRelationships

`func NewPaypalGatewayDataRelationships() *PaypalGatewayDataRelationships`

NewPaypalGatewayDataRelationships instantiates a new PaypalGatewayDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalGatewayDataRelationshipsWithDefaults

`func NewPaypalGatewayDataRelationshipsWithDefaults() *PaypalGatewayDataRelationships`

NewPaypalGatewayDataRelationshipsWithDefaults instantiates a new PaypalGatewayDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *PaypalGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaypalGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaypalGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaypalGatewayDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPaypalPayments

`func (o *PaypalGatewayDataRelationships) GetPaypalPayments() PaypalGatewayDataRelationshipsPaypalPayments`

GetPaypalPayments returns the PaypalPayments field if non-nil, zero value otherwise.

### GetPaypalPaymentsOk

`func (o *PaypalGatewayDataRelationships) GetPaypalPaymentsOk() (*PaypalGatewayDataRelationshipsPaypalPayments, bool)`

GetPaypalPaymentsOk returns a tuple with the PaypalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalPayments

`func (o *PaypalGatewayDataRelationships) SetPaypalPayments(v PaypalGatewayDataRelationshipsPaypalPayments)`

SetPaypalPayments sets PaypalPayments field to given value.

### HasPaypalPayments

`func (o *PaypalGatewayDataRelationships) HasPaypalPayments() bool`

HasPaypalPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


