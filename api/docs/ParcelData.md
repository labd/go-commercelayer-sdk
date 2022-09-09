# ParcelData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "parcels"]
**Attributes** | [**GETParcels200ResponseDataInnerAttributes**](GETParcels200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**ParcelDataRelationships**](ParcelDataRelationships.md) |  | [optional] 

## Methods

### NewParcelData

`func NewParcelData(type_ string, attributes GETParcels200ResponseDataInnerAttributes, ) *ParcelData`

NewParcelData instantiates a new ParcelData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelDataWithDefaults

`func NewParcelDataWithDefaults() *ParcelData`

NewParcelDataWithDefaults instantiates a new ParcelData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParcelData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParcelData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParcelData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ParcelData) GetAttributes() GETParcels200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParcelData) GetAttributesOk() (*GETParcels200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParcelData) SetAttributes(v GETParcels200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParcelData) GetRelationships() ParcelDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParcelData) GetRelationshipsOk() (*ParcelDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParcelData) SetRelationships(v ParcelDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ParcelData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


