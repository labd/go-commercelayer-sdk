# GeocoderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETBingGeocoders200ResponseDataInnerAttributes**](GETBingGeocoders200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**BingGeocoderDataRelationships**](BingGeocoderDataRelationships.md) |  | [optional] 

## Methods

### NewGeocoderData

`func NewGeocoderData(type_ interface{}, attributes GETBingGeocoders200ResponseDataInnerAttributes, ) *GeocoderData`

NewGeocoderData instantiates a new GeocoderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeocoderDataWithDefaults

`func NewGeocoderDataWithDefaults() *GeocoderData`

NewGeocoderDataWithDefaults instantiates a new GeocoderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeocoderData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeocoderData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeocoderData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *GeocoderData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GeocoderData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *GeocoderData) GetAttributes() GETBingGeocoders200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GeocoderData) GetAttributesOk() (*GETBingGeocoders200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GeocoderData) SetAttributes(v GETBingGeocoders200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GeocoderData) GetRelationships() BingGeocoderDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GeocoderData) GetRelationshipsOk() (*BingGeocoderDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GeocoderData) SetRelationships(v BingGeocoderDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GeocoderData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


