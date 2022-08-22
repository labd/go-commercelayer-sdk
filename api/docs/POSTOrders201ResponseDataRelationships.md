# POSTOrders201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets**](GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets.md) |  | [optional] 
**Customer** | Pointer to [**GETCouponRecipients200ResponseDataInnerRelationshipsCustomer**](GETCouponRecipients200ResponseDataInnerRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**GETBingGeocoders200ResponseDataInnerRelationshipsAddresses**](GETBingGeocoders200ResponseDataInnerRelationshipsAddresses.md) |  | [optional] 
**BillingAddress** | Pointer to [**GETBingGeocoders200ResponseDataInnerRelationshipsAddresses**](GETBingGeocoders200ResponseDataInnerRelationshipsAddresses.md) |  | [optional] 
**PaymentMethod** | Pointer to [**GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods**](GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource**](GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource.md) |  | [optional] 

## Methods

### NewPOSTOrders201ResponseDataRelationships

`func NewPOSTOrders201ResponseDataRelationships() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationships instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrders201ResponseDataRelationshipsWithDefaults

`func NewPOSTOrders201ResponseDataRelationshipsWithDefaults() *POSTOrders201ResponseDataRelationships`

NewPOSTOrders201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrders201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTOrders201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTOrders201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTOrders201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTOrders201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *POSTOrders201ResponseDataRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *POSTOrders201ResponseDataRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *POSTOrders201ResponseDataRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *POSTOrders201ResponseDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddress() GETBingGeocoders200ResponseDataInnerRelationshipsAddresses`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetShippingAddressOk() (*GETBingGeocoders200ResponseDataInnerRelationshipsAddresses, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetShippingAddress(v GETBingGeocoders200ResponseDataInnerRelationshipsAddresses)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddress() GETBingGeocoders200ResponseDataInnerRelationshipsAddresses`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *POSTOrders201ResponseDataRelationships) GetBillingAddressOk() (*GETBingGeocoders200ResponseDataInnerRelationshipsAddresses, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) SetBillingAddress(v GETBingGeocoders200ResponseDataInnerRelationshipsAddresses)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *POSTOrders201ResponseDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethod() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentMethodOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentMethod(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSource() GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *POSTOrders201ResponseDataRelationships) GetPaymentSourceOk() (*GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) SetPaymentSource(v GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *POSTOrders201ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


