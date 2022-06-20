# ShipmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "shipments"]
**Attributes** | [**ShipmentDataAttributes**](ShipmentDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShipmentDataRelationships**](ShipmentDataRelationships.md) |  | [optional] 

## Methods

### NewShipmentData

`func NewShipmentData(type_ string, attributes ShipmentDataAttributes, ) *ShipmentData`

NewShipmentData instantiates a new ShipmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDataWithDefaults

`func NewShipmentDataWithDefaults() *ShipmentData`

NewShipmentDataWithDefaults instantiates a new ShipmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShipmentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipmentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipmentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ShipmentData) GetAttributes() ShipmentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShipmentData) GetAttributesOk() (*ShipmentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShipmentData) SetAttributes(v ShipmentDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShipmentData) GetRelationships() ShipmentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShipmentData) GetRelationshipsOk() (*ShipmentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShipmentData) SetRelationships(v ShipmentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShipmentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


