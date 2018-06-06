package ghost

// eve metadata
type Link struct {
	Href  string `json:"href,omitempty"`
	Rel   string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
}

type EveItemMetadata struct {
	ID string `json:"_id,omitempty"`

	Etag *string `json:"_etag,omitempty"`

	Created *string `json:"_created,omitempty"`
	Updated *string `json:"_updated,omitempty"`

	Version       *int64 `json:"_version,omitempty"`
	LatestVersion *int64 `json:"_latest_version,omitempty"`

	Links *struct {
		Self Link `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type EveCollectionMetadata struct {
	Links *struct {
		Parent Link `json:"parent,omitempty"`
		Self   Link `json:"self,omitempty"`
	} `json:"_links,omitempty"`

	Meta *struct {
		MaxResults int64 `json:"max_results,omitempty"`
		Page       int64 `json:"page,omitempty"`
		Total      int64 `json:"total,omitempty"`
	} `json:"_meta,omitempty"`
}

// Ghost App's build_infos struct
type BuildInfos struct {
	SourceAmi            string `json:"source_ami"`
	SshUsername          string `json:"ssh_username"`
	SubnetID             string `json:"subnet_id"`
	AmiName              string `json:"ami_name,omitempty"`
	ContainerImage       string `json:"container_image,omitempty"`
	SourceContainerImage string `json:"source_container_image"`
}

// Ghost App's environment_infos structs
type OptionalVolume struct {
	DeviceName                string `json:"device_name"`
	VolumeType                string `json:"volume_type"`
	VolumeSize                int    `json:"volume_size"`
	Iops                      int    `json:"iops"`
	LaunchBlockDeviceMappings bool   `json:"launch_block_device_mappings"`
}

type RootBlockDevice struct {
	Size int    `json:"size"`
	Name string `json:"name"`
}

type InstanceTag struct {
	TagName  string `json:"tag_name"`
	TagValue string `json:"tag_value"`
}

type EnvironmentInfos struct {
	InstanceProfile string            `json:"instance_profile"`
	KeyName         string            `json:"key_name"`
	SecurityGroups  []string          `json:"security_groups"`
	SubnetIDs       []string          `json:"subnet_ids"`
	OptionalVolumes *[]OptionalVolume `json:"optional_volumes"`
	RootBlockDevice *RootBlockDevice  `json:"root_block_device"`
	PublicIpAddress bool              `json:"public_ip_address"`
	InstanceTags    *[]InstanceTag    `json:"instance_tags"`
}

// Ghost App's feature struct
type Feature struct {
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Provisioner string      `json:"provisioner"`
	Parameters  interface{} `json:"parameters"`
}

// Ghost App's module struct
type Module struct {
	Initialized *bool `json:"initialized,omitempty"`

	Name    string `json:"name"`
	GitRepo string `json:"git_repo"`
	Scope   string `json:"scope"`
	Path    string `json:"path"`

	UID int `json:"uid"`
	GID int `json:"gid"`

	// Scripts
	BuildPack      string `json:"build_pack"`
	PreDeploy      string `json:"pre_deploy"`
	PostDeploy     string `json:"post_deploy"`
	AfterAllDeploy string `json:"after_all_deploy"`
	LastDeployment string `json:"last_deployment,omitempty"`
}

type LifecycleHooks struct {
	PreBuildimage  string `json:"pre_buildimage"`
	PostBuildimage string `json:"post_buildimage"`
	PreBootstrap   string `json:"pre_bootstrap"`
	PostBootstrap  string `json:"post_bootstrap"`
}

type EnvironmentVariable struct {
	Key   string `json:"var_key"`
	Value string `json:"var_value"`
}

type Autoscale struct {
	Min           int    `json:"min"`
	Max           int    `json:"max"`
	EnableMetrics bool   `json:"enable_metrics"`
	Name          string `json:"name"`
}

type SafeDeployment struct {
	WaitBeforeDeploy int    `json:"wait_before_deploy"`
	WaitAfterDeploy  int    `json:"wait_after_deploy"`
	LoadBalancerType string `json:"load_balancer_type"`
	AppTagValue      string `json:"app_tag_value"`
	HaBackend        string `json:"ha_backend"`
	ApiPort          int    `json:"api_port"`
}

type PendingChange struct {
	Field   string `json:"field"`
	Updated string `json:"updated"`
	User    string `json:"user"`
}

// Ghost App struct
type App struct {
	EveItemMetadata
	User string `json:"user"`

	Name        string `json:"name"`
	Env         string `json:"env"`
	Role        string `json:"role"`
	Description string `json:"description"`

	Region             string `json:"region"`
	InstanceType       string `json:"instance_type"`
	InstanceMonitoring bool   `json:"instance_monitoring"`
	VpcID              string `json:"vpc_id"`

	LifecycleHooks *LifecycleHooks `json:"lifecycle_hooks"`

	LogNotifications []string `json:"log_notifications"`

	BuildInfos *BuildInfos `json:"build_infos"`

	EnvironmentInfos *EnvironmentInfos `json:"environment_infos"`

	EnvironmentVariables *[]EnvironmentVariable `json:"env_vars"`

	Features *[]Feature `json:"features"`

	Modules *[]Module `json:"modules"`

	Autoscale *Autoscale `json:"autoscale"`

	SafeDeployment *SafeDeployment `json:"safe-deployment"`

	PendingChanges *[]PendingChange `json:"pending_changes,omitempty"`
}

// Ghost Apps collection
type Apps struct {
	EveCollectionMetadata
	Items []App `json:"_items"`
}
