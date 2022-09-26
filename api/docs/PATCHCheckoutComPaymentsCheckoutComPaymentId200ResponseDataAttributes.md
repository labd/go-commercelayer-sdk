# PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes

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

### NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes

`func NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes() *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes`

NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes instantiates a new PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributesWithDefaults

`func NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributesWithDefaults() *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes`

NewPATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentType

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetDetails() bool`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetDetailsOk() (*bool, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetDetails(v bool)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


