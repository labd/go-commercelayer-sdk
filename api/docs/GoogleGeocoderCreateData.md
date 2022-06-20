# GoogleGeocoderCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "google_geocoders"]
**Attributes** | [**GoogleGeocoderCreateDataAttributes**](GoogleGeocoderCreateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGoogleGeocoderCreateData

`func NewGoogleGeocoderCreateData(type_ string, attributes GoogleGeocoderCreateDataAttributes, ) *GoogleGeocoderCreateData`

NewGoogleGeocoderCreateData instantiates a new GoogleGeocoderCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleGeocoderCreateDataWithDefaults

`func NewGoogleGeocoderCreateDataWithDefaults() *GoogleGeocoderCreateData`

NewGoogleGeocoderCreateDataWithDefaults instantiates a new GoogleGeocoderCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GoogleGeocoderCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GoogleGeocoderCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GoogleGeocoderCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GoogleGeocoderCreateData) GetAttributes() GoogleGeocoderCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GoogleGeocoderCreateData) GetAttributesOk() (*GoogleGeocoderCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GoogleGeocoderCreateData) SetAttributes(v GoogleGeocoderCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GoogleGeocoderCreateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GoogleGeocoderCreateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GoogleGeocoderCreateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GoogleGeocoderCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


