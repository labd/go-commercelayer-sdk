# ShipmentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAdyenPayments201ResponseDataAttributes**](POSTAdyenPayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShipmentCreateDataRelationships**](ShipmentCreateDataRelationships.md) |  | [optional] 

## Methods

### NewShipmentCreateData

`func NewShipmentCreateData(type_ interface{}, attributes POSTAdyenPayments201ResponseDataAttributes, ) *ShipmentCreateData`

NewShipmentCreateData instantiates a new ShipmentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentCreateDataWithDefaults

`func NewShipmentCreateDataWithDefaults() *ShipmentCreateData`

NewShipmentCreateDataWithDefaults instantiates a new ShipmentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShipmentCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipmentCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipmentCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShipmentCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShipmentCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ShipmentCreateData) GetAttributes() POSTAdyenPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShipmentCreateData) GetAttributesOk() (*POSTAdyenPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShipmentCreateData) SetAttributes(v POSTAdyenPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShipmentCreateData) GetRelationships() ShipmentCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShipmentCreateData) GetRelationshipsOk() (*ShipmentCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShipmentCreateData) SetRelationships(v ShipmentCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShipmentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


