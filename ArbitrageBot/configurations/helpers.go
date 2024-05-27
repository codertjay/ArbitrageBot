package configurations

type HelperInterface interface {
}

type Helper struct {
}

func NewHelper() HelperInterface {
	return &Helper{}
}

func (h Helper) name() {

}
