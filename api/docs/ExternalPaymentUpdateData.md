# ExternalPaymentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "external_payments"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes**](PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PATCHAdyenPaymentsAdyenPaymentId200ResponseDataRelationships**](PATCHAdyenPaymentsAdyenPaymentId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPaymentUpdateData

`func NewExternalPaymentUpdateData(type_ string, id string, attributes PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes, ) *ExternalPaymentUpdateData`

NewExternalPaymentUpdateData instantiates a new ExternalPaymentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentUpdateDataWithDefaults

`func NewExternalPaymentUpdateDataWithDefaults() *ExternalPaymentUpdateData`

NewExternalPaymentUpdateDataWithDefaults instantiates a new ExternalPaymentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPaymentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPaymentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPaymentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ExternalPaymentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalPaymentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalPaymentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ExternalPaymentUpdateData) GetAttributes() PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPaymentUpdateData) GetAttributesOk() (*PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPaymentUpdateData) SetAttributes(v PATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPaymentUpdateData) GetRelationships() PATCHAdyenPaymentsAdyenPaymentId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPaymentUpdateData) GetRelationshipsOk() (*PATCHAdyenPaymentsAdyenPaymentId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPaymentUpdateData) SetRelationships(v PATCHAdyenPaymentsAdyenPaymentId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPaymentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


