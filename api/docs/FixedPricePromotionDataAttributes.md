# FixedPricePromotionDataAttributes

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
**FixedAmountCents** | Pointer to **int32** | The price fixed amount to be applied on matching SKUs, in cents | [optional] 
**FixedAmountFloat** | Pointer to **float32** | The discount fixed amount to be applied, float. | [optional] 
**FormattedFixedAmount** | Pointer to **string** | The discount fixed amount to be applied, formatted. | [optional] 

## Methods

### NewFixedPricePromotionDataAttributes

`func NewFixedPricePromotionDataAttributes() *FixedPricePromotionDataAttributes`

NewFixedPricePromotionDataAttributes instantiates a new FixedPricePromotionDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionDataAttributesWithDefaults

`func NewFixedPricePromotionDataAttributesWithDefaults() *FixedPricePromotionDataAttributes`

NewFixedPricePromotionDataAttributesWithDefaults instantiates a new FixedPricePromotionDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FixedPricePromotionDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedPricePromotionDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedPricePromotionDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FixedPricePromotionDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *FixedPricePromotionDataAttributes) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FixedPricePromotionDataAttributes) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FixedPricePromotionDataAttributes) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FixedPricePromotionDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetStartsAt

`func (o *FixedPricePromotionDataAttributes) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *FixedPricePromotionDataAttributes) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *FixedPricePromotionDataAttributes) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *FixedPricePromotionDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *FixedPricePromotionDataAttributes) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FixedPricePromotionDataAttributes) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FixedPricePromotionDataAttributes) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FixedPricePromotionDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetTotalUsageLimit

`func (o *FixedPricePromotionDataAttributes) GetTotalUsageLimit() int32`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *FixedPricePromotionDataAttributes) GetTotalUsageLimitOk() (*int32, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *FixedPricePromotionDataAttributes) SetTotalUsageLimit(v int32)`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *FixedPricePromotionDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### GetTotalUsageCount

`func (o *FixedPricePromotionDataAttributes) GetTotalUsageCount() int32`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *FixedPricePromotionDataAttributes) GetTotalUsageCountOk() (*int32, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *FixedPricePromotionDataAttributes) SetTotalUsageCount(v int32)`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *FixedPricePromotionDataAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### GetActive

`func (o *FixedPricePromotionDataAttributes) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FixedPricePromotionDataAttributes) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FixedPricePromotionDataAttributes) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FixedPricePromotionDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *FixedPricePromotionDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedPricePromotionDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedPricePromotionDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FixedPricePromotionDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FixedPricePromotionDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FixedPricePromotionDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FixedPricePromotionDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FixedPricePromotionDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FixedPricePromotionDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FixedPricePromotionDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FixedPricePromotionDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FixedPricePromotionDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *FixedPricePromotionDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FixedPricePromotionDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FixedPricePromotionDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *FixedPricePromotionDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *FixedPricePromotionDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *FixedPricePromotionDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *FixedPricePromotionDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *FixedPricePromotionDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *FixedPricePromotionDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FixedPricePromotionDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FixedPricePromotionDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FixedPricePromotionDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFixedAmountCents

`func (o *FixedPricePromotionDataAttributes) GetFixedAmountCents() int32`

GetFixedAmountCents returns the FixedAmountCents field if non-nil, zero value otherwise.

### GetFixedAmountCentsOk

`func (o *FixedPricePromotionDataAttributes) GetFixedAmountCentsOk() (*int32, bool)`

GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountCents

`func (o *FixedPricePromotionDataAttributes) SetFixedAmountCents(v int32)`

SetFixedAmountCents sets FixedAmountCents field to given value.

### HasFixedAmountCents

`func (o *FixedPricePromotionDataAttributes) HasFixedAmountCents() bool`

HasFixedAmountCents returns a boolean if a field has been set.

### GetFixedAmountFloat

`func (o *FixedPricePromotionDataAttributes) GetFixedAmountFloat() float32`

GetFixedAmountFloat returns the FixedAmountFloat field if non-nil, zero value otherwise.

### GetFixedAmountFloatOk

`func (o *FixedPricePromotionDataAttributes) GetFixedAmountFloatOk() (*float32, bool)`

GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountFloat

`func (o *FixedPricePromotionDataAttributes) SetFixedAmountFloat(v float32)`

SetFixedAmountFloat sets FixedAmountFloat field to given value.

### HasFixedAmountFloat

`func (o *FixedPricePromotionDataAttributes) HasFixedAmountFloat() bool`

HasFixedAmountFloat returns a boolean if a field has been set.

### GetFormattedFixedAmount

`func (o *FixedPricePromotionDataAttributes) GetFormattedFixedAmount() string`

GetFormattedFixedAmount returns the FormattedFixedAmount field if non-nil, zero value otherwise.

### GetFormattedFixedAmountOk

`func (o *FixedPricePromotionDataAttributes) GetFormattedFixedAmountOk() (*string, bool)`

GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFixedAmount

`func (o *FixedPricePromotionDataAttributes) SetFormattedFixedAmount(v string)`

SetFormattedFixedAmount sets FormattedFixedAmount field to given value.

### HasFormattedFixedAmount

`func (o *FixedPricePromotionDataAttributes) HasFormattedFixedAmount() bool`

HasFormattedFixedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


