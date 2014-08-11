package buildpack

// CloudFoundry is a contract that satisfies the Cloud Foundry buildpack API.
type CloudFoundry interface {
	Detect(buildDir string) int
	Compile(buildDir string, cacheDir string) int
	Release(buildDir string)
	Package(mode PackageMode)
}
