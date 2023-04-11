# GETPercentageDiscountPromotions200ResponseDataInnerAttributes

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
**Percentage** | Pointer to **interface{}** | The discount percentage to be applied. | [optional] 

## Methods

### NewGETPercentageDiscountPromotions200ResponseDataInnerAttributes

`func NewGETPercentageDiscountPromotions200ResponseDataInnerAttributes() *GETPercentageDiscountPromotions200ResponseDataInnerAttributes`

NewGETPercentageDiscountPromotions200ResponseDataInnerAttributes instantiates a new GETPercentageDiscountPromotions200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPercentageDiscountPromotions200ResponseDataInnerAttributesWithDefaults

`func NewGETPercentageDiscountPromotions200ResponseDataInnerAttributesWithDefaults() *GETPercentageDiscountPromotions200ResponseDataInnerAttributes`

NewGETPercentageDiscountPromotions200ResponseDataInnerAttributesWithDefaults instantiates a new GETPercentageDiscountPromotions200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetStartsAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetTotalUsageCount

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetTotalUsageCount() interface{}`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetTotalUsageCountOk() (*interface{}, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetTotalUsageCount(v interface{})`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### SetTotalUsageCountNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetTotalUsageCountNil(b bool)`

 SetTotalUsageCountNil sets the value for TotalUsageCount to be an explicit nil

### UnsetTotalUsageCount
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetTotalUsageCount()`

UnsetTotalUsageCount ensures that no value is present for TotalUsageCount, not even an explicit nil
### GetActive

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetCreatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPercentage

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetPercentage() interface{}`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) GetPercentageOk() (*interface{}, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetPercentage(v interface{})`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *GETPercentageDiscountPromotions200ResponseDataInnerAttributes) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


