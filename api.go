package fragment

import "github.com/moisespsena-go/aorm"

type FragmentedModelInterface interface {
	aorm.VirtualFieldsGetter

	GetFragments() map[string]FragmentModelInterface
	GetFragment(id string) FragmentModelInterface
	SetFragment(super FragmentedModelInterface, id string, value FragmentModelInterface)
	GetFormFragments() map[string]FormFragmentModelInterface
	GetFormFragment(id string) FormFragmentModelInterface
	SetFormFragment(super FragmentedModelInterface, id string, value FormFragmentModelInterface)
	SetData(key, value interface{})
	GetData(key interface{}) (value interface{}, ok bool)
	HasData(key interface{}) (ok bool)
	DeleteData(key interface{}) (ok bool)
}

type FragmentModelInterface interface {
	FragmentedModelInterface
	Super() FragmentedModelInterface
	SetSuper(super FragmentedModelInterface)
}

type FormFragmentModelInterface interface {
	FragmentModelInterface
	Enabled() bool
	SetEnabled(v bool)
	Enable()
	Disable()
}