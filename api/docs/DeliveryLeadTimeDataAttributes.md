# DeliveryLeadTimeDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinHours** | Pointer to **int32** | The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method. | [optional] 
**MaxHours** | Pointer to **int32** | The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method. | [optional] 
**MinDays** | Pointer to **int32** | The delivery lead minimum time, in days (rounded) | [optional] 
**MaxDays** | Pointer to **int32** | The delivery lead maximun time, in days (rounded) | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewDeliveryLeadTimeDataAttributes

`func NewDeliveryLeadTimeDataAttributes() *DeliveryLeadTimeDataAttributes`

NewDeliveryLeadTimeDataAttributes instantiates a new DeliveryLeadTimeDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLeadTimeDataAttributesWithDefaults

`func NewDeliveryLeadTimeDataAttributesWithDefaults() *DeliveryLeadTimeDataAttributes`

NewDeliveryLeadTimeDataAttributesWithDefaults instantiates a new DeliveryLeadTimeDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinHours

`func (o *DeliveryLeadTimeDataAttributes) GetMinHours() int32`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *DeliveryLeadTimeDataAttributes) GetMinHoursOk() (*int32, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *DeliveryLeadTimeDataAttributes) SetMinHours(v int32)`

SetMinHours sets MinHours field to given value.

### HasMinHours

`func (o *DeliveryLeadTimeDataAttributes) HasMinHours() bool`

HasMinHours returns a boolean if a field has been set.

### GetMaxHours

`func (o *DeliveryLeadTimeDataAttributes) GetMaxHours() int32`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *DeliveryLeadTimeDataAttributes) GetMaxHoursOk() (*int32, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *DeliveryLeadTimeDataAttributes) SetMaxHours(v int32)`

SetMaxHours sets MaxHours field to given value.

### HasMaxHours

`func (o *DeliveryLeadTimeDataAttributes) HasMaxHours() bool`

HasMaxHours returns a boolean if a field has been set.

### GetMinDays

`func (o *DeliveryLeadTimeDataAttributes) GetMinDays() int32`

GetMinDays returns the MinDays field if non-nil, zero value otherwise.

### GetMinDaysOk

`func (o *DeliveryLeadTimeDataAttributes) GetMinDaysOk() (*int32, bool)`

GetMinDaysOk returns a tuple with the MinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDays

`func (o *DeliveryLeadTimeDataAttributes) SetMinDays(v int32)`

SetMinDays sets MinDays field to given value.

### HasMinDays

`func (o *DeliveryLeadTimeDataAttributes) HasMinDays() bool`

HasMinDays returns a boolean if a field has been set.

### GetMaxDays

`func (o *DeliveryLeadTimeDataAttributes) GetMaxDays() int32`

GetMaxDays returns the MaxDays field if non-nil, zero value otherwise.

### GetMaxDaysOk

`func (o *DeliveryLeadTimeDataAttributes) GetMaxDaysOk() (*int32, bool)`

GetMaxDaysOk returns a tuple with the MaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDays

`func (o *DeliveryLeadTimeDataAttributes) SetMaxDays(v int32)`

SetMaxDays sets MaxDays field to given value.

### HasMaxDays

`func (o *DeliveryLeadTimeDataAttributes) HasMaxDays() bool`

HasMaxDays returns a boolean if a field has been set.

### GetId

`func (o *DeliveryLeadTimeDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryLeadTimeDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryLeadTimeDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryLeadTimeDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeliveryLeadTimeDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeliveryLeadTimeDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeliveryLeadTimeDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeliveryLeadTimeDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeliveryLeadTimeDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeliveryLeadTimeDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeliveryLeadTimeDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeliveryLeadTimeDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *DeliveryLeadTimeDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DeliveryLeadTimeDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DeliveryLeadTimeDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DeliveryLeadTimeDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *DeliveryLeadTimeDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *DeliveryLeadTimeDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *DeliveryLeadTimeDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *DeliveryLeadTimeDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *DeliveryLeadTimeDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeliveryLeadTimeDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeliveryLeadTimeDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeliveryLeadTimeDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


