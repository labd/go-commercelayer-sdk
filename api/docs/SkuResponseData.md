# SkuResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**SkuResponseDataRelationships**](SkuResponseDataRelationships.md) |  | [optional] 

## Methods

### NewSkuResponseData

`func NewSkuResponseData() *SkuResponseData`

NewSkuResponseData instantiates a new SkuResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuResponseDataWithDefaults

`func NewSkuResponseDataWithDefaults() *SkuResponseData`

NewSkuResponseDataWithDefaults instantiates a new SkuResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SkuResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SkuResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SkuResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SkuResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *SkuResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SkuResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SkuResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SkuResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *SkuResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *SkuResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *SkuResponseData) GetRelationships() SkuResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuResponseData) GetRelationshipsOk() (*SkuResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuResponseData) SetRelationships(v SkuResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


