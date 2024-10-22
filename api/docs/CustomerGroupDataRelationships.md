# CustomerGroupDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**Markets** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewCustomerGroupDataRelationships

`func NewCustomerGroupDataRelationships() *CustomerGroupDataRelationships`

NewCustomerGroupDataRelationships instantiates a new CustomerGroupDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerGroupDataRelationshipsWithDefaults

`func NewCustomerGroupDataRelationshipsWithDefaults() *CustomerGroupDataRelationships`

NewCustomerGroupDataRelationshipsWithDefaults instantiates a new CustomerGroupDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *CustomerGroupDataRelationships) GetCustomers() CouponRecipientDataRelationshipsCustomer`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *CustomerGroupDataRelationships) GetCustomersOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *CustomerGroupDataRelationships) SetCustomers(v CouponRecipientDataRelationshipsCustomer)`

SetCustomers sets Customers field to given value.

### HasCustomers

`func (o *CustomerGroupDataRelationships) HasCustomers() bool`

HasCustomers returns a boolean if a field has been set.

### GetMarkets

`func (o *CustomerGroupDataRelationships) GetMarkets() AvalaraAccountDataRelationshipsMarkets`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *CustomerGroupDataRelationships) GetMarketsOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *CustomerGroupDataRelationships) SetMarkets(v AvalaraAccountDataRelationshipsMarkets)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *CustomerGroupDataRelationships) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetAttachments

`func (o *CustomerGroupDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CustomerGroupDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CustomerGroupDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CustomerGroupDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *CustomerGroupDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CustomerGroupDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CustomerGroupDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CustomerGroupDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


