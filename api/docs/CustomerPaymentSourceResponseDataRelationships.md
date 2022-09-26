# CustomerPaymentSourceResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientResponseDataRelationshipsCustomer**](CouponRecipientResponseDataRelationshipsCustomer.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceResponseDataRelationshipsPaymentSource**](CustomerPaymentSourceResponseDataRelationshipsPaymentSource.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceResponseDataRelationships

`func NewCustomerPaymentSourceResponseDataRelationships() *CustomerPaymentSourceResponseDataRelationships`

NewCustomerPaymentSourceResponseDataRelationships instantiates a new CustomerPaymentSourceResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceResponseDataRelationshipsWithDefaults

`func NewCustomerPaymentSourceResponseDataRelationshipsWithDefaults() *CustomerPaymentSourceResponseDataRelationships`

NewCustomerPaymentSourceResponseDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerPaymentSourceResponseDataRelationships) GetCustomer() CouponRecipientResponseDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPaymentSourceResponseDataRelationships) GetCustomerOk() (*CouponRecipientResponseDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPaymentSourceResponseDataRelationships) SetCustomer(v CouponRecipientResponseDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerPaymentSourceResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetPaymentSource

`func (o *CustomerPaymentSourceResponseDataRelationships) GetPaymentSource() CustomerPaymentSourceResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *CustomerPaymentSourceResponseDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *CustomerPaymentSourceResponseDataRelationships) SetPaymentSource(v CustomerPaymentSourceResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *CustomerPaymentSourceResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


