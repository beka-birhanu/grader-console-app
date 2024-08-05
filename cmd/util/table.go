package util

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/beka-birhanu/grader/app/common"
)

// DisplayAvarageResponse prints the subjects and their marks in a tabular format
// and displays the average mark.
func DisplayAvarageResponse(response common.AvarageResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', tabwriter.Debug)

	if len(response.Subjects) == 0 {
		fmt.Println("No subjects to display.")
		return
	}

	ownerName := response.Subjects[0].OwnerName()
	fmt.Printf("Hey %s, here is your summary:\n\n", ownerName)

	// Print headers with a separator line
	fmt.Fprintln(w, "Subject Name\tMark\t")
	fmt.Fprintln(w, "------------\t----\t")

	// Print each subject's details
	for _, subject := range response.Subjects {
		fmt.Fprintf(w, "%s\t%.2f\t\n", subject.Name(), subject.Mark())
	}

	w.Flush()

	// Print the average mark
	fmt.Printf("\nAverage\t%.2f\t\n\n", response.Avarage)
}

