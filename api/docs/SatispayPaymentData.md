# SatispayPaymentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETSatispayPayments200ResponseDataInnerAttributes**](GETSatispayPayments200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentDataRelationships**](AdyenPaymentDataRelationships.md) |  | [optional] 

## Methods

### NewSatispayPaymentData

`func NewSatispayPaymentData(type_ interface{}, attributes GETSatispayPayments200ResponseDataInnerAttributes, ) *SatispayPaymentData`

NewSatispayPaymentData instantiates a new SatispayPaymentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayPaymentDataWithDefaults

`func NewSatispayPaymentDataWithDefaults() *SatispayPaymentData`

NewSatispayPaymentDataWithDefaults instantiates a new SatispayPaymentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SatispayPaymentData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SatispayPaymentData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SatispayPaymentData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SatispayPaymentData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SatispayPaymentData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SatispayPaymentData) GetAttributes() GETSatispayPayments200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SatispayPaymentData) GetAttributesOk() (*GETSatispayPayments200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SatispayPaymentData) SetAttributes(v GETSatispayPayments200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SatispayPaymentData) GetRelationships() AdyenPaymentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SatispayPaymentData) GetRelationshipsOk() (*AdyenPaymentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SatispayPaymentData) SetRelationships(v AdyenPaymentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SatispayPaymentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


