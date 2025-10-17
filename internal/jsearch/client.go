package jsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Passenger7390/JobAgg/models"
)

const (
	NoExperience		 = "no_experience"
	NoDegree			 = "no_degree"
	Under3YearsExp 		 = "under_3_years_experience"
	Morethan3YearsExp	 = "more_than_3_years_experience"

	All			= "all"
	Today		= "today"
	ThreeDays	= "3days"
	Week   		= "week"
	Month  		= "month"

	FullTime	= "FULLTIME"
	Contractor	= "CONTRACTOR"
	PartTime	= "PARTTIME"
	Intern		= "INTERN"

	MinNumPages = 1
	MaxNumPages = 50
)

var allowedDatePosted = map[string]string{
	All:       All,
	Today:     Today,
	ThreeDays: ThreeDays,
	Week:      Week,
	Month:     Month,
}

type Client struct {
	baseURL    string
	apiKey     string
	apiHost    string
	httpClient *http.Client
	fields     string
	query      string
}

func NewClient(baseURL, apiKey, apiHost, fields string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		baseURL:	strings.TrimRight(baseURL, "/"),
		apiKey:		apiKey,
		apiHost:	apiHost,
		httpClient: httpClient,
		fields:		fields,
	}
}

func (c *Client) SearchJob(ctx context.Context, params models.SearchJobParams) (string, *models.JSearchResponse, error) {
	fmt.Println("Number of pages: ", params.NumPages)
	if params.Page <= 0 {
		params.Page = 1
	} else {
		c.query += "&page=" + string(rune(params.Page))
	}

	if params.NumPages < MinNumPages || params.NumPages > MaxNumPages {
		return "", nil, fmt.Errorf("num_pages must be between %d and %d", MinNumPages, MaxNumPages)
	}

	if params.DatePosted != "" {
		if _, ok := allowedDatePosted[params.DatePosted]; !ok {
			return "", nil, fmt.Errorf("invalid date_posted value: %s", params.DatePosted)
		}
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return "", nil, fmt.Errorf("invalid base url: %w", err)
	}

	u.Path = strings.TrimRight(u.Path, "/") + "/search"

	q := u.Query()
	if params.Query != "" {
		q.Set("query", params.Query)
	}

	q.Set("page", strconv.FormatUint(uint64(params.Page), 10))
	q.Set("num_pages", strconv.FormatUint(uint64(params.NumPages), 10))

	if params.Country != "" {
		q.Set("country", params.Country)
	}
	
	if params.Language != "" {
		q.Set("language", params.Language)
	}

	if params.WorkFromHome != "" {
		q.Set("work_from_home", params.WorkFromHome)
	}

	if len(params.EmploymentTypes) > 0 {
        q.Set("employment_types", strings.Join(params.EmploymentTypes, ","))
    }

    if len(params.ExperienceLevels) > 0 {
        q.Set("job_requirements", params.ExperienceLevels)
    }

    if len(params.ExcludeJobPublishers) > 0 {
        q.Set("exclude_job_publishers", strings.Join(params.ExcludeJobPublishers, ","))
    }

	if c.fields != "" {
		q.Set("fields", c.fields)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create request: %w", err)
	}

	if c.apiKey != "" {
		req.Header.Add("x-rapidapi-key", c.apiKey)
	}

	if c.apiHost != "" {
		req.Header.Add("x-rapidapi-host", c.apiHost)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	var out models.JSearchResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return "", nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return u.String(), &out, nil 
}



