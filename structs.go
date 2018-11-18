package schoology

//School Represents a school
type School struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	Website    string `json:"website"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	PictureURL string `json:"picture_url"`
}

//Building same as school
type Building struct {
	*School
}

//User user objcet
type User struct {
	ID                 string `json:"uid"`
	SchoolID           string `json:"school_id"`
	BuildingID         string `json:"building_id"`
	SchoolUID          string `json:"school_uid"`
	NameTitle          string `json:"name_title"`
	ShowNameTitle      int    `json:"name_title_show"`
	FirstName          string `json:"name_first"`
	PreferredFirstName string `json:"name_first_preferred"`
	MiddleName         string `json:"name_middle"`
	ShowMiddleName     int    `json:"middle_name_show"`
	LastName           string `json:"name_last"`
	DisplayName        string `json:"name_display"`
	Username           string `json:"username"`
	PrimaryEmail       string `json:"primary_email"`
	Position           string `json:"position"`
	Gender             string `json:"gender"`
	GraduationYear     int    `json:"grad_year"`
	Birthday           string `json:"birthday_date"`
	//Password           string `json:"password"`
	RoleID             int64  `json:"role_id"`
	EmailLoginInfo     int    `json:"email_login_info"`
	ProfileURL         string `json:"profile_url"`
	TimezoneName       string `json:"tz_name"`
	Parents            []User `json:"parents"`
	ParentUIDs         string `json:"parents_uids"`
	AdvisorUIDs        string `json:"advisor_uids"`
	ChildUIDs          string `json:"childs_uids"`
	SendMessage        int    `json:"send_message"`
	Synced             int    `json:"synced"`
	ProfilePictureFID  int    `json:"profile_picture_fid"`
	AdditionalBuilding string `json:"additional_buildings"`
}

//ExtendedUser user but with the extended fields
type ExtendedUser struct {
	*User
	SubjectsTaught string `json:"subjects_taught"`
	GradesTaught   string `json:"grades_taught"`
	Department     string `json:"department"`
	Bio            string `json:"bio"`
	Phone          string `json:"phone"`
	Website        string `json:"website"`
	Address        string `json:"address"`
	Interests      string `json:"interests"`
	Activites      string `json:"activites"`
}

//Group A group
type Group struct {
	ID           string `json:"id"`
	BuildingID   string `json:"building_id"`
	SchoolID     string `json:"school_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PictureURL   string `json:"picture_url"`
	Website      string `json:"website"`
	AccessCode   string `json:"access_code"`
	PrivacyLevel string `json:"privacy_level"`
	Category     string `json:"category"`
	Options      struct {
		InviteType           int `json:"invite_type"`
		CanMemberPost        int `json:"member_post"`
		CanMemberPostComment int `json:"member_post_comment"`
		CanCreateDiscussion  int `json:"create_discussion"`
		CanCreateFiles       int `json:"create_files"`
	} `json:"options"`
	GroupCode string `json:"group_code"`
}

//Course a course
type Course struct {
	ID                   string  `json:"id"`
	BuildingID           string  `json:"building_id"`
	Title                string  `json:"title"`
	CourseCode           string  `json:"course_code"`
	Department           string  `json:"department"`
	Description          string  `json:"descriptions"`
	Credits              float32 `json:"credits"`
	Synced               int     `json:"synced"`
	GradeLevelRangeStart int     `json:"grade_level_range_start"`
	GradeLevelRangeEnd   int     `json:"grade_level_range_end"`
	SubjectArea          int     `json:"subject_area"`
}

//CourseSection a section of a course(?)
type CourseSection struct {
	ID                string   `json:"id"`
	Title             string   `json:"section_title"`
	SectionCode       string   `json:"section_code"`
	SectionSchoolCode string   `json:"section_school_code"`
	AccessCode        string   `json:"access_code"`
	GradingPeriods    []int    `json:"grading_periods"`
	Description       string   `json:"description"`
	ProfileURL        string   `json:"profile_url"`
	Location          string   `json:"location"`
	MeetingDays       []string `json:"meeting_days"`
	StartTime         string   `json:"start_time"`
	EndTime           string   `json:"end_time"`
	ClassPeriods      []int    `json:"class_periods"`
	Options           struct {
		CourseFormat              int    `json:"course_format"`
		WeightedGradingCategories string `json:"weighted_grading_categories"`
		CanUploadDocument         string `json:"upload_document"`
		CanCreateDiscussion       string `json:"create_discussion"`
		CanMemberPost             string `json:"member_post"`
		CanMemberPostComment      string `json:"member_post_comment"`
	} `json:"options"`
	Synced string `json:"synced"`
}
