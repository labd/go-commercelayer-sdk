# POSTManualTaxCalculators201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "manual_tax_calculators"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTManualTaxCalculators201ResponseDataAttributes**](POSTManualTaxCalculators201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETManualTaxCalculators200ResponseDataInnerRelationships**](GETManualTaxCalculators200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTManualTaxCalculators201ResponseData

`func NewPOSTManualTaxCalculators201ResponseData() *POSTManualTaxCalculators201ResponseData`

NewPOSTManualTaxCalculators201ResponseData instantiates a new POSTManualTaxCalculators201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTManualTaxCalculators201ResponseDataWithDefaults

`func NewPOSTManualTaxCalculators201ResponseDataWithDefaults() *POSTManualTaxCalculators201ResponseData`

NewPOSTManualTaxCalculators201ResponseDataWithDefaults instantiates a new POSTManualTaxCalculators201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTManualTaxCalculators201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTManualTaxCalculators201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTManualTaxCalculators201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTManualTaxCalculators201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTManualTaxCalculators201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTManualTaxCalculators201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTManualTaxCalculators201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTManualTaxCalculators201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTManualTaxCalculators201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTManualTaxCalculators201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTManualTaxCalculators201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTManualTaxCalculators201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTManualTaxCalculators201ResponseData) GetAttributes() POSTManualTaxCalculators201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTManualTaxCalculators201ResponseData) GetAttributesOk() (*POSTManualTaxCalculators201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTManualTaxCalculators201ResponseData) SetAttributes(v POSTManualTaxCalculators201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTManualTaxCalculators201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTManualTaxCalculators201ResponseData) GetRelationships() GETManualTaxCalculators200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTManualTaxCalculators201ResponseData) GetRelationshipsOk() (*GETManualTaxCalculators200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTManualTaxCalculators201ResponseData) SetRelationships(v GETManualTaxCalculators200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTManualTaxCalculators201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


