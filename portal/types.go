package portal

const CLOUD_APPLICATION = "commerce-cloud"

type Build struct {
	Name             string `json:"name"`
	Branch           string `json:"branch"`
	ApplicationCode  string `json:"applicationCode"`
	SubscriptionCode string `json:"subscriptionCode"`
}

type BuildResponse struct {
	SubscriptionCode string
	Code             string
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
	SubscriptionCode             string
	ApplicationCode              string
	ApplicationDefinitionVersion string
	Code                         string
	Name                         string
	CreatedBy                    string
	Status                       string
	Branch                       string
	BuildVersion                 string
	BuildStartTimeStamp          string
	BuildEndTimeStamp            string
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
	Count int
	Value []BuildMeta
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
	EnvironmentCode    string `json:"environmentCode"`
	DatabaseUpdateMode string `json:"databaseUpdateMode"`
	Strategy           string `json:"strategy"`
	BuildCode          string `json:"buildCode"`
}

func NewDeployment(environment, migrationMode, updateMode, releaseCode string) Deployment {
	d := Deployment{}
	d.EnvironmentCode = environment
	d.DatabaseUpdateMode = migrationMode
	d.Strategy = updateMode
	d.BuildCode = releaseCode

	return d
}

type DeploymentResponse struct {
	SubscriptionCode string
	Code             string
}

type DeploymentMeta struct {
	Code                          string
	SubscriptionCode              string
	CreatedBy                     string
	CreatedTimestamp              string
	BuildCode                     string
	EnvironmentCode               string
	DatabaseUpdateMode            string
	Strategy                      string
	ScheduledTimestamp            string
	DeployedTimestamp             string
	FailedTimestamp               string
	UndeployedTimestamp           string
	Status                        string
	CanceledBy                    string
	CanceledTimestamp             string
	CancellationFinishedTimestamp string
	CancellationFailed            string
	Cancelation                   string
}

type DeploymentPage struct {
	Value []DeploymentMeta
	Count int
}

type InitialPasswordEntry struct {
	Uid      string
	Password string
}

type InitialPasswords struct {
	Key   string
	Value []InitialPasswordEntry
}
