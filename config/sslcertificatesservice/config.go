package sslcertificatesservice

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_ssl_certificates_service_certificate", func(r *config.Resource) {
		// upjet's name-derived defaults would be group "ssl" and Kind
		// "CertificatesServiceCertificate"; name both explicitly instead.
		r.ShortGroup = string(common.SSLCertificatesService)
		r.Kind = "Certificate"

		// "name" is deprecated in favour of "certificate_name" and "lang" is
		// deprecated upstream; both are dropped during code generation, see
		// generationOnlySchemaOmissions in config/schema_omissions.go, so the
		// CRD only exposes the current fields.

		// The certificate body is not flagged sensitive upstream, so upjet
		// would generate it as a plain string. Marked sensitive during code
		// generation instead; see generationOnlySensitiveFields in
		// config/schema_omissions.go, so it is supplied via a Secret
		// reference (certSecretRef), the same way the private key
		// (keySecretRef) is. The SM2 dual-certificate bodies (encrypt_cert,
		// sign_cert) can be made secret-backed the same way if needed.

		// Upstream marks certificate_name with
		// AtLeastOneOf: ["certificate_name", "name"], so dropping "name" above
		// leaves it effectively required. It stays Optional+Computed in the
		// schema, so without this the CRD carries no validation and a
		// Certificate missing certificateName fails at plan time instead of at
		// admission. MarkAsRequired records that without touching the schema,
		// which keeps the 2.0 port simpler.
		r.MarkAsRequired("certificate_name")
	})
}
