# GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes

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
**MaxQuantity** | Pointer to **interface{}** | The max quantity of free gifts globally applicable by the promotion. | [optional] 

## Methods

### NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes

`func NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes() *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes`

NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes instantiates a new GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributesWithDefaults

`func NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributesWithDefaults() *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes`

NewGETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributesWithDefaults instantiates a new GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCurrencyCode

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetExclusive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetTotalUsageCount

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{}`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{})`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### SetTotalUsageCountNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetTotalUsageCountNil(b bool)`

 SetTotalUsageCountNil sets the value for TotalUsageCount to be an explicit nil

### UnsetTotalUsageCount
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetTotalUsageCount()`

UnsetTotalUsageCount ensures that no value is present for TotalUsageCount, not even an explicit nil
### GetActive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetStatus

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDisabledAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMaxQuantity

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetMaxQuantity() interface{}`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) GetMaxQuantityOk() (*interface{}, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetMaxQuantity(v interface{})`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### SetMaxQuantityNil

`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) SetMaxQuantityNil(b bool)`

 SetMaxQuantityNil sets the value for MaxQuantity to be an explicit nil

### UnsetMaxQuantity
`func (o *GETFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) UnsetMaxQuantity()`

UnsetMaxQuantity ensures that no value is present for MaxQuantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


