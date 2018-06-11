package config

type ConfigCube struct {
	Size int
}

type ConfigDraw struct {
	View    string
	Title   string
	Spec    []ConfigDrawSpec
	PreAlg  string
	PostAlg string
	Colors  []string
}

type ConfigDrawSpec struct {
	Name        string
	Description string
	Algs        []string
}

type Configuration struct {
	Cube ConfigCube
	Draw []ConfigDraw
}
