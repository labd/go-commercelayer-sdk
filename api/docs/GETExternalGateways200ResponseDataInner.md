# GETExternalGateways200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "external_gateways"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETExternalGateways200ResponseDataInnerAttributes**](GETExternalGateways200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETExternalGateways200ResponseDataInnerRelationships**](GETExternalGateways200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETExternalGateways200ResponseDataInner

`func NewGETExternalGateways200ResponseDataInner() *GETExternalGateways200ResponseDataInner`

NewGETExternalGateways200ResponseDataInner instantiates a new GETExternalGateways200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETExternalGateways200ResponseDataInnerWithDefaults

`func NewGETExternalGateways200ResponseDataInnerWithDefaults() *GETExternalGateways200ResponseDataInner`

NewGETExternalGateways200ResponseDataInnerWithDefaults instantiates a new GETExternalGateways200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETExternalGateways200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETExternalGateways200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETExternalGateways200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETExternalGateways200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETExternalGateways200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETExternalGateways200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETExternalGateways200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETExternalGateways200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETExternalGateways200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETExternalGateways200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETExternalGateways200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETExternalGateways200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETExternalGateways200ResponseDataInner) GetAttributes() GETExternalGateways200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETExternalGateways200ResponseDataInner) GetAttributesOk() (*GETExternalGateways200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETExternalGateways200ResponseDataInner) SetAttributes(v GETExternalGateways200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETExternalGateways200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETExternalGateways200ResponseDataInner) GetRelationships() GETExternalGateways200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETExternalGateways200ResponseDataInner) GetRelationshipsOk() (*GETExternalGateways200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETExternalGateways200ResponseDataInner) SetRelationships(v GETExternalGateways200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETExternalGateways200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


