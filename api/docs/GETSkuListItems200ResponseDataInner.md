# GETSkuListItems200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "sku_list_items"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETSkuListItems200ResponseDataInnerAttributes**](GETSkuListItems200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETSkuListItems200ResponseDataInnerRelationships**](GETSkuListItems200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETSkuListItems200ResponseDataInner

`func NewGETSkuListItems200ResponseDataInner() *GETSkuListItems200ResponseDataInner`

NewGETSkuListItems200ResponseDataInner instantiates a new GETSkuListItems200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkuListItems200ResponseDataInnerWithDefaults

`func NewGETSkuListItems200ResponseDataInnerWithDefaults() *GETSkuListItems200ResponseDataInner`

NewGETSkuListItems200ResponseDataInnerWithDefaults instantiates a new GETSkuListItems200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETSkuListItems200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETSkuListItems200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETSkuListItems200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETSkuListItems200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETSkuListItems200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETSkuListItems200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETSkuListItems200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETSkuListItems200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETSkuListItems200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETSkuListItems200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETSkuListItems200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETSkuListItems200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETSkuListItems200ResponseDataInner) GetAttributes() GETSkuListItems200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETSkuListItems200ResponseDataInner) GetAttributesOk() (*GETSkuListItems200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETSkuListItems200ResponseDataInner) SetAttributes(v GETSkuListItems200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETSkuListItems200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETSkuListItems200ResponseDataInner) GetRelationships() GETSkuListItems200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETSkuListItems200ResponseDataInner) GetRelationshipsOk() (*GETSkuListItems200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETSkuListItems200ResponseDataInner) SetRelationships(v GETSkuListItems200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETSkuListItems200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


