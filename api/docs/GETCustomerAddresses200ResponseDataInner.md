# GETCustomerAddresses200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "customer_addresses"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETCustomerAddresses200ResponseDataInnerAttributes**](GETCustomerAddresses200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETCustomerAddresses200ResponseDataInnerRelationships**](GETCustomerAddresses200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETCustomerAddresses200ResponseDataInner

`func NewGETCustomerAddresses200ResponseDataInner() *GETCustomerAddresses200ResponseDataInner`

NewGETCustomerAddresses200ResponseDataInner instantiates a new GETCustomerAddresses200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerAddresses200ResponseDataInnerWithDefaults

`func NewGETCustomerAddresses200ResponseDataInnerWithDefaults() *GETCustomerAddresses200ResponseDataInner`

NewGETCustomerAddresses200ResponseDataInnerWithDefaults instantiates a new GETCustomerAddresses200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETCustomerAddresses200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCustomerAddresses200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCustomerAddresses200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCustomerAddresses200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETCustomerAddresses200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETCustomerAddresses200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETCustomerAddresses200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETCustomerAddresses200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETCustomerAddresses200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETCustomerAddresses200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETCustomerAddresses200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETCustomerAddresses200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETCustomerAddresses200ResponseDataInner) GetAttributes() GETCustomerAddresses200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETCustomerAddresses200ResponseDataInner) GetAttributesOk() (*GETCustomerAddresses200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETCustomerAddresses200ResponseDataInner) SetAttributes(v GETCustomerAddresses200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETCustomerAddresses200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETCustomerAddresses200ResponseDataInner) GetRelationships() GETCustomerAddresses200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETCustomerAddresses200ResponseDataInner) GetRelationshipsOk() (*GETCustomerAddresses200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETCustomerAddresses200ResponseDataInner) SetRelationships(v GETCustomerAddresses200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETCustomerAddresses200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


