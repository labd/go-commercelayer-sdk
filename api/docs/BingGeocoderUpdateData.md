# BingGeocoderUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "bing_geocoders"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes**](PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBingGeocoderUpdateData

`func NewBingGeocoderUpdateData(type_ string, id string, attributes PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes, ) *BingGeocoderUpdateData`

NewBingGeocoderUpdateData instantiates a new BingGeocoderUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBingGeocoderUpdateDataWithDefaults

`func NewBingGeocoderUpdateDataWithDefaults() *BingGeocoderUpdateData`

NewBingGeocoderUpdateDataWithDefaults instantiates a new BingGeocoderUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BingGeocoderUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BingGeocoderUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BingGeocoderUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BingGeocoderUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BingGeocoderUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BingGeocoderUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *BingGeocoderUpdateData) GetAttributes() PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BingGeocoderUpdateData) GetAttributesOk() (*PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BingGeocoderUpdateData) SetAttributes(v PATCHBingGeocodersBingGeocoderId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BingGeocoderUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BingGeocoderUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BingGeocoderUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BingGeocoderUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


