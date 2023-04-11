# PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentRequestData** | Pointer to **interface{}** | The Axerve payment request data, collected by client. | [optional] 
**Update** | Pointer to **interface{}** | Send this attribute if you want to update the payment with fresh order data. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributes

`func NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributes() *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes`

NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributes instantiates a new PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributesWithDefaults

`func NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributesWithDefaults() *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes`

NewPATCHAxervePaymentsAxervePaymentIdRequestDataAttributesWithDefaults instantiates a new PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentRequestData

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetPaymentRequestData() interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetPaymentRequestDataOk() (*interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetPaymentRequestData(v interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### SetPaymentRequestDataNil

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetPaymentRequestDataNil(b bool)`

 SetPaymentRequestDataNil sets the value for PaymentRequestData to be an explicit nil

### UnsetPaymentRequestData
`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) UnsetPaymentRequestData()`

UnsetPaymentRequestData ensures that no value is present for PaymentRequestData, not even an explicit nil
### GetUpdate

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetUpdate() interface{}`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetUpdateOk() (*interface{}, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetUpdate(v interface{})`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetReference

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHAxervePaymentsAxervePaymentIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


