# PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "klarna_payments"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes**](PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETAdyenPayments200ResponseDataInnerRelationships**](GETAdyenPayments200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseData

`func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseData() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData`

NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseData instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataWithDefaults

`func NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataWithDefaults() *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData`

NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataWithDefaults instantiates a new PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetAttributes() PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetAttributesOk() (*PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) SetAttributes(v PATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetRelationships() GETAdyenPayments200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) GetRelationshipsOk() (*GETAdyenPayments200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) SetRelationships(v GETAdyenPayments200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHKlarnaPaymentsKlarnaPaymentId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


