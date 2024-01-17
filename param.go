package wgapi

type ParamOption interface {
	apply(*paramOption)
}

// paramOption contains all basic info for a request
type paramOption struct {
	// only accept GET or POST so far, this will decide use which way when request to remote
	method string
	// the path of the request
	path string
	// the region of the request
	region string
	// the params of the request
	args map[string]string
}

type asyncOption struct {
	unMarshal func([]byte, error) (any, error)
}

type Param struct {
	popts paramOption
	async asyncOption
}

// params default values
func defaultParamOption() paramOption {
	return paramOption{
		method: "GET",
		path:   Host,
		region: "asia",
		args: map[string]string{
			"language": "zh-cn",
		},
	}
}

// returns the region of the request
func (p *Param) GetRegion() string {
	return p.popts.region
}

// returns the params of the request
func (p *Param) GetParam() map[string]string {
	return p.popts.args
}

// param Setter
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
