# POSTOrdersRequestDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBillingInfoValidationRulesRequestDataRelationshipsMarket**](POSTBillingInfoValidationRulesRequestDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**POSTCouponRecipientsRequestDataRelationshipsCustomer**](POSTCouponRecipientsRequestDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**POSTCustomerAddressesRequestDataRelationshipsAddress**](POSTCustomerAddressesRequestDataRelationshipsAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**POSTCustomerAddressesRequestDataRelationshipsAddress**](POSTCustomerAddressesRequestDataRelationshipsAddress.md) |  | [optional] 
**PaymentMethod** | Pointer to [**POSTOrdersRequestDataRelationshipsPaymentMethod**](POSTOrdersRequestDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource**](POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource.md) |  | [optional] 

## Methods

### NewPOSTOrdersRequestDataRelationships

`func NewPOSTOrdersRequestDataRelationships() *POSTOrdersRequestDataRelationships`

NewPOSTOrdersRequestDataRelationships instantiates a new POSTOrdersRequestDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrdersRequestDataRelationshipsWithDefaults

`func NewPOSTOrdersRequestDataRelationshipsWithDefaults() *POSTOrdersRequestDataRelationships`

NewPOSTOrdersRequestDataRelationshipsWithDefaults instantiates a new POSTOrdersRequestDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrdersRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrdersRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrdersRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrdersRequestDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTOrdersRequestDataRelationships) GetCustomer() POSTCouponRecipientsRequestDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTOrdersRequestDataRelationships) GetCustomerOk() (*POSTCouponRecipientsRequestDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTOrdersRequestDataRelationships) SetCustomer(v POSTCouponRecipientsRequestDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTOrdersRequestDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *POSTOrdersRequestDataRelationships) GetShippingAddress() POSTCustomerAddressesRequestDataRelationshipsAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *POSTOrdersRequestDataRelationships) GetShippingAddressOk() (*POSTCustomerAddressesRequestDataRelationshipsAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *POSTOrdersRequestDataRelationships) SetShippingAddress(v POSTCustomerAddressesRequestDataRelationshipsAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *POSTOrdersRequestDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *POSTOrdersRequestDataRelationships) GetBillingAddress() POSTCustomerAddressesRequestDataRelationshipsAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *POSTOrdersRequestDataRelationships) GetBillingAddressOk() (*POSTCustomerAddressesRequestDataRelationshipsAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *POSTOrdersRequestDataRelationships) SetBillingAddress(v POSTCustomerAddressesRequestDataRelationshipsAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *POSTOrdersRequestDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *POSTOrdersRequestDataRelationships) GetPaymentMethod() POSTOrdersRequestDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *POSTOrdersRequestDataRelationships) GetPaymentMethodOk() (*POSTOrdersRequestDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *POSTOrdersRequestDataRelationships) SetPaymentMethod(v POSTOrdersRequestDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *POSTOrdersRequestDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *POSTOrdersRequestDataRelationships) GetPaymentSource() POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *POSTOrdersRequestDataRelationships) GetPaymentSourceOk() (*POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *POSTOrdersRequestDataRelationships) SetPaymentSource(v POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *POSTOrdersRequestDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


