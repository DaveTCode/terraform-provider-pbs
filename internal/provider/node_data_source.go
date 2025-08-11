package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-pbs/internal/pbsclient"
)

func NewPbsNodeDataSource() datasource.DataSource {
	return &pbsNodeDataSource{}
}

type pbsNodeDataSource struct {
	client *pbsclient.PbsClient
}

func (d *pbsNodeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_node"
}

func (d *pbsNodeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeID,
			},
			"comment": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeComment,
			},
			"current_aoe": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeCurrentAoe,
			},
			"current_eoe": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeCurrentEoe,
			},
			"in_multi_node_host": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescNodeInMultiNodeHost,
			},
			"mom": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeMom,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: DescNodeName,
			},
			"no_multinode_jobs": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeNoMultinodeJobs,
			},
			"partition": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodePartition,
			},
			"p_names": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodePNames,
			},
			"port": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescNodePort,
			},
			"poweroff_eligible": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescNodePoweroffEligible,
			},
			"power_provisioning": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescNodePowerProvisioning,
			},
			"priority": schema.Int32Attribute{
				Computed:            true,
				MarkdownDescription: DescNodePriority,
				Validators: []validator.Int32{
					int32validator.Between(-1024, 1023),
				},
			},
			"provision_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeProvisionEnable,
			},
			"queue": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeQueue,
				DeprecationMessage:  "This attribute is deprecated",
			},
			"resources_available": schema.MapAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeResourcesAvailable,
				ElementType:         types.StringType,
			},
			"resv_enable": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: DescNodeResvEnable,
			},
		},
	}
}

func (d *pbsNodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sourceData := pbsNodeModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &sourceData)...)

	resultData, err := d.client.GetNode(sourceData.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to PBS server and get hook information", err.Error())
		return
	}

	model := createPbsNodeModel(resultData)

	diag := resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diag...)
}

func (d *pbsNodeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
