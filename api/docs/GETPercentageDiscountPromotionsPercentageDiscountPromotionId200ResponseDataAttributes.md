# GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The promotion&#39;s internal name. | [optional] 
**Type** | Pointer to **interface{}** | The promotion&#39;s type. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Exclusive** | Pointer to **interface{}** | Indicates if the promotion will be applied exclusively, based on its priority score. | [optional] 
**Priority** | Pointer to **interface{}** | The priority assigned to the promotion (lower means higher priority). | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. When &#39;null&#39; it means promotion can be applied infinite times. | [optional] 
**TotalUsageCount** | Pointer to **interface{}** | The number of times this promotion has been applied. | [optional] 
**Active** | Pointer to **interface{}** | Indicates if the promotion is active (enabled and not expired). | [optional] 
**Status** | Pointer to **interface{}** | The promotion status. One of &#39;disabled&#39;, &#39;expired&#39;, &#39;pending&#39;, &#39;active&#39;, or &#39;inactive&#39;. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which this resource was disabled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Percentage** | Pointer to **interface{}** | The discount percentage to be applied. | [optional] 

## Methods

### NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes

`func NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes() *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes`

NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes instantiates a new GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults

`func NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults() *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes`

NewGETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributesWithDefaults instantiates a new GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCurrencyCode

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExclusive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetTotalUsageCount

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{}`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{})`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### SetTotalUsageCountNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetTotalUsageCountNil(b bool)`

 SetTotalUsageCountNil sets the value for TotalUsageCount to be an explicit nil

### UnsetTotalUsageCount
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetTotalUsageCount()`

UnsetTotalUsageCount ensures that no value is present for TotalUsageCount, not even an explicit nil
### GetActive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetStatus

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDisabledAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPercentage

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPercentage() interface{}`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) GetPercentageOk() (*interface{}, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetPercentage(v interface{})`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *GETPercentageDiscountPromotionsPercentageDiscountPromotionId200ResponseDataAttributes) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


