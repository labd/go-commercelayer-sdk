# ExternalPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "external_payments"]
**Attributes** | [**ExternalPaymentDataAttributes**](ExternalPaymentDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPaymentDataRelationships**](ExternalPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPaymentData

`func NewExternalPaymentData(type_ string, attributes ExternalPaymentDataAttributes, ) *ExternalPaymentData`

NewExternalPaymentData instantiates a new ExternalPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentDataWithDefaults

`func NewExternalPaymentDataWithDefaults() *ExternalPaymentData`

NewExternalPaymentDataWithDefaults instantiates a new ExternalPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPaymentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPaymentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPaymentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ExternalPaymentData) GetAttributes() ExternalPaymentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPaymentData) GetAttributesOk() (*ExternalPaymentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPaymentData) SetAttributes(v ExternalPaymentDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPaymentData) GetRelationships() ExternalPaymentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPaymentData) GetRelationshipsOk() (*ExternalPaymentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPaymentData) SetRelationships(v ExternalPaymentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPaymentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


