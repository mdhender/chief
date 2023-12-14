# Turn Report Parser

## Setup
I open the `.docx` file, copy the entire document, and paste it into a text document.

I name the text file {clanNo}.{turn}.Turn-Report.txt.
For example, `0138.900-04.Turn-Report.txt`.

## Tweaks

The form-feed character is used as a sentinel to end tribe/unit sections.
I ensure that the character is present and that there's one at the end of the last tribe/unit.

### Initial Set-Up Report
The initial set-up report seems to be hand generated.
It needs to have some things added to parse successfully.

### Regular turns

You must place a form-feed after the last unit report, before Transfers and before Settlements.
