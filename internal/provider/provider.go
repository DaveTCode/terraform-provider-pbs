package provider

import (
	"context"
	"fmt"
	"net"
	"os"
	"terraform-provider-pbs/internal/pbsclient"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/crypto/ssh"
)

var (
	_ provider.Provider = &pbsProvider{}
)

type pbsProviderModel struct {
	Server        types.String `tfsdk:"server"`
	SshPort       types.String `tfsdk:"sshport"`
	Username      types.String `tfsdk:"username"`
	Password      types.String `tfsdk:"password"`
	SshPrivateKey types.String `tfsdk:"ssh_private_key"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &pbsProvider{
			version: version,
		}
	}
}

type pbsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *pbsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pbs"
	resp.Version = p.version
}

func (p *pbsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"server": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The PBS server address",
			},
			"sshport": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The PBS server SSH port",
			},
			"username": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "An SSH username with access to run qmgr commands on the PBS server",
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "The password for the SSH username",
			},
			"ssh_private_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "The SSH private key content for authentication (alternative to password)",
			},
		},
	}
}

func (p *pbsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config pbsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Server.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("server"),
			"Unknown PBS server",
			"The provider cannot create the PBS client as there is an unknown configuration value for the PBS server. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PBS_SERVER environment variable.",
		)
	}

	if config.SshPort.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sshport"),
			"Unknown PBS sshport",
			"The provider cannot create the PBS client as there is an unknown configuration value for the PBS sshport. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PBS_SSH_PORT environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown PBS Username",
			"The provider cannot create the PBS client as there is an unknown configuration value for the PBS username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PBS_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown PBS Password",
			"The provider cannot create the PBS API client as there is an unknown configuration value for the PBS API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PBS_PASSWORD environment variable.",
		)
	}

	if config.SshPrivateKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("ssh_private_key"),
			"Unknown SSH Private Key",
			"The provider cannot create the PBS client as there is an unknown configuration value for the SSH private key.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	server := os.Getenv("PBS_SERVER")
	sshPort := os.Getenv("PBS_SSH_PORT")
	username := os.Getenv("PBS_USERNAME")
	password := os.Getenv("PBS_PASSWORD")
	sshPrivateKey := os.Getenv("PBS_SSH_PRIVATE_KEY")

	if !config.Server.IsNull() {
		server = config.Server.ValueString()
	}

	if !config.SshPort.IsNull() {
		sshPort = config.SshPort.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if !config.SshPrivateKey.IsNull() {
		sshPrivateKey = config.SshPrivateKey.ValueString()
	}

	if server == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("server"),
			"Missing PBS server",
			"The provider cannot create the PBS client as there is a missing or empty value for the PBS server. "+
				"Set the server value in the configuration or use the PBS_SERVER environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if sshPort == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sshport"),
			"Missing PBS sshport",
			"The provider cannot create the PBS client as there is a missing or empty value for the PBS SSH port. "+
				"Set the server value in the configuration or use the PBS_SSH_PORT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing PBS Username",
			"The provider cannot create the PBS client as there is a missing or empty value for the PBS username. "+
				"Set the username value in the configuration or use the PBS_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	// Validate authentication methods - either password OR SSH key must be provided
	hasPassword := password != ""
	hasSshKey := sshPrivateKey != ""

	if !hasPassword && !hasSshKey {
		resp.Diagnostics.AddError(
			"Missing authentication credentials",
			"The provider requires either password or SSH key authentication. Provide either:\n"+
				"- password (via 'password' configuration or PBS_PASSWORD environment variable)\n"+
				"- ssh_private_key (via 'ssh_private_key' configuration or PBS_SSH_PRIVATE_KEY environment variable)",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Build SSH authentication methods
	var authMethods []ssh.AuthMethod

	// Add password authentication if provided
	if hasPassword {
		authMethods = append(authMethods, ssh.Password(password))
	}

	// Add SSH key authentication if provided
	if hasSshKey {
		// Parse the private key
		signer, err := ssh.ParsePrivateKey([]byte(sshPrivateKey))
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("ssh_private_key"),
				"Invalid SSH private key",
				fmt.Sprintf("Cannot parse SSH private key: %v", err),
			)
			return
		}

		authMethods = append(authMethods, ssh.PublicKeys(signer))
	}

	// Create a new SSH client using the configuration values
	sshConfig := &ssh.ClientConfig{
		User:            username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	pbsClient := &pbsclient.PbsClient{
		SshClientConfig: sshConfig,
		Address:         net.JoinHostPort(server, sshPort),
	}

	// Make the pbs client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = pbsClient
	resp.ResourceData = pbsClient
}

func (p *pbsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewQueueDataSource,
		NewPbsResourceDataSource,
		NewPbsHookDataSource,
		NewPbsNodeDataSource,
		NewServerDataSource,
	}
}

func (p *pbsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewQueueResource,
		NewPbsResourceResource,
		NewPbsHookResource,
		NewPbsNodeResource,
		NewServerResource,
	}
}
