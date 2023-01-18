# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PaymentMethodData**](PaymentMethodData.md) |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentMethod) GetData() PaymentMethodData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentMethod) GetDataOk() (*PaymentMethodData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentMethod) SetData(v PaymentMethodData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentMethod) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


