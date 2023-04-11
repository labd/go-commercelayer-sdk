# PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **interface{}** | The frequency of the subscription. Use one of the supported within &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, &#39;yearly&#39;, or provide your custom crontab expression (min unit is hour). Must be supported by existing associated subscription_model. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**NextRunAt** | Pointer to **interface{}** | The date/time of the subscription next run. Can be updated but only in the future, to copy with frequency changes. | [optional] 
**Activate** | Pointer to **interface{}** | Send this attribute if you want to mark this subscription as active. | [optional] 
**Deactivate** | Pointer to **interface{}** | Send this attribute if you want to mark this subscription as inactive. | [optional] 
**Cancel** | Pointer to **interface{}** | Send this attribute if you want to mark this subscription as cancelled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes

`func NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes() *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes`

NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes instantiates a new PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributesWithDefaults

`func NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributesWithDefaults() *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes`

NewPATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributesWithDefaults instantiates a new PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetExpiresAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetNextRunAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetNextRunAt() interface{}`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetNextRunAtOk() (*interface{}, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetNextRunAt(v interface{})`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetActivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetActivate() interface{}`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetActivateOk() (*interface{}, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetActivate(v interface{})`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### SetActivateNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetActivateNil(b bool)`

 SetActivateNil sets the value for Activate to be an explicit nil

### UnsetActivate
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetActivate()`

UnsetActivate ensures that no value is present for Activate, not even an explicit nil
### GetDeactivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetDeactivate() interface{}`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetDeactivateOk() (*interface{}, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetDeactivate(v interface{})`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### SetDeactivateNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetDeactivateNil(b bool)`

 SetDeactivateNil sets the value for Deactivate to be an explicit nil

### UnsetDeactivate
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetDeactivate()`

UnsetDeactivate ensures that no value is present for Deactivate, not even an explicit nil
### GetCancel

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetCancel() interface{}`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetCancelOk() (*interface{}, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetCancel(v interface{})`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### SetCancelNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetCancelNil(b bool)`

 SetCancelNil sets the value for Cancel to be an explicit nil

### UnsetCancel
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetCancel()`

UnsetCancel ensures that no value is present for Cancel, not even an explicit nil
### GetReference

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHOrderSubscriptionsOrderSubscriptionIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


