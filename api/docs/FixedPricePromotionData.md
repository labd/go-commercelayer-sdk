# FixedPricePromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETFixedPricePromotions200ResponseDataInnerAttributes**](GETFixedPricePromotions200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionDataRelationships**](FixedPricePromotionDataRelationships.md) |  | [optional] 

## Methods

### NewFixedPricePromotionData

`func NewFixedPricePromotionData(type_ interface{}, attributes GETFixedPricePromotions200ResponseDataInnerAttributes, ) *FixedPricePromotionData`

NewFixedPricePromotionData instantiates a new FixedPricePromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionDataWithDefaults

`func NewFixedPricePromotionDataWithDefaults() *FixedPricePromotionData`

NewFixedPricePromotionDataWithDefaults instantiates a new FixedPricePromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedPricePromotionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedPricePromotionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedPricePromotionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FixedPricePromotionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FixedPricePromotionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FixedPricePromotionData) GetAttributes() GETFixedPricePromotions200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedPricePromotionData) GetAttributesOk() (*GETFixedPricePromotions200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedPricePromotionData) SetAttributes(v GETFixedPricePromotions200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedPricePromotionData) GetRelationships() FixedPricePromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedPricePromotionData) GetRelationshipsOk() (*FixedPricePromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedPricePromotionData) SetRelationships(v FixedPricePromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedPricePromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


