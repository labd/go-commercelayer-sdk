# ParcelLineItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETParcelLineItemsParcelLineItemId200ResponseDataAttributes**](GETParcelLineItemsParcelLineItemId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ParcelLineItemDataRelationships**](ParcelLineItemDataRelationships.md) |  | [optional] 

## Methods

### NewParcelLineItemData

`func NewParcelLineItemData(type_ interface{}, attributes GETParcelLineItemsParcelLineItemId200ResponseDataAttributes, ) *ParcelLineItemData`

NewParcelLineItemData instantiates a new ParcelLineItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelLineItemDataWithDefaults

`func NewParcelLineItemDataWithDefaults() *ParcelLineItemData`

NewParcelLineItemDataWithDefaults instantiates a new ParcelLineItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParcelLineItemData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParcelLineItemData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParcelLineItemData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ParcelLineItemData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ParcelLineItemData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ParcelLineItemData) GetAttributes() GETParcelLineItemsParcelLineItemId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParcelLineItemData) GetAttributesOk() (*GETParcelLineItemsParcelLineItemId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParcelLineItemData) SetAttributes(v GETParcelLineItemsParcelLineItemId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParcelLineItemData) GetRelationships() ParcelLineItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParcelLineItemData) GetRelationshipsOk() (*ParcelLineItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParcelLineItemData) SetRelationships(v ParcelLineItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ParcelLineItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


