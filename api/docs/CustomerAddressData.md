# CustomerAddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customer_addresses"]
**Attributes** | [**GETCustomerAddresses200ResponseDataInnerAttributes**](GETCustomerAddresses200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerAddressDataRelationships**](CustomerAddressDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerAddressData

`func NewCustomerAddressData(type_ string, attributes GETCustomerAddresses200ResponseDataInnerAttributes, ) *CustomerAddressData`

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

`func (o *CustomerAddressData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddressData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddressData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CustomerAddressData) GetAttributes() GETCustomerAddresses200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerAddressData) GetAttributesOk() (*GETCustomerAddresses200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerAddressData) SetAttributes(v GETCustomerAddresses200ResponseDataInnerAttributes)`

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


