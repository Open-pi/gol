package gol

import "fmt"

type Subject struct {
	Container
	name  string
	works []Work
}

// GetSubject returns the chosen subject's information
func GetSubject(subject string) (sbj Subject, err error) {
	return getSubject(subject, false)
}

// GetSubject returns the chosen subject's information with more details
func GetSubjectDetails(subject string) (sbj Subject, err error) {
	return getSubject(subject, true)
}

// GeneralGetSubject is a general function for GetSubject and GetSubjectDetails functions
func getSubject(s string, detail bool) (sbj Subject, err error) {
	if !detail {
		sbj.Container, err = MakeSubjectRequest(s)
	} else {
		sbj.Container, err = MakeDetailedSubjectRequest(s)
	}

	if err != nil {
		return sbj, err
	}

	// verify if an error field is present in the returned data
	if err := HasError(sbj.Container); err != nil {
		return sbj, err
	}

	if v, ok := sbj.Path("name").Data().(string); ok {
		sbj.name = v
	}
	return
}

// Works returns the works related to that subject.
// Note that the work fields won't get loaded (using the Work.load() method);
// If the works were loaded before, it will return sbj.works
func (sbj *Subject) Works() ([]Work, error) {
	if len(sbj.works) > 0 {
		return sbj.works, nil
	}

	for _, child := range sbj.Path("works").Children() {
		sbj.works = append(sbj.works, Work{Container: child})
	}
	if len(sbj.works) == 0 {
		return sbj.works, fmt.Errorf("Could not find works")
	}
	return sbj.works, nil
}
