package shared

type ProjectRunStatusEnum string

const (
	ProjectRunStatusEnumPending                ProjectRunStatusEnum = "PENDING"
	ProjectRunStatusEnumRunning                ProjectRunStatusEnum = "RUNNING"
	ProjectRunStatusEnumErrored                ProjectRunStatusEnum = "ERRORED"
	ProjectRunStatusEnumCompleted              ProjectRunStatusEnum = "COMPLETED"
	ProjectRunStatusEnumKilled                 ProjectRunStatusEnum = "KILLED"
	ProjectRunStatusEnumUnableToAllocateKernel ProjectRunStatusEnum = "UNABLE_TO_ALLOCATE_KERNEL"
)
