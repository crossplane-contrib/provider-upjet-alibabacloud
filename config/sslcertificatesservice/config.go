package sslcertificatesservice

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ssl_certificates_service_certificate", func(r *config.Resource) {
		// Override the default group that upjet would generate for this
		// resource (which would be "sslcertificatesservice") and shorten the
		// Kind from "ssl_certificates_service_certificate".
		r.ShortGroup = string(common.SSLCertificatesService)
		r.Kind = "Certificate"

		// "name" is deprecated in favour of "certificate_name" and "lang" is
		// deprecated upstream; drop them so the CRD only exposes the current
		// fields.
		delete(r.TerraformResource.Schema, "name")
		delete(r.TerraformResource.Schema, "lang")

		// The certificate body is not flagged sensitive upstream, so upjet
		// generates it as a plain string. Mark it sensitive so it is supplied
		// via a Secret reference (certSecretRef), the same way the private key
		// (keySecretRef) is. The SM2 dual-certificate bodies (encrypt_cert,
		// sign_cert) can be made secret-backed the same way if needed.
		r.TerraformResource.Schema["cert"].Sensitive = true
	})
}
