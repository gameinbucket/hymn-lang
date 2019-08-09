package main

func prefixSign(me *parser, op string) *node {
	node := nodeInit(getPrefixName(op))
	me.eat(op)
	right := me.calc(getPrefixPrecedence(op))
	node.push(right)
	node.copyType(right)
	return node
}

func prefixGroup(me *parser, op string) *node {
	me.eat("(")
	node := me.calc(0)
	node.attributes["parenthesis"] = "true"
	me.eat(")")
	return node
}

func prefixPrimitive(me *parser, op string) *node {
	node := nodeInit(op)
	node.vdata = me.hmfile.typeToVarData(op)
	node.value = me.token.value
	me.eat(op)
	return node
}

func prefixIdent(me *parser, op string) *node {
	useStack := false
	if me.token.is == "$" {
		me.eat("$")
		useStack = true
	}

	name := me.token.value
	if _, ok := me.hmfile.functions[name]; ok {
		return me.call(me.hmfile)
	}
	if _, ok := me.hmfile.types[name]; ok {
		if _, ok := me.hmfile.classes[name]; ok {
			data := &allocData{}
			data.useStack = useStack
			return me.allocClass(me.hmfile, data)
		}
		if _, ok := me.hmfile.enums[name]; ok {
			data := &allocData{}
			data.useStack = useStack
			return me.allocEnum(me.hmfile, data)
		}
		if def, ok := me.hmfile.defs[name]; ok {
			return me.exprDef(name, def)
		}
		panic(me.fail() + "bad type \"" + name + "\" definition")
	}
	if _, ok := me.hmfile.imports[name]; ok {
		return me.extern()
	}
	v := me.hmfile.getvar(name)
	if me.peek().is == ":=" {
		if v != nil && v.mutable == false {
			panic(me.fail() + "variable not mutable")
		}
	} else if v == nil {
		panic(me.fail() + "variable out of scope")
	}
	return me.eatvar(me.hmfile)
}

func prefixArray(me *parser, op string) *node {
	me.eat("[")
	size := me.calc(0)
	if size.getType() != "int" {
		panic(me.fail() + "array size must be integer")
	}
	me.eat("]")
	node := nodeInit("array")
	alloc := &allocData{}
	alloc.isArray = true
	node.vdata = me.buildAnyType(alloc)
	node.push(size)
	return node
}

func prefixNot(me *parser, op string) *node {
	if me.token.is == "!" {
		me.eat("!")
	} else {
		me.eat("not")
	}
	node := nodeInit("not")
	node.vdata = me.hmfile.typeToVarData("bool")
	node.push(me.calcBool())
	return node
}

func prefixNone(me *parser, op string) *node {
	me.eat("none")
	me.eat("<")
	option := me.declareType(true)
	me.eat(">")
	typed := "none<" + option.typed + ">"
	data := me.hmfile.typeToVarData(typed)
	me.defineMaybeImpl(typed)

	node := nodeInit("none")
	node.vdata = data
	return node
}

func prefixMaybe(me *parser, op string) *node {
	me.eat("maybe")
	me.eat("<")
	option := me.declareType(true)
	me.eat(">")
	typed := "maybe<" + option.typed + ">"
	data := me.hmfile.typeToVarData(typed)
	me.defineMaybeImpl(typed)

	n := nodeInit("maybe")
	n.vdata = data
	return n
}
