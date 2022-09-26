# POSTImports201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTImports201ResponseDataAttributes**](POSTImports201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETExports200ResponseDataInnerRelationships**](GETExports200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTImports201ResponseData

`func NewPOSTImports201ResponseData() *POSTImports201ResponseData`

NewPOSTImports201ResponseData instantiates a new POSTImports201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTImports201ResponseDataWithDefaults

`func NewPOSTImports201ResponseDataWithDefaults() *POSTImports201ResponseData`

NewPOSTImports201ResponseDataWithDefaults instantiates a new POSTImports201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTImports201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTImports201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTImports201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTImports201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTImports201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTImports201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTImports201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTImports201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTImports201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTImports201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTImports201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTImports201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTImports201ResponseData) GetAttributes() POSTImports201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTImports201ResponseData) GetAttributesOk() (*POSTImports201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTImports201ResponseData) SetAttributes(v POSTImports201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTImports201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTImports201ResponseData) GetRelationships() GETExports200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTImports201ResponseData) GetRelationshipsOk() (*GETExports200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTImports201ResponseData) SetRelationships(v GETExports200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTImports201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


