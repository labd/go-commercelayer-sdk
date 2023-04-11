# PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentRequestData** | Pointer to **interface{}** | The Adyen payment request data, collected by client. | [optional] 
**PaymentRequestDetails** | Pointer to **interface{}** | The Adyen additional details request data, collected by client. | [optional] 
**PaymentResponse** | Pointer to **interface{}** | The Adyen payment response, used by client (includes &#39;resultCode&#39; and &#39;action&#39;). | [optional] 
**Details** | Pointer to **interface{}** | Send this attribute if you want to send additional details the payment request. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes

`func NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes() *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes`

NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes instantiates a new PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributesWithDefaults

`func NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributesWithDefaults() *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes`

NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentRequestData

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentRequestData() interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentRequestDataOk() (*interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentRequestData(v interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### SetPaymentRequestDataNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentRequestDataNil(b bool)`

 SetPaymentRequestDataNil sets the value for PaymentRequestData to be an explicit nil

### UnsetPaymentRequestData
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetPaymentRequestData()`

UnsetPaymentRequestData ensures that no value is present for PaymentRequestData, not even an explicit nil
### GetPaymentRequestDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentRequestDetails() interface{}`

GetPaymentRequestDetails returns the PaymentRequestDetails field if non-nil, zero value otherwise.

### GetPaymentRequestDetailsOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentRequestDetailsOk() (*interface{}, bool)`

GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentRequestDetails(v interface{})`

SetPaymentRequestDetails sets PaymentRequestDetails field to given value.

### HasPaymentRequestDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasPaymentRequestDetails() bool`

HasPaymentRequestDetails returns a boolean if a field has been set.

### SetPaymentRequestDetailsNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentRequestDetailsNil(b bool)`

 SetPaymentRequestDetailsNil sets the value for PaymentRequestDetails to be an explicit nil

### UnsetPaymentRequestDetails
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetPaymentRequestDetails()`

UnsetPaymentRequestDetails ensures that no value is present for PaymentRequestDetails, not even an explicit nil
### GetPaymentResponse

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentResponse() interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetPaymentResponseOk() (*interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentResponse(v interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### SetPaymentResponseNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetPaymentResponseNil(b bool)`

 SetPaymentResponseNil sets the value for PaymentResponse to be an explicit nil

### UnsetPaymentResponse
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetPaymentResponse()`

UnsetPaymentResponse ensures that no value is present for PaymentResponse, not even an explicit nil
### GetDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetReference

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


