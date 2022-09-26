# GETDeliveryLeadTimes200ResponseDataInnerAttributes

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

### NewGETDeliveryLeadTimes200ResponseDataInnerAttributes

`func NewGETDeliveryLeadTimes200ResponseDataInnerAttributes() *GETDeliveryLeadTimes200ResponseDataInnerAttributes`

NewGETDeliveryLeadTimes200ResponseDataInnerAttributes instantiates a new GETDeliveryLeadTimes200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETDeliveryLeadTimes200ResponseDataInnerAttributesWithDefaults

`func NewGETDeliveryLeadTimes200ResponseDataInnerAttributesWithDefaults() *GETDeliveryLeadTimes200ResponseDataInnerAttributes`

NewGETDeliveryLeadTimes200ResponseDataInnerAttributesWithDefaults instantiates a new GETDeliveryLeadTimes200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinHours() int32`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinHoursOk() (*int32, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMinHours(v int32)`

SetMinHours sets MinHours field to given value.

### HasMinHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMinHours() bool`

HasMinHours returns a boolean if a field has been set.

### GetMaxHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxHours() int32`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxHoursOk() (*int32, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMaxHours(v int32)`

SetMaxHours sets MaxHours field to given value.

### HasMaxHours

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMaxHours() bool`

HasMaxHours returns a boolean if a field has been set.

### GetMinDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinDays() int32`

GetMinDays returns the MinDays field if non-nil, zero value otherwise.

### GetMinDaysOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMinDaysOk() (*int32, bool)`

GetMinDaysOk returns a tuple with the MinDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMinDays(v int32)`

SetMinDays sets MinDays field to given value.

### HasMinDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMinDays() bool`

HasMinDays returns a boolean if a field has been set.

### GetMaxDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxDays() int32`

GetMaxDays returns the MaxDays field if non-nil, zero value otherwise.

### GetMaxDaysOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMaxDaysOk() (*int32, bool)`

GetMaxDaysOk returns a tuple with the MaxDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMaxDays(v int32)`

SetMaxDays sets MaxDays field to given value.

### HasMaxDays

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMaxDays() bool`

HasMaxDays returns a boolean if a field has been set.

### GetId

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETDeliveryLeadTimes200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


