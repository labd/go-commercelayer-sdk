# GETOrderValidationRules200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETBillingInfoValidationRules200ResponseDataInnerAttributes**](GETBillingInfoValidationRules200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETBillingInfoValidationRules200ResponseDataInnerRelationships**](GETBillingInfoValidationRules200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETOrderValidationRules200ResponseDataInner

`func NewGETOrderValidationRules200ResponseDataInner() *GETOrderValidationRules200ResponseDataInner`

NewGETOrderValidationRules200ResponseDataInner instantiates a new GETOrderValidationRules200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderValidationRules200ResponseDataInnerWithDefaults

`func NewGETOrderValidationRules200ResponseDataInnerWithDefaults() *GETOrderValidationRules200ResponseDataInner`

NewGETOrderValidationRules200ResponseDataInnerWithDefaults instantiates a new GETOrderValidationRules200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETOrderValidationRules200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETOrderValidationRules200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETOrderValidationRules200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETOrderValidationRules200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETOrderValidationRules200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETOrderValidationRules200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETOrderValidationRules200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETOrderValidationRules200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETOrderValidationRules200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETOrderValidationRules200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETOrderValidationRules200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETOrderValidationRules200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETOrderValidationRules200ResponseDataInner) GetAttributes() GETBillingInfoValidationRules200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETOrderValidationRules200ResponseDataInner) GetAttributesOk() (*GETBillingInfoValidationRules200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETOrderValidationRules200ResponseDataInner) SetAttributes(v GETBillingInfoValidationRules200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETOrderValidationRules200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETOrderValidationRules200ResponseDataInner) GetRelationships() GETBillingInfoValidationRules200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETOrderValidationRules200ResponseDataInner) GetRelationshipsOk() (*GETBillingInfoValidationRules200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETOrderValidationRules200ResponseDataInner) SetRelationships(v GETBillingInfoValidationRules200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETOrderValidationRules200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


