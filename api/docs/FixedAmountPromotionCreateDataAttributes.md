# FixedAmountPromotionCreateDataAttributes

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
**FixedAmountCents** | **int32** | The discount fixed amount to be applied, in cents | 

## Methods

### NewFixedAmountPromotionCreateDataAttributes

`func NewFixedAmountPromotionCreateDataAttributes(name string, startsAt string, expiresAt string, totalUsageLimit int32, fixedAmountCents int32, ) *FixedAmountPromotionCreateDataAttributes`

NewFixedAmountPromotionCreateDataAttributes instantiates a new FixedAmountPromotionCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAmountPromotionCreateDataAttributesWithDefaults

`func NewFixedAmountPromotionCreateDataAttributesWithDefaults() *FixedAmountPromotionCreateDataAttributes`

NewFixedAmountPromotionCreateDataAttributesWithDefaults instantiates a new FixedAmountPromotionCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FixedAmountPromotionCreateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedAmountPromotionCreateDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetCurrencyCode

`func (o *FixedAmountPromotionCreateDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FixedAmountPromotionCreateDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FixedAmountPromotionCreateDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *FixedAmountPromotionCreateDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *FixedAmountPromotionCreateDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.


### GetExpiresAt

`func (o *FixedAmountPromotionCreateDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FixedAmountPromotionCreateDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetTotalUsageLimit

`func (o *FixedAmountPromotionCreateDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *FixedAmountPromotionCreateDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.


### GetReference

`func (o *FixedAmountPromotionCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FixedAmountPromotionCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *FixedAmountPromotionCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *FixedAmountPromotionCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *FixedAmountPromotionCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *FixedAmountPromotionCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *FixedAmountPromotionCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FixedAmountPromotionCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FixedAmountPromotionCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFixedAmountCents

`func (o *FixedAmountPromotionCreateDataAttributes) GetFixedAmountCents() int32`

GetFixedAmountCents returns the FixedAmountCents field if non-nil, zero value otherwise.

### GetFixedAmountCentsOk

`func (o *FixedAmountPromotionCreateDataAttributes) GetFixedAmountCentsOk() (*int32, bool)`

GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountCents

`func (o *FixedAmountPromotionCreateDataAttributes) SetFixedAmountCents(v int32)`

SetFixedAmountCents sets FixedAmountCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


