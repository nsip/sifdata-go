package environment

import (
    "encoding/xml"
    "log"
)

type EnvironmentRequest struct {
	XMLName              xml.Name `xml:"environment"`
	Text                 string   `xml:",chardata"`
	Xmlns                string   `xml:"xmlns,attr"`
	SolutionId           string   `xml:"solutionId"`
	AuthenticationMethod string   `xml:"authenticationMethod"`
	UserToken            string   `xml:"userToken"`
	ConsumerName         string   `xml:"consumerName"`
	ApplicationInfo      struct {
		Text                           string `xml:",chardata"`
		ApplicationKey                 string `xml:"applicationKey"`
		SupportedInfrastructureVersion string `xml:"supportedInfrastructureVersion"`
		DataModelNamespace             string `xml:"dataModelNamespace"`
		Transport                      string `xml:"transport"`
		ApplicationProduct             struct {
			Text           string `xml:",chardata"`
			VendorName     string `xml:"vendorName"`
			ProductName    string `xml:"productName"`
			ProductVersion string `xml:"productVersion"`
		} `xml:"applicationProduct"`
	} `xml:"applicationInfo"`
}

type EnvironmentResponse struct {
	XMLName      xml.Name `xml:"environment"`
	Text         string   `xml:",chardata"`
	Type         string   `xml:"type,attr"`
	ID           string   `xml:"id,attr"`
	Xmlns        string   `xml:"xmlns,attr"`
	Fingerprint  string   `xml:"fingerprint"`
	SessionToken string   `xml:"sessionToken"`
	SolutionId   string   `xml:"solutionId"`
	DefaultZone  struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		Description string `xml:"description"`
	} `xml:"defaultZone"`
	AuthenticationMethod string `xml:"authenticationMethod"`
	UserToken            string `xml:"userToken"`
	ConsumerName         string `xml:"consumerName"`
	ApplicationInfo      struct {
		Text                           string `xml:",chardata"`
		ApplicationKey                 string `xml:"applicationKey"`
		SupportedInfrastructureVersion string `xml:"supportedInfrastructureVersion"`
		DataModelNamespace             string `xml:"dataModelNamespace"`
		Transport                      string `xml:"transport"`
		ApplicationProduct             struct {
			Text           string `xml:",chardata"`
			VendorName     string `xml:"vendorName"`
			ProductName    string `xml:"productName"`
			ProductVersion string `xml:"productVersion"`
		} `xml:"applicationProduct"`
	} `xml:"applicationInfo"`
	InfrastructureServices struct {
		Text                  string `xml:",chardata"`
		InfrastructureService []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"infrastructureService"`
	} `xml:"infrastructureServices"`
	ProvisionedZones struct {
		Text            string `xml:",chardata"`
		ProvisionedZone struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Services struct {
				Text    string `xml:",chardata"`
				Service []struct {
					Text      string `xml:",chardata"`
					Name      string `xml:"name,attr"`
					ContextId string `xml:"contextId,attr"`
					Type      string `xml:"type,attr"`
					Rights    struct {
						Text  string `xml:",chardata"`
						Right []struct {
							Text string `xml:",chardata"`
							Type string `xml:"type,attr"`
						} `xml:"right"`
					} `xml:"rights"`
				} `xml:"service"`
			} `xml:"services"`
		} `xml:"provisionedZone"`
	} `xml:"provisionedZones"`
}

// Return an EnvironmentRequest
func NewEnvironmentRequest() EnvironmentRequest {
    e:= EnvironmentRequest{}
    e.SolutionId = "HITS"
    e.AuthenticationMethod = "Basic"
    return e
}

func ExampleNewEnvironmentRequest() {
}

func NewEnvironmentResponse() EnvironmentResponse {
    e:= EnvironmentResponse{}
    return e
}

func ParseEnvironmentResponse(in []byte) EnvironmentResponse {
    e := NewEnvironmentResponse()
	if err := xml.Unmarshal(in, &e); err != nil {
		log.Fatal(err)
	}
    return e
}
