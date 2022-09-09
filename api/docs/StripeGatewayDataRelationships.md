# StripeGatewayDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**StripePayments** | Pointer to [**StripeGatewayDataRelationshipsStripePayments**](StripeGatewayDataRelationshipsStripePayments.md) |  | [optional] 

## Methods

### NewStripeGatewayDataRelationships

`func NewStripeGatewayDataRelationships() *StripeGatewayDataRelationships`

NewStripeGatewayDataRelationships instantiates a new StripeGatewayDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeGatewayDataRelationshipsWithDefaults

`func NewStripeGatewayDataRelationshipsWithDefaults() *StripeGatewayDataRelationships`

NewStripeGatewayDataRelationshipsWithDefaults instantiates a new StripeGatewayDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *StripeGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *StripeGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *StripeGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *StripeGatewayDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetStripePayments

`func (o *StripeGatewayDataRelationships) GetStripePayments() StripeGatewayDataRelationshipsStripePayments`

GetStripePayments returns the StripePayments field if non-nil, zero value otherwise.

### GetStripePaymentsOk

`func (o *StripeGatewayDataRelationships) GetStripePaymentsOk() (*StripeGatewayDataRelationshipsStripePayments, bool)`

GetStripePaymentsOk returns a tuple with the StripePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePayments

`func (o *StripeGatewayDataRelationships) SetStripePayments(v StripeGatewayDataRelationshipsStripePayments)`

SetStripePayments sets StripePayments field to given value.

### HasStripePayments

`func (o *StripeGatewayDataRelationships) HasStripePayments() bool`

HasStripePayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


