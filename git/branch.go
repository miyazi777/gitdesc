package git

import (
	"regexp"
)

type Branch interface {
	DescriptionMap() (map[string]string, error)
	Description() (string, error)
	SetDescription(desc string) error
	Page() (string, error)
	SetPage(desc string) error
}

type BranchImpl struct {
	Git Git
}

func (b *BranchImpl) DescriptionMap() (map[string]string, error) {
	configList, err := b.Git.GetConfigList()
	if err != nil {
		return nil, err
	}

	descMap := buildDescriptionMap(configList)
	return descMap, nil
}

func buildDescriptionMap(configList []string) map[string]string {
	descLineReg := regexp.MustCompile(`^branch.*description=`)
	descMap := make(map[string]string)
	for _, configLine := range configList {
		if descLineReg.MatchString(configLine) {
			desc := extractDescription(configLine)
			branchName := extractBranchName(configLine)
			descMap[branchName] = desc
		}
	}

	return descMap
}

func extractDescription(line string) string {
	descReg := regexp.MustCompile(`^branch.*description=`)
	return descReg.ReplaceAllString(line, "")
}

func extractBranchName(line string) string {
	reg := regexp.MustCompile(`(branch\.|\.description|=.+)`)
	return reg.ReplaceAllString(line, "")
}

func (b *BranchImpl) Description() (string, error) {
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildDescriptionKey(branchName)
	description, err := b.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b *BranchImpl) SetDescription(desc string) error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildDescriptionKey(branchName)
	err = b.Git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildDescriptionKey(branchName string) string {
	return "branch." + branchName + ".description"
}

func (b *BranchImpl) Page() (string, error) {
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := buildPageKey(branchName)
	description, err := b.Git.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return description, nil
}

func (b *BranchImpl) SetPage(desc string) error {
	var err error
	branchName, err := b.Git.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := buildPageKey(branchName)
	err = b.Git.SetConfigValue(key, desc)
	if err != nil {
		return err
	}

	return nil
}

func buildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}