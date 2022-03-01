# Summary
Display API documentation.

# Arguments
	REFERENCE

# Description
The **doc** command displays documentation for a given API or topic.

Topics are stored in a tree, where each topic may have sub-topics. A
"reference", similar to a file path, can refer to a topic, or a section within a
topic. The presence of a `:` separator causes the reference to begin drilling
into the sections of the topic to the left of the separator. The topic portion
of a reference is case-insensitive, while the section portion is case-sensitive.
For example,

	rbxmk doc types/Instance:Methods/FindFirstChild

This reference will start at the top-level "types" topic, then drill into the
"Instance" sub-topic, then drill into the "Method" section, then the
"FindFirstChild" sub-section. The doc command will then display the content of
this sub-section only.

If the referred topic has no content, but has sub-topics, then these sub-topics
are listed instead (see --list). If no reference is specified, then all
top-level topics are listed.

The following top-level topics are available:
{{Topics}}

# Flags
## list
Returns a list of sub-topics and sub-sections for the given reference. If no
reference is specified, then all top-level topics are listed.

The top-level sections of a topic will be prefixed with ":". Sub-topics and
sub-sections of sections will be prefixed with "/". That is, the whole name of a
listed item, including the prefix, can be appended to the current reference to
query that item.
