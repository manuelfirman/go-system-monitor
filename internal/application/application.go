package application

// Application is the interface that wraps the basic methods of an application.
type Application interface {
	// SetUp is the method that sets up the application.
	SetUp()
	// Run is the method that runs the application.
	Run() error
}
