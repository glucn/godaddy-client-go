# Go API client for swagger


## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ote-godaddy.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*V1Api* | [**Available**](docs/V1Api.md#available) | **Get** /v1/domains/available | Determine whether or not the specified domain is available for purchase
*V1Api* | [**AvailableBulk**](docs/V1Api.md#availablebulk) | **Post** /v1/domains/available | Determine whether or not the specified domains are available for purchase
*V1Api* | [**Cancel**](docs/V1Api.md#cancel) | **Delete** /v1/domains/{domain} | Cancel a purchased domain
*V1Api* | [**CancelPrivacy**](docs/V1Api.md#cancelprivacy) | **Delete** /v1/domains/{domain}/privacy | Submit a privacy cancellation request for the given domain
*V1Api* | [**ContactsValidate**](docs/V1Api.md#contactsvalidate) | **Post** /v1/domains/contacts/validate | Validate the request body using the Domain Contact Validation Schema for specified domains.
*V1Api* | [**Get**](docs/V1Api.md#get) | **Get** /v1/domains/{domain} | Retrieve details for the specified Domain
*V1Api* | [**GetAgreement**](docs/V1Api.md#getagreement) | **Get** /v1/domains/agreements | Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons
*V1Api* | [**List**](docs/V1Api.md#list) | **Get** /v1/domains | Retrieve a list of Domains for the specified Shopper
*V1Api* | [**Purchase**](docs/V1Api.md#purchase) | **Post** /v1/domains/purchase | Purchase and register the specified Domain
*V1Api* | [**PurchasePrivacy**](docs/V1Api.md#purchaseprivacy) | **Post** /v1/domains/{domain}/privacy/purchase | Purchase privacy for a specified domain
*V1Api* | [**RecordAdd**](docs/V1Api.md#recordadd) | **Patch** /v1/domains/{domain}/records | Add the specified DNS Records to the specified Domain
*V1Api* | [**RecordGet**](docs/V1Api.md#recordget) | **Get** /v1/domains/{domain}/records/{type}/{name} | Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name
*V1Api* | [**RecordReplace**](docs/V1Api.md#recordreplace) | **Put** /v1/domains/{domain}/records | Replace all DNS Records for the specified Domain
*V1Api* | [**RecordReplaceType**](docs/V1Api.md#recordreplacetype) | **Put** /v1/domains/{domain}/records/{type} | Replace all DNS Records for the specified Domain with the specified Type
*V1Api* | [**RecordReplaceTypeName**](docs/V1Api.md#recordreplacetypename) | **Put** /v1/domains/{domain}/records/{type}/{name} | Replace all DNS Records for the specified Domain with the specified Type and Name
*V1Api* | [**Renew**](docs/V1Api.md#renew) | **Post** /v1/domains/{domain}/renew | Renew the specified Domain
*V1Api* | [**Schema**](docs/V1Api.md#schema) | **Get** /v1/domains/purchase/schema/{tld} | Retrieve the schema to be submitted when registering a Domain for the specified TLD
*V1Api* | [**Suggest**](docs/V1Api.md#suggest) | **Get** /v1/domains/suggest | Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper&#39;s purchase history
*V1Api* | [**Tlds**](docs/V1Api.md#tlds) | **Get** /v1/domains/tlds | Retrieves a list of TLDs supported and enabled for sale
*V1Api* | [**TransferIn**](docs/V1Api.md#transferin) | **Post** /v1/domains/{domain}/transfer | Purchase and start or restart transfer process
*V1Api* | [**Update**](docs/V1Api.md#update) | **Patch** /v1/domains/{domain} | Update details for the specified Domain
*V1Api* | [**UpdateContacts**](docs/V1Api.md#updatecontacts) | **Patch** /v1/domains/{domain}/contacts | Update domain
*V1Api* | [**Validate**](docs/V1Api.md#validate) | **Post** /v1/domains/purchase/validate | Validate the request body using the Domain Purchase Schema for the specified TLD
*V1Api* | [**VerifyEmail**](docs/V1Api.md#verifyemail) | **Post** /v1/domains/{domain}/verifyRegistrantEmail | Re-send Contact E-mail Verification for specified Domain
*V2Api* | [**DomainsForwardsDelete**](docs/V2Api.md#domainsforwardsdelete) | **Delete** /v2/customers/{customerId}/domains/forwards/{fqdn} | Submit a forwarding cancellation request for the given fqdn
*V2Api* | [**DomainsForwardsGet**](docs/V2Api.md#domainsforwardsget) | **Get** /v2/customers/{customerId}/domains/forwards/{fqdn} | Retrieve the forwarding information for the given fqdn
*V2Api* | [**DomainsForwardsPut**](docs/V2Api.md#domainsforwardsput) | **Put** /v2/customers/{customerId}/domains/forwards/{fqdn} | Modify the forwarding information for the given fqdn


## Documentation For Models

 - [Address](docs/Address.md)
 - [ArrayOfDnsRecord](docs/ArrayOfDnsRecord.md)
 - [Consent](docs/Consent.md)
 - [Contact](docs/Contact.md)
 - [DnsRecord](docs/DnsRecord.md)
 - [DnsRecordCreateType](docs/DnsRecordCreateType.md)
 - [DnsRecordCreateTypeName](docs/DnsRecordCreateTypeName.md)
 - [Domain](docs/Domain.md)
 - [DomainAvailableBulk](docs/DomainAvailableBulk.md)
 - [DomainAvailableBulkMixed](docs/DomainAvailableBulkMixed.md)
 - [DomainAvailableError](docs/DomainAvailableError.md)
 - [DomainAvailableResponse](docs/DomainAvailableResponse.md)
 - [DomainContacts](docs/DomainContacts.md)
 - [DomainDetail](docs/DomainDetail.md)
 - [DomainForwarding](docs/DomainForwarding.md)
 - [DomainForwardingCreate](docs/DomainForwardingCreate.md)
 - [DomainForwardingMask](docs/DomainForwardingMask.md)
 - [DomainPurchase](docs/DomainPurchase.md)
 - [DomainPurchaseResponse](docs/DomainPurchaseResponse.md)
 - [DomainRenew](docs/DomainRenew.md)
 - [DomainSuggestion](docs/DomainSuggestion.md)
 - [DomainSummary](docs/DomainSummary.md)
 - [DomainTransferIn](docs/DomainTransferIn.md)
 - [DomainUpdate](docs/DomainUpdate.md)
 - [DomainsContactsBulk](docs/DomainsContactsBulk.md)
 - [ErrorDomainContactsValidate](docs/ErrorDomainContactsValidate.md)
 - [ErrorField](docs/ErrorField.md)
 - [ErrorFieldDomainContactsValidate](docs/ErrorFieldDomainContactsValidate.md)
 - [ErrorLimit](docs/ErrorLimit.md)
 - [JsonDataType](docs/JsonDataType.md)
 - [JsonProperty](docs/JsonProperty.md)
 - [JsonSchema](docs/JsonSchema.md)
 - [LegalAgreement](docs/LegalAgreement.md)
 - [ModelError](docs/ModelError.md)
 - [PrivacyPurchase](docs/PrivacyPurchase.md)
 - [RealNameValidation](docs/RealNameValidation.md)
 - [TldSummary](docs/TldSummary.md)
 - [VerificationDomainName](docs/VerificationDomainName.md)
 - [VerificationRealName](docs/VerificationRealName.md)
 - [VerificationsDomain](docs/VerificationsDomain.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



