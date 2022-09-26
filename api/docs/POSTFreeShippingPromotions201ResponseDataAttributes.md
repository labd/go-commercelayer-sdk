# POSTFreeShippingPromotions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The promotion&#39;s internal name. | 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**StartsAt** | **string** | The activation date/time of this promotion. | 
**ExpiresAt** | **string** | The expiration date/time of this promotion (must be after starts_at). | 
**TotalUsageLimit** | **int32** | The total number of times this promotion can be applied. | 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTFreeShippingPromotions201ResponseDataAttributes

`func NewPOSTFreeShippingPromotions201ResponseDataAttributes(name string, startsAt string, expiresAt string, totalUsageLimit int32, ) *POSTFreeShippingPromotions201ResponseDataAttributes`

NewPOSTFreeShippingPromotions201ResponseDataAttributes instantiates a new POSTFreeShippingPromotions201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults

`func NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults() *POSTFreeShippingPromotions201ResponseDataAttributes`

NewPOSTFreeShippingPromotions201ResponseDataAttributesWithDefaults instantiates a new POSTFreeShippingPromotions201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.


### GetExpiresAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetTotalUsageLimit

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.


### GetReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTFreeShippingPromotions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


