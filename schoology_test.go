package schoology

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSchool(t *testing.T) {
	t.Logf("SetCredentials('%s', '%s')\nGetSchool('%s')", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("SCHOOL_ID"))
	SetCredentials(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	s, code, err := GetSchool(os.Getenv("SCHOOL_ID"))
	if err != nil {
		t.Errorf("failed to get school (status %d): %v\n", code, err)
		t.Fail()
	}
	fmt.Printf("school (status %d): %v\n", code, s)
}

func TestGetBuildings(t *testing.T) {
	t.Logf("SetCredentials('%s', '%s')\nGetSchoolBuildings('%s')", os.Getenv("API_KEY"), os.Getenv("API_SECRET"), os.Getenv("SCHOOL_ID"))
	SetCredentials(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	buildings, code, err := GetSchoolBuildings(os.Getenv("SCHOOL_ID"))
	if err != nil {
		t.Errorf("failed to get buildings (status %d): %v\n", code, err)
		t.Fail()
	}
	fmt.Printf("school buildings (code %d): %v\n", code, buildings)
}
