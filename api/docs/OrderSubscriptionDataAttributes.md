# OrderSubscriptionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Unique identifier for the subscription (numeric) | [optional] 
**Status** | Pointer to **string** | The subscription status. One of &#39;draft&#39; (default), &#39;inactive&#39;, &#39;active&#39;, or &#39;cancelled&#39;. | [optional] 
**Frequency** | Pointer to **string** | The frequency of the subscription. One of &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, or &#39;yearly&#39;. | [optional] 
**ActivateBySourceOrder** | Pointer to **bool** | Indicates if the subscription will be activated considering the placed source order as its first run, default to true. | [optional] 
**CustomerEmail** | Pointer to **string** | The email address of the customer, if any, associated to the source order. | [optional] 
**StartsAt** | Pointer to **string** | The activation date/time of this subscription. | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**NextRunAt** | Pointer to **string** | The date/time of the subscription next run. | [optional] 
**Occurrencies** | Pointer to **int32** | The number of times this subscription has run. | [optional] 
**ErrorsCount** | Pointer to **int32** | Indicates the number of subscription errors, if any. | [optional] 
**SucceededOnLastRun** | Pointer to **bool** | Indicates if the subscription has succeeded on its last run. | [optional] 
**Options** | Pointer to **map[string]interface{}** | The subscription options used to create the order copy (check order_copies for more information). For subscriptions the &#x60;place_target_order&#x60; is enabled by default, specify custom options to overwrite it. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewOrderSubscriptionDataAttributes

`func NewOrderSubscriptionDataAttributes() *OrderSubscriptionDataAttributes`

NewOrderSubscriptionDataAttributes instantiates a new OrderSubscriptionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionDataAttributesWithDefaults

`func NewOrderSubscriptionDataAttributesWithDefaults() *OrderSubscriptionDataAttributes`

NewOrderSubscriptionDataAttributesWithDefaults instantiates a new OrderSubscriptionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *OrderSubscriptionDataAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderSubscriptionDataAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderSubscriptionDataAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OrderSubscriptionDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *OrderSubscriptionDataAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderSubscriptionDataAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderSubscriptionDataAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderSubscriptionDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFrequency

`func (o *OrderSubscriptionDataAttributes) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *OrderSubscriptionDataAttributes) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *OrderSubscriptionDataAttributes) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *OrderSubscriptionDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetActivateBySourceOrder

`func (o *OrderSubscriptionDataAttributes) GetActivateBySourceOrder() bool`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *OrderSubscriptionDataAttributes) GetActivateBySourceOrderOk() (*bool, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *OrderSubscriptionDataAttributes) SetActivateBySourceOrder(v bool)`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *OrderSubscriptionDataAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *OrderSubscriptionDataAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *OrderSubscriptionDataAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *OrderSubscriptionDataAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *OrderSubscriptionDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetStartsAt

`func (o *OrderSubscriptionDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *OrderSubscriptionDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *OrderSubscriptionDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *OrderSubscriptionDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrderSubscriptionDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrderSubscriptionDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrderSubscriptionDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrderSubscriptionDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNextRunAt

`func (o *OrderSubscriptionDataAttributes) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *OrderSubscriptionDataAttributes) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *OrderSubscriptionDataAttributes) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *OrderSubscriptionDataAttributes) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetOccurrencies

`func (o *OrderSubscriptionDataAttributes) GetOccurrencies() int32`

GetOccurrencies returns the Occurrencies field if non-nil, zero value otherwise.

### GetOccurrenciesOk

`func (o *OrderSubscriptionDataAttributes) GetOccurrenciesOk() (*int32, bool)`

GetOccurrenciesOk returns a tuple with the Occurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrencies

`func (o *OrderSubscriptionDataAttributes) SetOccurrencies(v int32)`

SetOccurrencies sets Occurrencies field to given value.

### HasOccurrencies

`func (o *OrderSubscriptionDataAttributes) HasOccurrencies() bool`

HasOccurrencies returns a boolean if a field has been set.

### GetErrorsCount

`func (o *OrderSubscriptionDataAttributes) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *OrderSubscriptionDataAttributes) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *OrderSubscriptionDataAttributes) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *OrderSubscriptionDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### GetSucceededOnLastRun

`func (o *OrderSubscriptionDataAttributes) GetSucceededOnLastRun() bool`

GetSucceededOnLastRun returns the SucceededOnLastRun field if non-nil, zero value otherwise.

### GetSucceededOnLastRunOk

`func (o *OrderSubscriptionDataAttributes) GetSucceededOnLastRunOk() (*bool, bool)`

GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededOnLastRun

`func (o *OrderSubscriptionDataAttributes) SetSucceededOnLastRun(v bool)`

SetSucceededOnLastRun sets SucceededOnLastRun field to given value.

### HasSucceededOnLastRun

`func (o *OrderSubscriptionDataAttributes) HasSucceededOnLastRun() bool`

HasSucceededOnLastRun returns a boolean if a field has been set.

### GetOptions

`func (o *OrderSubscriptionDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrderSubscriptionDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrderSubscriptionDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OrderSubscriptionDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetId

`func (o *OrderSubscriptionDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderSubscriptionDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderSubscriptionDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderSubscriptionDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderSubscriptionDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderSubscriptionDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderSubscriptionDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderSubscriptionDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderSubscriptionDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderSubscriptionDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderSubscriptionDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderSubscriptionDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *OrderSubscriptionDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderSubscriptionDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderSubscriptionDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderSubscriptionDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderSubscriptionDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderSubscriptionDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderSubscriptionDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderSubscriptionDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderSubscriptionDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderSubscriptionDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderSubscriptionDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderSubscriptionDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


