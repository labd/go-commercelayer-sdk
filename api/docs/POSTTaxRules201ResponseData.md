# POSTTaxRules201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTTaxRules201ResponseDataAttributes**](POSTTaxRules201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETTaxRules200ResponseDataInnerRelationships**](GETTaxRules200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTTaxRules201ResponseData

`func NewPOSTTaxRules201ResponseData() *POSTTaxRules201ResponseData`

NewPOSTTaxRules201ResponseData instantiates a new POSTTaxRules201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTTaxRules201ResponseDataWithDefaults

`func NewPOSTTaxRules201ResponseDataWithDefaults() *POSTTaxRules201ResponseData`

NewPOSTTaxRules201ResponseDataWithDefaults instantiates a new POSTTaxRules201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTTaxRules201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTTaxRules201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTTaxRules201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTTaxRules201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTTaxRules201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTTaxRules201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTTaxRules201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTTaxRules201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTTaxRules201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTTaxRules201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTTaxRules201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTTaxRules201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTTaxRules201ResponseData) GetAttributes() POSTTaxRules201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTTaxRules201ResponseData) GetAttributesOk() (*POSTTaxRules201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTTaxRules201ResponseData) SetAttributes(v POSTTaxRules201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTTaxRules201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTTaxRules201ResponseData) GetRelationships() GETTaxRules200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTTaxRules201ResponseData) GetRelationshipsOk() (*GETTaxRules200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTTaxRules201ResponseData) SetRelationships(v GETTaxRules200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTTaxRules201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


