# GETAxervePaymentsAxervePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **interface{}** | The merchant login code. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The URL where the payer is redirected after they approve the payment. | [optional] 
**PaymentRequestData** | Pointer to **interface{}** | The Axerve payment request data, collected by client. | [optional] 
**ClientIp** | Pointer to **interface{}** | The IP adress of the client creating the payment. | [optional] 
**BuyerDetails** | Pointer to **interface{}** | The details of the buyer creating the payment. | [optional] 
**RequestToken** | Pointer to **interface{}** | Requires the creation of a token to represent this payment, mandatory to use customer&#39;s wallet and order subscriptions. | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributes

`func NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributes() *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes`

NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributes instantiates a new GETAxervePaymentsAxervePaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributesWithDefaults

`func NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributesWithDefaults() *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes`

NewGETAxervePaymentsAxervePaymentId200ResponseDataAttributesWithDefaults instantiates a new GETAxervePaymentsAxervePaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetLogin() interface{}`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetLoginOk() (*interface{}, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetLogin(v interface{})`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetReturnUrl

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetPaymentRequestData

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentRequestData() interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentRequestDataOk() (*interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetPaymentRequestData(v interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### SetPaymentRequestDataNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetPaymentRequestDataNil(b bool)`

 SetPaymentRequestDataNil sets the value for PaymentRequestData to be an explicit nil

### UnsetPaymentRequestData
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetPaymentRequestData()`

UnsetPaymentRequestData ensures that no value is present for PaymentRequestData, not even an explicit nil
### GetClientIp

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetClientIp() interface{}`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetClientIpOk() (*interface{}, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetClientIp(v interface{})`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetBuyerDetails

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetBuyerDetails() interface{}`

GetBuyerDetails returns the BuyerDetails field if non-nil, zero value otherwise.

### GetBuyerDetailsOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetBuyerDetailsOk() (*interface{}, bool)`

GetBuyerDetailsOk returns a tuple with the BuyerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerDetails

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetBuyerDetails(v interface{})`

SetBuyerDetails sets BuyerDetails field to given value.

### HasBuyerDetails

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasBuyerDetails() bool`

HasBuyerDetails returns a boolean if a field has been set.

### SetBuyerDetailsNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetBuyerDetailsNil(b bool)`

 SetBuyerDetailsNil sets the value for BuyerDetails to be an explicit nil

### UnsetBuyerDetails
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetBuyerDetails()`

UnsetBuyerDetails ensures that no value is present for BuyerDetails, not even an explicit nil
### GetRequestToken

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetRequestToken() interface{}`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetRequestTokenOk() (*interface{}, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetRequestToken(v interface{})`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetPaymentInstrument

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAxervePaymentsAxervePaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


