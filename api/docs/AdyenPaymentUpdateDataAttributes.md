# AdyenPaymentUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentRequestData** | Pointer to **map[string]interface{}** | The Adyen payment request data, collected by client. | [optional] 
**PaymentRequestDetails** | Pointer to **map[string]interface{}** | The Adyen additional details request data, collected by client. | [optional] 
**PaymentResponse** | Pointer to **map[string]interface{}** | The Adyen payment response, used by client (includes &#39;resultCode&#39; and &#39;action&#39;). | [optional] 
**Details** | Pointer to **bool** | Send this attribute if you want to send additional details the payment request. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewAdyenPaymentUpdateDataAttributes

`func NewAdyenPaymentUpdateDataAttributes() *AdyenPaymentUpdateDataAttributes`

NewAdyenPaymentUpdateDataAttributes instantiates a new AdyenPaymentUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenPaymentUpdateDataAttributesWithDefaults

`func NewAdyenPaymentUpdateDataAttributesWithDefaults() *AdyenPaymentUpdateDataAttributes`

NewAdyenPaymentUpdateDataAttributesWithDefaults instantiates a new AdyenPaymentUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentRequestData

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestData() map[string]interface{}`

GetPaymentRequestData returns the PaymentRequestData field if non-nil, zero value otherwise.

### GetPaymentRequestDataOk

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDataOk() (*map[string]interface{}, bool)`

GetPaymentRequestDataOk returns a tuple with the PaymentRequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestData

`func (o *AdyenPaymentUpdateDataAttributes) SetPaymentRequestData(v map[string]interface{})`

SetPaymentRequestData sets PaymentRequestData field to given value.

### HasPaymentRequestData

`func (o *AdyenPaymentUpdateDataAttributes) HasPaymentRequestData() bool`

HasPaymentRequestData returns a boolean if a field has been set.

### GetPaymentRequestDetails

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDetails() map[string]interface{}`

GetPaymentRequestDetails returns the PaymentRequestDetails field if non-nil, zero value otherwise.

### GetPaymentRequestDetailsOk

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentRequestDetailsOk() (*map[string]interface{}, bool)`

GetPaymentRequestDetailsOk returns a tuple with the PaymentRequestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestDetails

`func (o *AdyenPaymentUpdateDataAttributes) SetPaymentRequestDetails(v map[string]interface{})`

SetPaymentRequestDetails sets PaymentRequestDetails field to given value.

### HasPaymentRequestDetails

`func (o *AdyenPaymentUpdateDataAttributes) HasPaymentRequestDetails() bool`

HasPaymentRequestDetails returns a boolean if a field has been set.

### GetPaymentResponse

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentResponse() map[string]interface{}`

GetPaymentResponse returns the PaymentResponse field if non-nil, zero value otherwise.

### GetPaymentResponseOk

`func (o *AdyenPaymentUpdateDataAttributes) GetPaymentResponseOk() (*map[string]interface{}, bool)`

GetPaymentResponseOk returns a tuple with the PaymentResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentResponse

`func (o *AdyenPaymentUpdateDataAttributes) SetPaymentResponse(v map[string]interface{})`

SetPaymentResponse sets PaymentResponse field to given value.

### HasPaymentResponse

`func (o *AdyenPaymentUpdateDataAttributes) HasPaymentResponse() bool`

HasPaymentResponse returns a boolean if a field has been set.

### GetDetails

`func (o *AdyenPaymentUpdateDataAttributes) GetDetails() bool`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdyenPaymentUpdateDataAttributes) GetDetailsOk() (*bool, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdyenPaymentUpdateDataAttributes) SetDetails(v bool)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdyenPaymentUpdateDataAttributes) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetReference

`func (o *AdyenPaymentUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdyenPaymentUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdyenPaymentUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *AdyenPaymentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *AdyenPaymentUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *AdyenPaymentUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *AdyenPaymentUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AdyenPaymentUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AdyenPaymentUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AdyenPaymentUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


