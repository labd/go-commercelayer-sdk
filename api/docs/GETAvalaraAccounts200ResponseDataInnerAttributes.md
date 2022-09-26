# GETAvalaraAccounts200ResponseDataInnerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The tax calculator&#39;s internal name. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the resource (hash). | [optional] 
**CreatedAt** | Pointer to **string** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Username** | Pointer to **string** | The Avalara account username. | [optional] 
**CompanyCode** | Pointer to **string** | The Avalara company code. | [optional] 
**CommitInvoice** | Pointer to **string** | Indicates if the transaction will be recorded and visible on the Avalara website. | [optional] 
**Ddp** | Pointer to **string** | Indicates if the seller is responsible for paying/remitting the customs duty &amp; import tax to the customs authorities. | [optional] 

## Methods

### NewGETAvalaraAccounts200ResponseDataInnerAttributes

`func NewGETAvalaraAccounts200ResponseDataInnerAttributes() *GETAvalaraAccounts200ResponseDataInnerAttributes`

NewGETAvalaraAccounts200ResponseDataInnerAttributes instantiates a new GETAvalaraAccounts200ResponseDataInnerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAvalaraAccounts200ResponseDataInnerAttributesWithDefaults

`func NewGETAvalaraAccounts200ResponseDataInnerAttributesWithDefaults() *GETAvalaraAccounts200ResponseDataInnerAttributes`

NewGETAvalaraAccounts200ResponseDataInnerAttributesWithDefaults instantiates a new GETAvalaraAccounts200ResponseDataInnerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReference

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsername

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCompanyCode

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCompanyCode() string`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCompanyCodeOk() (*string, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCompanyCode(v string)`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### GetCommitInvoice

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCommitInvoice() string`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetCommitInvoiceOk() (*string, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetCommitInvoice(v string)`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### GetDdp

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetDdp() string`

GetDdp returns the Ddp field if non-nil, zero value otherwise.

### GetDdpOk

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) GetDdpOk() (*string, bool)`

GetDdpOk returns a tuple with the Ddp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdp

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) SetDdp(v string)`

SetDdp sets Ddp field to given value.

### HasDdp

`func (o *GETAvalaraAccounts200ResponseDataInnerAttributes) HasDdp() bool`

HasDdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


