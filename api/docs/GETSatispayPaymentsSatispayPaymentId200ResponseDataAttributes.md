# GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **interface{}** | The payment unique identifier. | [optional] 
**Flow** | Pointer to **interface{}** | The Satispay payment flow, inspect gateway API details for more information. | [optional] 
**Status** | Pointer to **interface{}** | The Satispay payment status. | [optional] 
**RedirectUrl** | Pointer to **interface{}** | The url to redirect the customer after the payment flow is completed. | [optional] 
**PaymentUrl** | Pointer to **interface{}** | Redirect url to the payment page. | [optional] 
**PaymentResponse** | Pointer to **interface{}** | The Satispay payment response, used to fetch internal data. | [optional] 
**PaymentInstrument** | Pointer to **interface{}** | Information about the payment instrument used in the transaction. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes

`func NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes() *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes`

NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes instantiates a new GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults

`func NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults() *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes`

NewGETSatispayPaymentsSatispayPaymentId200ResponseDataAttributesWithDefaults instantiates a new GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentId() interface{}`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentIdOk() (*interface{}, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentId(v interface{})`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetFlow

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetFlow() interface{}`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetFlowOk() (*interface{}, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetFlow(v interface{})`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### SetFlowNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetFlowNil(b bool)`

 SetFlowNil sets the value for Flow to be an explicit nil

### UnsetFlow
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetFlow()`

UnsetFlow ensures that no value is present for Flow, not even an explicit nil
### GetStatus

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRedirectUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrl() interface{}`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetRedirectUrlOk() (*interface{}, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetRedirectUrl(v interface{})`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetPaymentUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentUrl() interface{}`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentUrlOk() (*interface{}, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentUrl(v interface{})`

SetPaymentUrl sets PaymentUrl field to given value.

### HasPaymentUrl

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasPaymentUrl() bool`

HasPaymentUrl returns a boolean if a field has been set.

### SetPaymentUrlNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentUrlNil(b bool)`

 SetPaymentUrlNil sets the value for PaymentUrl to be an explicit nil

### UnsetPaymentUrl
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetPaymentUrl()`

UnsetPaymentUrl ensures that no value is present for PaymentUrl, not even an explicit nil
### GetPaymentResponse

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentResponse() interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentResponseOk() (*interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentResponse(v interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### SetPaymentResponseNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentResponseNil(b bool)`

 SetPaymentResponseNil sets the value for PaymentResponse to be an explicit nil

### UnsetPaymentResponse
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetPaymentResponse()`

UnsetPaymentResponse ensures that no value is present for PaymentResponse, not even an explicit nil
### GetPaymentInstrument

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentInstrument() interface{}`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetPaymentInstrumentOk() (*interface{}, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentInstrument(v interface{})`

SetPaymentInstrument sets PaymentInstrument field to given value.

### HasPaymentInstrument

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasPaymentInstrument() bool`

HasPaymentInstrument returns a boolean if a field has been set.

### SetPaymentInstrumentNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetPaymentInstrumentNil(b bool)`

 SetPaymentInstrumentNil sets the value for PaymentInstrument to be an explicit nil

### UnsetPaymentInstrument
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetPaymentInstrument()`

UnsetPaymentInstrument ensures that no value is present for PaymentInstrument, not even an explicit nil
### GetCreatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETSatispayPaymentsSatispayPaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


