package version

const Unversioned = "unversioned"

var (
	version = Unversioned
)

type Version interface {
	String() string
}

type _default struct{}

func NewDefaultVersion() Version {
	return &_default{}
}

func (r *_default) String() string {
	return version
}
