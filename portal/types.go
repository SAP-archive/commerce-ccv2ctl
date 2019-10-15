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

// .../v1/deploymentmodes
//{"deploymentMode":["ROLLING_UPDATE","RECREATE"],"dataMigrationMode":["NONE","UPDATE","INITIALIZE"]}
var AllowedDeploymentModes = map[string]struct{}{
	"ROLLING_UPDATE": {},
	"RECREATE":       {},
}
var AllowedMigrationModes = map[string]struct{}{
	"NONE":       {},
	"UPDATE":     {},
	"INITIALIZE": {},
}

//{"environmentCode":"d8","databaseUpdateMode":"NONE","strategy":"RECREATE","customerReleaseCode":"20180912.2","applicationCode":"commerce-cloud"}
type Deployment struct {
	EnvironmentCode     string `json:"environmentCode"`
	DatabaseUpdateMode  string `json:"databaseUpdateMode"`
	Strategy            string `json:"strategy"`
	CustomerReleaseCode string `json:"customerReleaseCode"`
	ApplicationCode     string `json:"applicationCode"`
}

func NewDeployment(environment, migrationMode, updateMode, releaseCode string) Deployment {
	d := Deployment{}
	d.EnvironmentCode = environment
	d.DatabaseUpdateMode = migrationMode
	d.Strategy = updateMode
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
