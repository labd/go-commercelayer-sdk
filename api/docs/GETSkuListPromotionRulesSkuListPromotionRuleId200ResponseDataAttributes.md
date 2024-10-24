# GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** | The promotion rule&#39;s type. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**AllSkus** | Pointer to **interface{}** | Indicates if the rule is activated only when all of the SKUs of the list is also part of the order. | [optional] 
**MinQuantity** | Pointer to **interface{}** | The min quantity of SKUs of the list that must be also part of the order. If positive, overwrites the &#39;all_skus&#39; option. When the SKU list is manual, its items quantities are honoured. | [optional] 

## Methods

### NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes

`func NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes() *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes`

NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes instantiates a new GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributesWithDefaults

`func NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributesWithDefaults() *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes`

NewGETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributesWithDefaults instantiates a new GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAllSkus

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetAllSkus() interface{}`

GetAllSkus returns the AllSkus field if non-nil, zero value otherwise.

### GetAllSkusOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetAllSkusOk() (*interface{}, bool)`

GetAllSkusOk returns a tuple with the AllSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSkus

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetAllSkus(v interface{})`

SetAllSkus sets AllSkus field to given value.

### HasAllSkus

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasAllSkus() bool`

HasAllSkus returns a boolean if a field has been set.

### SetAllSkusNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetAllSkusNil(b bool)`

 SetAllSkusNil sets the value for AllSkus to be an explicit nil

### UnsetAllSkus
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetAllSkus()`

UnsetAllSkus ensures that no value is present for AllSkus, not even an explicit nil
### GetMinQuantity

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetMinQuantity() interface{}`

GetMinQuantity returns the MinQuantity field if non-nil, zero value otherwise.

### GetMinQuantityOk

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) GetMinQuantityOk() (*interface{}, bool)`

GetMinQuantityOk returns a tuple with the MinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQuantity

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetMinQuantity(v interface{})`

SetMinQuantity sets MinQuantity field to given value.

### HasMinQuantity

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) HasMinQuantity() bool`

HasMinQuantity returns a boolean if a field has been set.

### SetMinQuantityNil

`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) SetMinQuantityNil(b bool)`

 SetMinQuantityNil sets the value for MinQuantity to be an explicit nil

### UnsetMinQuantity
`func (o *GETSkuListPromotionRulesSkuListPromotionRuleId200ResponseDataAttributes) UnsetMinQuantity()`

UnsetMinQuantity ensures that no value is present for MinQuantity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


