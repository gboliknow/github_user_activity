package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// createTableWriter creates a new tabwriter with consistent formatting
func createTableWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
}

// printTableRow prints a row in the table with proper formatting
func printTableRow(w *tabwriter.Writer, format string, a ...interface{}) {
	fmt.Fprintf(w, format+"\n", a...)
}

// printTableHeader prints a header row with proper formatting
func printTableHeader(w *tabwriter.Writer, headers ...string) {
	// Print headers
	for i, header := range headers {
		if i > 0 {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprint(w, formatHeader(header))
	}
	fmt.Fprintln(w)

	// Print separator
	for i, header := range headers {
		if i > 0 {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprint(w, formatSecondary(generateSeparator(header)))
	}
	fmt.Fprintln(w)
}

// generateSeparator generates a separator line for a header
func generateSeparator(header string) string {
	separator := ""
	for i := 0; i < len(header); i++ {
		separator += "-"
	}
	return separator
}
