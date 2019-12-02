/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainForwarding struct {
	// The fqdn (domain or sub domain) to forward (ex somedomain.com or sub.somedomain.com)
	Fqdn string `json:"fqdn"`
	// The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>
	Type_ string `json:"type"`
	// Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)
	Url string `json:"url"`
	// Additional configuration that can be provided when type = 'MASKED'
	Mask *DomainForwardingMask `json:"mask,omitempty"`
}
