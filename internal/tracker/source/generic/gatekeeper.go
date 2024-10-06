package generic

// Based on: https://github.com/open-policy-agent/gatekeeper/blob/a1f01f40ed89db0ae0cb0b84dcc498b76c6fa448/pkg/gator/verify/suite.go#L8
type GKSuite struct {
	Tests []GKTest `yaml:"tests"`
}

// Based on: https://github.com/open-policy-agent/gatekeeper/blob/a1f01f40ed89db0ae0cb0b84dcc498b76c6fa448/pkg/gator/verify/suite.go#L27
type GKTest struct {
	Name       string    `yaml:"name"`
	Constraint string    `yaml:"constraint"`
	Cases      []*GKCase `yaml:"cases"`
}

// Based on: https://github.com/open-policy-agent/gatekeeper/blob/a1f01f40ed89db0ae0cb0b84dcc498b76c6fa448/pkg/gator/verify/suite.go#L53
type GKCase struct {
	Name   string `yaml:"name"`
	Object string `yaml:"object"`
}

// GKExample represents an example in a Gatekeeper policy.
type GKExample struct {
	Name  string           `json:"name"`
	Cases []*GKExampleCase `json:"cases"`
}

// GKExampleCase represents an example's case in a Gatekeeper policy.
type GKExampleCase struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Content string `json:"content"`
}

type ComposeFile struct {
	Version  string             `json:"version"`
	Name  string             `json:"name"`
	Services map[string]Service  `json:"services"`
	Networks map[string]Network  `json:"networks,omitempty"`
	Volumes  map[string]Volume   `json:"volumes,omitempty"`
}

// Service represents a service in Docker Compose.
type Service struct {
	Image       string             `json:"image,omitempty"`
	ContainerName string           `json:"container_name,omitempty"`
	Ports       []string           `json:"ports,omitempty"`
	Volumes     []string           `json:"volumes,omitempty"`
	Networks    []string           `json:"networks,omitempty"`
	Environment map[string]string  `json:"environment,omitempty"`
}

// Network represents a network configuration in Docker Compose.
type Network struct {
	Driver string `json:"driver,omitempty"`
}

// Volume represents a volume configuration in Docker Compose.
type Volume struct {
	Driver string `json:"driver,omitempty"`
}