# BingGeocoderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETBingGeocodersBingGeocoderId200ResponseDataAttributes**](GETBingGeocodersBingGeocoderId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BingGeocoderDataRelationships**](BingGeocoderDataRelationships.md) |  | [optional] 

## Methods

### NewBingGeocoderData

`func NewBingGeocoderData(type_ interface{}, attributes GETBingGeocodersBingGeocoderId200ResponseDataAttributes, ) *BingGeocoderData`

NewBingGeocoderData instantiates a new BingGeocoderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBingGeocoderDataWithDefaults

`func NewBingGeocoderDataWithDefaults() *BingGeocoderData`

NewBingGeocoderDataWithDefaults instantiates a new BingGeocoderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BingGeocoderData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BingGeocoderData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BingGeocoderData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BingGeocoderData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BingGeocoderData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *BingGeocoderData) GetAttributes() GETBingGeocodersBingGeocoderId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BingGeocoderData) GetAttributesOk() (*GETBingGeocodersBingGeocoderId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BingGeocoderData) SetAttributes(v GETBingGeocodersBingGeocoderId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BingGeocoderData) GetRelationships() BingGeocoderDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BingGeocoderData) GetRelationshipsOk() (*BingGeocoderDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BingGeocoderData) SetRelationships(v BingGeocoderDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BingGeocoderData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


