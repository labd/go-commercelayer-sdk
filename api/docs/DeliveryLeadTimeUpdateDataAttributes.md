# DeliveryLeadTimeUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinHours** | Pointer to **int32** | The delivery lead minimum time (in hours) when shipping from the associated stock location with the associated shipping method. | [optional] 
**MaxHours** | Pointer to **int32** | The delivery lead maximun time (in hours) when shipping from the associated stock location with the associated shipping method. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewDeliveryLeadTimeUpdateDataAttributes

`func NewDeliveryLeadTimeUpdateDataAttributes() *DeliveryLeadTimeUpdateDataAttributes`

NewDeliveryLeadTimeUpdateDataAttributes instantiates a new DeliveryLeadTimeUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLeadTimeUpdateDataAttributesWithDefaults

`func NewDeliveryLeadTimeUpdateDataAttributesWithDefaults() *DeliveryLeadTimeUpdateDataAttributes`

NewDeliveryLeadTimeUpdateDataAttributesWithDefaults instantiates a new DeliveryLeadTimeUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMinHours() int32`

GetMinHours returns the MinHours field if non-nil, zero value otherwise.

### GetMinHoursOk

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMinHoursOk() (*int32, bool)`

GetMinHoursOk returns a tuple with the MinHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) SetMinHours(v int32)`

SetMinHours sets MinHours field to given value.

### HasMinHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) HasMinHours() bool`

HasMinHours returns a boolean if a field has been set.

### GetMaxHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMaxHours() int32`

GetMaxHours returns the MaxHours field if non-nil, zero value otherwise.

### GetMaxHoursOk

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMaxHoursOk() (*int32, bool)`

GetMaxHoursOk returns a tuple with the MaxHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) SetMaxHours(v int32)`

SetMaxHours sets MaxHours field to given value.

### HasMaxHours

`func (o *DeliveryLeadTimeUpdateDataAttributes) HasMaxHours() bool`

HasMaxHours returns a boolean if a field has been set.

### GetReference

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DeliveryLeadTimeUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *DeliveryLeadTimeUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *DeliveryLeadTimeUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *DeliveryLeadTimeUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeliveryLeadTimeUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeliveryLeadTimeUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeliveryLeadTimeUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


