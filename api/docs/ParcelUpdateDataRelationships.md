# ParcelUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Shipment** | Pointer to [**ParcelCreateDataRelationshipsShipment**](ParcelCreateDataRelationshipsShipment.md) |  | [optional] 
**Package** | Pointer to [**ParcelCreateDataRelationshipsPackage**](ParcelCreateDataRelationshipsPackage.md) |  | [optional] 

## Methods

### NewParcelUpdateDataRelationships

`func NewParcelUpdateDataRelationships() *ParcelUpdateDataRelationships`

NewParcelUpdateDataRelationships instantiates a new ParcelUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelUpdateDataRelationshipsWithDefaults

`func NewParcelUpdateDataRelationshipsWithDefaults() *ParcelUpdateDataRelationships`

NewParcelUpdateDataRelationshipsWithDefaults instantiates a new ParcelUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipment

`func (o *ParcelUpdateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment`

GetShipment returns the Shipment field if non-nil, zero value otherwise.

### GetShipmentOk

`func (o *ParcelUpdateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool)`

GetShipmentOk returns a tuple with the Shipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipment

`func (o *ParcelUpdateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment)`

SetShipment sets Shipment field to given value.

### HasShipment

`func (o *ParcelUpdateDataRelationships) HasShipment() bool`

HasShipment returns a boolean if a field has been set.

### GetPackage

`func (o *ParcelUpdateDataRelationships) GetPackage() ParcelCreateDataRelationshipsPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ParcelUpdateDataRelationships) GetPackageOk() (*ParcelCreateDataRelationshipsPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ParcelUpdateDataRelationships) SetPackage(v ParcelCreateDataRelationshipsPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ParcelUpdateDataRelationships) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


