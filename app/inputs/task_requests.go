package inputs

type SaveTaskRequest struct {
	Name               string            `json:"name" binding:"required"`
	Description        string            `json:"description"`
	IntervalValue      string            `json:"interval_value"`
	IntervalUnit       string            `json:"interval_unit"`
	ScheduleExpression string            `json:"schedule_expression"`
	Timezone           string            `json:"timezone" binding:"required"`
	Endpoint           string            `json:"endpoint" binding:"required,url"`
	Method             string            `json:"method" binding:"required,oneof=GET POST PUT DELETE"`
	Headers            map[string]string `json:"headers"`
	Body               string            `json:"body"`
	Timeout            string            `json:"timeout"`
	RetryCount         string            `json:"retry_count"`
	RetryInterval      string            `json:"retry_interval"`
}
