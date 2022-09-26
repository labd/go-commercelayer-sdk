# DeliveryLeadTimeCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | 
**Attributes** | [**DeliveryLeadTimeCreateDataAttributes**](DeliveryLeadTimeCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**DeliveryLeadTimeCreateDataRelationships**](DeliveryLeadTimeCreateDataRelationships.md) |  | [optional] 

## Methods

### NewDeliveryLeadTimeCreateData

`func NewDeliveryLeadTimeCreateData(type_ string, attributes DeliveryLeadTimeCreateDataAttributes, ) *DeliveryLeadTimeCreateData`

NewDeliveryLeadTimeCreateData instantiates a new DeliveryLeadTimeCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLeadTimeCreateDataWithDefaults

`func NewDeliveryLeadTimeCreateDataWithDefaults() *DeliveryLeadTimeCreateData`

NewDeliveryLeadTimeCreateDataWithDefaults instantiates a new DeliveryLeadTimeCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeliveryLeadTimeCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeliveryLeadTimeCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeliveryLeadTimeCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *DeliveryLeadTimeCreateData) GetAttributes() DeliveryLeadTimeCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DeliveryLeadTimeCreateData) GetAttributesOk() (*DeliveryLeadTimeCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DeliveryLeadTimeCreateData) SetAttributes(v DeliveryLeadTimeCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *DeliveryLeadTimeCreateData) GetRelationships() DeliveryLeadTimeCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *DeliveryLeadTimeCreateData) GetRelationshipsOk() (*DeliveryLeadTimeCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *DeliveryLeadTimeCreateData) SetRelationships(v DeliveryLeadTimeCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *DeliveryLeadTimeCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


