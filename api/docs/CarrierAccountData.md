# CarrierAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "carrier_accounts"]
**Attributes** | [**GETCarrierAccounts200ResponseDataInnerAttributes**](GETCarrierAccounts200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**CarrierAccountDataRelationships**](CarrierAccountDataRelationships.md) |  | [optional] 

## Methods

### NewCarrierAccountData

`func NewCarrierAccountData(type_ string, attributes GETCarrierAccounts200ResponseDataInnerAttributes, ) *CarrierAccountData`

NewCarrierAccountData instantiates a new CarrierAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierAccountDataWithDefaults

`func NewCarrierAccountDataWithDefaults() *CarrierAccountData`

NewCarrierAccountDataWithDefaults instantiates a new CarrierAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CarrierAccountData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CarrierAccountData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CarrierAccountData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CarrierAccountData) GetAttributes() GETCarrierAccounts200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CarrierAccountData) GetAttributesOk() (*GETCarrierAccounts200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CarrierAccountData) SetAttributes(v GETCarrierAccounts200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CarrierAccountData) GetRelationships() CarrierAccountDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CarrierAccountData) GetRelationshipsOk() (*CarrierAccountDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CarrierAccountData) SetRelationships(v CarrierAccountDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CarrierAccountData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


