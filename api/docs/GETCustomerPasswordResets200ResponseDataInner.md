# GETCustomerPasswordResets200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "customer_password_resets"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETCustomerPasswordResets200ResponseDataInnerAttributes**](GETCustomerPasswordResets200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETCustomerPasswordResets200ResponseDataInnerRelationships**](GETCustomerPasswordResets200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETCustomerPasswordResets200ResponseDataInner

`func NewGETCustomerPasswordResets200ResponseDataInner() *GETCustomerPasswordResets200ResponseDataInner`

NewGETCustomerPasswordResets200ResponseDataInner instantiates a new GETCustomerPasswordResets200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCustomerPasswordResets200ResponseDataInnerWithDefaults

`func NewGETCustomerPasswordResets200ResponseDataInnerWithDefaults() *GETCustomerPasswordResets200ResponseDataInner`

NewGETCustomerPasswordResets200ResponseDataInnerWithDefaults instantiates a new GETCustomerPasswordResets200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETCustomerPasswordResets200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETCustomerPasswordResets200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETCustomerPasswordResets200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETCustomerPasswordResets200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETCustomerPasswordResets200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETCustomerPasswordResets200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetAttributes() GETCustomerPasswordResets200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetAttributesOk() (*GETCustomerPasswordResets200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETCustomerPasswordResets200ResponseDataInner) SetAttributes(v GETCustomerPasswordResets200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETCustomerPasswordResets200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetRelationships() GETCustomerPasswordResets200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETCustomerPasswordResets200ResponseDataInner) GetRelationshipsOk() (*GETCustomerPasswordResets200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETCustomerPasswordResets200ResponseDataInner) SetRelationships(v GETCustomerPasswordResets200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETCustomerPasswordResets200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


