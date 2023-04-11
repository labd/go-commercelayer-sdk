# GETOrderSubscriptions200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the subscription (numeric) | [optional] 
**Status** | Pointer to **interface{}** | The subscription status. One of &#39;draft&#39; (default), &#39;inactive&#39;, &#39;active&#39;, or &#39;cancelled&#39;. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency of the subscription. Use one of the supported within &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model. | [optional] 
**ActivateBySourceOrder** | Pointer to **interface{}** | Indicates if the subscription will be activated considering the placed source order as its first run. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the customer, if any, associated to the source order. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this subscription. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**NextRunAt** | Pointer to **interface{}** | The date/time of the subscription next run. Can be updated but only in the future, to copy with frequency changes. | [optional] 
**Occurrencies** | Pointer to **interface{}** | The number of times this subscription has run. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of subscription errors, if any. | [optional] 
**SucceededOnLastRun** | Pointer to **interface{}** | Indicates if the subscription has succeeded on its last run. | [optional] 
**Options** | Pointer to **interface{}** | The subscription options used to create the order (check order_factories for more information). For subscriptions the &#x60;place_target_order&#x60; is enabled by default, specify custom options to overwrite it. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrderSubscriptions200ResponseDataInnerAttributes

`func NewGETOrderSubscriptions200ResponseDataInnerAttributes() *GETOrderSubscriptions200ResponseDataInnerAttributes`

NewGETOrderSubscriptions200ResponseDataInnerAttributes instantiates a new GETOrderSubscriptions200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderSubscriptions200ResponseDataInnerAttributesWithDefaults

`func NewGETOrderSubscriptions200ResponseDataInnerAttributesWithDefaults() *GETOrderSubscriptions200ResponseDataInnerAttributes`

NewGETOrderSubscriptions200ResponseDataInnerAttributesWithDefaults instantiates a new GETOrderSubscriptions200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetActivateBySourceOrder() interface{}`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetActivateBySourceOrderOk() (*interface{}, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetActivateBySourceOrder(v interface{})`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### SetActivateBySourceOrderNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetActivateBySourceOrderNil(b bool)`

 SetActivateBySourceOrderNil sets the value for ActivateBySourceOrder to be an explicit nil

### UnsetActivateBySourceOrder
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetActivateBySourceOrder()`

UnsetActivateBySourceOrder ensures that no value is present for ActivateBySourceOrder, not even an explicit nil
### GetCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNextRunAt() interface{}`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNextRunAtOk() (*interface{}, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNextRunAt(v interface{})`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOccurrencies() interface{}`

GetOccurrencies returns the Occurrencies field if non-nil, zero value otherwise.

### GetOccurrenciesOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOccurrenciesOk() (*interface{}, bool)`

GetOccurrenciesOk returns a tuple with the Occurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOccurrencies(v interface{})`

SetOccurrencies sets Occurrencies field to given value.

### HasOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasOccurrencies() bool`

HasOccurrencies returns a boolean if a field has been set.

### SetOccurrenciesNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOccurrenciesNil(b bool)`

 SetOccurrenciesNil sets the value for Occurrencies to be an explicit nil

### UnsetOccurrencies
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetOccurrencies()`

UnsetOccurrencies ensures that no value is present for Occurrencies, not even an explicit nil
### GetErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetSucceededOnLastRun() interface{}`

GetSucceededOnLastRun returns the SucceededOnLastRun field if non-nil, zero value otherwise.

### GetSucceededOnLastRunOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetSucceededOnLastRunOk() (*interface{}, bool)`

GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetSucceededOnLastRun(v interface{})`

SetSucceededOnLastRun sets SucceededOnLastRun field to given value.

### HasSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasSucceededOnLastRun() bool`

HasSucceededOnLastRun returns a boolean if a field has been set.

### SetSucceededOnLastRunNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetSucceededOnLastRunNil(b bool)`

 SetSucceededOnLastRunNil sets the value for SucceededOnLastRun to be an explicit nil

### UnsetSucceededOnLastRun
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetSucceededOnLastRun()`

UnsetSucceededOnLastRun ensures that no value is present for SucceededOnLastRun, not even an explicit nil
### GetOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


