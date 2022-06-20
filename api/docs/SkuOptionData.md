# SkuOptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "sku_options"]
**Attributes** | [**SkuOptionDataAttributes**](SkuOptionDataAttributes.md) |  | 
**Relationships** | Pointer to [**CarrierAccountDataRelationships**](CarrierAccountDataRelationships.md) |  | [optional] 

## Methods

### NewSkuOptionData

`func NewSkuOptionData(type_ string, attributes SkuOptionDataAttributes, ) *SkuOptionData`

NewSkuOptionData instantiates a new SkuOptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuOptionDataWithDefaults

`func NewSkuOptionDataWithDefaults() *SkuOptionData`

NewSkuOptionDataWithDefaults instantiates a new SkuOptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuOptionData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuOptionData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuOptionData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuOptionData) GetAttributes() SkuOptionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuOptionData) GetAttributesOk() (*SkuOptionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuOptionData) SetAttributes(v SkuOptionDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuOptionData) GetRelationships() CarrierAccountDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuOptionData) GetRelationshipsOk() (*CarrierAccountDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuOptionData) SetRelationships(v CarrierAccountDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuOptionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


