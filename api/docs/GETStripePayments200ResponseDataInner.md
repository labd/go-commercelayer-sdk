# GETStripePayments200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETStripePayments200ResponseDataInnerAttributes**](GETStripePayments200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationships**](GETAdyenPayments200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETStripePayments200ResponseDataInner

`func NewGETStripePayments200ResponseDataInner() *GETStripePayments200ResponseDataInner`

NewGETStripePayments200ResponseDataInner instantiates a new GETStripePayments200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStripePayments200ResponseDataInnerWithDefaults

`func NewGETStripePayments200ResponseDataInnerWithDefaults() *GETStripePayments200ResponseDataInner`

NewGETStripePayments200ResponseDataInnerWithDefaults instantiates a new GETStripePayments200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETStripePayments200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETStripePayments200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETStripePayments200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETStripePayments200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETStripePayments200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETStripePayments200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETStripePayments200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETStripePayments200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETStripePayments200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETStripePayments200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETStripePayments200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETStripePayments200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETStripePayments200ResponseDataInner) GetAttributes() GETStripePayments200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETStripePayments200ResponseDataInner) GetAttributesOk() (*GETStripePayments200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETStripePayments200ResponseDataInner) SetAttributes(v GETStripePayments200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETStripePayments200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETStripePayments200ResponseDataInner) GetRelationships() GETAdyenPayments200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETStripePayments200ResponseDataInner) GetRelationshipsOk() (*GETAdyenPayments200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETStripePayments200ResponseDataInner) SetRelationships(v GETAdyenPayments200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETStripePayments200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


