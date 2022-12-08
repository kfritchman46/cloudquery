// Code generated by codegen; DO NOT EDIT.

package billing

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/billing/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/billing/apiv1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "gcp_billing_services",
		Resolver:  fetchServices,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceId"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "business_entity_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BusinessEntityName"),
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListServicesRequest{}
	gcpClient, err := billing.NewCloudCatalogClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp

	}
	return nil
}
