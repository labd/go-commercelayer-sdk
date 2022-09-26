# OrderSubscriptionUpdateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**Activate** | Pointer to **bool** | Send this attribute if you want to mark this subscription as active. | [optional] 
**Deactivate** | Pointer to **bool** | Send this attribute if you want to mark this subscription as inactive. | [optional] 
**Cancel** | Pointer to **bool** | Send this attribute if you want to mark this subscription as cancelled. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewOrderSubscriptionUpdateDataAttributes

`func NewOrderSubscriptionUpdateDataAttributes() *OrderSubscriptionUpdateDataAttributes`

NewOrderSubscriptionUpdateDataAttributes instantiates a new OrderSubscriptionUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionUpdateDataAttributesWithDefaults

`func NewOrderSubscriptionUpdateDataAttributesWithDefaults() *OrderSubscriptionUpdateDataAttributes`

NewOrderSubscriptionUpdateDataAttributesWithDefaults instantiates a new OrderSubscriptionUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *OrderSubscriptionUpdateDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrderSubscriptionUpdateDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrderSubscriptionUpdateDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetActivate

`func (o *OrderSubscriptionUpdateDataAttributes) GetActivate() bool`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetActivateOk() (*bool, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *OrderSubscriptionUpdateDataAttributes) SetActivate(v bool)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *OrderSubscriptionUpdateDataAttributes) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *OrderSubscriptionUpdateDataAttributes) GetDeactivate() bool`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetDeactivateOk() (*bool, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *OrderSubscriptionUpdateDataAttributes) SetDeactivate(v bool)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *OrderSubscriptionUpdateDataAttributes) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetCancel

`func (o *OrderSubscriptionUpdateDataAttributes) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *OrderSubscriptionUpdateDataAttributes) SetCancel(v bool)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *OrderSubscriptionUpdateDataAttributes) HasCancel() bool`

HasCancel returns a boolean if a field has been set.

### GetReference

`func (o *OrderSubscriptionUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OrderSubscriptionUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OrderSubscriptionUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *OrderSubscriptionUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *OrderSubscriptionUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *OrderSubscriptionUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderSubscriptionUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderSubscriptionUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderSubscriptionUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderSubscriptionUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


