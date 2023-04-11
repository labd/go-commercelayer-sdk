# POSTDeliveryLeadTimesRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinHours** | **interface{}** | The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**MaxHours** | **interface{}** | The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method. | 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTDeliveryLeadTimesRequestDataAttributes

`func NewPOSTDeliveryLeadTimesRequestDataAttributes(minHours interface{}, maxHours interface{}, ) *POSTDeliveryLeadTimesRequestDataAttributes`

NewPOSTDeliveryLeadTimesRequestDataAttributes instantiates a new POSTDeliveryLeadTimesRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTDeliveryLeadTimesRequestDataAttributesWithDefaults

`func NewPOSTDeliveryLeadTimesRequestDataAttributesWithDefaults() *POSTDeliveryLeadTimesRequestDataAttributes`

NewPOSTDeliveryLeadTimesRequestDataAttributesWithDefaults instantiates a new POSTDeliveryLeadTimesRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinHours

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMinHours() interface{}`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMinHoursOk() (*interface{}, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMinHours(v interface{})`

SetMinHours sets MinHours field to given value.


### SetMinHoursNil

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMinHoursNil(b bool)`

 SetMinHoursNil sets the value for MinHours to be an explicit nil

### UnsetMinHours
`func (o *POSTDeliveryLeadTimesRequestDataAttributes) UnsetMinHours()`

UnsetMinHours ensures that no value is present for MinHours, not even an explicit nil
### GetMaxHours

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMaxHours() interface{}`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMaxHoursOk() (*interface{}, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMaxHours(v interface{})`

SetMaxHours sets MaxHours field to given value.


### SetMaxHoursNil

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMaxHoursNil(b bool)`

 SetMaxHoursNil sets the value for MaxHours to be an explicit nil

### UnsetMaxHours
`func (o *POSTDeliveryLeadTimesRequestDataAttributes) UnsetMaxHours()`

UnsetMaxHours ensures that no value is present for MaxHours, not even an explicit nil
### GetReference

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTDeliveryLeadTimesRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTDeliveryLeadTimesRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTDeliveryLeadTimesRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTDeliveryLeadTimesRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


