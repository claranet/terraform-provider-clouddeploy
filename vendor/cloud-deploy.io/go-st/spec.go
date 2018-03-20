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
	SourceContainerImage string `json:"source_container_image,omitempty"`
}

// Ghost App's environment_infos structs
type OptionalVolume struct {
	DeviceName                string `json:"device_name"`
	VolumeType                string `json:"volume_type"`
	VolumeSize                int    `json:"volume_size"`
	Iops                      int    `json:"iops,omitempty"`
	LaunchBlockDeviceMappings bool   `json:"launch_block_device_mappings,omitempty"`
}

type RootBlockDevice struct {
	Size int    `json:"size,omitempty"`
	Name string `json:"name,omitempty"`
}

type InstanceTag struct {
	TagName  string `json:"tag_name"`
	TagValue string `json:"tag_value"`
}

type EnvironmentInfos struct {
	InstanceProfile string            `json:"instance_profile,omitempty"`
	KeyName         string            `json:"key_name,omitempty"`
	SecurityGroups  []string          `json:"security_groups,omitempty"`
	SubnetIDs       []string          `json:"subnet_ids,omitempty"`
	OptionalVolumes *[]OptionalVolume `json:"optional_volumes,omitempty"`
	RootBlockDevice *RootBlockDevice  `json:"root_block_device,omitempty"`
	PublicIpAddress bool              `json:"public_ip_address"`
	InstanceTags    *[]InstanceTag    `json:"instance_tags,omitempty"`
}

// Ghost App's feature struct
type Feature struct {
	Name        string `json:"name"`
	Version     string `json:"version,omitempty"`
	Provisioner string `json:"provisioner,omitempty"`
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
	BuildPack      string `json:"build_pack,omitempty"`
	PreDeploy      string `json:"pre_deploy,omitempty"`
	PostDeploy     string `json:"post_deploy,omitempty"`
	AfterAllDeploy string `json:"after_all_deploy,omitempty"`
	LastDeployment string `json:"last_deployment,omitempty"`
}

type LifecycleHooks struct {
	PreBuildimage  string `json:"pre_buildimage,omitempty"`
	PostBuildimage string `json:"post_buildimage,omitempty"`
	PreBootstrap   string `json:"pre_bootstrap,omitempty"`
	PostBootstrap  string `json:"post_bootstrap,omitempty"`
}

type EnvironmentVariable struct {
	Key   string `json:"var_key,omitempty"`
	Value string `json:"var_value,omitempty"`
}

type Autoscale struct {
	Min           int    `json:"min"`
	Max           int    `json:"max"`
	EnableMetrics bool   `json:"enable_metrics"`
	Name          string `json:"name,omitempty"`
}

type SafeDeployment struct {
	WaitBeforeDeploy int    `json:"wait_before_deploy"`
	WaitAfterDeploy  int    `json:"wait_after_deploy"`
	LoadBalancerType string `json:"load_balancer_type,omitempty"`
	AppTagValue      string `json:"app_tag_value,omitempty"`
	HaBackend        string `json:"ha_backend,omitempty"`
	ApiPort          string `json:"api_port,omitempty"`
}

type PendingChange struct {
	Field   string `json:"field,omitempty"`
	Updated string `json:"updated,omitempty"`
	User    string `json:"user,omitempty"`
}

// Ghost App struct
type App struct {
	EveItemMetadata
	User string `json:"user"`

	Name string `json:"name"`
	Env  string `json:"env"`
	Role string `json:"role"`

	Region             string `json:"region"`
	InstanceType       string `json:"instance_type"`
	InstanceMonitoring bool   `json:"instance_monitoring"`
	VpcID              string `json:"vpc_id"`

	LifecycleHooks *LifecycleHooks `json:"lifecycle_hooks,omitempty"`

	LogNotifications []string `json:"log_notifications,omitempty"`

	BuildInfos *BuildInfos `json:"build_infos,omitempty"`

	EnvironmentInfos *EnvironmentInfos `json:"environment_infos,omitempty"`

	EnvironmentVariables *[]EnvironmentVariable `json:"env_vars,omitempty"`

	Features *[]Feature `json:"features,omitempty"`

	Modules *[]Module `json:"modules"`

	Autoscale *Autoscale `json:"autoscale,omitempty"`

	SafeDeployment *SafeDeployment `json:"safe-deployment,omitempty"`

	PendingChanges *[]PendingChange `json:"pending_changes,omitempty"`
}

// Ghost Apps collection
type Apps struct {
	EveCollectionMetadata
	Items []App `json:"_items"`
}
