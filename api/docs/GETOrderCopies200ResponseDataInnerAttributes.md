# GETOrderCopies200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **interface{}** | The order factory status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;failed&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the order copy was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the order copy was completed. | [optional] 
**FailedAt** | Pointer to **interface{}** | Time at which the order copy has failed. | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the order copy errors, if any. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of copy errors, if any. | [optional] 
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the target order must be placed upon copy. | [optional] 
**ReuseWallet** | Pointer to **interface{}** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CancelSourceOrder** | Pointer to **interface{}** | Indicates if the source order must be cancelled upon copy. | [optional] 

## Methods

### NewGETOrderCopies200ResponseDataInnerAttributes

`func NewGETOrderCopies200ResponseDataInnerAttributes() *GETOrderCopies200ResponseDataInnerAttributes`

NewGETOrderCopies200ResponseDataInnerAttributes instantiates a new GETOrderCopies200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderCopies200ResponseDataInnerAttributesWithDefaults

`func NewGETOrderCopies200ResponseDataInnerAttributesWithDefaults() *GETOrderCopies200ResponseDataInnerAttributes`

NewGETOrderCopies200ResponseDataInnerAttributesWithDefaults instantiates a new GETOrderCopies200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFailedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetFailedAt() interface{}`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetFailedAtOk() (*interface{}, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetFailedAt(v interface{})`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetErrorsLog

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetErrorsCount

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetPlaceTargetOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetReuseWallet

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReuseWallet() interface{}`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReuseWalletOk() (*interface{}, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReuseWallet(v interface{})`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### SetReuseWalletNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReuseWalletNil(b bool)`

 SetReuseWalletNil sets the value for ReuseWallet to be an explicit nil

### UnsetReuseWallet
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetReuseWallet()`

UnsetReuseWallet ensures that no value is present for ReuseWallet, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCancelSourceOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCancelSourceOrder() interface{}`

GetCancelSourceOrder returns the CancelSourceOrder field if non-nil, zero value otherwise.

### GetCancelSourceOrderOk

`func (o *GETOrderCopies200ResponseDataInnerAttributes) GetCancelSourceOrderOk() (*interface{}, bool)`

GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCancelSourceOrder(v interface{})`

SetCancelSourceOrder sets CancelSourceOrder field to given value.

### HasCancelSourceOrder

`func (o *GETOrderCopies200ResponseDataInnerAttributes) HasCancelSourceOrder() bool`

HasCancelSourceOrder returns a boolean if a field has been set.

### SetCancelSourceOrderNil

`func (o *GETOrderCopies200ResponseDataInnerAttributes) SetCancelSourceOrderNil(b bool)`

 SetCancelSourceOrderNil sets the value for CancelSourceOrder to be an explicit nil

### UnsetCancelSourceOrder
`func (o *GETOrderCopies200ResponseDataInnerAttributes) UnsetCancelSourceOrder()`

UnsetCancelSourceOrder ensures that no value is present for CancelSourceOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


