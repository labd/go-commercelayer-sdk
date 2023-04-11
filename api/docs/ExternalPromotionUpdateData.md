# ExternalPromotionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes**](PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionCreateDataRelationships**](ExternalPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPromotionUpdateData

`func NewExternalPromotionUpdateData(type_ interface{}, id interface{}, attributes PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes, ) *ExternalPromotionUpdateData`

NewExternalPromotionUpdateData instantiates a new ExternalPromotionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPromotionUpdateDataWithDefaults

`func NewExternalPromotionUpdateDataWithDefaults() *ExternalPromotionUpdateData`

NewExternalPromotionUpdateDataWithDefaults instantiates a new ExternalPromotionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPromotionUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPromotionUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPromotionUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalPromotionUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalPromotionUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ExternalPromotionUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalPromotionUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalPromotionUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalPromotionUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalPromotionUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ExternalPromotionUpdateData) GetAttributes() PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPromotionUpdateData) GetAttributesOk() (*PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPromotionUpdateData) SetAttributes(v PATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPromotionUpdateData) GetRelationships() ExternalPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPromotionUpdateData) GetRelationshipsOk() (*ExternalPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPromotionUpdateData) SetRelationships(v ExternalPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPromotionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


