# PaymentMethodResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BillingInfoValidationRuleResponseDataRelationshipsMarket**](BillingInfoValidationRuleResponseDataRelationshipsMarket.md) |  | [optional] 
**PaymentGateway** | Pointer to [**AdyenPaymentResponseDataRelationshipsPaymentGateway**](AdyenPaymentResponseDataRelationshipsPaymentGateway.md) |  | [optional] 
**Attachments** | Pointer to [**AvalaraAccountResponseDataRelationshipsAttachments**](AvalaraAccountResponseDataRelationshipsAttachments.md) |  | [optional] 

## Methods

### NewPaymentMethodResponseDataRelationships

`func NewPaymentMethodResponseDataRelationships() *PaymentMethodResponseDataRelationships`

NewPaymentMethodResponseDataRelationships instantiates a new PaymentMethodResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodResponseDataRelationshipsWithDefaults

`func NewPaymentMethodResponseDataRelationshipsWithDefaults() *PaymentMethodResponseDataRelationships`

NewPaymentMethodResponseDataRelationshipsWithDefaults instantiates a new PaymentMethodResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *PaymentMethodResponseDataRelationships) GetMarket() BillingInfoValidationRuleResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PaymentMethodResponseDataRelationships) GetMarketOk() (*BillingInfoValidationRuleResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PaymentMethodResponseDataRelationships) SetMarket(v BillingInfoValidationRuleResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *PaymentMethodResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *PaymentMethodResponseDataRelationships) GetPaymentGateway() AdyenPaymentResponseDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *PaymentMethodResponseDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentResponseDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *PaymentMethodResponseDataRelationships) SetPaymentGateway(v AdyenPaymentResponseDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *PaymentMethodResponseDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetAttachments

`func (o *PaymentMethodResponseDataRelationships) GetAttachments() AvalaraAccountResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PaymentMethodResponseDataRelationships) GetAttachmentsOk() (*AvalaraAccountResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PaymentMethodResponseDataRelationships) SetAttachments(v AvalaraAccountResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PaymentMethodResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


