# ExternalGatewayDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**ExternalPayments** | Pointer to [**ExternalGatewayDataRelationshipsExternalPayments**](ExternalGatewayDataRelationshipsExternalPayments.md) |  | [optional] 

## Methods

### NewExternalGatewayDataRelationships

`func NewExternalGatewayDataRelationships() *ExternalGatewayDataRelationships`

NewExternalGatewayDataRelationships instantiates a new ExternalGatewayDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayDataRelationshipsWithDefaults

`func NewExternalGatewayDataRelationshipsWithDefaults() *ExternalGatewayDataRelationships`

NewExternalGatewayDataRelationshipsWithDefaults instantiates a new ExternalGatewayDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *ExternalGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *ExternalGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *ExternalGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *ExternalGatewayDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetVersions

`func (o *ExternalGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ExternalGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ExternalGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ExternalGatewayDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetExternalPayments

`func (o *ExternalGatewayDataRelationships) GetExternalPayments() ExternalGatewayDataRelationshipsExternalPayments`

GetExternalPayments returns the ExternalPayments field if non-nil, zero value otherwise.

### GetExternalPaymentsOk

`func (o *ExternalGatewayDataRelationships) GetExternalPaymentsOk() (*ExternalGatewayDataRelationshipsExternalPayments, bool)`

GetExternalPaymentsOk returns a tuple with the ExternalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayments

`func (o *ExternalGatewayDataRelationships) SetExternalPayments(v ExternalGatewayDataRelationshipsExternalPayments)`

SetExternalPayments sets ExternalPayments field to given value.

### HasExternalPayments

`func (o *ExternalGatewayDataRelationships) HasExternalPayments() bool`

HasExternalPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


