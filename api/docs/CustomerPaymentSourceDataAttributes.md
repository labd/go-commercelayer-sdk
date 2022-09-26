# CustomerPaymentSourceDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Returns the associated payment source&#39;s name | [optional] 
**CustomerToken** | Pointer to **string** | Returns the customer gateway token stored in the gateway | [optional] 
**PaymentSourceToken** | Pointer to **string** | Returns the payment source token stored in the gateway | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewCustomerPaymentSourceDataAttributes

`func NewCustomerPaymentSourceDataAttributes() *CustomerPaymentSourceDataAttributes`

NewCustomerPaymentSourceDataAttributes instantiates a new CustomerPaymentSourceDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceDataAttributesWithDefaults

`func NewCustomerPaymentSourceDataAttributesWithDefaults() *CustomerPaymentSourceDataAttributes`

NewCustomerPaymentSourceDataAttributesWithDefaults instantiates a new CustomerPaymentSourceDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomerPaymentSourceDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerPaymentSourceDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerPaymentSourceDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerPaymentSourceDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerToken

`func (o *CustomerPaymentSourceDataAttributes) GetCustomerToken() string`

GetCustomerToken returns the CustomerToken field if non-nil, zero value otherwise.

### GetCustomerTokenOk

`func (o *CustomerPaymentSourceDataAttributes) GetCustomerTokenOk() (*string, bool)`

GetCustomerTokenOk returns a tuple with the CustomerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerToken

`func (o *CustomerPaymentSourceDataAttributes) SetCustomerToken(v string)`

SetCustomerToken sets CustomerToken field to given value.

### HasCustomerToken

`func (o *CustomerPaymentSourceDataAttributes) HasCustomerToken() bool`

HasCustomerToken returns a boolean if a field has been set.

### GetPaymentSourceToken

`func (o *CustomerPaymentSourceDataAttributes) GetPaymentSourceToken() string`

GetPaymentSourceToken returns the PaymentSourceToken field if non-nil, zero value otherwise.

### GetPaymentSourceTokenOk

`func (o *CustomerPaymentSourceDataAttributes) GetPaymentSourceTokenOk() (*string, bool)`

GetPaymentSourceTokenOk returns a tuple with the PaymentSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceToken

`func (o *CustomerPaymentSourceDataAttributes) SetPaymentSourceToken(v string)`

SetPaymentSourceToken sets PaymentSourceToken field to given value.

### HasPaymentSourceToken

`func (o *CustomerPaymentSourceDataAttributes) HasPaymentSourceToken() bool`

HasPaymentSourceToken returns a boolean if a field has been set.

### GetId

`func (o *CustomerPaymentSourceDataAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPaymentSourceDataAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPaymentSourceDataAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerPaymentSourceDataAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerPaymentSourceDataAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerPaymentSourceDataAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerPaymentSourceDataAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerPaymentSourceDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerPaymentSourceDataAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerPaymentSourceDataAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerPaymentSourceDataAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerPaymentSourceDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *CustomerPaymentSourceDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CustomerPaymentSourceDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CustomerPaymentSourceDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CustomerPaymentSourceDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *CustomerPaymentSourceDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *CustomerPaymentSourceDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *CustomerPaymentSourceDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *CustomerPaymentSourceDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerPaymentSourceDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerPaymentSourceDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerPaymentSourceDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerPaymentSourceDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


