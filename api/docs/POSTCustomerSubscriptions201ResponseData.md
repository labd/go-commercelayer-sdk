# POSTCustomerSubscriptions201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTCustomerSubscriptions201ResponseDataAttributes**](POSTCustomerSubscriptions201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETCustomerPasswordResets200ResponseDataInnerRelationships**](GETCustomerPasswordResets200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTCustomerSubscriptions201ResponseData

`func NewPOSTCustomerSubscriptions201ResponseData() *POSTCustomerSubscriptions201ResponseData`

NewPOSTCustomerSubscriptions201ResponseData instantiates a new POSTCustomerSubscriptions201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomerSubscriptions201ResponseDataWithDefaults

`func NewPOSTCustomerSubscriptions201ResponseDataWithDefaults() *POSTCustomerSubscriptions201ResponseData`

NewPOSTCustomerSubscriptions201ResponseDataWithDefaults instantiates a new POSTCustomerSubscriptions201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTCustomerSubscriptions201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTCustomerSubscriptions201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTCustomerSubscriptions201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTCustomerSubscriptions201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTCustomerSubscriptions201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCustomerSubscriptions201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCustomerSubscriptions201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTCustomerSubscriptions201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTCustomerSubscriptions201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTCustomerSubscriptions201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTCustomerSubscriptions201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTCustomerSubscriptions201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTCustomerSubscriptions201ResponseData) GetAttributes() POSTCustomerSubscriptions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCustomerSubscriptions201ResponseData) GetAttributesOk() (*POSTCustomerSubscriptions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCustomerSubscriptions201ResponseData) SetAttributes(v POSTCustomerSubscriptions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTCustomerSubscriptions201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTCustomerSubscriptions201ResponseData) GetRelationships() GETCustomerPasswordResets200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCustomerSubscriptions201ResponseData) GetRelationshipsOk() (*GETCustomerPasswordResets200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCustomerSubscriptions201ResponseData) SetRelationships(v GETCustomerPasswordResets200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCustomerSubscriptions201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


