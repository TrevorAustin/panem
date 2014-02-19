# Panem
Panem is a program for selecting the presenter for Monday All-Staff meetings.
# Selecting a New Tribute
`go run main.go`
# Recording a Presenter
After someone presents, add a new line to the bottom of weeks.csv. The first column is the employee's ID number (which is just the line number they appear on in employees.csv minus 1). The second and third columns are their name and the date they led the meeting.  Those columns aren't used by the program, but make bookkeeping easier.
# Adding New Hires
Add a new line to the bottom of employees.csv. The first column is the employee's name and the second column is the week they started (counting from Jan 1 2014)
