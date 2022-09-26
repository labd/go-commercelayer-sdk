# ExternalGatewayResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayResponseDataRelationshipsPaymentMethods**](AdyenGatewayResponseDataRelationshipsPaymentMethods.md) |  | [optional] 
**ExternalPayments** | Pointer to [**ExternalGatewayResponseDataRelationshipsExternalPayments**](ExternalGatewayResponseDataRelationshipsExternalPayments.md) |  | [optional] 

## Methods

### NewExternalGatewayResponseDataRelationships

`func NewExternalGatewayResponseDataRelationships() *ExternalGatewayResponseDataRelationships`

NewExternalGatewayResponseDataRelationships instantiates a new ExternalGatewayResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayResponseDataRelationshipsWithDefaults

`func NewExternalGatewayResponseDataRelationshipsWithDefaults() *ExternalGatewayResponseDataRelationships`

NewExternalGatewayResponseDataRelationshipsWithDefaults instantiates a new ExternalGatewayResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *ExternalGatewayResponseDataRelationships) GetPaymentMethods() AdyenGatewayResponseDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *ExternalGatewayResponseDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayResponseDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *ExternalGatewayResponseDataRelationships) SetPaymentMethods(v AdyenGatewayResponseDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *ExternalGatewayResponseDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetExternalPayments

`func (o *ExternalGatewayResponseDataRelationships) GetExternalPayments() ExternalGatewayResponseDataRelationshipsExternalPayments`

GetExternalPayments returns the ExternalPayments field if non-nil, zero value otherwise.

### GetExternalPaymentsOk

`func (o *ExternalGatewayResponseDataRelationships) GetExternalPaymentsOk() (*ExternalGatewayResponseDataRelationshipsExternalPayments, bool)`

GetExternalPaymentsOk returns a tuple with the ExternalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayments

`func (o *ExternalGatewayResponseDataRelationships) SetExternalPayments(v ExternalGatewayResponseDataRelationshipsExternalPayments)`

SetExternalPayments sets ExternalPayments field to given value.

### HasExternalPayments

`func (o *ExternalGatewayResponseDataRelationships) HasExternalPayments() bool`

HasExternalPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


