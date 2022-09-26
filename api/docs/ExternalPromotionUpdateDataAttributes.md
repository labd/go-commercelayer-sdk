# ExternalPromotionUpdateDataAttributes

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
**PromotionUrl** | Pointer to **string** | The URL to the service that will compute the discount. | [optional] 

## Methods

### NewExternalPromotionUpdateDataAttributes

`func NewExternalPromotionUpdateDataAttributes() *ExternalPromotionUpdateDataAttributes`

NewExternalPromotionUpdateDataAttributes instantiates a new ExternalPromotionUpdateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPromotionUpdateDataAttributesWithDefaults

`func NewExternalPromotionUpdateDataAttributesWithDefaults() *ExternalPromotionUpdateDataAttributes`

NewExternalPromotionUpdateDataAttributesWithDefaults instantiates a new ExternalPromotionUpdateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalPromotionUpdateDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalPromotionUpdateDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalPromotionUpdateDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalPromotionUpdateDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ExternalPromotionUpdateDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ExternalPromotionUpdateDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ExternalPromotionUpdateDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ExternalPromotionUpdateDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *ExternalPromotionUpdateDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ExternalPromotionUpdateDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ExternalPromotionUpdateDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *ExternalPromotionUpdateDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ExternalPromotionUpdateDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExternalPromotionUpdateDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExternalPromotionUpdateDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ExternalPromotionUpdateDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTotalUsageLimit

`func (o *ExternalPromotionUpdateDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *ExternalPromotionUpdateDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *ExternalPromotionUpdateDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *ExternalPromotionUpdateDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### GetReference

`func (o *ExternalPromotionUpdateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ExternalPromotionUpdateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ExternalPromotionUpdateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ExternalPromotionUpdateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ExternalPromotionUpdateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ExternalPromotionUpdateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ExternalPromotionUpdateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ExternalPromotionUpdateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ExternalPromotionUpdateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ExternalPromotionUpdateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ExternalPromotionUpdateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ExternalPromotionUpdateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPromotionUrl

`func (o *ExternalPromotionUpdateDataAttributes) GetPromotionUrl() string`

GetPromotionUrl returns the PromotionUrl field if non-nil, zero value otherwise.

### GetPromotionUrlOk

`func (o *ExternalPromotionUpdateDataAttributes) GetPromotionUrlOk() (*string, bool)`

GetPromotionUrlOk returns a tuple with the PromotionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionUrl

`func (o *ExternalPromotionUpdateDataAttributes) SetPromotionUrl(v string)`

SetPromotionUrl sets PromotionUrl field to given value.

### HasPromotionUrl

`func (o *ExternalPromotionUpdateDataAttributes) HasPromotionUrl() bool`

HasPromotionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


