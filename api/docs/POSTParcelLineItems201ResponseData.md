# POSTParcelLineItems201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "parcel_line_items"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTParcelLineItems201ResponseDataAttributes**](POSTParcelLineItems201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETParcelLineItems200ResponseDataInnerRelationships**](GETParcelLineItems200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTParcelLineItems201ResponseData

`func NewPOSTParcelLineItems201ResponseData() *POSTParcelLineItems201ResponseData`

NewPOSTParcelLineItems201ResponseData instantiates a new POSTParcelLineItems201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTParcelLineItems201ResponseDataWithDefaults

`func NewPOSTParcelLineItems201ResponseDataWithDefaults() *POSTParcelLineItems201ResponseData`

NewPOSTParcelLineItems201ResponseDataWithDefaults instantiates a new POSTParcelLineItems201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTParcelLineItems201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTParcelLineItems201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTParcelLineItems201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTParcelLineItems201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTParcelLineItems201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTParcelLineItems201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTParcelLineItems201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTParcelLineItems201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTParcelLineItems201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTParcelLineItems201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTParcelLineItems201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTParcelLineItems201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTParcelLineItems201ResponseData) GetAttributes() POSTParcelLineItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTParcelLineItems201ResponseData) GetAttributesOk() (*POSTParcelLineItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTParcelLineItems201ResponseData) SetAttributes(v POSTParcelLineItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTParcelLineItems201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTParcelLineItems201ResponseData) GetRelationships() GETParcelLineItems200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTParcelLineItems201ResponseData) GetRelationshipsOk() (*GETParcelLineItems200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTParcelLineItems201ResponseData) SetRelationships(v GETParcelLineItems200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTParcelLineItems201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


