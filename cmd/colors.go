package cmd

import "github.com/fatih/color"

var (
	// Headers and titles
	headerStyle = color.New(color.FgCyan, color.Bold).SprintfFunc()
	titleStyle  = color.New(color.FgGreen, color.Bold).SprintfFunc()

	// Status colors
	successStyle = color.New(color.FgGreen).SprintfFunc()


	// Data colors
	primaryStyle   = color.New(color.FgBlue).SprintfFunc()
	secondaryStyle = color.New(color.FgMagenta).SprintfFunc()
)

// formatHeader formats a section header
func formatHeader(format string, a ...interface{}) string {
	return headerStyle(format, a...)
}

// formatTitle formats a title
func formatTitle(format string, a ...interface{}) string {
	return titleStyle(format, a...)
}

// formatSuccess formats success messages
func formatSuccess(format string, a ...interface{}) string {
	return successStyle(format, a...)
}


// formatPrimary formats primary data
func formatPrimary(format string, a ...interface{}) string {
	return primaryStyle(format, a...)
}

// formatSecondary formats secondary data
func formatSecondary(format string, a ...interface{}) string {
	return secondaryStyle(format, a...)
}
