# POSTDeliveryLeadTimes201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinHours** | **int32** | The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**MaxHours** | **int32** | The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTDeliveryLeadTimes201ResponseDataAttributes

`func NewPOSTDeliveryLeadTimes201ResponseDataAttributes(minHours int32, maxHours int32, ) *POSTDeliveryLeadTimes201ResponseDataAttributes`

NewPOSTDeliveryLeadTimes201ResponseDataAttributes instantiates a new POSTDeliveryLeadTimes201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTDeliveryLeadTimes201ResponseDataAttributesWithDefaults

`func NewPOSTDeliveryLeadTimes201ResponseDataAttributesWithDefaults() *POSTDeliveryLeadTimes201ResponseDataAttributes`

NewPOSTDeliveryLeadTimes201ResponseDataAttributesWithDefaults instantiates a new POSTDeliveryLeadTimes201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMinHours() int32`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMinHoursOk() (*int32, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMinHours(v int32)`

SetMinHours sets MinHours field to given value.


### GetMaxHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMaxHours() int32`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMaxHoursOk() (*int32, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMaxHours(v int32)`

SetMaxHours sets MaxHours field to given value.


### GetReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


