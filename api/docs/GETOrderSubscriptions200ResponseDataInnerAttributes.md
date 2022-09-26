# GETOrderSubscriptions200ResponseDataInnerAttributes

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

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetActivateBySourceOrder() bool`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetActivateBySourceOrderOk() (*bool, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetActivateBySourceOrder(v bool)`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOccurrencies() int32`

GetOccurrencies returns the Occurrencies field if non-nil, zero value otherwise.

### GetOccurrenciesOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOccurrenciesOk() (*int32, bool)`

GetOccurrenciesOk returns a tuple with the Occurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOccurrencies(v int32)`

SetOccurrencies sets Occurrencies field to given value.

### HasOccurrencies

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasOccurrencies() bool`

HasOccurrencies returns a boolean if a field has been set.

### GetErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetErrorsCount() int32`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetErrorsCountOk() (*int32, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetErrorsCount(v int32)`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### GetSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetSucceededOnLastRun() bool`

GetSucceededOnLastRun returns the SucceededOnLastRun field if non-nil, zero value otherwise.

### GetSucceededOnLastRunOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetSucceededOnLastRunOk() (*bool, bool)`

GetSucceededOnLastRunOk returns a tuple with the SucceededOnLastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetSucceededOnLastRun(v bool)`

SetSucceededOnLastRun sets SucceededOnLastRun field to given value.

### HasSucceededOnLastRun

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasSucceededOnLastRun() bool`

HasSucceededOnLastRun returns a boolean if a field has been set.

### GetOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetId

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderSubscriptions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


