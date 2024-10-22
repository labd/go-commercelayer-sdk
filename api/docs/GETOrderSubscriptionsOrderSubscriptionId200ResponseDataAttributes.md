# GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | Unique identifier for the subscription (numeric). | [optional] 
**Status** | Pointer to **interface{}** | The subscription status. One of &#39;draft&#39; (default), &#39;inactive&#39;, &#39;active&#39;, or &#39;cancelled&#39;. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency of the subscription. Use one of the supported within &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model. | [optional] 
**ActivateBySourceOrder** | Pointer to **interface{}** | Indicates if the subscription will be activated considering the placed source order as its first run. | [optional] 
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the subscription created orders are automatically placed at the end of the copy. | [optional] 
**RenewalAlertPeriod** | Pointer to **interface{}** | Indicates the number of hours the renewal alert will be triggered before the subscription next run. Must be included between 1 and 720 hours. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the customer, if any, associated to the source order. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this subscription. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**LastRunAt** | Pointer to **interface{}** | The date/time of the subscription last run. | [optional] 
**NextRunAt** | Pointer to **interface{}** | The date/time of the subscription next run. Can be updated but only in the future, to copy with frequency changes. | [optional] 
**Occurrencies** | Pointer to **interface{}** | The number of times this subscription has run. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of subscription errors, if any. | [optional] 
**SucceededOnLastRun** | Pointer to **interface{}** | Indicates if the subscription has succeeded on its last run. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes

`func NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes() *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes`

NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes instantiates a new GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults

`func NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults() *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes`

NewGETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributesWithDefaults instantiates a new GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFrequency

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetActivateBySourceOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivateBySourceOrder() interface{}`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetActivateBySourceOrderOk() (*interface{}, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetActivateBySourceOrder(v interface{})`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### SetActivateBySourceOrderNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetActivateBySourceOrderNil(b bool)`

 SetActivateBySourceOrderNil sets the value for ActivateBySourceOrder to be an explicit nil

### UnsetActivateBySourceOrder
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetActivateBySourceOrder()`

UnsetActivateBySourceOrder ensures that no value is present for ActivateBySourceOrder, not even an explicit nil
### GetPlaceTargetOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetRenewalAlertPeriod

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetRenewalAlertPeriod() interface{}`

GetRenewalAlertPeriod returns the RenewalAlertPeriod field if non-nil, zero value otherwise.

### GetRenewalAlertPeriodOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetRenewalAlertPeriodOk() (*interface{}, bool)`

GetRenewalAlertPeriodOk returns a tuple with the RenewalAlertPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAlertPeriod

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetRenewalAlertPeriod(v interface{})`

SetRenewalAlertPeriod sets RenewalAlertPeriod field to given value.

### HasRenewalAlertPeriod

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasRenewalAlertPeriod() bool`

HasRenewalAlertPeriod returns a boolean if a field has been set.

### SetRenewalAlertPeriodNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetRenewalAlertPeriodNil(b bool)`

 SetRenewalAlertPeriodNil sets the value for RenewalAlertPeriod to be an explicit nil

### UnsetRenewalAlertPeriod
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetRenewalAlertPeriod()`

UnsetRenewalAlertPeriod ensures that no value is present for RenewalAlertPeriod, not even an explicit nil
### GetCustomerEmail

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetStartsAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetLastRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetLastRunAt() interface{}`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetLastRunAtOk() (*interface{}, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetLastRunAt(v interface{})`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetNextRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNextRunAt() interface{}`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetNextRunAtOk() (*interface{}, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNextRunAt(v interface{})`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetOccurrencies

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOccurrencies() interface{}`

GetOccurrencies returns the Occurrencies field if non-nil, zero value otherwise.

### GetOccurrenciesOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetOccurrenciesOk() (*interface{}, bool)`

GetOccurrenciesOk returns a tuple with the Occurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrencies

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetOccurrencies(v interface{})`

SetOccurrencies sets Occurrencies field to given value.

### HasOccurrencies

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasOccurrencies() bool`

HasOccurrencies returns a boolean if a field has been set.

### SetOccurrenciesNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetOccurrenciesNil(b bool)`

 SetOccurrenciesNil sets the value for Occurrencies to be an explicit nil

### UnsetOccurrencies
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetOccurrencies()`

UnsetOccurrencies ensures that no value is present for Occurrencies, not even an explicit nil
### GetErrorsCount

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetSucceededOnLastRun

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetSucceededOnLastRun() interface{}`

GetSucceededOnLastRun returns the SucceededOnLastRun field if non-nil, zero value otherwise.

### GetSucceededOnLastRunOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetSucceededOnLastRunOk() (*interface{}, bool)`

GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededOnLastRun

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetSucceededOnLastRun(v interface{})`

SetSucceededOnLastRun sets SucceededOnLastRun field to given value.

### HasSucceededOnLastRun

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasSucceededOnLastRun() bool`

HasSucceededOnLastRun returns a boolean if a field has been set.

### SetSucceededOnLastRunNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetSucceededOnLastRunNil(b bool)`

 SetSucceededOnLastRunNil sets the value for SucceededOnLastRun to be an explicit nil

### UnsetSucceededOnLastRun
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetSucceededOnLastRun()`

UnsetSucceededOnLastRun ensures that no value is present for SucceededOnLastRun, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderSubscriptionsOrderSubscriptionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


