# AdyenPaymentDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | The public key linked to your API credential. | [optional] 
**PaymentMethods** | Pointer to **map[string]interface{}** | The merchant available payment methods for the assoiated order (i.e. country and amount). Required by the Adyen JS SDK. | [optional] 
**PaymentRequestData** | Pointer to **map[string]interface{}** | The Adyen payment request data, collected by client. | [optional] 
**PaymentRequestDetails** | Pointer to **map[string]interface{}** | The Adyen additional details request data, collected by client. | [optional] 
**PaymentResponse** | Pointer to **map[string]interface{}** | The Adyen payment response, used by client (includes &#39;resultCode&#39; and &#39;action&#39;). | [optional] 
**MismatchedAmounts** | Pointer to **bool** | Indicates if the order current amount differs form the one of the associated authorization. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewAdyenPaymentDataAttributes

`func NewAdyenPaymentDataAttributes() *AdyenPaymentDataAttributes`

NewAdyenPaymentDataAttributes instantiates a new AdyenPaymentDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenPaymentDataAttributesWithDefaults

`func NewAdyenPaymentDataAttributesWithDefaults() *AdyenPaymentDataAttributes`

NewAdyenPaymentDataAttributesWithDefaults instantiates a new AdyenPaymentDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *AdyenPaymentDataAttributes) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AdyenPaymentDataAttributes) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AdyenPaymentDataAttributes) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AdyenPaymentDataAttributes) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *AdyenPaymentDataAttributes) GetPaymentMethods() map[string]interface{}`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *AdyenPaymentDataAttributes) GetPaymentMethodsOk() (*map[string]interface{}, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *AdyenPaymentDataAttributes) SetPaymentMethods(v map[string]interface{})`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *AdyenPaymentDataAttributes) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPaymentRequestData

`func (o *AdyenPaymentDataAttributes) GetPaymentRequestData() map[string]interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *AdyenPaymentDataAttributes) GetPaymentRequestDataOk() (*map[string]interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *AdyenPaymentDataAttributes) SetPaymentRequestData(v map[string]interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *AdyenPaymentDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### GetPaymentRequestDetails

`func (o *AdyenPaymentDataAttributes) GetPaymentRequestDetails() map[string]interface{}`

GetPaymentRequestDetails returns the PaymentRequestDetails field if non-nil, zero value otherwise.

### GetPaymentRequestDetailsOk

`func (o *AdyenPaymentDataAttributes) GetPaymentRequestDetailsOk() (*map[string]interface{}, bool)`

GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestDetails

`func (o *AdyenPaymentDataAttributes) SetPaymentRequestDetails(v map[string]interface{})`

SetPaymentRequestDetails sets PaymentRequestDetails field to given value.

### HasPaymentRequestDetails

`func (o *AdyenPaymentDataAttributes) HasPaymentRequestDetails() bool`

HasPaymentRequestDetails returns a boolean if a field has been set.

### GetPaymentResponse

`func (o *AdyenPaymentDataAttributes) GetPaymentResponse() map[string]interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *AdyenPaymentDataAttributes) GetPaymentResponseOk() (*map[string]interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *AdyenPaymentDataAttributes) SetPaymentResponse(v map[string]interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *AdyenPaymentDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### GetMismatchedAmounts

`func (o *AdyenPaymentDataAttributes) GetMismatchedAmounts() bool`

GetMismatchedAmounts returns the MismatchedAmounts field if non-nil, zero value otherwise.

### GetMismatchedAmountsOk

`func (o *AdyenPaymentDataAttributes) GetMismatchedAmountsOk() (*bool, bool)`

GetMismatchedAmountsOk returns a tuple with the MismatchedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismatchedAmounts

`func (o *AdyenPaymentDataAttributes) SetMismatchedAmounts(v bool)`

SetMismatchedAmounts sets MismatchedAmounts field to given value.

### HasMismatchedAmounts

`func (o *AdyenPaymentDataAttributes) HasMismatchedAmounts() bool`

HasMismatchedAmounts returns a boolean if a field has been set.

### GetId

`func (o *AdyenPaymentDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdyenPaymentDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdyenPaymentDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdyenPaymentDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdyenPaymentDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdyenPaymentDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdyenPaymentDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdyenPaymentDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdyenPaymentDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdyenPaymentDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdyenPaymentDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdyenPaymentDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *AdyenPaymentDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdyenPaymentDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdyenPaymentDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdyenPaymentDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AdyenPaymentDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AdyenPaymentDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AdyenPaymentDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AdyenPaymentDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AdyenPaymentDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdyenPaymentDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdyenPaymentDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdyenPaymentDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


