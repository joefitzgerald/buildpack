package buildpack

// Releaser is responsible for providing metadata about how an application
// should be executed.
type Releaser interface {
	// Release provides feedback metadata back to Cloud Foundry indicating how the
	// application should be executed. It is called with one argument, the build location
	// of the application.
	Release(buildDir string)
}
