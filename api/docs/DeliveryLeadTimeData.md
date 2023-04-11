# DeliveryLeadTimeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes**](GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**DeliveryLeadTimeDataRelationships**](DeliveryLeadTimeDataRelationships.md) |  | [optional] 

## Methods

### NewDeliveryLeadTimeData

`func NewDeliveryLeadTimeData(type_ interface{}, attributes GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, ) *DeliveryLeadTimeData`

NewDeliveryLeadTimeData instantiates a new DeliveryLeadTimeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLeadTimeDataWithDefaults

`func NewDeliveryLeadTimeDataWithDefaults() *DeliveryLeadTimeData`

NewDeliveryLeadTimeDataWithDefaults instantiates a new DeliveryLeadTimeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeliveryLeadTimeData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeliveryLeadTimeData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeliveryLeadTimeData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *DeliveryLeadTimeData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DeliveryLeadTimeData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *DeliveryLeadTimeData) GetAttributes() GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeliveryLeadTimeData) GetAttributesOk() (*GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeliveryLeadTimeData) SetAttributes(v GETDeliveryLeadTimesDeliveryLeadTimeId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *DeliveryLeadTimeData) GetRelationships() DeliveryLeadTimeDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *DeliveryLeadTimeData) GetRelationshipsOk() (*DeliveryLeadTimeDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *DeliveryLeadTimeData) SetRelationships(v DeliveryLeadTimeDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *DeliveryLeadTimeData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


