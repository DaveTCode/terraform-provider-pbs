package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewPbsResourceDataSource() datasource.DataSource {
	return &pbsResourceDataSource{}
}

type pbsResourceDataSource struct {
	client *pbsclient.PbsClient
}

type pbsResourceDataSourceModel struct {
	Resources []pbsResourceModel `tfsdk:"resources"`
}

func (d *pbsResourceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

func (d *pbsResourceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"flag": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *pbsResourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	pbsResourceData := pbsResourceDataSourceModel{}
	pbsResources, err := d.client.GetResources()
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get resource information", err.Error())
		return
	}

	for _, resource := range pbsResources {
		pbsResourceData.Resources = append(pbsResourceData.Resources, createPbsResoureModel(resource))
	}
}

func (d *pbsResourceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pbsclient.PbsClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *PbsClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
