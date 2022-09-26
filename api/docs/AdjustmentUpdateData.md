# AdjustmentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AdjustmentUpdateDataAttributes**](AdjustmentUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAdjustmentUpdateData

`func NewAdjustmentUpdateData(type_ string, id string, attributes AdjustmentUpdateDataAttributes, ) *AdjustmentUpdateData`

NewAdjustmentUpdateData instantiates a new AdjustmentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentUpdateDataWithDefaults

`func NewAdjustmentUpdateDataWithDefaults() *AdjustmentUpdateData`

NewAdjustmentUpdateDataWithDefaults instantiates a new AdjustmentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdjustmentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdjustmentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdjustmentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AdjustmentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdjustmentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdjustmentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AdjustmentUpdateData) GetAttributes() AdjustmentUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdjustmentUpdateData) GetAttributesOk() (*AdjustmentUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdjustmentUpdateData) SetAttributes(v AdjustmentUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdjustmentUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdjustmentUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdjustmentUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdjustmentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


