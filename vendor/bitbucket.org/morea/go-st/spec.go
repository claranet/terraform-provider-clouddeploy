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
	SourceAmi   string `json:"source_ami"`
	SshUsername string `json:"ssh_username"`
	SubnetID    string `json:"subnet_id"`
}

// Ghost App's environment_infos structs
type OptionalVolume struct {
	InstanceProfile string `json:"device_name"`
	KeyName         string `json:"volume_type"`
	VolumeSize      int    `json:"volume_size"`
	IOPS            int    `json:"iops"`
}

type RootBlockDevice struct {
	Size int    `json:"size"`
	Name string `json:"name"`
}

type EnvironmentInfos struct {
	InstanceProfile string           `json:"instance_profile"`
	KeyName         string           `json:"key_name"`
	OptionalVolumes []OptionalVolume `json:"optional_volumes"`
	RootBlockDevice RootBlockDevice  `json:"root_block_device"`
	SecurityGroups  []string         `json:"security_groups"`
	SubnetIDs       []string         `json:"subnet_ids"`
}

// Ghost App's feature struct
type Feature struct {
	Name    string `json:"name"`
	Version string `json:"version"`
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
}

// Ghost App struct
type App struct {
	EveItemMetadata
	User string `json:"user"`

	Name string `json:"name"`
	Env  string `json:"env"`
	Role string `json:"role"`

	Region       string `json:"region"`
	InstanceType string `json:"instance_type"`
	VpcID        string `json:"vpc_id"`

	LogNotifications []string `json:"log_notifications"`

	BuildInfos BuildInfos `json:"build_infos"`

	EnvironmentInfos EnvironmentInfos `json:"environment_infos"`

	Features []Feature `json:"features"`

	Modules []Module `json:"modules"`
}

// Ghost Apps collection
type Apps struct {
	EveCollectionMetadata
	Items []App `json:"_items"`
}
