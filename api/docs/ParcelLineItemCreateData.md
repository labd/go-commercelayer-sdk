# ParcelLineItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "parcel_line_items"]
**Attributes** | [**POSTParcelLineItems201ResponseDataAttributes**](POSTParcelLineItems201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTParcelLineItems201ResponseDataRelationships**](POSTParcelLineItems201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewParcelLineItemCreateData

`func NewParcelLineItemCreateData(type_ string, attributes POSTParcelLineItems201ResponseDataAttributes, ) *ParcelLineItemCreateData`

NewParcelLineItemCreateData instantiates a new ParcelLineItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelLineItemCreateDataWithDefaults

`func NewParcelLineItemCreateDataWithDefaults() *ParcelLineItemCreateData`

NewParcelLineItemCreateDataWithDefaults instantiates a new ParcelLineItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParcelLineItemCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParcelLineItemCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParcelLineItemCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ParcelLineItemCreateData) GetAttributes() POSTParcelLineItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParcelLineItemCreateData) GetAttributesOk() (*POSTParcelLineItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParcelLineItemCreateData) SetAttributes(v POSTParcelLineItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParcelLineItemCreateData) GetRelationships() POSTParcelLineItems201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParcelLineItemCreateData) GetRelationshipsOk() (*POSTParcelLineItems201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParcelLineItemCreateData) SetRelationships(v POSTParcelLineItems201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ParcelLineItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


