# GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The price list scheduler&#39;s internal name. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this price list scheduler. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this price list scheduler (must be after starts_at). | [optional] 
**Active** | Pointer to **interface{}** | Indicates if the price list scheduler is active (enabled and not expired). | [optional] 
**Status** | Pointer to **interface{}** | The price list scheduler status. One of &#39;disabled&#39;, &#39;expired&#39;, &#39;pending&#39;, or &#39;active&#39;. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which this resource was disabled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes

`func NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes() *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes`

NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes instantiates a new GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributesWithDefaults

`func NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributesWithDefaults() *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes`

NewGETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributesWithDefaults instantiates a new GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartsAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetActive

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetStatus

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDisabledAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


