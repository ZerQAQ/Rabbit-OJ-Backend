package config

var (
	Secret string
)

const (
	PageSize            = 20
	DefaultExchangeName = "rabbit.oj"

	JudgeQueueName        = "judge"
	JudgeRoutingKey       = "judge"
	JudgeResultQueueName  = "judge_result"
	JudgeResultRoutingKey = "judge_result"
)
