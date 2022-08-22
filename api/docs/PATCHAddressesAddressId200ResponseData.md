# PATCHAddressesAddressId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "addresses"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHAddressesAddressId200ResponseDataAttributes**](PATCHAddressesAddressId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETAddresses200ResponseDataInnerRelationships**](GETAddresses200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPATCHAddressesAddressId200ResponseData

`func NewPATCHAddressesAddressId200ResponseData() *PATCHAddressesAddressId200ResponseData`

NewPATCHAddressesAddressId200ResponseData instantiates a new PATCHAddressesAddressId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAddressesAddressId200ResponseDataWithDefaults

`func NewPATCHAddressesAddressId200ResponseDataWithDefaults() *PATCHAddressesAddressId200ResponseData`

NewPATCHAddressesAddressId200ResponseDataWithDefaults instantiates a new PATCHAddressesAddressId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHAddressesAddressId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHAddressesAddressId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHAddressesAddressId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHAddressesAddressId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHAddressesAddressId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHAddressesAddressId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHAddressesAddressId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHAddressesAddressId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHAddressesAddressId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHAddressesAddressId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHAddressesAddressId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHAddressesAddressId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHAddressesAddressId200ResponseData) GetAttributes() PATCHAddressesAddressId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHAddressesAddressId200ResponseData) GetAttributesOk() (*PATCHAddressesAddressId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHAddressesAddressId200ResponseData) SetAttributes(v PATCHAddressesAddressId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHAddressesAddressId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHAddressesAddressId200ResponseData) GetRelationships() GETAddresses200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHAddressesAddressId200ResponseData) GetRelationshipsOk() (*GETAddresses200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHAddressesAddressId200ResponseData) SetRelationships(v GETAddresses200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHAddressesAddressId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


