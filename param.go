package wgapi

type ParamOption interface {
	apply(*paramOption)
}

type paramOption struct {
	path   string
	region string
	args   map[string]string
}

type Param struct {
	popts paramOption
}

func (p *Param) GetRegion() string {
	return p.popts.region
}
func (p *Param) GetParam() map[string]string {
	return p.popts.args
}

func defaultParamOption() paramOption {
	return paramOption{
		path:   Host,
		region: "asia",
		args: map[string]string{
			"language": "zh-cn",
		},
	}
}

type FuncParam struct {
	F func(*paramOption)
}

func (fp *FuncParam) apply(po *paramOption) {
	fp.F(po)
}

func newParam(ops ...ParamOption) (param *Param) {
	ret := &Param{
		popts: defaultParamOption(),
	}
	for _, o := range ops {
		o.apply(&ret.popts)
	}
	return ret
}

func WithPath(path string) ParamOption {
	return &FuncParam{
		F: func(po *paramOption) {
			po.path = path
		},
	}
}

// WithRegion will set request target region
func WithRegion(region string) ParamOption {
	return &FuncParam{
		F: func(po *paramOption) {
			po.region = region
		},
	}
}

func WithParam(key string, value string) ParamOption {
	return &FuncParam{
		F: func(po *paramOption) {
			po.args[key] = value
		},
	}
}

func InsertBefore(orig *[]ParamOption, opts ...ParamOption) /*[]ParamOption*/ {
	*orig = append(opts, *orig...)
}
