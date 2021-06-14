# Summary
List of document fragment paths.

# Description
The **fragments** dump format lists the fragment references made by the
program.

A fragment is a chunk of text that documents some particular element.
Internally, fragments are represented by a collection of files. A fragment
reference is a combination of a file path, which drills into the file tree, and
a "section" path, which drills into the sections of the referred file.

The following top-level directories are defined:

- **Libraries**: Descriptions of items within libraries.
- **Types**: Descriptions of data types.
- **Commands**: Descriptions of sub-commands.
- **DumpFormats**: Descriptions of dump formats (like this one).
- **Flags**: Descriptions of flags that apply to multi sub-commands.
