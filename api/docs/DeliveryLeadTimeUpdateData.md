# DeliveryLeadTimeUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes**](PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**DeliveryLeadTimeUpdateDataRelationships**](DeliveryLeadTimeUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewDeliveryLeadTimeUpdateData

`func NewDeliveryLeadTimeUpdateData(type_ interface{}, id interface{}, attributes PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, ) *DeliveryLeadTimeUpdateData`

NewDeliveryLeadTimeUpdateData instantiates a new DeliveryLeadTimeUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLeadTimeUpdateDataWithDefaults

`func NewDeliveryLeadTimeUpdateDataWithDefaults() *DeliveryLeadTimeUpdateData`

NewDeliveryLeadTimeUpdateDataWithDefaults instantiates a new DeliveryLeadTimeUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeliveryLeadTimeUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeliveryLeadTimeUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeliveryLeadTimeUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *DeliveryLeadTimeUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeliveryLeadTimeUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *DeliveryLeadTimeUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryLeadTimeUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryLeadTimeUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *DeliveryLeadTimeUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeliveryLeadTimeUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *DeliveryLeadTimeUpdateData) GetAttributes() PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeliveryLeadTimeUpdateData) GetAttributesOk() (*PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeliveryLeadTimeUpdateData) SetAttributes(v PATCHDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *DeliveryLeadTimeUpdateData) GetRelationships() DeliveryLeadTimeUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *DeliveryLeadTimeUpdateData) GetRelationshipsOk() (*DeliveryLeadTimeUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *DeliveryLeadTimeUpdateData) SetRelationships(v DeliveryLeadTimeUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *DeliveryLeadTimeUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


