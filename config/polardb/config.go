package polardb

import (
	"strings"

	"github.com/aliyun/terraform-provider-alicloud/alicloud"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-alibabacloud/config/common"
)

// clusterChildren are the resources whose upstream Read retries for minutes
// while their cluster is not found. Their ids all start with the cluster id.
var clusterChildren = []string{
	"alicloud_polardb_endpoint",
	"alicloud_polardb_cluster_endpoint",
	"alicloud_polardb_primary_endpoint",
	"alicloud_polardb_database",
}

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
		r.UseAsync = true
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
	for _, name := range clusterChildren {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.TerraformResource.Read = goneWithCluster(r.TerraformResource.Read, clusterNotFound)
		})
	}
}

// goneWithCluster reports a cluster child as gone, without calling upstream's
// Read, once its cluster no longer exists.
//
// Upstream's describe calls for these resources retry InvalidDBClusterId.NotFound
// for five or ten minutes before giving up, presumably to ride out eventual
// consistency right after the cluster is created. Crossplane does not order
// deletes between managed resources, so when a Composition (or anything else)
// deletes a cluster together with its endpoints or databases, the cluster often
// goes first. Every Read of a child then outlasts the reconcile timeout, the
// finalizer update that would follow runs out of time, and the child is stuck
// for good.
func goneWithCluster(read schema.ReadFunc, gone func(meta any, clusterID string) bool) schema.ReadFunc { //nolint:staticcheck // upstream implements these resources with the legacy Read field.
	if read == nil {
		// The code generation schema carries no CRUD functions.
		return nil
	}
	return func(d *schema.ResourceData, meta any) error {
		clusterID, _, ok := strings.Cut(d.Id(), ":")
		if ok && !d.IsNewResource() && gone(meta, clusterID) {
			d.SetId("")
			return nil
		}
		return read(d, meta)
	}
}

// clusterNotFound asks once whether the cluster exists, with the call
// upstream's describes retry on. Any other outcome, including an error, leaves
// the decision to upstream's Read.
func clusterNotFound(meta any, clusterID string) bool {
	client, ok := meta.(*connectivity.AliyunClient)
	if !ok {
		return false
	}
	_, err := client.RpcPost("polardb", "2017-08-01", "DescribeDBClusterEndpoints", nil, map[string]any{
		"RegionId":    client.RegionId,
		"DBClusterId": clusterID,
	}, true)
	return alicloud.IsExpectedErrors(err, []string{"InvalidDBClusterId.NotFound"})
}
