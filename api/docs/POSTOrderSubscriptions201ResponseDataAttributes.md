# POSTOrderSubscriptions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | **string** | The frequency of the subscription. One of &#39;hourly&#39;, &#39;daily&#39;, &#39;weekly&#39;, &#39;monthly&#39;, &#39;two-month&#39;, &#39;three-month&#39;, &#39;four-month&#39;, &#39;six-month&#39;, or &#39;yearly&#39;. | 
**ActivateBySourceOrder** | Pointer to **bool** | Indicates if the subscription will be activated considering the placed source order as its first run, default to true. | [optional] 
**StartsAt** | Pointer to **string** | The activation date/time of this subscription. | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration date/time of this subscription (must be after starts_at). | [optional] 
**Options** | Pointer to **map[string]interface{}** | The subscription options used to create the order copy (check order_copies for more information). For subscriptions the &#x60;place_target_order&#x60; is enabled by default, specify custom options to overwrite it. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTOrderSubscriptions201ResponseDataAttributes

`func NewPOSTOrderSubscriptions201ResponseDataAttributes(frequency string, ) *POSTOrderSubscriptions201ResponseDataAttributes`

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

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrder() bool`

GetActivateBySourceOrder returns the ActivateBySourceOrder field if non-nil, zero value otherwise.

### GetActivateBySourceOrderOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetActivateBySourceOrderOk() (*bool, bool)`

GetActivateBySourceOrderOk returns a tuple with the ActivateBySourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetActivateBySourceOrder(v bool)`

SetActivateBySourceOrder sets ActivateBySourceOrder field to given value.

### HasActivateBySourceOrder

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasActivateBySourceOrder() bool`

HasActivateBySourceOrder returns a boolean if a field has been set.

### GetStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetOptions

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderSubscriptions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


