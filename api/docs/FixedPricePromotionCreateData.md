# FixedPricePromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFixedPricePromotions201ResponseDataAttributes**](POSTFixedPricePromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**FixedPricePromotionCreateDataRelationships**](FixedPricePromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFixedPricePromotionCreateData

`func NewFixedPricePromotionCreateData(type_ interface{}, attributes POSTFixedPricePromotions201ResponseDataAttributes, ) *FixedPricePromotionCreateData`

NewFixedPricePromotionCreateData instantiates a new FixedPricePromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedPricePromotionCreateDataWithDefaults

`func NewFixedPricePromotionCreateDataWithDefaults() *FixedPricePromotionCreateData`

NewFixedPricePromotionCreateDataWithDefaults instantiates a new FixedPricePromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedPricePromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedPricePromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedPricePromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FixedPricePromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FixedPricePromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FixedPricePromotionCreateData) GetAttributes() POSTFixedPricePromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedPricePromotionCreateData) GetAttributesOk() (*POSTFixedPricePromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedPricePromotionCreateData) SetAttributes(v POSTFixedPricePromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedPricePromotionCreateData) GetRelationships() FixedPricePromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedPricePromotionCreateData) GetRelationshipsOk() (*FixedPricePromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedPricePromotionCreateData) SetRelationships(v FixedPricePromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedPricePromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


