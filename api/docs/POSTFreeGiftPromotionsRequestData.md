# POSTFreeGiftPromotionsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFreeGiftPromotionsRequestDataAttributes**](POSTFreeGiftPromotionsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTFixedPricePromotionsRequestDataRelationships**](POSTFixedPricePromotionsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTFreeGiftPromotionsRequestData

`func NewPOSTFreeGiftPromotionsRequestData(type_ interface{}, attributes POSTFreeGiftPromotionsRequestDataAttributes, ) *POSTFreeGiftPromotionsRequestData`

NewPOSTFreeGiftPromotionsRequestData instantiates a new POSTFreeGiftPromotionsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTFreeGiftPromotionsRequestDataWithDefaults

`func NewPOSTFreeGiftPromotionsRequestDataWithDefaults() *POSTFreeGiftPromotionsRequestData`

NewPOSTFreeGiftPromotionsRequestDataWithDefaults instantiates a new POSTFreeGiftPromotionsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTFreeGiftPromotionsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTFreeGiftPromotionsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTFreeGiftPromotionsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTFreeGiftPromotionsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTFreeGiftPromotionsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTFreeGiftPromotionsRequestData) GetAttributes() POSTFreeGiftPromotionsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTFreeGiftPromotionsRequestData) GetAttributesOk() (*POSTFreeGiftPromotionsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTFreeGiftPromotionsRequestData) SetAttributes(v POSTFreeGiftPromotionsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTFreeGiftPromotionsRequestData) GetRelationships() POSTFixedPricePromotionsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTFreeGiftPromotionsRequestData) GetRelationshipsOk() (*POSTFixedPricePromotionsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTFreeGiftPromotionsRequestData) SetRelationships(v POSTFixedPricePromotionsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTFreeGiftPromotionsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


