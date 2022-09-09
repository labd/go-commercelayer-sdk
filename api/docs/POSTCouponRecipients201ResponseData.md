# POSTCouponRecipients201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "coupon_recipients"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTCouponRecipients201ResponseDataAttributes**](POSTCouponRecipients201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETCouponRecipients200ResponseDataInnerRelationships**](GETCouponRecipients200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewPOSTCouponRecipients201ResponseData

`func NewPOSTCouponRecipients201ResponseData() *POSTCouponRecipients201ResponseData`

NewPOSTCouponRecipients201ResponseData instantiates a new POSTCouponRecipients201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCouponRecipients201ResponseDataWithDefaults

`func NewPOSTCouponRecipients201ResponseDataWithDefaults() *POSTCouponRecipients201ResponseData`

NewPOSTCouponRecipients201ResponseDataWithDefaults instantiates a new POSTCouponRecipients201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTCouponRecipients201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTCouponRecipients201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTCouponRecipients201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTCouponRecipients201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTCouponRecipients201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCouponRecipients201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCouponRecipients201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTCouponRecipients201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTCouponRecipients201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTCouponRecipients201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTCouponRecipients201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTCouponRecipients201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTCouponRecipients201ResponseData) GetAttributes() POSTCouponRecipients201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCouponRecipients201ResponseData) GetAttributesOk() (*POSTCouponRecipients201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCouponRecipients201ResponseData) SetAttributes(v POSTCouponRecipients201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTCouponRecipients201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTCouponRecipients201ResponseData) GetRelationships() GETCouponRecipients200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCouponRecipients201ResponseData) GetRelationshipsOk() (*GETCouponRecipients200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCouponRecipients201ResponseData) SetRelationships(v GETCouponRecipients200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCouponRecipients201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


