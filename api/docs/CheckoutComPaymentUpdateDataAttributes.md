# CheckoutComPaymentUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | Pointer to **string** | The payment source type. | [optional] 
**Token** | Pointer to **string** | The Checkout.com card or digital wallet token. | [optional] 
**SessionId** | Pointer to **string** | A payment session ID used to obtain the details. | [optional] 
**SuccessUrl** | Pointer to **string** | The URL to redirect your customer upon 3DS succeeded authentication. | [optional] 
**FailureUrl** | Pointer to **string** | The URL to redirect your customer upon 3DS failed authentication. | [optional] 
**Details** | Pointer to **bool** | Send this attribute if you want to send additional details the payment request (i.e. upon 3DS check). | [optional] 
**Refresh** | Pointer to **bool** | Send this attribute if you want to refresh all the pending transactions, can be used as webhooks fallback logic. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewCheckoutComPaymentUpdateDataAttributes

`func NewCheckoutComPaymentUpdateDataAttributes() *CheckoutComPaymentUpdateDataAttributes`

NewCheckoutComPaymentUpdateDataAttributes instantiates a new CheckoutComPaymentUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComPaymentUpdateDataAttributesWithDefaults

`func NewCheckoutComPaymentUpdateDataAttributesWithDefaults() *CheckoutComPaymentUpdateDataAttributes`

NewCheckoutComPaymentUpdateDataAttributesWithDefaults instantiates a new CheckoutComPaymentUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *CheckoutComPaymentUpdateDataAttributes) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *CheckoutComPaymentUpdateDataAttributes) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *CheckoutComPaymentUpdateDataAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetToken

`func (o *CheckoutComPaymentUpdateDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckoutComPaymentUpdateDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CheckoutComPaymentUpdateDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSessionId

`func (o *CheckoutComPaymentUpdateDataAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CheckoutComPaymentUpdateDataAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CheckoutComPaymentUpdateDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *CheckoutComPaymentUpdateDataAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetDetails

`func (o *CheckoutComPaymentUpdateDataAttributes) GetDetails() bool`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetDetailsOk() (*bool, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CheckoutComPaymentUpdateDataAttributes) SetDetails(v bool)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CheckoutComPaymentUpdateDataAttributes) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRefresh

`func (o *CheckoutComPaymentUpdateDataAttributes) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *CheckoutComPaymentUpdateDataAttributes) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *CheckoutComPaymentUpdateDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetReference

`func (o *CheckoutComPaymentUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CheckoutComPaymentUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CheckoutComPaymentUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CheckoutComPaymentUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CheckoutComPaymentUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckoutComPaymentUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutComPaymentUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutComPaymentUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutComPaymentUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


