package config

import (
	"github.com/miyazi777/git-desc/git"
)

type Page interface {
	GetPage() (string, error)
	SetPage(page string) error
	DeletePage(branchName string) error
}

type PageImpl struct {
	Command git.Command
}

func (p *PageImpl) GetPage() (string, error) {
	branchName, err := p.Command.GetCurrentBranch()
	if err != nil {
		return "", err
	}

	key := BuildPageKey(branchName)
	page, err := p.Command.GetConfigValue(key)
	if err != nil {
		return "", nil
	}

	return page, nil
}

func (p *PageImpl) SetPage(page string) error {
	var err error
	branchName, err := p.Command.GetCurrentBranch()
	if err != nil {
		return err
	}

	key := BuildPageKey(branchName)
	err = p.Command.SetConfigValue(key, page)
	if err != nil {
		return err
	}

	return nil
}

func (p *PageImpl) DeletePage(branchName string) error {
	pageKey := BuildPageKey(branchName)
	err := p.Command.DeleteConfigValue(pageKey)
	if err != nil {
		return err
	}

	return nil
}

func BuildPageKey(branchName string) string {
	return "branch." + branchName + ".page"
}
