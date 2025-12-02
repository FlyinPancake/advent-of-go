package aoc

import (
	"fmt"
	"os"
	"text/template"
)

const templatesLoc = "solutions/template"
const solutionTeplate = templatesLoc + "/dayXX.go.tpl"
const testTemplate = templatesLoc + "/dayXX_test.go.tpl"

type SolutionTemplateData struct {
	Day       int
	DayPadded string
	Year      int
}

func formatDayPadded(day int) string {
	return fmt.Sprintf("%02d", day)
}

func NewSolutionTemplateData(year, day int) SolutionTemplateData {
	return SolutionTemplateData{
		Day:       day,
		DayPadded: formatDayPadded(day),
		Year:      year,
	}
}

func CreateNewSolutionScaffold(year, day int) error {
	data := NewSolutionTemplateData(year, day)
	err := os.MkdirAll(fmt.Sprintf("solutions/%d/day%s", year, data.DayPadded), 0755)
	if err != nil {
		return err
	}
	err = createFileFromTemplate(solutionTeplate, fmt.Sprintf("solutions/%d/day%s/day%02d.go", year, data.DayPadded, day), data)
	if err != nil {
		return err
	}
	err = createFileFromTemplate(testTemplate, fmt.Sprintf("solutions/%d/day%s/day%02d_test.go", year, data.DayPadded, day), data)
	if err != nil {
		return err
	}
	_, err = os.Create(fmt.Sprintf("inputs/%d/day%s.txt", year, data.DayPadded))
	if err != nil {
		return err
	}
	return nil
}

func createFileFromTemplate(templatePath, outputPath string, data SolutionTemplateData) error {
	templateTextBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}
	tmpl, err := template.New("tpl").Parse(string(templateTextBytes))
	if err != nil {
		return err
	}
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = tmpl.Execute(outFile, data)
	if err != nil {
		return err
	}
	// Successfully created file from
	return nil
}
