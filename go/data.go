package main

import (
	"strings"
)

type idData struct {
	module *hmfile
	name   string
}

type varData struct {
	module      *hmfile
	typed       string
	full        string
	mutable     bool
	onStack     bool
	isptr       bool
	heap        bool
	array       bool
	none        bool
	maybe       bool
	some        string
	noneType    string
	typeInArray *varData
	en          *enum
	un          *union
	cl          *class
	fn          *fnSig
}

func (me *varData) copy() *varData {
	v := &varData{}
	v.module = me.module
	v.typed = me.typed
	v.full = me.full
	v.mutable = me.mutable
	v.onStack = me.onStack
	v.isptr = me.isptr
	v.heap = me.heap
	v.array = me.array
	v.none = me.none
	v.maybe = me.maybe
	v.some = me.some
	v.noneType = me.noneType
	v.typeInArray = me.typeInArray
	v.en = me.en
	v.un = me.un
	v.cl = me.cl
	v.fn = me.fn
	return v
}

func (me *hmfile) typeToVarDataWithAttributes(typed string, attributes map[string]string) *varData {
	data := me.typeToVarData(typed)

	if _, ok := attributes["use-stack"]; ok {
		data.onStack = true
	}

	return data
}

func (me *hmfile) typeToVarData(typed string) *varData {
	data := &varData{}
	data.full = typed
	data.mutable = true
	data.isptr = true
	data.heap = true

	if strings.HasPrefix(typed, "maybe") {
		data.maybe = true
		data.some = typed[6 : len(typed)-1]

	} else if strings.HasPrefix(typed, "none") {
		data.none = true
		data.noneType = typed[5 : len(typed)-1]
	}

	data.array = checkIsArray(typed)
	if data.array {
		typed = typeOfArray(typed)
		data.typeInArray = me.typeToVarData(typed)
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

func (me *varData) asVariable() *variable {
	v := &variable{}
	v.vdat = me
	return v
}

func (me *varData) merge(alloc *allocData) {
	if alloc == nil {
		return
	}

	me.array = alloc.isArray
	me.heap = !alloc.useStack

	if me.array {
		typeInArray := me.copy()
		typeInArray.array = false

		me.typeInArray = typeInArray
		me.full = "[]" + typeInArray.full
		me.typed = "[]" + typeInArray.typed
	}
}

func (me *varData) checkIsArray() bool {
	return strings.HasPrefix(me.full, "[]")
}

func (me *varData) checkIsFunction() (*fnSig, bool) {
	return me.fn, me.fn != nil
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

	} else if me.maybe {
		if other.maybe {
			if me.some == other.some {
				return true
			}
		} else if other.none {
			if me.some == other.noneType {
				return true
			}
		} else if me.some == other.full {
			return true
		}

	} else if me.none {
		if other.maybe {
			if me.noneType == other.some {
				return true
			}
		} else if other.none {
			if me.noneType == other.noneType {
				return true
			}
		} else if me.noneType == other.full {
			return true
		}

	} else if other.maybe {
		if other.some == me.full {
			return true
		}
	} else if other.none {
		if other.noneType == me.full {
			return true
		}
	}

	return false
}

func (me *varData) notEqual(other *varData) bool {
	return !me.equal(other)
}

func (me *varData) postfixConst() bool {
	if me.array {
		return true
	}
	if me.maybe {
		return me.module.typeToVarData(me.some).postfixConst()
	}
	if me.none {
		return me.module.typeToVarData(me.noneType).postfixConst()
	}
	if _, ok := me.checkIsClass(); ok {
		return true
	}
	if _, _, ok := me.checkIsEnum(); ok {
		return true
	}
	return false
}

func (me *varData) typeSigOf(name string, mutable bool) string {
	code := ""
	if _, ok := me.checkIsFunction(); ok {
		sig := me.fn
		code += fmtassignspace(sig.typed.typeSig())
		code += "(*"
		if !mutable {
			code += "const "
		}
		code += name
		code += ")("
		for ix, arg := range sig.args {
			if ix > 0 {
				code += ", "
			}
			code += arg.vdat.typeSig()
		}
		code += ")"

	} else {
		sig := fmtassignspace(me.typeSig())
		if mutable {
			code += sig
		} else if me.postfixConst() {
			code += sig + "const "
		} else {
			code += "const " + sig
		}
		code += name
	}
	return code
}

func (me *varData) typeSig() string {
	if me.array {
		return fmtptr(me.typeInArray.typeSig())
	}
	if me.maybe {
		return me.module.typeToVarData(me.some).typeSig()
	}
	if me.none {
		return me.module.typeToVarData(me.noneType).typeSig()
	}
	if _, ok := me.checkIsClass(); ok {
		sig := me.module.classNameSpace(me.typed)
		if !me.onStack && me.isptr {
			sig += " *"
		}
		return sig
	} else if en, _, ok := me.checkIsEnum(); ok {
		return en.typeSig()
	} else if me.full == "string" {
		return "char *"
	}
	return me.full
}

func (me *varData) noMallocTypeSig() string {
	if me.array {
		return fmtptr(me.typeInArray.noMallocTypeSig())
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

func (me *varData) memPtr() string {
	if me.isptr {
		return "->"
	}
	return "."
}
