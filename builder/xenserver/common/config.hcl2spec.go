// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Username                  *string           `mapstructure:"remote_username" cty:"remote_username" hcl:"remote_username"`
	Password                  *string           `mapstructure:"remote_password" cty:"remote_password" hcl:"remote_password"`
	HostIp                    *string           `mapstructure:"remote_host" cty:"remote_host" hcl:"remote_host"`
	HostSshPort               *uint             `mapstructure:"remote_ssh_port" cty:"remote_ssh_port" hcl:"remote_ssh_port"`
	VMName                    *string           `mapstructure:"vm_name" cty:"vm_name" hcl:"vm_name"`
	VMDescription             *string           `mapstructure:"vm_description" cty:"vm_description" hcl:"vm_description"`
	SrName                    *string           `mapstructure:"sr_name" cty:"sr_name" hcl:"sr_name"`
	SrISOName                 *string           `mapstructure:"sr_iso_name" required:"false" cty:"sr_iso_name" hcl:"sr_iso_name"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files" hcl:"floppy_files"`
	NetworkNames              []string          `mapstructure:"network_names" cty:"network_names" hcl:"network_names"`
	ExportNetworkNames        []string          `mapstructure:"export_network_names" cty:"export_network_names" hcl:"export_network_names"`
	VMTags                    []string          `mapstructure:"vm_tags" cty:"vm_tags" hcl:"vm_tags"`
	HostPortMin               *uint             `mapstructure:"host_port_min" cty:"host_port_min" hcl:"host_port_min"`
	HostPortMax               *uint             `mapstructure:"host_port_max" cty:"host_port_max" hcl:"host_port_max"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command" hcl:"boot_command"`
	ShutdownCommand           *string           `mapstructure:"shutdown_command" cty:"shutdown_command" hcl:"shutdown_command"`
	RawBootWait               *string           `mapstructure:"boot_wait" cty:"boot_wait" hcl:"boot_wait"`
	ToolsIsoName              *string           `mapstructure:"tools_iso_name" cty:"tools_iso_name" hcl:"tools_iso_name"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory" hcl:"http_directory"`
	HTTPPortMin               *uint             `mapstructure:"http_port_min" cty:"http_port_min" hcl:"http_port_min"`
	HTTPPortMax               *uint             `mapstructure:"http_port_max" cty:"http_port_max" hcl:"http_port_max"`
	SSHKeyPath                *string           `mapstructure:"ssh_key_path" cty:"ssh_key_path" hcl:"ssh_key_path"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHPort                   *uint             `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUser                   *string           `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string           `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int              `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string          `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string          `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string           `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string           `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string           `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte            `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool             `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	SSHHostPortMin            *uint             `mapstructure:"ssh_host_port_min" cty:"ssh_host_port_min" hcl:"ssh_host_port_min"`
	SSHHostPortMax            *uint             `mapstructure:"ssh_host_port_max" cty:"ssh_host_port_max" hcl:"ssh_host_port_max"`
	SSHSkipNatMapping         *bool             `mapstructure:"ssh_skip_nat_mapping" cty:"ssh_skip_nat_mapping" hcl:"ssh_skip_nat_mapping"`
	OutputDir                 *string           `mapstructure:"output_directory" cty:"output_directory" hcl:"output_directory"`
	Format                    *string           `mapstructure:"format" cty:"format" hcl:"format"`
	KeepVM                    *string           `mapstructure:"keep_vm" cty:"keep_vm" hcl:"keep_vm"`
	IPGetter                  *string           `mapstructure:"ip_getter" cty:"ip_getter" hcl:"ip_getter"`
	VCPUsMax                  *uint             `mapstructure:"vcpus_max" cty:"vcpus_max" hcl:"vcpus_max"`
	VCPUsAtStartup            *uint             `mapstructure:"vcpus_atstartup" cty:"vcpus_atstartup" hcl:"vcpus_atstartup"`
	VMMemory                  *uint             `mapstructure:"vm_memory" cty:"vm_memory" hcl:"vm_memory"`
	DiskName                  *string           `mapstructure:"disk_name" cty:"disk_name" hcl:"disk_name"`
	DiskSize                  *uint             `mapstructure:"disk_size" cty:"disk_size" hcl:"disk_size"`
	CloneTemplate             *string           `mapstructure:"clone_template" cty:"clone_template" hcl:"clone_template"`
	VMOtherConfig             map[string]string `mapstructure:"vm_other_config" cty:"vm_other_config" hcl:"vm_other_config"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" cty:"iso_checksum" hcl:"iso_checksum"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls" hcl:"iso_urls"`
	ISOUrl                    *string           `mapstructure:"iso_url" cty:"iso_url" hcl:"iso_url"`
	ISOName                   *string           `mapstructure:"iso_name" cty:"iso_name" hcl:"iso_name"`
	PlatformArgs              map[string]string `mapstructure:"platform_args" cty:"platform_args" hcl:"platform_args"`
	RawInstallTimeout         *string           `mapstructure:"install_timeout" cty:"install_timeout" hcl:"install_timeout"`
	SourcePath                *string           `mapstructure:"source_path" cty:"source_path" hcl:"source_path"`
	Firmware                  *string           `mapstructure:"firmware" cty:"firmware" hcl:"firmware"`
	SkipSetTemplate           *bool             `mapstructure:"skip_set_template" cty:"skip_set_template" hcl:"skip_set_template"`
	InstallationDoneFile			*string           `mapstructure:"installation_done_file" cty:"installation_done_file" hcl:"installation_done_file"`
	RawInstallationDoneTimeout		*string           `mapstructure:"installation_done_timeout" cty:"installation_done_timeout" hcl:"installation_done_timeout"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"remote_username":              &hcldec.AttrSpec{Name: "remote_username", Type: cty.String, Required: false},
		"remote_password":              &hcldec.AttrSpec{Name: "remote_password", Type: cty.String, Required: false},
		"remote_host":                  &hcldec.AttrSpec{Name: "remote_host", Type: cty.String, Required: false},
		"remote_ssh_port":              &hcldec.AttrSpec{Name: "remote_ssh_port", Type: cty.Number, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"vm_description":               &hcldec.AttrSpec{Name: "vm_description", Type: cty.String, Required: false},
		"sr_name":                      &hcldec.AttrSpec{Name: "sr_name", Type: cty.String, Required: false},
		"sr_iso_name":                  &hcldec.AttrSpec{Name: "sr_iso_name", Type: cty.String, Required: false},
		"floppy_files":                 &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"network_names":                &hcldec.AttrSpec{Name: "network_names", Type: cty.List(cty.String), Required: false},
		"export_network_names":         &hcldec.AttrSpec{Name: "export_network_names", Type: cty.List(cty.String), Required: false},
		"vm_tags":                      &hcldec.AttrSpec{Name: "vm_tags", Type: cty.List(cty.String), Required: false},
		"host_port_min":                &hcldec.AttrSpec{Name: "host_port_min", Type: cty.Number, Required: false},
		"host_port_max":                &hcldec.AttrSpec{Name: "host_port_max", Type: cty.Number, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"shutdown_command":             &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"tools_iso_name":               &hcldec.AttrSpec{Name: "tools_iso_name", Type: cty.String, Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"ssh_key_path":                 &hcldec.AttrSpec{Name: "ssh_key_path", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"ssh_host_port_min":            &hcldec.AttrSpec{Name: "ssh_host_port_min", Type: cty.Number, Required: false},
		"ssh_host_port_max":            &hcldec.AttrSpec{Name: "ssh_host_port_max", Type: cty.Number, Required: false},
		"ssh_skip_nat_mapping":         &hcldec.AttrSpec{Name: "ssh_skip_nat_mapping", Type: cty.Bool, Required: false},
		"output_directory":             &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"format":                       &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"keep_vm":                      &hcldec.AttrSpec{Name: "keep_vm", Type: cty.String, Required: false},
		"ip_getter":                    &hcldec.AttrSpec{Name: "ip_getter", Type: cty.String, Required: false},
		"vcpus_max":                    &hcldec.AttrSpec{Name: "vcpus_max", Type: cty.Number, Required: false},
		"vcpus_atstartup":              &hcldec.AttrSpec{Name: "vcpus_atstartup", Type: cty.Number, Required: false},
		"vm_memory":                    &hcldec.AttrSpec{Name: "vm_memory", Type: cty.Number, Required: false},
		"disk_name":                    &hcldec.AttrSpec{Name: "disk_name", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"clone_template":               &hcldec.AttrSpec{Name: "clone_template", Type: cty.String, Required: false},
		"vm_other_config":              &hcldec.AttrSpec{Name: "vm_other_config", Type: cty.Map(cty.String), Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_name":                     &hcldec.AttrSpec{Name: "iso_name", Type: cty.String, Required: false},
		"platform_args":                &hcldec.AttrSpec{Name: "platform_args", Type: cty.Map(cty.String), Required: false},
		"install_timeout":              &hcldec.AttrSpec{Name: "install_timeout", Type: cty.String, Required: false},
		"source_path":                  &hcldec.AttrSpec{Name: "source_path", Type: cty.String, Required: false},
		"firmware":                     &hcldec.AttrSpec{Name: "firmware", Type: cty.String, Required: false},
		"skip_set_template":            &hcldec.AttrSpec{Name: "skip_set_template", Type: cty.Bool, Required: false},
		"installation_done_file":       &hcldec.AttrSpec{Name: "installation_done_file", Type: cty.String, Required: false},
		"installation_done_timeout": 		&hcldec.AttrSpec{Name: "installation_done_timeout", Type: cty.String, Required: false},
	}
	return s
}
