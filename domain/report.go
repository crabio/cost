package domain

type Report struct {
	// Map with by node reports
	// Key - node id
	// Value - node resources report
	NodeReports map[string]*NodeReport
}

func NewReport() *Report {
	r := new(Report)

	r.NodeReports = make(map[string]*NodeReport)

	return r
}
