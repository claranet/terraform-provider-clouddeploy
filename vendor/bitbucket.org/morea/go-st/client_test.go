package ghost

import (
	"crypto/rand"
	"fmt"
	"testing"
)

// Global vars reused between tests
var c = NewClient("https://demo.ghost.morea.fr", "demo", "***REMOVED***")
var eveMetadata EveItemMetadata

func pseudo_uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err == nil {
		uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	}
	return
}

func TestClientGetApps(t *testing.T) {
	fmt.Println("Testing Ghost client get all apps")
	apps, err := c.GetApps()
	if err == nil {
		fmt.Println("All apps retrieved:")
		fmt.Println(apps)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

func TestClientCreateApp(t *testing.T) {
	fmt.Println("Testing Ghost client create app")
	newapp := App{
		Name: "test-" + pseudo_uuid(),
		Env:  "test",
		Role: "webfront",

		Region:       "eu-west-1",
		InstanceType: "t2.nano",
		VpcID:        "vpc-123456",

		LogNotifications: []string{"ghost-demo@domain.com"},

		BuildInfos: BuildInfos{
			SourceAmi:   "ami-123456",
			SshUsername: "admin",
			SubnetID:    "subnet-123456",
		},

		EnvironmentInfos: EnvironmentInfos{
			InstanceProfile: "test-instance-profile",
			KeyName:         "test-key-name",
			OptionalVolumes: []OptionalVolume{},
			RootBlockDevice: RootBlockDevice{Name: "/dev/xvda"},
			SecurityGroups:  []string{"sg-123456"},
			SubnetIDs:       []string{"subnet-123456"},
		},

		Features: []Feature{
			{
				Name:    "nginx",
				Version: "1.10",
			},
		},
		Modules: []Module{
			{
				Name:    "testmod",
				GitRepo: "git@bitbucket.org/morea/testmod",
				Scope:   "system",
				Path:    "/tmp/path",
			},
		},
	}

	var err error
	eveMetadata, err = c.CreateApp(newapp)
	if err == nil {
		fmt.Println("App created: " + eveMetadata.ID)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

func TestClientGetApp(t *testing.T) {
	fmt.Println("Testing Ghost client get single app")
	app, err := c.GetApp(eveMetadata.ID)
	if err == nil {
		fmt.Println("App retrieved: " + app.ID)
		fmt.Println(app)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

func TestClientUpdateApp(t *testing.T) {
	fmt.Println("Testing Ghost client update app")

	app, err := c.GetApp(eveMetadata.ID)

	// Remove read only fields
	app.Etag = nil
	app.Links = nil
	app.Created = nil
	app.Updated = nil
	app.Version = nil
	app.LatestVersion = nil
	app.Modules[0].Initialized = nil

	// Add module
	app.Modules = append(app.Modules, Module{
		Name:    "testmod2",
		GitRepo: "git@bitbucket.org/morea/testmod2",
		Scope:   "system",
		Path:    "/tmp/path2",
	})

	eveMetadata, err = c.UpdateApp(&app, eveMetadata.ID, *eveMetadata.Etag)
	if err == nil {
		fmt.Println("App updated: " + eveMetadata.ID)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}

	app, err = c.GetApp(eveMetadata.ID)
	if len(app.Modules) != 2 {
		t.Fatalf("Assertion error: added module is missing")
	}
}

func TestClientDeleteApp(t *testing.T) {
	fmt.Println("Testing Ghost client delete app")

	err := c.DeleteApp(eveMetadata.ID, *eveMetadata.Etag)
	if err == nil {
		fmt.Println("App deleted: " + eveMetadata.ID)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}
