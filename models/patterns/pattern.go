package patterns

import (
	"github.com/Masterminds/semver/v3"
	"github.com/layer5io/meshkit/utils"
	"github.com/meshery/schemas/models/v1alpha2"
	"gopkg.in/yaml.v2"
)

func GetNextVersion(p *v1alpha2.PatternFile) (string, error) {
	// Existing patterns do not have version hence when trying to assign next version for such patterns, it will fail the validation.
	// Hence, if version is not present, start versioning for those afresh.
	if p.Version == "" {
		AssignVersion(p)
		return p.Version, nil
	}
	version, err := semver.NewVersion(p.Version)
	if err != nil {
		return "", err
		// return ErrInvalidVersion(err) // send meshkit error
	}

	nextVersion := version.IncPatch().String()
	return nextVersion, nil
}

func AssignVersion(p *v1alpha2.PatternFile) {
	p.Version = semver.New(0, 0, 1, "", "").String()
}

func GetPatternFormat(patternFile string) (*v1alpha2.PatternFile, error) {
	pattern := v1alpha2.PatternFile{}
	err := yaml.Unmarshal([]byte(patternFile), &pattern)
	if err != nil {
		err = utils.ErrDecodeYaml(err)
		return nil, err
	}
	return &pattern, nil
}
