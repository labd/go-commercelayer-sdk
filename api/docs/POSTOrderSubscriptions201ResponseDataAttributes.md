# POSTOrderSubscriptions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **interface{}** | The frequency of the subscription. Use one of the supported within &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model. | 
**ActivateBySourceOrder** | Pointer to **interface{}** | Indicates if the subscription will be activated considering the placed source order as its first run. | [optional] 
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the subscription created orders are automatically placed at the end of the copy. | [optional] 
**RenewalAlertPeriod** | Pointer to **interface{}** | Indicates the number of hours the renewal alert will be triggered before the subscription next run. Must be included between 1 and 720 hours. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this subscription. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTOrderSubscriptions201ResponseDataAttributes

`func NewPOSTOrderSubscriptions201ResponseDataAttributes(frequency interface{}, ) *POSTOrderSubscriptions201ResponseDataAttributes`

NewPOSTOrderSubscriptions201ResponseDataAttributes instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults

`func NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults() *POSTOrderSubscriptions201ResponseDataAttributes`

NewPOSTOrderSubscriptions201ResponseDataAttributesWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.


### SetFrequencyNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrder() interface{}`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrderOk() (*interface{}, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetActivateBySourceOrder(v interface{})`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### SetActivateBySourceOrderNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetActivateBySourceOrderNil(b bool)`

 SetActivateBySourceOrderNil sets the value for ActivateBySourceOrder to be an explicit nil

### UnsetActivateBySourceOrder
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetActivateBySourceOrder()`

UnsetActivateBySourceOrder ensures that no value is present for ActivateBySourceOrder, not even an explicit nil
### GetPlaceTargetOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetRenewalAlertPeriod

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetRenewalAlertPeriod() interface{}`

GetRenewalAlertPeriod returns the RenewalAlertPeriod field if non-nil, zero value otherwise.

### GetRenewalAlertPeriodOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetRenewalAlertPeriodOk() (*interface{}, bool)`

GetRenewalAlertPeriodOk returns a tuple with the RenewalAlertPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalAlertPeriod

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetRenewalAlertPeriod(v interface{})`

SetRenewalAlertPeriod sets RenewalAlertPeriod field to given value.

### HasRenewalAlertPeriod

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasRenewalAlertPeriod() bool`

HasRenewalAlertPeriod returns a boolean if a field has been set.

### SetRenewalAlertPeriodNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetRenewalAlertPeriodNil(b bool)`

 SetRenewalAlertPeriodNil sets the value for RenewalAlertPeriod to be an explicit nil

### UnsetRenewalAlertPeriod
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetRenewalAlertPeriod()`

UnsetRenewalAlertPeriod ensures that no value is present for RenewalAlertPeriod, not even an explicit nil
### GetStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTOrderSubscriptions201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


