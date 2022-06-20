# OrderCopyDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The order copy status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;failed&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **string** | Time at which the order copy was started. | [optional] 
**CompletedAt** | Pointer to **string** | Time at which the order copy was completed. | [optional] 
**FailedAt** | Pointer to **string** | Time at which the order copy has failed. | [optional] 
**PlaceTargetOrder** | Pointer to **bool** | Indicates if the target order must be placed upon copy. | [optional] 
**CancelSourceOrder** | Pointer to **bool** | Indicates if the source order must be cancelled upon copy. | [optional] 
**ReuseWallet** | Pointer to **bool** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**ErrorsLog** | Pointer to **map[string]interface{}** | Contains the order copy errors, if any. | [optional] 
**ErrorsCount** | Pointer to **int32** | Indicates the number of copy errors, if any. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewOrderCopyDataAttributes

`func NewOrderCopyDataAttributes() *OrderCopyDataAttributes`

NewOrderCopyDataAttributes instantiates a new OrderCopyDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCopyDataAttributesWithDefaults

`func NewOrderCopyDataAttributesWithDefaults() *OrderCopyDataAttributes`

NewOrderCopyDataAttributesWithDefaults instantiates a new OrderCopyDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OrderCopyDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderCopyDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderCopyDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderCopyDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *OrderCopyDataAttributes) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *OrderCopyDataAttributes) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *OrderCopyDataAttributes) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *OrderCopyDataAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *OrderCopyDataAttributes) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *OrderCopyDataAttributes) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *OrderCopyDataAttributes) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *OrderCopyDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetFailedAt

`func (o *OrderCopyDataAttributes) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *OrderCopyDataAttributes) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *OrderCopyDataAttributes) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *OrderCopyDataAttributes) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetPlaceTargetOrder

`func (o *OrderCopyDataAttributes) GetPlaceTargetOrder() bool`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *OrderCopyDataAttributes) GetPlaceTargetOrderOk() (*bool, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *OrderCopyDataAttributes) SetPlaceTargetOrder(v bool)`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *OrderCopyDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### GetCancelSourceOrder

`func (o *OrderCopyDataAttributes) GetCancelSourceOrder() bool`

GetCancelSourceOrder returns the CancelSourceOrder field if non-nil, zero value otherwise.

### GetCancelSourceOrderOk

`func (o *OrderCopyDataAttributes) GetCancelSourceOrderOk() (*bool, bool)`

GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceOrder

`func (o *OrderCopyDataAttributes) SetCancelSourceOrder(v bool)`

SetCancelSourceOrder sets CancelSourceOrder field to given value.

### HasCancelSourceOrder

`func (o *OrderCopyDataAttributes) HasCancelSourceOrder() bool`

HasCancelSourceOrder returns a boolean if a field has been set.

### GetReuseWallet

`func (o *OrderCopyDataAttributes) GetReuseWallet() bool`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *OrderCopyDataAttributes) GetReuseWalletOk() (*bool, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *OrderCopyDataAttributes) SetReuseWallet(v bool)`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *OrderCopyDataAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### GetErrorsLog

`func (o *OrderCopyDataAttributes) GetErrorsLog() map[string]interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *OrderCopyDataAttributes) GetErrorsLogOk() (*map[string]interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *OrderCopyDataAttributes) SetErrorsLog(v map[string]interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *OrderCopyDataAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### GetErrorsCount

`func (o *OrderCopyDataAttributes) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *OrderCopyDataAttributes) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *OrderCopyDataAttributes) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *OrderCopyDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### GetId

`func (o *OrderCopyDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderCopyDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderCopyDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderCopyDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderCopyDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderCopyDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderCopyDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderCopyDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderCopyDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderCopyDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderCopyDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderCopyDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *OrderCopyDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderCopyDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderCopyDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderCopyDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderCopyDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderCopyDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderCopyDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderCopyDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderCopyDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderCopyDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderCopyDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderCopyDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


