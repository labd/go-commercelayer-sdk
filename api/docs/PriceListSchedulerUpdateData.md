# PriceListSchedulerUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes**](PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceListSchedulerUpdateDataRelationships**](PriceListSchedulerUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceListSchedulerUpdateData

`func NewPriceListSchedulerUpdateData(type_ interface{}, id interface{}, attributes PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes, ) *PriceListSchedulerUpdateData`

NewPriceListSchedulerUpdateData instantiates a new PriceListSchedulerUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListSchedulerUpdateDataWithDefaults

`func NewPriceListSchedulerUpdateDataWithDefaults() *PriceListSchedulerUpdateData`

NewPriceListSchedulerUpdateDataWithDefaults instantiates a new PriceListSchedulerUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceListSchedulerUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceListSchedulerUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceListSchedulerUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceListSchedulerUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceListSchedulerUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PriceListSchedulerUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceListSchedulerUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceListSchedulerUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PriceListSchedulerUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PriceListSchedulerUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PriceListSchedulerUpdateData) GetAttributes() PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceListSchedulerUpdateData) GetAttributesOk() (*PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceListSchedulerUpdateData) SetAttributes(v PATCHPriceListSchedulersPriceListSchedulerId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceListSchedulerUpdateData) GetRelationships() PriceListSchedulerUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceListSchedulerUpdateData) GetRelationshipsOk() (*PriceListSchedulerUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceListSchedulerUpdateData) SetRelationships(v PriceListSchedulerUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceListSchedulerUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


