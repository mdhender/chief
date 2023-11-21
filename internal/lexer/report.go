// chief - a TribeNet player aid
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package lexer

// Report is the representation of a single parsed template.
type Report struct {
	FileName string // name of the report document.
	ClanId   string // clan identifier (4 digits, starts with '0')
}

// NewReport allocates a new report with the given name and clan id.
func NewReport(filename string) *Report {
	return &Report{
		FileName: filename,
	}
}

// Copy returns a copy of the Report. Any parsing state is discarded.
func (r *Report) Copy() *Report {
	if r == nil {
		return nil
	}
	return &Report{
		FileName: r.FileName,
		ClanId:   r.ClanId,
	}
}
