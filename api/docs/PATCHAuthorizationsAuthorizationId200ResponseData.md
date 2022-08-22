# PATCHAuthorizationsAuthorizationId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "authorizations"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHAuthorizationsAuthorizationId200ResponseDataAttributes**](PATCHAuthorizationsAuthorizationId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPATCHAuthorizationsAuthorizationId200ResponseData

`func NewPATCHAuthorizationsAuthorizationId200ResponseData() *PATCHAuthorizationsAuthorizationId200ResponseData`

NewPATCHAuthorizationsAuthorizationId200ResponseData instantiates a new PATCHAuthorizationsAuthorizationId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAuthorizationsAuthorizationId200ResponseDataWithDefaults

`func NewPATCHAuthorizationsAuthorizationId200ResponseDataWithDefaults() *PATCHAuthorizationsAuthorizationId200ResponseData`

NewPATCHAuthorizationsAuthorizationId200ResponseDataWithDefaults instantiates a new PATCHAuthorizationsAuthorizationId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetAttributes() PATCHAuthorizationsAuthorizationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetAttributesOk() (*PATCHAuthorizationsAuthorizationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) SetAttributes(v PATCHAuthorizationsAuthorizationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHAuthorizationsAuthorizationId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


