# POSTExternalTaxCalculators201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "external_tax_calculators"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTExternalTaxCalculators201ResponseDataAttributes**](POSTExternalTaxCalculators201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTAvalaraAccounts201ResponseDataRelationships**](POSTAvalaraAccounts201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTExternalTaxCalculators201ResponseData

`func NewPOSTExternalTaxCalculators201ResponseData() *POSTExternalTaxCalculators201ResponseData`

NewPOSTExternalTaxCalculators201ResponseData instantiates a new POSTExternalTaxCalculators201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalTaxCalculators201ResponseDataWithDefaults

`func NewPOSTExternalTaxCalculators201ResponseDataWithDefaults() *POSTExternalTaxCalculators201ResponseData`

NewPOSTExternalTaxCalculators201ResponseDataWithDefaults instantiates a new POSTExternalTaxCalculators201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTExternalTaxCalculators201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTExternalTaxCalculators201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTExternalTaxCalculators201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTExternalTaxCalculators201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTExternalTaxCalculators201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTExternalTaxCalculators201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTExternalTaxCalculators201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTExternalTaxCalculators201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTExternalTaxCalculators201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTExternalTaxCalculators201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTExternalTaxCalculators201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTExternalTaxCalculators201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTExternalTaxCalculators201ResponseData) GetAttributes() POSTExternalTaxCalculators201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTExternalTaxCalculators201ResponseData) GetAttributesOk() (*POSTExternalTaxCalculators201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTExternalTaxCalculators201ResponseData) SetAttributes(v POSTExternalTaxCalculators201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTExternalTaxCalculators201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTExternalTaxCalculators201ResponseData) GetRelationships() POSTAvalaraAccounts201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTExternalTaxCalculators201ResponseData) GetRelationshipsOk() (*POSTAvalaraAccounts201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTExternalTaxCalculators201ResponseData) SetRelationships(v POSTAvalaraAccounts201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTExternalTaxCalculators201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


