# Summary
Display documentation.

# Arguments
	[ MODE ] QUERY

# Description
The **doc** command displays documentation for a given topic.

Topics are stored in a tree as "fragments", which is much like a file system,
where each topic may have sub-topics. A "reference" is similar to a file path,
but enables referring to the contents within a topic. The presence of a `:`
separator causes the reference to begin drilling into the sections of the topic
to the left of the separator.

For example, the query "libraries/base:Fields/print/Description" will display
the description of the print function, which is within the base library topic.

The topic portion of a reference is case-insensitive, while the section portion
is case-sensitive.

The doc command has several modes for querying topics. Only one mode can be used
at a time.

Mode | Name   | Description
-----|--------|------------
Raw  | (none) | Displays the content of queried topic.
List | `list` | Lists the sub-topics and sub-sections of a queried topic.

## Raw mode
With no mode specified (that is, one argument), the doc command receives a
fragment reference, and displays the content of that topic.

	rbxmk doc types/Instance:Methods/FindFirstChild

If the topic has no content, but has sub-topics, then these sub-topics are
listed instead (see List mode). If no topic is specified, then all top-level
topics are listed.

## List mode
List mode returns a list of sub-topics and sub-sections for the given reference.

	rbxmk doc list libraries/base

If "." is given as a query, this will return a list of top-level topics.

The top-level sections of a topic will be prefixed with ":". Sub-topics and
sub-sections of sections will be prefixed with "/". That is, the whole name of a
listed item, including the prefix, can be appended to the current reference to
query that item.
