# ExternalPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETExternalPaymentsExternalPaymentId200ResponseDataAttributes**](GETExternalPaymentsExternalPaymentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalPaymentDataRelationships**](ExternalPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPaymentData

`func NewExternalPaymentData(type_ interface{}, attributes GETExternalPaymentsExternalPaymentId200ResponseDataAttributes, ) *ExternalPaymentData`

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

`func (o *ExternalPaymentData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPaymentData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPaymentData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalPaymentData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalPaymentData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalPaymentData) GetAttributes() GETExternalPaymentsExternalPaymentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPaymentData) GetAttributesOk() (*GETExternalPaymentsExternalPaymentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPaymentData) SetAttributes(v GETExternalPaymentsExternalPaymentId200ResponseDataAttributes)`

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


