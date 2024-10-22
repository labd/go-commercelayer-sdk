# PATCHStripePaymentsStripePaymentId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **interface{}** | Stripe payment options: &#39;customer&#39;, &#39;payment_method&#39;, &#39;return_url&#39;, etc. Check Stripe payment intent API for more details. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The URL to return to when a redirect payment is completed. | [optional] 
**Update** | Pointer to **interface{}** | Send this attribute if you want to update the created payment intent with fresh order data. | [optional] 
**Refresh** | Pointer to **interface{}** | Send this attribute if you want to refresh the payment status, can be used as webhooks fallback logic. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes

`func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes`

NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributes instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults

`func NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults() *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes`

NewPATCHStripePaymentsStripePaymentId200ResponseDataAttributesWithDefaults instantiates a new PATCHStripePaymentsStripePaymentId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetReturnUrl

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetUpdate

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdate() interface{}`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetUpdateOk() (*interface{}, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdate(v interface{})`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### SetUpdateNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetUpdateNil(b bool)`

 SetUpdateNil sets the value for Update to be an explicit nil

### UnsetUpdate
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetUpdate()`

UnsetUpdate ensures that no value is present for Update, not even an explicit nil
### GetRefresh

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefresh() interface{}`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetRefreshOk() (*interface{}, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetRefresh(v interface{})`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### SetRefreshNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetRefreshNil(b bool)`

 SetRefreshNil sets the value for Refresh to be an explicit nil

### UnsetRefresh
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetRefresh()`

UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
### GetReference

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHStripePaymentsStripePaymentId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


