package domain

type ErrorDescription struct {
	HttpCode int
	Code     string
}

type Error struct {
	ErrorDescription

	Stack           Stack
	Child           *Error                 `json:",omitempty"`
	AdditionnalData map[string]interface{} `json:",omitempty"`
	Fingerprint     string
}

var (
	ErrorNotFound = ErrorDescription{HttpCode: 404, Code: "PageNotFound"}
	ErrorNoMethod = ErrorDescription{HttpCode: 405, Code: "MethodNotAllowed"}
)
