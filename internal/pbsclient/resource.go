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
	out, err := c.runCommand("/opt/pbs/bin/qmgr -c 'list resource @default'")
	if err != nil {
		return nil, err
	}

	return parseResourceOutput(out)
}

func (c *PbsClient) CreateResource(newResource PbsResource) (PbsResource, error) {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'create resource %s type=%s'", newResource.Name, newResource.Type)
	_, err := c.runCommand(cmd)
	if err != nil {
		return PbsResource{}, err
	}

	if newResource.Flag != nil {
		cmd = fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s flag=%s'", newResource.Name, *newResource.Flag)
		_, err = c.runCommand(cmd)
		if err != nil {
			return PbsResource{}, err
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
		_, err = c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s type=%s'", r.Name, r.Type))
		if err != nil {
			return PbsResource{}, err
		}
	}

	if oldResource.Flag != nil && r.Flag == nil {
		_, err = c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'unset resource %s flag'", r.Name))
		if err != nil {
			return PbsResource{}, err
		}
	} else {
		_, err = c.runCommand(fmt.Sprintf("/opt/pbs/bin/qmgr -c 'set resource %s flag=%s'", r.Name, *r.Flag))
		if err != nil {
			return PbsResource{}, err
		}
	}

	return c.GetResource(r.Name)
}

func (c *PbsClient) DeleteResource(name string) error {
	cmd := fmt.Sprintf("/opt/pbs/bin/qmgr -c 'delete resource %s'", name)
	_, err := c.runCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}
