package advancedError

import "Datearn/pkg/logger"

// observer defines related and needed set of methods
type observer interface {
	Update(msg string)
}

// fileLogger defines custom observer for logging
type fileLogger struct{}

// Update is called when new error's error method is going to get called
func (f *fileLogger) Update(msg string) {
	logger.Factory("").Log(msg)
}
