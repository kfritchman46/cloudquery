package downtimes

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Downtimes() *schema.Table {
	return &schema.Table{
		Name:      "datadog_downtimes",
		Resolver:  fetchDowntimes,
		Multiplex: client.AccountMultiplex,
		Transform: transformers.TransformWithStruct(&datadogV1.Downtime{}),
		Columns: []schema.Column{
			{
				Name:       "account_name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAccountName,
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("Id"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchDowntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.DowntimesAPI.ListDowntimes(ctx)
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
