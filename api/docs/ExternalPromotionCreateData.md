# ExternalPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTExternalPromotions201ResponseDataAttributes**](POSTExternalPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BuyXPayYPromotionUpdateDataRelationships**](BuyXPayYPromotionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPromotionCreateData

`func NewExternalPromotionCreateData(type_ interface{}, attributes POSTExternalPromotions201ResponseDataAttributes, ) *ExternalPromotionCreateData`

NewExternalPromotionCreateData instantiates a new ExternalPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPromotionCreateDataWithDefaults

`func NewExternalPromotionCreateDataWithDefaults() *ExternalPromotionCreateData`

NewExternalPromotionCreateDataWithDefaults instantiates a new ExternalPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalPromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalPromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalPromotionCreateData) GetAttributes() POSTExternalPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPromotionCreateData) GetAttributesOk() (*POSTExternalPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPromotionCreateData) SetAttributes(v POSTExternalPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPromotionCreateData) GetRelationships() BuyXPayYPromotionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPromotionCreateData) SetRelationships(v BuyXPayYPromotionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


