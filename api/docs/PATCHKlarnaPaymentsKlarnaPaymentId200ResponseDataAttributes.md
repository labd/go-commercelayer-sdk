# PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **string** | The token returned by a successful client authorization, mandatory to place the order. | [optional] 
**Update** | Pointer to **bool** | Send this attribute if you want to update the payment session with fresh order data. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes

`func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults

`func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


