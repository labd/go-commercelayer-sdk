# POSTDeliveryLeadTimes201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinHours** | **interface{}** | The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**MaxHours** | **interface{}** | The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTDeliveryLeadTimes201ResponseDataAttributes

`func NewPOSTDeliveryLeadTimes201ResponseDataAttributes(minHours interface{}, maxHours interface{}, ) *POSTDeliveryLeadTimes201ResponseDataAttributes`

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

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMinHours() interface{}`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMinHoursOk() (*interface{}, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMinHours(v interface{})`

SetMinHours sets MinHours field to given value.


### SetMinHoursNil

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMinHoursNil(b bool)`

 SetMinHoursNil sets the value for MinHours to be an explicit nil

### UnsetMinHours
`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) UnsetMinHours()`

UnsetMinHours ensures that no value is present for MinHours, not even an explicit nil
### GetMaxHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMaxHours() interface{}`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMaxHoursOk() (*interface{}, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMaxHours(v interface{})`

SetMaxHours sets MaxHours field to given value.


### SetMaxHoursNil

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMaxHoursNil(b bool)`

 SetMaxHoursNil sets the value for MaxHours to be an explicit nil

### UnsetMaxHours
`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) UnsetMaxHours()`

UnsetMaxHours ensures that no value is present for MaxHours, not even an explicit nil
### GetReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTDeliveryLeadTimes201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


