// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao/internal"
)

// failedJobsDao is the data access object for table yjs_failed_jobs.
// You can define custom methods on it to extend its functionality as you wish.
type failedJobsDao struct {
	*internal.FailedJobsDao
}

var (
	// FailedJobs is globally public accessible object for table yjs_failed_jobs operations.
	FailedJobs = failedJobsDao{
		internal.NewFailedJobsDao(),
	}
)

// Fill with you ideas below.
