package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewPbsHookDataSource() datasource.DataSource {
	return &pbsHookDataSource{}
}

type pbsHookDataSource struct {
	client *pbsclient.PbsClient
}

type pbsHookDataSourceModel struct {
	Hooks []pbsHookModel `tfsdk:"hooks"`
}

func (d *pbsHookDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hook"
}

func (d *pbsHookDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"alarm": schema.Int32Attribute{
				Computed: true,
			},
			"debug": schema.BoolAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"event": schema.StringAttribute{
				Computed: true,
			},
			"fail_action": schema.StringAttribute{
				Computed: true,
			},
			"freq": schema.Int32Attribute{
				Computed: true,
			},
			"order": schema.Int32Attribute{
				Computed: true,
			},
			"user": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *pbsHookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	pbsHookData := pbsHookDataSourceModel{}
	pbsHooks, err := d.client.GetHooks()
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get hook information", err.Error())
		return
	}

	for _, hook := range pbsHooks {
		pbsHookData.Hooks = append(pbsHookData.Hooks, createPbsHookModel(hook))
	}
}

func (d *pbsHookDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
