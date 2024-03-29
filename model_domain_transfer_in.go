/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainTransferIn struct {
	// Authorization code from registrar for transferring a domain
	AuthCode string `json:"authCode"`
	// Required agreements can be retrieved via the GET ./domains/agreements endpoint
	Consent *Consent `json:"consent"`
	// Can be more than 1 but no more than 10 years total including current registration length
	Period int32 `json:"period,omitempty"`
	// Whether or not privacy has been requested
	Privacy bool `json:"privacy,omitempty"`
	// Whether or not the domain should be configured to automatically renew
	RenewAuto bool `json:"renewAuto,omitempty"`
	// The contact to use for the domain admin contact. Depending on the tld of the domain being transferred, this field may be required.
	ContactAdmin *Contact `json:"contactAdmin,omitempty"`
	// The contact to use for the domain billing contact. Depending on the tld of the domain being transferred, this field may be required.
	ContactBilling *Contact `json:"contactBilling,omitempty"`
	// The contact to use for the domain registrant contact. Depending on the tld of the domain being transferred, this field may be required.
	ContactRegistrant *Contact `json:"contactRegistrant,omitempty"`
	// The contact to use for the domain tech contact. Depending on the tld of the domain being transferred, this field may be required.
	ContactTech *Contact `json:"contactTech,omitempty"`
}
