package models



type JSearchResponse struct {
	Status  	string `json:"status"`
	RequestID	string `json:"request_id"`
	Parameters	struct {
		Query 				string `json:"query"`
		Page 				int    `json:"page"`
		Num_Pages 			int    `json:"num_pages"`
		Date_posted 		string `json:"date_posted"`
		Employment_types	[]string `json:"employment_types"`
		Job_requirements 	[]string `json:"job_requirements"`
		Country 			string `json:"country"`
		Language 			string `json:"language"`
	} `json:"parameters"`
	Data 		[]JobPost `json:"data"`
}

type JobPost struct {
	JobID 		string   `json:"job_id"`
	JobTitle 	string   `json:"job_title"`
	EmployerName string   `json:"employer_name"`
	EmployerWebsite string `json:"employer_website"`
	JobPulisher string   `json:"job_publisher"`
	JobEmploymentType string `json:"job_employment_type"`
	JobApplyLink string `json:"job_apply_link"`
	ApplyOptions []struct {
		Publisher 	string `json:"publisher"`
		ApplyLink 	string `json:"apply_link"`
		IsDirect 	bool   `json:"is_direct"`
	} `json:"apply_options"`
	JobDescription string `json:"job_description"`
	JobIsRemote		bool	`json:"job_is_remote"`
	JobPostedAtTimestamp	any		`json:"job_posted_at_timestamp"`
	JobCity					string		`json:"job_city"`
	JobState				string		`json:"job_state"`
	JobCountry				string		`json:"job_country"`
	JobMinSalary			any			`json:"job_min_salary"`
	JobMaxSalary			any			`json:"job_max_salary"`
	JobSalaryPeriod			any			`json:"job_salary_period"`
}

type SearchJobParams struct {
	Query					string		`json:"query"`
	Page					uint		`json:"page"`
	NumPages				uint		`json:"num_pages"`
	Country					string		`json:"country"`
	Language				string		`json:"language"`
	DatePosted				string		`json:"date_posted"`
	WorkFromHome			string		`json:"work_from_home"`
	EmploymentTypes			[]string	`json:"employment_types"`
	ExperienceLevels		string		`json:"experience_levels"`
	ExcludeJobPublishers 	[]string	`json:"exclude_job_publishers"`
}



