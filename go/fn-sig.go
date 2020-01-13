package main

type fnSig struct {
	module      *hmfile
	args        []*funcArg
	argVariadic *funcArg
	returns     *varData
}

func fnSigInit(module *hmfile) *fnSig {
	f := &fnSig{}
	f.module = module
	f.args = make([]*funcArg, 0)
	return f
}

func (me *fnSig) print() string {
	sig := "("
	for ix, arg := range me.args {
		if ix > 0 {
			sig += ", "
		}
		sig += arg.data().print()
	}
	if me.argVariadic != nil {
		if len(me.args) > 0 {
			sig += ", "
		}
		sig += "..." + me.argVariadic.data().print()
	}
	sig += ")"
	if !me.returns.isVoid() {
		sig += " "
		sig += me.returns.print()
	}
	return sig
}

func (me *fnSig) data() *varData {
	sig := me.print()
	d := &varData{}
	d.fn = me
	d.module = me.module
	d.dtype = getdatatype(nil, sig)
	return d
}
