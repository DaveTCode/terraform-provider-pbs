package pbsclient

import (
	"fmt"
	"regexp"
	"strings"
)

type PbsResource struct {
	Name string
	Type string
	Flag *string
}

var (
	resourceNameRegex      = regexp.MustCompile(`Resource\s+(\w+)`)
	resourceAttributeRegex = regexp.MustCompile(`(\w+)\s+=\s+(.+)`)
)

func parseResourceOutput(output []byte) ([]PbsResource, error) {
	var currentResource PbsResource
	var resources []PbsResource
	for line := range strings.SplitSeq(string(output), "\n") {
		if resourceNameRegex.MatchString(line) {
			if currentResource.Name != "" { // Is there a resource currently being processed? If so add it to the completed list
				resources = append(resources, currentResource)
			}

			currentResource = PbsResource{
				Name: resourceNameRegex.FindStringSubmatch(line)[1],
			}
		} else if resourceAttributeRegex.MatchString(line) {
			subMatch := resourceAttributeRegex.FindStringSubmatch(line)
			attribute := subMatch[1]
			value := subMatch[2]

			switch strings.ToLower(attribute) {
			case "type":
				currentResource.Type = value
			case "flag":
				currentResource.Flag = &value
			default:
				// TODO - What to do with attributes we don't recognise?
			}
		}
	}

	if currentResource.Name != "" {
		resources = append(resources, currentResource)
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
