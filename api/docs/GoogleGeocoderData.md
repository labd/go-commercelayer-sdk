# GoogleGeocoderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "google_geocoders"]
**Attributes** | [**BingGeocoderDataAttributes**](BingGeocoderDataAttributes.md) |  | 
**Relationships** | Pointer to [**BingGeocoderDataRelationships**](BingGeocoderDataRelationships.md) |  | [optional] 

## Methods

### NewGoogleGeocoderData

`func NewGoogleGeocoderData(type_ string, attributes BingGeocoderDataAttributes, ) *GoogleGeocoderData`

NewGoogleGeocoderData instantiates a new GoogleGeocoderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleGeocoderDataWithDefaults

`func NewGoogleGeocoderDataWithDefaults() *GoogleGeocoderData`

NewGoogleGeocoderDataWithDefaults instantiates a new GoogleGeocoderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GoogleGeocoderData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleGeocoderData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleGeocoderData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GoogleGeocoderData) GetAttributes() BingGeocoderDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GoogleGeocoderData) GetAttributesOk() (*BingGeocoderDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GoogleGeocoderData) SetAttributes(v BingGeocoderDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GoogleGeocoderData) GetRelationships() BingGeocoderDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GoogleGeocoderData) GetRelationshipsOk() (*BingGeocoderDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GoogleGeocoderData) SetRelationships(v BingGeocoderDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GoogleGeocoderData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


