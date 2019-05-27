package portal

const CLOUD_APPLICATION = "commerce-cloud"

type Build struct {
	Name             string `json:"name"`
	Branch           string `json:"branch"`
	ApplicationCode  string `json:"applicationCode"`
	SubscriptionCode string `json:"subscriptionCode"`
}

type Properties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewProperties(key, value string) Properties {
	p := Properties{}
	p.Key = key
	p.Value = value

	return p
}

type BuildMeta struct {
	Code                         string
	Name                         string
	CreatedBy                    string
	Status                       string
	Branch                       string
	BuildVersion                 string
	BuildStartTime               string
	BuildEndTime                 string
	ApplicationName              string
	ApplicationCode              string
	ApplicationDefinitionVersion string
	CreatedOn                    string
	ProgressData                 interface{}
	Properties                   []Properties
}

type PageParams struct {
	TotalPages       int
	TotalElements    int
	Number           int
	Size             int
	NumberOfElements int
	First            bool
	Last             bool
	Sort             string
}

type BuildPage struct {
	Params  PageParams
	Content []BuildMeta
}

func NewBuild(subscription, name, branch string) Build {
	b := Build{}
	b.Name = name
	b.Branch = branch
	b.ApplicationCode = CLOUD_APPLICATION
	b.SubscriptionCode = subscription

	return b
}

type Deployment struct {
	EnvironmentCode     string `json:"environmentCode"`
	Mode                string `json:"mode"`
	CustomerReleaseCode string `json:"customerReleaseCode"`
	ApplicationCode     string `json:"applicationCode"`
}

func NewDeployment(environment, mode, releaseCode string) Deployment {
	d := Deployment{}
	d.EnvironmentCode = environment
	d.Mode = mode
	d.CustomerReleaseCode = releaseCode
	d.ApplicationCode = CLOUD_APPLICATION

	return d
}

type RunningDeployment struct {
	CustomerReleaseCode          string
	EnvironmentCode              string
	Mode                         string
	ApplicationCode              string
	ScheduledDate                string
	RollingDeployment            bool
	ApplicationDefinitionVersion string
	Status                       string
}

type InitialPasswordEntry struct {
	Uid      string
	Password string
}

type InitialPasswords struct {
	Key   string
	Value []InitialPasswordEntry
}
