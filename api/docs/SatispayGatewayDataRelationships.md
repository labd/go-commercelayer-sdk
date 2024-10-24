# SatispayGatewayDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**SatispayPayments** | Pointer to [**SatispayGatewayDataRelationshipsSatispayPayments**](SatispayGatewayDataRelationshipsSatispayPayments.md) |  | [optional] 

## Methods

### NewSatispayGatewayDataRelationships

`func NewSatispayGatewayDataRelationships() *SatispayGatewayDataRelationships`

NewSatispayGatewayDataRelationships instantiates a new SatispayGatewayDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayGatewayDataRelationshipsWithDefaults

`func NewSatispayGatewayDataRelationshipsWithDefaults() *SatispayGatewayDataRelationships`

NewSatispayGatewayDataRelationshipsWithDefaults instantiates a new SatispayGatewayDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *SatispayGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *SatispayGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *SatispayGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *SatispayGatewayDataRelationships) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetVersions

`func (o *SatispayGatewayDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *SatispayGatewayDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *SatispayGatewayDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *SatispayGatewayDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetSatispayPayments

`func (o *SatispayGatewayDataRelationships) GetSatispayPayments() SatispayGatewayDataRelationshipsSatispayPayments`

GetSatispayPayments returns the SatispayPayments field if non-nil, zero value otherwise.

### GetSatispayPaymentsOk

`func (o *SatispayGatewayDataRelationships) GetSatispayPaymentsOk() (*SatispayGatewayDataRelationshipsSatispayPayments, bool)`

GetSatispayPaymentsOk returns a tuple with the SatispayPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatispayPayments

`func (o *SatispayGatewayDataRelationships) SetSatispayPayments(v SatispayGatewayDataRelationshipsSatispayPayments)`

SetSatispayPayments sets SatispayPayments field to given value.

### HasSatispayPayments

`func (o *SatispayGatewayDataRelationships) HasSatispayPayments() bool`

HasSatispayPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


