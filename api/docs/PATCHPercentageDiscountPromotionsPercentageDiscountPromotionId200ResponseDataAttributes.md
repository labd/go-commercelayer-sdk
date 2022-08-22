# PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The promotion&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**StartsAt** | Pointer to **string** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **int32** | The total number of times this promotion can be applied. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Percentage** | Pointer to **int32** | The discount percentage to be applied. | [optional] 

## Methods

### NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes

`func NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes() *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes`

NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes instantiates a new PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults

`func NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults() *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes`

NewPATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults instantiates a new PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTotalUsageLimit

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### GetReference

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPercentage

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PATCHPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


