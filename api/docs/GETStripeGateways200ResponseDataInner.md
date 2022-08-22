# GETStripeGateways200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "stripe_gateways"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETStripeGateways200ResponseDataInnerAttributes**](GETStripeGateways200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETStripeGateways200ResponseDataInnerRelationships**](GETStripeGateways200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETStripeGateways200ResponseDataInner

`func NewGETStripeGateways200ResponseDataInner() *GETStripeGateways200ResponseDataInner`

NewGETStripeGateways200ResponseDataInner instantiates a new GETStripeGateways200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETStripeGateways200ResponseDataInnerWithDefaults

`func NewGETStripeGateways200ResponseDataInnerWithDefaults() *GETStripeGateways200ResponseDataInner`

NewGETStripeGateways200ResponseDataInnerWithDefaults instantiates a new GETStripeGateways200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETStripeGateways200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETStripeGateways200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETStripeGateways200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETStripeGateways200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETStripeGateways200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETStripeGateways200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETStripeGateways200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETStripeGateways200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETStripeGateways200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETStripeGateways200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETStripeGateways200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETStripeGateways200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETStripeGateways200ResponseDataInner) GetAttributes() GETStripeGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETStripeGateways200ResponseDataInner) GetAttributesOk() (*GETStripeGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETStripeGateways200ResponseDataInner) SetAttributes(v GETStripeGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETStripeGateways200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETStripeGateways200ResponseDataInner) GetRelationships() GETStripeGateways200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETStripeGateways200ResponseDataInner) GetRelationshipsOk() (*GETStripeGateways200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETStripeGateways200ResponseDataInner) SetRelationships(v GETStripeGateways200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETStripeGateways200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


