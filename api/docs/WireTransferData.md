# WireTransferData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETWireTransfersWireTransferId200ResponseDataAttributes**](GETWireTransfersWireTransferId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**WireTransferDataRelationships**](WireTransferDataRelationships.md) |  | [optional] 

## Methods

### NewWireTransferData

`func NewWireTransferData(type_ interface{}, attributes GETWireTransfersWireTransferId200ResponseDataAttributes, ) *WireTransferData`

NewWireTransferData instantiates a new WireTransferData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransferDataWithDefaults

`func NewWireTransferDataWithDefaults() *WireTransferData`

NewWireTransferDataWithDefaults instantiates a new WireTransferData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WireTransferData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireTransferData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireTransferData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *WireTransferData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WireTransferData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *WireTransferData) GetAttributes() GETWireTransfersWireTransferId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WireTransferData) GetAttributesOk() (*GETWireTransfersWireTransferId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WireTransferData) SetAttributes(v GETWireTransfersWireTransferId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WireTransferData) GetRelationships() WireTransferDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WireTransferData) GetRelationshipsOk() (*WireTransferDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WireTransferData) SetRelationships(v WireTransferDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WireTransferData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


