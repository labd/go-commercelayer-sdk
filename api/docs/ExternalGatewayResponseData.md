# ExternalGatewayResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] 
**Links** | Pointer to [**AddressResponseDataLinks**](AddressResponseDataLinks.md) |  | [optional] 
**Attributes** | Pointer to [**Attributes**](Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**ExternalGatewayResponseDataRelationships**](ExternalGatewayResponseDataRelationships.md) |  | [optional] 

## Methods

### NewExternalGatewayResponseData

`func NewExternalGatewayResponseData() *ExternalGatewayResponseData`

NewExternalGatewayResponseData instantiates a new ExternalGatewayResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGatewayResponseDataWithDefaults

`func NewExternalGatewayResponseDataWithDefaults() *ExternalGatewayResponseData`

NewExternalGatewayResponseDataWithDefaults instantiates a new ExternalGatewayResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalGatewayResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalGatewayResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalGatewayResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalGatewayResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ExternalGatewayResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalGatewayResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalGatewayResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalGatewayResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ExternalGatewayResponseData) GetLinks() AddressResponseDataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalGatewayResponseData) GetLinksOk() (*AddressResponseDataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalGatewayResponseData) SetLinks(v AddressResponseDataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExternalGatewayResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *ExternalGatewayResponseData) GetAttributes() Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalGatewayResponseData) GetAttributesOk() (*Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalGatewayResponseData) SetAttributes(v Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ExternalGatewayResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ExternalGatewayResponseData) GetRelationships() ExternalGatewayResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalGatewayResponseData) GetRelationshipsOk() (*ExternalGatewayResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalGatewayResponseData) SetRelationships(v ExternalGatewayResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalGatewayResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


