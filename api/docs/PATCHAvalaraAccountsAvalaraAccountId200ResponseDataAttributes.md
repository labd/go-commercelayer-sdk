# PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The tax calculator&#39;s internal name. | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**Username** | Pointer to **string** | The Avalara account username. | [optional] 
**Password** | Pointer to **string** | The Avalara account password. | [optional] 
**CompanyCode** | Pointer to **string** | The Avalara company code. | [optional] 
**CommitInvoice** | Pointer to **string** | Indicates if the transaction will be recorded and visible on the Avalara website. | [optional] 
**Ddp** | Pointer to **string** | Indicates if the seller is responsible for paying/remitting the customs duty &amp; import tax to the customs authorities. | [optional] 

## Methods

### NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes

`func NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes() *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes`

NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes instantiates a new PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributesWithDefaults

`func NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributesWithDefaults() *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes`

NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributesWithDefaults instantiates a new PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCompanyCode() string`

GetCompanyCode returns the CompanyCode field if non-nil, zero value otherwise.

### GetCompanyCodeOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCompanyCodeOk() (*string, bool)`

GetCompanyCodeOk returns a tuple with the CompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetCompanyCode(v string)`

SetCompanyCode sets CompanyCode field to given value.

### HasCompanyCode

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasCompanyCode() bool`

HasCompanyCode returns a boolean if a field has been set.

### GetCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCommitInvoice() string`

GetCommitInvoice returns the CommitInvoice field if non-nil, zero value otherwise.

### GetCommitInvoiceOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetCommitInvoiceOk() (*string, bool)`

GetCommitInvoiceOk returns a tuple with the CommitInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetCommitInvoice(v string)`

SetCommitInvoice sets CommitInvoice field to given value.

### HasCommitInvoice

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasCommitInvoice() bool`

HasCommitInvoice returns a boolean if a field has been set.

### GetDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetDdp() string`

GetDdp returns the Ddp field if non-nil, zero value otherwise.

### GetDdpOk

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) GetDdpOk() (*string, bool)`

GetDdpOk returns a tuple with the Ddp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) SetDdp(v string)`

SetDdp sets Ddp field to given value.

### HasDdp

`func (o *PATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes) HasDdp() bool`

HasDdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


