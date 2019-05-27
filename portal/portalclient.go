package portal

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	persistentjar "github.com/monster1025/persistent-cookiejar"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

var PORTAL, _ = url.Parse("https://portal.commerce.ondemand.com/")

type Client struct {
	hc           *http.Client
	subscription string
	jar          *persistentjar.Jar
}

func fail(e error) {
	if e != nil {
		panic(e)
	}
}
func failWithBodyDump(e error, resp *http.Response) {
	if e != nil {
		dumpBody(resp.Body)
		panic(e)
	}
}

func chainingResolver(base *url.URL) func(action string) *url.URL {
	b := base
	return func(action string) *url.URL {
		a, err := url.Parse(action)
		fail(err)
		a = b.ResolveReference(a)
		b = a
		return a
	}
}

func getAttr(attr []html.Attribute, name string) string {
	for _, a := range attr {
		if a.Key == name {
			return a.Val
		}
	}
	return ""
}

func parseForm(r io.Reader) (string, url.Values, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return "", nil, err
	}

	var action string
	params := make(url.Values)

	var recurse func(node *html.Node) error
	recurse = func(n *html.Node) error {
		if n.Type == html.ElementNode {
			if n.Data == "form" {
				if action != "" {
					return errors.New("found more than one <form>")
				}
				action = getAttr(n.Attr, "action")
			}
			if n.Data == "input" && getAttr(n.Attr, "type") == "hidden" {
				params[getAttr(n.Attr, "name")] = []string{getAttr(n.Attr, "value")}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := recurse(c)
			if err != nil {
				return err
			}
		}
		return nil
	}
	err = recurse(doc)
	return action, params, err
}

func parseBody(b io.ReadCloser) (string, url.Values) {
	defer b.Close()
	action, params, err := parseForm(b)
	fail(err)
	return action, params
}

func NewClient(subscription string, certPEMBlock, keyPEMBlock []byte, jarfile string) Client {

	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	fail(err)
	tlsconf := tls.Config{Certificates: []tls.Certificate{cert}}
	tlsconf.BuildNameToCertificate()

	tr := http.Transport{
		TLSClientConfig: &tlsconf,
	}

	jar, err := persistentjar.New(&persistentjar.Options{Filename: jarfile})
	fail(err)

	c := &http.Client{Jar: jar, Transport: &tr}

	resolve := chainingResolver(PORTAL)

	//simulate frontend SSO
	resp, err := c.Get("https://portal.commerce.ondemand.com/")
	fail(err)

	if needLogin(resp) {
		fmt.Fprintf(os.Stderr, "Session expired, logging in...\n")
		action, params := parseBody(resp.Body)
		actionUrl := resolve(action)
		resp, err = c.PostForm(actionUrl.String(), params)
		fail(err)

		action, params = parseBody(resp.Body)
		actionUrl = resolve(action)
		resp, err = c.PostForm(actionUrl.String(), params)
		fail(err)

		action, params = parseBody(resp.Body)
		actionUrl = resolve(action)
		resp, err = c.PostForm(actionUrl.String(), params)
		fail(err)

		action, params = parseBody(resp.Body)
		actionUrl = resolve(action)
		resp, err = c.PostForm(actionUrl.String(), params)
		fail(err)
	}

	return Client{hc: c, subscription: subscription, jar: jar}
}

func needLogin(resp *http.Response) bool {

	if resp.StatusCode >= 300 {
		failWithBodyDump(fmt.Errorf("Login: Expected HTTP 2XX status, got %d", resp.StatusCode), resp)
	}
	return resp.Header.Get("com.sap.cloud.security.login") != ""
}

func (pc *Client) SaveCookieJar() error {
	return pc.jar.Save()
}

func (pc *Client) getOrFail(u *url.URL) *http.Response {

	resp, err := pc.hc.Get(u.String())
	fail(err)
	if resp.StatusCode >= 300 {
		failWithBodyDump(fmt.Errorf("GET %s: Expected HTTP 2XX status, got %d", u.String(), resp.StatusCode), resp)
	}
	return resp
}

func (pc *Client) postJSONorFail(u *url.URL, payload interface{}) *http.Response {
	j, err := json.Marshal(payload)
	request, err := http.NewRequest("POST", u.String(), bytes.NewReader(j))
	fail(err)
	request.Header.Add("content-type", "application/json")
	for _, c := range pc.hc.Jar.Cookies(u) {
		if c.Name == "XSRF-TOKEN" {
			request.Header.Add("x-xsrf-token", c.Value)
		}
	}
	resp, err := pc.hc.Do(request)
	fail(err)
	if resp.StatusCode >= 300 {
		failWithBodyDump(errors.New(fmt.Sprintf("POST %s: Expected HTTP 2XX status, got %d", u.String(), resp.StatusCode)), resp)
	}

	return resp
}

func (pc *Client) putJSONorFail(u *url.URL, payload interface{}) *http.Response {
	j, err := json.Marshal(payload)
	request, err := http.NewRequest("PUT", u.String(), bytes.NewReader(j))
	fail(err)
	request.Header.Add("content-type", "application/json")
	for _, c := range pc.hc.Jar.Cookies(u) {
		if c.Name == "XSRF-TOKEN" {
			request.Header.Add("x-xsrf-token", c.Value)
		}
	}
	resp, err := pc.hc.Do(request)
	fail(err)
	if resp.StatusCode >= 300 {
		failWithBodyDump(errors.New(fmt.Sprintf("PUT %s: Expected HTTP 2XX status, got %d", u.String(), resp.StatusCode)), resp)
	}

	return resp
}

func readJson(r io.ReadCloser, j interface{}) {
	defer r.Close()
	dec := json.NewDecoder(r)
	fail(dec.Decode(j))
}

func resolveAPI(a string) *url.URL {
	action, err := url.Parse(a)
	fail(err)
	action = PORTAL.ResolveReference(action)

	return action
}

func dumpBody(r io.ReadCloser) {
	defer r.Close()
	all, err := ioutil.ReadAll(r)
	fail(err)
	fmt.Fprintln(os.Stderr, string(all))
}

const builds = "/v1/subscriptions/%s/builds/"
const buildLogs = builds + "%s/logs/"
const allBuilds = "/v1/subscriptions/%s/applications/commerce-cloud/builds/"
const deployment = "/v1/subscriptions/%s/build"
const running = "/v1/subscriptions/%s/environments/%s/runningdeployments/"
const deployments = "/v1/subscriptions/%s/environments/%s/deployments"
const passwords = "/v1/subscriptions/%s/environments/%s/serviceconfiguration/hcs_admin/property/initialpassword"
const properties = "/v1/subscriptions/%s/environments/%s/serviceconfiguration/%s/property/customer-properties"

func (pc *Client) GetAllBuilds() (meta []BuildMeta) {

	action := resolveAPI(fmt.Sprintf(allBuilds, pc.subscription) + "?page=0&pageSize=20&limit=20&sort=desc(buildStartTime)")

	var page BuildPage
	resp := pc.getOrFail(action)
	readJson(resp.Body, &page)

	return page.Content
}

func (pc *Client) GetBuild(code string) (meta BuildMeta) {

	action := resolveAPI(fmt.Sprintf(builds, pc.subscription) + code)
	resp := pc.getOrFail(action)
	readJson(resp.Body, &meta)

	return meta
}

func (pc *Client) CreateBuild(name, branch string) (newBuild BuildMeta) {

	action := resolveAPI(fmt.Sprintf(builds, pc.subscription))

	build := NewBuild(pc.subscription, name, branch)

	pc.postJSONorFail(action, build)
	//unfortunately, the response of the POST request is NOT the build meta data, but sth else.
	//-> wait a bit, fetch all builds and find the matching one
	time.Sleep(1 * time.Second)
	all := pc.GetAllBuilds()
	found := false
	for _, b := range all {
		if b.Name == name && b.Branch == branch {
			if found {
				fmt.Fprintf(os.Stderr, "WARN: more than one build with name=%s and branch=%s found. Using the first one\n", name, branch)
			} else {
				newBuild = b
			}
			found = true
		}
	}
	return newBuild
}

func (pc *Client) GetBuildLogReader(code string) io.ReadCloser {
	logs := resolveAPI(fmt.Sprintf(buildLogs, pc.subscription, code))
	resp := pc.getOrFail(logs)

	//unzip the whole response in-memory, wasn't able to find a streaming unzip for go
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	fail(err)
	zipReader, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	fail(err)
	log, err := zipReader.File[0].Open()
	fail(err)

	return log
}

func (pc *Client) CreateDeployment(environment, mode, release string) (r RunningDeployment) {

	api := resolveAPI(fmt.Sprintf(deployment, pc.subscription))
	d := NewDeployment(environment, mode, release)
	resp := pc.postJSONorFail(api, d)
	readJson(resp.Body, &r)

	return r
}

func (pc *Client) GetRunningDeployments(environment string) (r []RunningDeployment) {

	api := resolveAPI(fmt.Sprintf(running, pc.subscription, environment))
	resp := pc.getOrFail(api)
	readJson(resp.Body, &r)

	return r
}

func (pc *Client) GetDeployments(environment string) (r []RunningDeployment) {

	api := resolveAPI(fmt.Sprintf(deployments, pc.subscription, environment))
	resp := pc.getOrFail(api)
	readJson(resp.Body, &r)

	return r
}

func (pc *Client) GetInitialPasswords(environment string) (p InitialPasswords) {

	api := resolveAPI(fmt.Sprintf(passwords, pc.subscription, environment))
	resp := pc.getOrFail(api)
	readJson(resp.Body, &p)

	return p
}

func (pc *Client) SetCustomerProperties(environment, aspect, filename string) (p Properties) {

	var value string

	if filename == "-" {
		data, err := ioutil.ReadAll(os.Stdin)
		fail(err)
		value = string(data)
	} else {
		f, err := os.Open(filename)
		fail(err)
		data, err := ioutil.ReadAll(f)
		fail(err)
		value = string(data)
		f.Close()
	}

	api := resolveAPI(fmt.Sprintf(properties, pc.subscription, environment, aspect))
	np := NewProperties("customer-properties", value)
	resp := pc.putJSONorFail(api, np)
	readJson(resp.Body, &p)

	return p
}

func (pc *Client) GetCustomerProperties(environment, aspect string) (p Properties) {

	api := resolveAPI(fmt.Sprintf(properties, pc.subscription, environment, aspect))
	resp := pc.getOrFail(api)
	readJson(resp.Body, &p)

	return p
}
