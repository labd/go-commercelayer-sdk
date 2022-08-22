# GETCustomerGroups200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "customer_groups"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETCustomerGroups200ResponseDataInnerAttributes**](GETCustomerGroups200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETCustomerGroups200ResponseDataInnerRelationships**](GETCustomerGroups200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETCustomerGroups200ResponseDataInner

`func NewGETCustomerGroups200ResponseDataInner() *GETCustomerGroups200ResponseDataInner`

NewGETCustomerGroups200ResponseDataInner instantiates a new GETCustomerGroups200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerGroups200ResponseDataInnerWithDefaults

`func NewGETCustomerGroups200ResponseDataInnerWithDefaults() *GETCustomerGroups200ResponseDataInner`

NewGETCustomerGroups200ResponseDataInnerWithDefaults instantiates a new GETCustomerGroups200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETCustomerGroups200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCustomerGroups200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCustomerGroups200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCustomerGroups200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETCustomerGroups200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETCustomerGroups200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETCustomerGroups200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETCustomerGroups200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETCustomerGroups200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETCustomerGroups200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETCustomerGroups200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETCustomerGroups200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETCustomerGroups200ResponseDataInner) GetAttributes() GETCustomerGroups200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETCustomerGroups200ResponseDataInner) GetAttributesOk() (*GETCustomerGroups200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETCustomerGroups200ResponseDataInner) SetAttributes(v GETCustomerGroups200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETCustomerGroups200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETCustomerGroups200ResponseDataInner) GetRelationships() GETCustomerGroups200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETCustomerGroups200ResponseDataInner) GetRelationshipsOk() (*GETCustomerGroups200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETCustomerGroups200ResponseDataInner) SetRelationships(v GETCustomerGroups200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETCustomerGroups200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


