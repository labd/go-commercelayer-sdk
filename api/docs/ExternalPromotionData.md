# ExternalPromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETExternalPromotionsExternalPromotionId200ResponseDataAttributes**](GETExternalPromotionsExternalPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPromotionDataRelationships**](ExternalPromotionDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPromotionData

`func NewExternalPromotionData(type_ interface{}, attributes GETExternalPromotionsExternalPromotionId200ResponseDataAttributes, ) *ExternalPromotionData`

NewExternalPromotionData instantiates a new ExternalPromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPromotionDataWithDefaults

`func NewExternalPromotionDataWithDefaults() *ExternalPromotionData`

NewExternalPromotionDataWithDefaults instantiates a new ExternalPromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPromotionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPromotionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPromotionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalPromotionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalPromotionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalPromotionData) GetAttributes() GETExternalPromotionsExternalPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPromotionData) GetAttributesOk() (*GETExternalPromotionsExternalPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPromotionData) SetAttributes(v GETExternalPromotionsExternalPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPromotionData) GetRelationships() ExternalPromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPromotionData) GetRelationshipsOk() (*ExternalPromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPromotionData) SetRelationships(v ExternalPromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


