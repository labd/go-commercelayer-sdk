# PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentType** | Pointer to **interface{}** | The payment source type. | [optional] 
**Token** | Pointer to **interface{}** | The Checkout.com card or digital wallet token. | [optional] 
**SessionId** | Pointer to **interface{}** | A payment session ID used to obtain the details. | [optional] 
**SuccessUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS succeeded authentication. | [optional] 
**FailureUrl** | Pointer to **interface{}** | The URL to redirect your customer upon 3DS failed authentication. | [optional] 
**Details** | Pointer to **interface{}** | Send this attribute if you want to send additional details the payment request (i.e. upon 3DS check). | [optional] 
**Refresh** | Pointer to **interface{}** | Send this attribute if you want to refresh all the pending transactions, can be used as webhooks fallback logic. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

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

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentType() interface{}`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetPaymentTypeOk() (*interface{}, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentType(v interface{})`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### SetPaymentTypeNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetPaymentTypeNil(b bool)`

 SetPaymentTypeNil sets the value for PaymentType to be an explicit nil

### UnsetPaymentType
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetPaymentType()`

UnsetPaymentType ensures that no value is present for PaymentType, not even an explicit nil
### GetToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetToken(v interface{})`

SetToken sets Token field to given value.

### HasToken

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionId() interface{}`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSessionIdOk() (*interface{}, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSessionId(v interface{})`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrl() interface{}`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetSuccessUrlOk() (*interface{}, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSuccessUrl(v interface{})`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrl() interface{}`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetFailureUrlOk() (*interface{}, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetFailureUrl(v interface{})`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### SetFailureUrlNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetFailureUrlNil(b bool)`

 SetFailureUrlNil sets the value for FailureUrl to be an explicit nil

### UnsetFailureUrl
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetFailureUrl()`

UnsetFailureUrl ensures that no value is present for FailureUrl, not even an explicit nil
### GetDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRefresh() interface{}`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetRefreshOk() (*interface{}, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetRefresh(v interface{})`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


