# POSTSkuListsRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTSkuListsRequestDataAttributes**](POSTSkuListsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCouponRecipientsRequestDataRelationships**](POSTCouponRecipientsRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTSkuListsRequestData

`func NewPOSTSkuListsRequestData(type_ interface{}, attributes POSTSkuListsRequestDataAttributes, ) *POSTSkuListsRequestData`

NewPOSTSkuListsRequestData instantiates a new POSTSkuListsRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuListsRequestDataWithDefaults

`func NewPOSTSkuListsRequestDataWithDefaults() *POSTSkuListsRequestData`

NewPOSTSkuListsRequestDataWithDefaults instantiates a new POSTSkuListsRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTSkuListsRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTSkuListsRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTSkuListsRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTSkuListsRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTSkuListsRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTSkuListsRequestData) GetAttributes() POSTSkuListsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTSkuListsRequestData) GetAttributesOk() (*POSTSkuListsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTSkuListsRequestData) SetAttributes(v POSTSkuListsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTSkuListsRequestData) GetRelationships() POSTCouponRecipientsRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTSkuListsRequestData) GetRelationshipsOk() (*POSTCouponRecipientsRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTSkuListsRequestData) SetRelationships(v POSTCouponRecipientsRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTSkuListsRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


