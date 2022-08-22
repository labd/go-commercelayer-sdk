# ShippingMethodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "shipping_methods"]
**Attributes** | [**GETShippingMethods200ResponseDataInnerAttributes**](GETShippingMethods200ResponseDataInnerAttributes.md) |  | 
**Relationships** | Pointer to [**GETShippingMethods200ResponseDataInnerRelationships**](GETShippingMethods200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewShippingMethodData

`func NewShippingMethodData(type_ string, attributes GETShippingMethods200ResponseDataInnerAttributes, ) *ShippingMethodData`

NewShippingMethodData instantiates a new ShippingMethodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodDataWithDefaults

`func NewShippingMethodDataWithDefaults() *ShippingMethodData`

NewShippingMethodDataWithDefaults instantiates a new ShippingMethodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingMethodData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingMethodData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingMethodData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ShippingMethodData) GetAttributes() GETShippingMethods200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingMethodData) GetAttributesOk() (*GETShippingMethods200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingMethodData) SetAttributes(v GETShippingMethods200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingMethodData) GetRelationships() GETShippingMethods200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingMethodData) GetRelationshipsOk() (*GETShippingMethods200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingMethodData) SetRelationships(v GETShippingMethods200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingMethodData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


