# Summary
Enter interactive mode.

# Description
Enters interactive mode. Each prompt executes a chunk of Lua code.

If a prompt begins with `=`, then the comma-separated list of expressions that
follow are evaluated and printed to standard output.

The environment contains the **os.exit** function. When called, interactive mode
is terminated, and the program exits.

Within supported terminals, the following shortcuts are available:

Shortcuts             | Description
----------------------|------------
Ctrl-A, Home          | Move cursor to beginning of line.
Ctrl-E, End           | Move cursor to end of line
Ctrl-B, Left          | Move cursor one character left.
Ctrl-F, Right         | Move cursor one character right.
Ctrl-Left, Alt-B      | Move cursor to previous word.
Ctrl-Right, Alt-F     | Move cursor to next word
Ctrl-D, Del           | If line is not empty, delete character under cursor.
Ctrl-D                | If line is empty, end of file.
Ctrl-C                | Reset input (create new empty prompt).
Ctrl-L                | Clear screen (line is unmodified).
Ctrl-T                | Transpose previous character with current character.
Ctrl-H, BackSpace     | Delete character before cursor.
Ctrl-W, Alt-BackSpace | Delete word leading up to cursor.
Alt-D                 | Delete word following cursor.
Ctrl-K                | Delete from cursor to end of line.
Ctrl-U                | Delete from start of line to cursor.
Ctrl-P, Up            | Previous match from history.
Ctrl-N, Down          | Next match from history.
Ctrl-R                | Reverse Search history (Ctrl-S forward, Ctrl-G cancel).
Ctrl-Y                | Paste from Yank buffer (Alt-Y to paste next yank instead).
Tab                   | Next completion.
Shift-Tab             | (after Tab) Previous completion.
