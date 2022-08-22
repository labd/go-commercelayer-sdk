# ParcelUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "parcels"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHParcelsParcelId200ResponseDataAttributes**](PATCHParcelsParcelId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHParcelsParcelId200ResponseDataRelationships**](PATCHParcelsParcelId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewParcelUpdateData

`func NewParcelUpdateData(type_ string, id string, attributes PATCHParcelsParcelId200ResponseDataAttributes, ) *ParcelUpdateData`

NewParcelUpdateData instantiates a new ParcelUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelUpdateDataWithDefaults

`func NewParcelUpdateDataWithDefaults() *ParcelUpdateData`

NewParcelUpdateDataWithDefaults instantiates a new ParcelUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParcelUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParcelUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParcelUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ParcelUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParcelUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParcelUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ParcelUpdateData) GetAttributes() PATCHParcelsParcelId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParcelUpdateData) GetAttributesOk() (*PATCHParcelsParcelId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParcelUpdateData) SetAttributes(v PATCHParcelsParcelId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParcelUpdateData) GetRelationships() PATCHParcelsParcelId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParcelUpdateData) GetRelationshipsOk() (*PATCHParcelsParcelId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParcelUpdateData) SetRelationships(v PATCHParcelsParcelId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ParcelUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


