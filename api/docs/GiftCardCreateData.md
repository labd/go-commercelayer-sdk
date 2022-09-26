# GiftCardCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**GiftCardCreateDataAttributes**](GiftCardCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**GiftCardCreateDataRelationships**](GiftCardCreateDataRelationships.md) |  | [optional] 

## Methods

### NewGiftCardCreateData

`func NewGiftCardCreateData(type_ string, attributes GiftCardCreateDataAttributes, ) *GiftCardCreateData`

NewGiftCardCreateData instantiates a new GiftCardCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiftCardCreateDataWithDefaults

`func NewGiftCardCreateDataWithDefaults() *GiftCardCreateData`

NewGiftCardCreateDataWithDefaults instantiates a new GiftCardCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GiftCardCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GiftCardCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GiftCardCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GiftCardCreateData) GetAttributes() GiftCardCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GiftCardCreateData) GetAttributesOk() (*GiftCardCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GiftCardCreateData) SetAttributes(v GiftCardCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GiftCardCreateData) GetRelationships() GiftCardCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GiftCardCreateData) GetRelationshipsOk() (*GiftCardCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GiftCardCreateData) SetRelationships(v GiftCardCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GiftCardCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


