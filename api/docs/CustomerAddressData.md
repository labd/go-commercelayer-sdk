# CustomerAddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETCustomerAddressesCustomerAddressId200ResponseDataAttributes**](GETCustomerAddressesCustomerAddressId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerAddressDataRelationships**](CustomerAddressDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerAddressData

`func NewCustomerAddressData(type_ interface{}, attributes GETCustomerAddressesCustomerAddressId200ResponseDataAttributes, ) *CustomerAddressData`

NewCustomerAddressData instantiates a new CustomerAddressData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressDataWithDefaults

`func NewCustomerAddressDataWithDefaults() *CustomerAddressData`

NewCustomerAddressDataWithDefaults instantiates a new CustomerAddressData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerAddressData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddressData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddressData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerAddressData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerAddressData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomerAddressData) GetAttributes() GETCustomerAddressesCustomerAddressId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerAddressData) GetAttributesOk() (*GETCustomerAddressesCustomerAddressId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerAddressData) SetAttributes(v GETCustomerAddressesCustomerAddressId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerAddressData) GetRelationships() CustomerAddressDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerAddressData) GetRelationshipsOk() (*CustomerAddressDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerAddressData) SetRelationships(v CustomerAddressDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerAddressData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


