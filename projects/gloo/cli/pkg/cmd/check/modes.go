package check

const (
	checkAllResourcesMode = iota
	checkFileSpecMode
)

var modeName map[int]string

func init() {
	modeName = make(map[int]string)
	modeName[checkAllResourcesMode] = "all resources"
	modeName[checkFileSpecMode] = "file specification"
}
