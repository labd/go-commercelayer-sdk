# GETFixedAmountPromotions200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The promotion&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. | [optional] 
**TotalUsageCount** | Pointer to **interface{}** | The number of times this promotion has been applied. | [optional] 
**Active** | Pointer to **interface{}** | Indicates if the promotion is active. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**FixedAmountCents** | Pointer to **interface{}** | The discount fixed amount to be applied, in cents | [optional] 
**FixedAmountFloat** | Pointer to **interface{}** | The discount fixed amount to be applied, float. | [optional] 
**FormattedFixedAmount** | Pointer to **interface{}** | The discount fixed amount to be applied, formatted. | [optional] 

## Methods

### NewGETFixedAmountPromotions200ResponseDataInnerAttributes

`func NewGETFixedAmountPromotions200ResponseDataInnerAttributes() *GETFixedAmountPromotions200ResponseDataInnerAttributes`

NewGETFixedAmountPromotions200ResponseDataInnerAttributes instantiates a new GETFixedAmountPromotions200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETFixedAmountPromotions200ResponseDataInnerAttributesWithDefaults

`func NewGETFixedAmountPromotions200ResponseDataInnerAttributesWithDefaults() *GETFixedAmountPromotions200ResponseDataInnerAttributes`

NewGETFixedAmountPromotions200ResponseDataInnerAttributesWithDefaults instantiates a new GETFixedAmountPromotions200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetStartsAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetTotalUsageCount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetTotalUsageCount() interface{}`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetTotalUsageCountOk() (*interface{}, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetTotalUsageCount(v interface{})`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### SetTotalUsageCountNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetTotalUsageCountNil(b bool)`

 SetTotalUsageCountNil sets the value for TotalUsageCount to be an explicit nil

### UnsetTotalUsageCount
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetTotalUsageCount()`

UnsetTotalUsageCount ensures that no value is present for TotalUsageCount, not even an explicit nil
### GetActive

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetCreatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFixedAmountCents

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFixedAmountCents() interface{}`

GetFixedAmountCents returns the FixedAmountCents field if non-nil, zero value otherwise.

### GetFixedAmountCentsOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFixedAmountCentsOk() (*interface{}, bool)`

GetFixedAmountCentsOk returns a tuple with the FixedAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountCents

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFixedAmountCents(v interface{})`

SetFixedAmountCents sets FixedAmountCents field to given value.

### HasFixedAmountCents

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasFixedAmountCents() bool`

HasFixedAmountCents returns a boolean if a field has been set.

### SetFixedAmountCentsNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFixedAmountCentsNil(b bool)`

 SetFixedAmountCentsNil sets the value for FixedAmountCents to be an explicit nil

### UnsetFixedAmountCents
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetFixedAmountCents()`

UnsetFixedAmountCents ensures that no value is present for FixedAmountCents, not even an explicit nil
### GetFixedAmountFloat

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFixedAmountFloat() interface{}`

GetFixedAmountFloat returns the FixedAmountFloat field if non-nil, zero value otherwise.

### GetFixedAmountFloatOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFixedAmountFloatOk() (*interface{}, bool)`

GetFixedAmountFloatOk returns a tuple with the FixedAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedAmountFloat

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFixedAmountFloat(v interface{})`

SetFixedAmountFloat sets FixedAmountFloat field to given value.

### HasFixedAmountFloat

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasFixedAmountFloat() bool`

HasFixedAmountFloat returns a boolean if a field has been set.

### SetFixedAmountFloatNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFixedAmountFloatNil(b bool)`

 SetFixedAmountFloatNil sets the value for FixedAmountFloat to be an explicit nil

### UnsetFixedAmountFloat
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetFixedAmountFloat()`

UnsetFixedAmountFloat ensures that no value is present for FixedAmountFloat, not even an explicit nil
### GetFormattedFixedAmount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFormattedFixedAmount() interface{}`

GetFormattedFixedAmount returns the FormattedFixedAmount field if non-nil, zero value otherwise.

### GetFormattedFixedAmountOk

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) GetFormattedFixedAmountOk() (*interface{}, bool)`

GetFormattedFixedAmountOk returns a tuple with the FormattedFixedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedFixedAmount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFormattedFixedAmount(v interface{})`

SetFormattedFixedAmount sets FormattedFixedAmount field to given value.

### HasFormattedFixedAmount

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) HasFormattedFixedAmount() bool`

HasFormattedFixedAmount returns a boolean if a field has been set.

### SetFormattedFixedAmountNil

`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) SetFormattedFixedAmountNil(b bool)`

 SetFormattedFixedAmountNil sets the value for FormattedFixedAmount to be an explicit nil

### UnsetFormattedFixedAmount
`func (o *GETFixedAmountPromotions200ResponseDataInnerAttributes) UnsetFormattedFixedAmount()`

UnsetFormattedFixedAmount ensures that no value is present for FormattedFixedAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


