/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package structs

// QueryParams defines the struct of query by page params
type QueryParams struct {
	PageNum   int   // page num
	PageLimit int   // page limit
	StartTime int64 // start time
	EndTime   int64 // end time
	Total     int64 // total number of query results
}
