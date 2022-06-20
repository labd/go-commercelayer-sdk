# CustomerPaymentSourceCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | 
**PaymentSource** | [**CustomerPaymentSourceDataRelationshipsPaymentSource**](CustomerPaymentSourceDataRelationshipsPaymentSource.md) |  | 

## Methods

### NewCustomerPaymentSourceCreateDataRelationships

`func NewCustomerPaymentSourceCreateDataRelationships(customer CouponRecipientDataRelationshipsCustomer, paymentSource CustomerPaymentSourceDataRelationshipsPaymentSource, ) *CustomerPaymentSourceCreateDataRelationships`

NewCustomerPaymentSourceCreateDataRelationships instantiates a new CustomerPaymentSourceCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceCreateDataRelationshipsWithDefaults

`func NewCustomerPaymentSourceCreateDataRelationshipsWithDefaults() *CustomerPaymentSourceCreateDataRelationships`

NewCustomerPaymentSourceCreateDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPaymentSourceCreateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.


### GetPaymentSource

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *CustomerPaymentSourceCreateDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


