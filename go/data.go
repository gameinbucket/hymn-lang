package main

import (
	"strings"
)

type variable struct {
	typed   string
	name    string
	dfault  string
	mutable bool
	isptr   bool
	cName   string
	vdat    *varData
}

func (me *hmfile) varInit(typed, name string, mutable, isptr bool) *variable {
	v := &variable{}
	v.typed = typed
	v.name = name
	v.cName = name
	v.mutable = mutable
	v.isptr = isptr
	v.vdat = me.typeToVarData(typed)
	return v
}

func (me *hmfile) varWithDefaultInit(typed, name string, mutable, isptr bool, dfault string) *variable {
	v := me.varInit(typed, name, mutable, isptr)
	v.dfault = dfault
	return v
}

func (me *variable) copy() *variable {
	v := &variable{}
	v.typed = me.typed
	v.name = me.name
	v.cName = me.name
	v.mutable = me.mutable
	v.isptr = me.isptr
	v.vdat = me.vdat
	return v
}

func (me *variable) update(module *hmfile, typed string) {
	me.typed = typed
	me.vdat = module.typeToVarData(typed)
}

func (me *variable) memget() string {
	if me.isptr {
		return "->"
	}
	return "."
}

type varData struct {
	module      *hmfile
	typed       string
	full        string
	mutable     bool
	pointer     bool
	heap        bool
	array       bool
	typeInArray string
	en          *enum
	un          *union
	cl          *class
}

func dataInit(module *hmfile, typed string, mutable, pointer, heap bool) *varData {
	d := &varData{}
	d.module = module
	d.typed = typed
	d.mutable = mutable
	d.pointer = pointer
	d.heap = heap
	return d
}

func (me *hmfile) typeToVarData(typed string) *varData {
	data := &varData{}
	data.full = typed
	data.mutable = true
	data.pointer = true
	data.heap = true

	data.array = checkIsArray(typed)
	if data.array {
		typed = typeOfArray(typed)
		data.typeInArray = typed
	}

	if typed[0] == '$' {
		data.pointer = false
		typed = typed[1:]
	} else if typed[0] == '\\' {
		data.heap = false
		typed = typed[1:]
	}

	data.module = me
	data.typed = typed

	dot := strings.Split(typed, ".")
	if len(dot) != 1 {
		if module, ok := me.program.hmfiles[dot[0]]; ok {
			data.module = module
			if len(dot) > 2 {
				if _, ok := me.enums[dot[1]]; ok {
					data.typed = dot[1] + "." + dot[2]
				} else {
					panic("unknown type \"" + typed + "\"")
				}
			} else {
				data.typed = dot[1]
			}
		} else if _, ok := me.enums[dot[0]]; ok {
			data.typed = dot[0] + "." + dot[1]
		} else {
			panic("unknown type \"" + typed + "\"")
		}
	}

	return data
}

func (me *varData) checkIsArray() bool {
	return strings.HasPrefix(me.full, "[]")
}

func (me *varData) checkIsClass() (*class, bool) {
	cl, ok := me.module.classes[me.typed]
	return cl, ok
}

func (me *varData) checkIsEnum() (*enum, *union, bool) {
	dot := strings.Split(me.typed, ".")
	if len(dot) != 1 {
		en, ok := me.module.enums[dot[0]]
		un, _ := en.types[dot[1]]
		return en, un, ok
	}
	en, ok := me.module.enums[me.typed]
	return en, nil, ok
}

func (me *varData) checkIsUnion() (*enum, bool) {
	dot := strings.Split(me.typed, ".")
	if len(dot) != 1 {
		en, ok := me.module.enums[dot[0]]
		if ok && en.simple {
			return en, true
		}
		return nil, false
	}
	en, ok := me.module.enums[me.typed]
	if ok && en.simple {
		return en, true
	}
	return nil, false
}

func (me *varData) postfixConst() bool {
	if me.checkIsArray() {
		return true
	}
	if _, ok := me.checkIsClass(); ok {
		return true
	}
	if _, ok := me.checkIsUnion(); ok {
		return true
	}
	return false
}

func (me *varData) equal(other *varData) bool {
	if me.full == other.full {
		return true
	}
	if en, _, ok := me.checkIsEnum(); ok {
		if en2, _, ok2 := other.checkIsEnum(); ok2 {
			if en.name == en2.name {
				return true
			}
		}
	}
	return false
}

func (me *varData) notEqual(other *varData) bool {
	return !me.equal(other)
}

func (me *varData) typeSig() string {
	if me.array {
		return fmtptr(me.module.typeToVarData(me.typeInArray).typeSig())
	}
	if _, ok := me.checkIsClass(); ok {
		return me.module.classNameSpace(me.typed) + " *"
	} else if en, _, ok := me.checkIsEnum(); ok {
		return en.typeSig()
	} else if me.full == "string" {
		return "char *"
	}
	return me.full
}

func (me *varData) noMallocTypeSig() string {
	if me.array {
		return fmtptr(me.module.typeToVarData(me.typeInArray).noMallocTypeSig())
	}
	if _, ok := me.checkIsClass(); ok {
		return me.module.classNameSpace(me.typed)
	} else if en, _, ok := me.checkIsEnum(); ok {
		return en.noMallocTypeSig()
	} else if me.full == "string" {
		return "char *"
	}
	return me.full
}
