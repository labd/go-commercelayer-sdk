# CustomerPaymentSourceCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [**CouponRecipientCreateDataRelationshipsCustomer**](CouponRecipientCreateDataRelationshipsCustomer.md) |  | 
**PaymentMethod** | Pointer to [**CustomerPaymentSourceCreateDataRelationshipsPaymentMethod**](CustomerPaymentSourceCreateDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceCreateDataRelationshipsPaymentSource**](CustomerPaymentSourceCreateDataRelationshipsPaymentSource.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceCreateDataRelationships

`func NewCustomerPaymentSourceCreateDataRelationships(customer CouponRecipientCreateDataRelationshipsCustomer, ) *CustomerPaymentSourceCreateDataRelationships`

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

`func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPaymentSourceCreateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPaymentSourceCreateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.


### GetPaymentMethod

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentMethod() CustomerPaymentSourceCreateDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentMethodOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CustomerPaymentSourceCreateDataRelationships) SetPaymentMethod(v CustomerPaymentSourceCreateDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CustomerPaymentSourceCreateDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSource() CustomerPaymentSourceCreateDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *CustomerPaymentSourceCreateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *CustomerPaymentSourceCreateDataRelationships) SetPaymentSource(v CustomerPaymentSourceCreateDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *CustomerPaymentSourceCreateDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


