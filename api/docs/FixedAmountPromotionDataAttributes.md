# FixedAmountPromotionDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The promotion&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **string** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**StartsAt** | Pointer to **string** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **string** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **int32** | The total number of times this promotion can be applied. | [optional] 
**TotalUsageCount** | Pointer to **int32** | The number of times this promotion has been applied. | [optional] 
**Active** | Pointer to **bool** | Indicates if the promotion is active. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**FixedAmountCents** | Pointer to **int32** | The discount fixed amount to be applied, in cents | [optional] 
**FixedAmountFloat** | Pointer to **float32** | The discount fixed amount to be applied, float. | [optional] 
**FormattedFixedAmount** | Pointer to **string** | The discount fixed amount to be applied, formatted. | [optional] 

## Methods

### NewFixedAmountPromotionDataAttributes

`func NewFixedAmountPromotionDataAttributes() *FixedAmountPromotionDataAttributes`

NewFixedAmountPromotionDataAttributes instantiates a new FixedAmountPromotionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAmountPromotionDataAttributesWithDefaults

`func NewFixedAmountPromotionDataAttributesWithDefaults() *FixedAmountPromotionDataAttributes`

NewFixedAmountPromotionDataAttributesWithDefaults instantiates a new FixedAmountPromotionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FixedAmountPromotionDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedAmountPromotionDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedAmountPromotionDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FixedAmountPromotionDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *FixedAmountPromotionDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FixedAmountPromotionDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FixedAmountPromotionDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FixedAmountPromotionDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *FixedAmountPromotionDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *FixedAmountPromotionDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *FixedAmountPromotionDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *FixedAmountPromotionDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *FixedAmountPromotionDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FixedAmountPromotionDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FixedAmountPromotionDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FixedAmountPromotionDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTotalUsageLimit

`func (o *FixedAmountPromotionDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *FixedAmountPromotionDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *FixedAmountPromotionDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *FixedAmountPromotionDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### GetTotalUsageCount

`func (o *FixedAmountPromotionDataAttributes) GetTotalUsageCount() int32`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *FixedAmountPromotionDataAttributes) GetTotalUsageCountOk() (*int32, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *FixedAmountPromotionDataAttributes) SetTotalUsageCount(v int32)`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *FixedAmountPromotionDataAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### GetActive

`func (o *FixedAmountPromotionDataAttributes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FixedAmountPromotionDataAttributes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FixedAmountPromotionDataAttributes) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FixedAmountPromotionDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *FixedAmountPromotionDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedAmountPromotionDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedAmountPromotionDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FixedAmountPromotionDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FixedAmountPromotionDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FixedAmountPromotionDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FixedAmountPromotionDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FixedAmountPromotionDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FixedAmountPromotionDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FixedAmountPromotionDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FixedAmountPromotionDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FixedAmountPromotionDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *FixedAmountPromotionDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FixedAmountPromotionDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FixedAmountPromotionDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *FixedAmountPromotionDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *FixedAmountPromotionDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *FixedAmountPromotionDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *FixedAmountPromotionDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *FixedAmountPromotionDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *FixedAmountPromotionDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FixedAmountPromotionDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FixedAmountPromotionDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FixedAmountPromotionDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFixedAmountCents

`func (o *FixedAmountPromotionDataAttributes) GetFixedAmountCents() int32`

GetFixedAmountCents returns the FixedAmountCents field if non-nil, zero value otherwise.

### GetFixedAmountCentsOk

`func (o *FixedAmountPromotionDataAttributes) GetFixedAmountCentsOk() (*int32, bool)`

GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountCents

`func (o *FixedAmountPromotionDataAttributes) SetFixedAmountCents(v int32)`

SetFixedAmountCents sets FixedAmountCents field to given value.

### HasFixedAmountCents

`func (o *FixedAmountPromotionDataAttributes) HasFixedAmountCents() bool`

HasFixedAmountCents returns a boolean if a field has been set.

### GetFixedAmountFloat

`func (o *FixedAmountPromotionDataAttributes) GetFixedAmountFloat() float32`

GetFixedAmountFloat returns the FixedAmountFloat field if non-nil, zero value otherwise.

### GetFixedAmountFloatOk

`func (o *FixedAmountPromotionDataAttributes) GetFixedAmountFloatOk() (*float32, bool)`

GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountFloat

`func (o *FixedAmountPromotionDataAttributes) SetFixedAmountFloat(v float32)`

SetFixedAmountFloat sets FixedAmountFloat field to given value.

### HasFixedAmountFloat

`func (o *FixedAmountPromotionDataAttributes) HasFixedAmountFloat() bool`

HasFixedAmountFloat returns a boolean if a field has been set.

### GetFormattedFixedAmount

`func (o *FixedAmountPromotionDataAttributes) GetFormattedFixedAmount() string`

GetFormattedFixedAmount returns the FormattedFixedAmount field if non-nil, zero value otherwise.

### GetFormattedFixedAmountOk

`func (o *FixedAmountPromotionDataAttributes) GetFormattedFixedAmountOk() (*string, bool)`

GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFixedAmount

`func (o *FixedAmountPromotionDataAttributes) SetFormattedFixedAmount(v string)`

SetFormattedFixedAmount sets FormattedFixedAmount field to given value.

### HasFormattedFixedAmount

`func (o *FixedAmountPromotionDataAttributes) HasFormattedFixedAmount() bool`

HasFormattedFixedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


