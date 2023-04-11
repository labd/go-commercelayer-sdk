# PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The promotion&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**MaxQuantity** | Pointer to **interface{}** | The max quantity of free gifts globally applicable by the promotion. | [optional] 

## Methods

### NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes

`func NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes() *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes`

NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributesWithDefaults

`func NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributesWithDefaults() *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes`

NewPATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributesWithDefaults instantiates a new PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetStartsAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetReference

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMaxQuantity

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetMaxQuantity() interface{}`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) GetMaxQuantityOk() (*interface{}, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetMaxQuantity(v interface{})`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### SetMaxQuantityNil

`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) SetMaxQuantityNil(b bool)`

 SetMaxQuantityNil sets the value for MaxQuantity to be an explicit nil

### UnsetMaxQuantity
`func (o *PATCHFreeGiftPromotionsFreeGiftPromotionIdRequestDataAttributes) UnsetMaxQuantity()`

UnsetMaxQuantity ensures that no value is present for MaxQuantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


