# SatispayPaymentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTSatispayPayments201ResponseDataAttributes**](POSTSatispayPayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentCreateDataRelationships**](AdyenPaymentCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSatispayPaymentCreateData

`func NewSatispayPaymentCreateData(type_ interface{}, attributes POSTSatispayPayments201ResponseDataAttributes, ) *SatispayPaymentCreateData`

NewSatispayPaymentCreateData instantiates a new SatispayPaymentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSatispayPaymentCreateDataWithDefaults

`func NewSatispayPaymentCreateDataWithDefaults() *SatispayPaymentCreateData`

NewSatispayPaymentCreateDataWithDefaults instantiates a new SatispayPaymentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SatispayPaymentCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SatispayPaymentCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SatispayPaymentCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SatispayPaymentCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SatispayPaymentCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *SatispayPaymentCreateData) GetAttributes() POSTSatispayPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SatispayPaymentCreateData) GetAttributesOk() (*POSTSatispayPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SatispayPaymentCreateData) SetAttributes(v POSTSatispayPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SatispayPaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SatispayPaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SatispayPaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SatispayPaymentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


