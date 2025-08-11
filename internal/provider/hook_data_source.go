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

func (d *pbsHookDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hook"
}

func (d *pbsHookDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescHookID,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescHookName,
			},
			"type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescHookType,
			},
			"alarm": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescHookAlarm,
			},
			"debug": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescHookDebug,
			},
			"enabled": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescHookEnabled,
			},
			"event": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescHookEvent,
			},
			"fail_action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescHookFailAction,
			},
			"freq": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescHookFreq,
			},
			"order": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescHookOrder,
			},
			"user": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescHookUser,
			},
		},
	}
}

func (d *pbsHookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sourceData := pbsHookModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &sourceData)...)

	resultData, err := d.client.GetHook(sourceData.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get hook information", err.Error())
		return
	}

	model := createPbsHookModel(resultData)

	diag := resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diag...)
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
