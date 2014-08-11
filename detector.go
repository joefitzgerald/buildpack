package buildpack

type Detector interface {
	// Detect is used to determine whether or not to apply a buildpack to an
	// application. It is called with one argument: the build directory for the
	// application. It returns an exit code of 0 if the application can be supported
	// by the buildpack. If it does return 0, it should also print a framework name
	// to STDOUT.
	Detect(buildDir string) int
}
