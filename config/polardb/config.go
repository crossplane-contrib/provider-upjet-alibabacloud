package polardb

import (
	"github.com/crossplane-contrib/provider-upjet-alibabacloud/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("alicloud_polardb", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
	})
	p.AddResourceConfigurator("alicloud_polardb_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_account_privilege", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
		r.References["account_name"] = config.Reference{
			TerraformName: "alicloud_polardb_account",
			Extractor:     common.PathAccountNameExtractor,
		}
		// r.References["db_names"] = config.Reference{
		//	 TerraformName: "alicloud_polardb_database",
		//	 Extractor:     common.PathDBNameExtractor,
		// }
	})
	p.AddResourceConfigurator("alicloud_polardb_backup_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_cluster", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["vswitch_id"] = config.Reference{
			TerraformName: "alicloud_vswitch",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_cluster_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_database", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_endpoint_address", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
		r.References["db_endpoint_id"] = config.Reference{
			TerraformName: "alicloud_polardb_endpoint",
			Extractor:     common.PathDBEndpointIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_global_database_network", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
	p.AddResourceConfigurator("alicloud_polardb_primary_endpoint", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "polardb"
		r.ShortGroup = string(common.POLARDB)
		r.References["db_cluster_id"] = config.Reference{
			TerraformName: "alicloud_polardb_cluster",
			Extractor:     common.PathIdExtractor,
		}
	})
}
