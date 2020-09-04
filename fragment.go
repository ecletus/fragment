package fragment

import (
	"reflect"

	"github.com/moisespsena-go/bid"

	"github.com/moisespsena-go/aorm"
)

type FragmentedModel struct {
	aorm.VirtualFields

	Fragments     map[string]FragmentModelInterface     `sql:"-";gorm:"-"`
	FormFragments map[string]FormFragmentModelInterface `sql:"-";gorm:"-"`
	data          map[interface{}]interface{}
}

func (f *FragmentedModel) SetData(key, value interface{}) {
	if f.data == nil {
		f.data = map[interface{}]interface{}{}
	}
	f.data[key] = value
}

func (f *FragmentedModel) GetData(key interface{}) (value interface{}, ok bool) {
	if f.data != nil {
		value, ok = f.data[key]
	}
	return
}

func (f *FragmentedModel) HasData(key interface{}) (ok bool) {
	if f.data != nil {
		_, ok = f.data[key]
	}
	return
}

func (f *FragmentedModel) DeleteData(key interface{}) (ok bool) {
	if f.data != nil {
		if _, ok = f.data[key]; ok {
			delete(f.data, key)
		}
	}
	return
}

func (f *FragmentedModel) GetFragments() map[string]FragmentModelInterface {
	if f.Fragments == nil {
		f.Fragments = make(map[string]FragmentModelInterface)
	}
	return f.Fragments
}

func (f *FragmentedModel) GetFragment(id string) FragmentModelInterface {
	if f.Fragments == nil {
		return nil
	}
	return f.Fragments[id]
}

func (f *FragmentedModel) SetFragment(super FragmentedModelInterface, id string, value FragmentModelInterface) {
	if value == nil {
		if f.Fragments != nil {
			if old, ok := f.Fragments[id]; ok {
				old.SetSuper(nil)
				delete(f.Fragments, id)
			}
		}
	} else {
		if f.Fragments == nil {
			f.Fragments = make(map[string]FragmentModelInterface)
		}
		f.Fragments[id] = value
		value.SetSuper(super)
	}
}

func (f *FragmentedModel) GetFormFragments() map[string]FormFragmentModelInterface {
	if f.FormFragments == nil {
		f.FormFragments = make(map[string]FormFragmentModelInterface)
	}
	return f.FormFragments
}

func (f *FragmentedModel) GetFormFragment(id string) FormFragmentModelInterface {
	if f.FormFragments == nil {
		return nil
	}
	return f.FormFragments[id]
}

func (f *FragmentedModel) SetFormFragment(super FragmentedModelInterface, id string, value FormFragmentModelInterface) {
	if value == nil {
		if f.FormFragments != nil {
			if old, ok := f.FormFragments[id]; ok {
				old.SetSuper(nil)
				delete(f.FormFragments, id)
			}
		}
	} else {
		if f.FormFragments == nil {
			f.FormFragments = make(map[string]FormFragmentModelInterface)
		}
		f.FormFragments[id] = value
		value.SetSuper(super)
	}
}

func (f *FragmentedModel) GetVirtualField(name string) (interface{}, bool) {
	if f.Fragments != nil {
		for _, v := range f.Fragments {
			if f := reflect.ValueOf(v).Elem().FieldByName(name); f.IsValid() {
				return f.Interface(), true
			}
		}
	}

	if f.FormFragments != nil {
		if v, ok := f.FormFragments[name]; ok {
			return v, true
		}
	}

	return f.VirtualFields.GetVirtualField(name)
}

type FragmentModel struct {
	ID bid.BID `aorm:"primary_key"`
	SingletonFragmentModel
}

func (f *FragmentModel) SetSuper(super FragmentedModelInterface) {
	f.super = super
	if super != nil {
		f.ID = aorm.RawOfId(aorm.IdOf(super)).(bid.BID)
	}
}

type SingletonFragmentModel struct {
	FragmentedModel
	super FragmentedModelInterface
}

func (f *SingletonFragmentModel) Super() FragmentedModelInterface {
	return f.super
}

func (f *SingletonFragmentModel) SetSuper(super FragmentedModelInterface) {
	f.super = super
}

type FragmentEnabledAttribute struct {
	FragmentEnabled bool
}

func (f *FragmentEnabledAttribute) Enabled() bool {
	return f.FragmentEnabled
}

func (f *FragmentEnabledAttribute) SetEnabled(v bool) {
	f.FragmentEnabled = v
}

func (f *FragmentEnabledAttribute) Enable() {
	f.FragmentEnabled = true
}

func (f *FragmentEnabledAttribute) Disable() {
	f.FragmentEnabled = false
}

type SingletonFormFragmentModel struct {
	SingletonFragmentModel
	FragmentEnabledAttribute
}

type FormFragmentModel struct {
	FragmentModel
	FragmentEnabledAttribute
}
