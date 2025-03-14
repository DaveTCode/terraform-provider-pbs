package pbsclient

import (
	"fmt"
	"strings"
)

type PbsResource struct {
	Name string
	Type string
	Flag *string
}

func parseResourceOutput(output []byte) ([]PbsResource, error) {
	parsedOutput := parseGenericQmgrOutput(string(output))
	var resources []PbsResource

	for _, r := range parsedOutput {
		if r.objType == "Resource" {
			current := PbsResource{
				Name: r.name,
			}

			for k, v := range r.attributes {
				if s, ok := v.(string); ok {
					switch strings.ToLower(k) {
					case "type":
						current.Type = s
					case "flag":
						current.Flag = &s
					}
				}
			}

			resources = append(resources, current)
		}
	}

	return resources, nil
}

func (c *PbsClient) GetResource(name string) (PbsResource, error) {
	allResources, err := c.GetResources()
	if err != nil {
		return PbsResource{}, err
	}

	for _, r := range allResources {
		if r.Name == name {
			return r, nil
		}
	}

	return PbsResource{}, nil
}

func (c *PbsClient) GetResources() ([]PbsResource, error) {
	out, errOutput, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list resource @default'")
	if err != nil {
		return nil, fmt.Errorf("%s %s", err, errOutput)
	}

	return parseResourceOutput(out)
}

func (c *PbsClient) CreateResource(newResource PbsResource) (PbsResource, error) {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create resource %s type=%s'", newResource.Name, newResource.Type)
	_, errOutput, err := c.runCommand(cmd)
	if err != nil {
		return PbsResource{}, fmt.Errorf("%s %s", err, errOutput)
	}

	if newResource.Flag != nil {
		cmd = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s flag=%s'", newResource.Name, *newResource.Flag)
		_, errOutput, err := c.runCommand(cmd)
		if err != nil {
			return PbsResource{}, fmt.Errorf("%s %s", err, errOutput)
		}
	}

	return c.GetResource(newResource.Name)
}

func (c *PbsClient) UpdateResource(r PbsResource) (PbsResource, error) {
	oldResource, err := c.GetResource(r.Name)
	if err != nil {
		return PbsResource{}, err
	}

	if oldResource.Type != r.Type {
		_, errOutput, err := c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s type=%s'", r.Name, r.Type))
		if err != nil {
			return PbsResource{}, fmt.Errorf("%s %s", err, errOutput)
		}
	}

	if oldResource.Flag != nil && r.Flag == nil {
		_, errOutput, err := c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset resource %s flag'", r.Name))
		if err != nil {
			return PbsResource{}, fmt.Errorf("%s %s", err, errOutput)
		}
	} else {
		_, errOutput, err := c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s flag=%s'", r.Name, *r.Flag))
		if err != nil {
			return PbsResource{}, fmt.Errorf("%s %s", err, errOutput)
		}
	}

	return c.GetResource(r.Name)
}

func (c *PbsClient) DeleteResource(name string) error {
	_, errOutput, err := c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete resource %s'", name))
	if err != nil {
		return fmt.Errorf("%s %s", err, errOutput)
	}

	return nil
}
