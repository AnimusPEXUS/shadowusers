package shadowusers

import (
	"errors"
	"io/ioutil"
	"path"
)

type Ctl struct {
	Pth string

	Passwds  *Passwds
	Groups   *Groups
	Shadows  *Shadows
	GShadows *GShadows
}

func NewCtl(pth string) *Ctl {
	self := new(Ctl)
	self.Pth = pth
	return self
}

func (self *Ctl) PasswdsPath() string {
	return path.Join(self.Pth, "passwd")
}

func (self *Ctl) GroupsPath() string {
	return path.Join(self.Pth, "group")
}

func (self *Ctl) ShadowsPath() string {
	return path.Join(self.Pth, "shadow")
}

func (self *Ctl) GShadowsPath() string {
	return path.Join(self.Pth, "gshadow")
}

func (self *Ctl) ReadPasswds() error {

	t, err := ioutil.ReadFile(self.PasswdsPath())
	if err != nil {
		return err
	}

	o, err := NewPasswdsFromString(string(t))
	if err != nil {
		return err
	}

	self.Passwds = o

	return nil
}

func (self *Ctl) WritePasswds() error {
	if self.Passwds == nil {
		return errors.New("Passwd not loaded")
	}

	txt, err := self.Passwds.Render()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(self.PasswdsPath(), []byte(txt), 755)
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) ReadGroups() error {

	t, err := ioutil.ReadFile(self.GroupsPath())
	if err != nil {
		return err
	}

	o, err := NewGroupsFromString(string(t))
	if err != nil {
		return err
	}

	self.Groups = o

	return nil
}

func (self *Ctl) WriteGroups() error {
	if self.Groups == nil {
		return errors.New("Groups not loaded")
	}

	txt, err := self.Groups.Render()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(self.GroupsPath(), []byte(txt), 755)
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) ReadShadows() error {

	t, err := ioutil.ReadFile(self.ShadowsPath())
	if err != nil {
		return err
	}

	o, err := NewShadowsFromString(string(t))
	if err != nil {
		return err
	}

	self.Shadows = o

	return nil
}

func (self *Ctl) WriteShadows() error {
	if self.Shadows == nil {
		return errors.New("Shadows not loaded")
	}

	txt, err := self.Shadows.Render()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(self.ShadowsPath(), []byte(txt), 700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) ReadGShadows() error {

	t, err := ioutil.ReadFile(self.GShadowsPath())
	if err != nil {
		return err
	}

	o, err := NewGShadowsFromString(string(t))
	if err != nil {
		return err
	}

	self.GShadows = o

	return nil
}

func (self *Ctl) WriteGShadows() error {
	if self.GShadows == nil {
		return errors.New("GShadows not loaded")
	}

	txt, err := self.GShadows.Render()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(self.GShadowsPath(), []byte(txt), 700)
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) ReadAll() error {
	err := self.ReadPasswds()
	if err != nil {
		return err
	}

	err = self.ReadGroups()
	if err != nil {
		return err
	}

	err = self.ReadShadows()
	if err != nil {
		return err
	}

	err = self.ReadGShadows()
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) WriteAll() error {
	err := self.WritePasswds()
	if err != nil {
		return err
	}

	err = self.WriteGroups()
	if err != nil {
		return err
	}

	err = self.WriteShadows()
	if err != nil {
		return err
	}

	err = self.WriteGShadows()
	if err != nil {
		return err
	}

	return nil
}

func (self *Ctl) NewPasswds() {
	self.Passwds = &Passwds{}
}

func (self *Ctl) NewGroups() {
	self.Groups = &Groups{}
}

func (self *Ctl) NewShadows() {
	self.Shadows = &Shadows{}
}

func (self *Ctl) NewGShadow() {
	self.GShadows = &GShadows{}
}
