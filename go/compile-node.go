package main

type cnode struct {
	is     string
	value  string
	has    []*cnode
	typed  string
	_vdata *varData
	code   string
}

func (me *cnode) data() *varData {
	return me._vdata
}

func (me *cnode) copyData(data *varData) {
	if data == nil {
		me._vdata = nil
	} else {
		me._vdata = data.copy()
	}
}

func (me *cnode) copyType(other *cnode) {
	me.typed = other.typed
	me._vdata = other.data().copy()
}

func (me *cnode) copyTypeFromVar(other *variable) {
	me._vdata = other.data().copy()
}

func (me *cnode) getType() string {
	if me.data() != nil {
		return me.data().full
	}
	return me.typed
}

func (me *cnode) asVar(module *hmfile) *varData {
	if me.data() != nil {
		return me.data()
	}
	me.copyData(module.typeToVarData(me.typed))
	return me.data()
}

type codeblock struct {
	pre     *codeblock
	current *cnode
}

func (me *codeblock) flatten() []*cnode {
	flat := make([]*cnode, 0)
	if me.pre != nil {
		for _, p := range me.pre.flatten() {
			flat = append(flat, p)
		}
	}
	return append(flat, me.current)
}

func (me *codeblock) preCode() string {
	if me.pre != nil {
		return me.pre.code()
	}
	return ""
}

func (me *codeblock) code() string {
	ls := me.flatten()
	code := ""
	for _, n := range ls {
		code += n.code
	}
	return code
}

func (me *codeblock) data() *varData {
	return me.current.data()
}

func (me *codeblock) getType() string {
	return me.current.getType()
}

func (me *codeblock) asVar(module *hmfile) *varData {
	return me.current.asVar(module)
}

func (me *codeblock) is() string {
	return me.current.is
}

func codeBlockOne(n *node, code string) *codeblock {
	me := &codeblock{}
	me.current = codeNode(n, code)
	return me
}

func codeNodeUpgrade(cn *cnode) *codeblock {
	me := &codeblock{}
	me.current = cn
	return me
}