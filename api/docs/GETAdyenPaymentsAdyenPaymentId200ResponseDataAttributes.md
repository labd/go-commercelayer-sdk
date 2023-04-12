# GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **interface{}** | The public key linked to your API credential. | [optional] 
**PaymentMethods** | Pointer to **interface{}** | The merchant available payment methods for the assoiated order (i.e. country and amount). Required by the Adyen JS SDK. | [optional] 
**PaymentRequestData** | Pointer to **interface{}** | The Adyen payment request data, collected by client. | [optional] 
**PaymentRequestDetails** | Pointer to **interface{}** | The Adyen additional details request data, collected by client. | [optional] 
**PaymentResponse** | Pointer to **interface{}** | The Adyen payment response, used by client (includes &#39;resultCode&#39; and &#39;action&#39;). | [optional] 
**MismatchedAmounts** | Pointer to **interface{}** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes

`func NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes() *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes`

NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes instantiates a new GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributesWithDefaults

`func NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributesWithDefaults() *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes`

NewGETAdyenPaymentsAdyenPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPublicKey() interface{}`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPublicKeyOk() (*interface{}, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPublicKey(v interface{})`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPaymentMethods

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentMethods() interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentMethodsOk() (*interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentMethods(v interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### SetPaymentMethodsNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentMethodsNil(b bool)`

 SetPaymentMethodsNil sets the value for PaymentMethods to be an explicit nil

### UnsetPaymentMethods
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPaymentMethods()`

UnsetPaymentMethods ensures that no value is present for PaymentMethods, not even an explicit nil
### GetPaymentRequestData

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentRequestData() interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentRequestDataOk() (*interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentRequestData(v interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### SetPaymentRequestDataNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentRequestDataNil(b bool)`

 SetPaymentRequestDataNil sets the value for PaymentRequestData to be an explicit nil

### UnsetPaymentRequestData
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPaymentRequestData()`

UnsetPaymentRequestData ensures that no value is present for PaymentRequestData, not even an explicit nil
### GetPaymentRequestDetails

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentRequestDetails() interface{}`

GetPaymentRequestDetails returns the PaymentRequestDetails field if non-nil, zero value otherwise.

### GetPaymentRequestDetailsOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentRequestDetailsOk() (*interface{}, bool)`

GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestDetails

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentRequestDetails(v interface{})`

SetPaymentRequestDetails sets PaymentRequestDetails field to given value.

### HasPaymentRequestDetails

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPaymentRequestDetails() bool`

HasPaymentRequestDetails returns a boolean if a field has been set.

### SetPaymentRequestDetailsNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentRequestDetailsNil(b bool)`

 SetPaymentRequestDetailsNil sets the value for PaymentRequestDetails to be an explicit nil

### UnsetPaymentRequestDetails
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPaymentRequestDetails()`

UnsetPaymentRequestDetails ensures that no value is present for PaymentRequestDetails, not even an explicit nil
### GetPaymentResponse

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentResponse() interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentResponseOk() (*interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentResponse(v interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### SetPaymentResponseNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentResponseNil(b bool)`

 SetPaymentResponseNil sets the value for PaymentResponse to be an explicit nil

### UnsetPaymentResponse
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPaymentResponse()`

UnsetPaymentResponse ensures that no value is present for PaymentResponse, not even an explicit nil
### GetMismatchedAmounts

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetMismatchedAmounts() interface{}`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetMismatchedAmountsOk() (*interface{}, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetMismatchedAmounts(v interface{})`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### SetMismatchedAmountsNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetMismatchedAmountsNil(b bool)`

 SetMismatchedAmountsNil sets the value for MismatchedAmounts to be an explicit nil

### UnsetMismatchedAmounts
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetMismatchedAmounts()`

UnsetMismatchedAmounts ensures that no value is present for MismatchedAmounts, not even an explicit nil
### GetPaymentInstrument

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETAdyenPaymentsAdyenPaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


