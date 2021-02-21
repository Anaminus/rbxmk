package formats

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

type l10nCSVEntry struct {
	Key      string            `json:"key,omitempty"`
	Context  string            `json:"context,omitempty"`
	Examples string            `json:"examples,omitempty"`
	Source   string            `json:"source,omitempty"`
	Values   map[string]string `json:"values"`
}

type orderedMap map[string]string

func (m orderedMap) UnmarshalJSON(b []byte) error {
	return nil
}

// Number of index headers.
const l10nIndexHeaders = 4

func decodeL10nCSV(r io.Reader) (j []byte, err error) {
	cr := csv.NewReader(r)
	headers, err := cr.Read()
	if err != nil {
		if err == io.EOF {
			return []byte("[]"), nil
		}
		return nil, err
	}
	// Headers that have been mapped.
	mappedHeaders := map[string]int{}
	// Mapping of record index to entry index.
	headerMap := map[int]int{}
	// List of locale headers.
	headerNames := []string{
		"Key",
		"Context",
		"Example",
		"Source",
	}
	// Current entry index of locale header.
	n := l10nIndexHeaders

	// Map unordered headers to ordered indices.
loop:
	for i, header := range headers {
		for j, name := range headerNames[:l10nIndexHeaders] {
			if header != name {
				continue
			}
			if _, ok := mappedHeaders[header]; !ok {
				mappedHeaders[header] = i
				headerMap[i] = j
				continue loop
			}
		}
		lheader := strings.ToLower(header)
		if j, ok := mappedHeaders[lheader]; ok {
			return nil, fmt.Errorf("column %d (%q) conflicts with column %d", i, header, j)
		}
		mappedHeaders[lheader] = i
		headerMap[i] = n
		headerNames = append(headerNames, lheader)
		n++

	}
	if _, ok := mappedHeaders["Key"]; !ok {
		if _, ok := mappedHeaders["Source"]; !ok {
			return nil, errors.New("missing Key or Source header column")
		}
	}

	type tuple struct {
		context string
		source  string
	}
	kIndex := map[string]int{}
	csIndex := map[tuple]int{}

	// Scan remaining entries into JSON structure.
	var entries []l10nCSVEntry
	cr.ReuseRecord = true
	entry := make([]string, n)
	for {
		rec, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		for i, v := range rec {
			if j, ok := headerMap[i]; ok {
				entry[j] = v
			}
		}

		var jentry l10nCSVEntry
		jentry.Key = entry[0]
		jentry.Context = entry[1]
		jentry.Examples = entry[2]
		jentry.Source = entry[3]

		if jentry.Key == "" && jentry.Source == "" {
			return nil, fmt.Errorf("record %d has empty key and source", len(entries))
		} else if jentry.Key != "" {
			if j, ok := kIndex[jentry.Key]; ok {
				return nil, fmt.Errorf("record %d conflicts with %d by key %q", len(entries), j, jentry.Key)
			}
			kIndex[jentry.Key] = len(entries) + 1
		} else {
			t := tuple{jentry.Context, jentry.Source}
			if j, ok := csIndex[t]; ok {
				return nil, fmt.Errorf("record %d conflicts %d by (context,source) (%q, %q)", len(entries), j, t.context, t.source)
			}
			csIndex[t] = len(entries) + 1
		}

		jentry.Values = make(map[string]string, n-l10nIndexHeaders)
		for j := l10nIndexHeaders; j < len(entry); j++ {
			key := headerNames[j]
			jentry.Values[key] = entry[j]
		}
		entries = append(entries, jentry)
	}

	if len(entries) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(entries)
}

func encodeL10nCSV(w io.Writer, b []byte) error {
	var entries []l10nCSVEntry
	if err := json.Unmarshal(b, &entries); err != nil {
		return err
	}

	type tuple struct {
		context string
		source  string
	}
	kIndex := map[string]int{}
	csIndex := map[tuple]int{}

	headers := []string{
		"Key",
		"Context",
		"Example",
		"Source",
	}
	mappedHeaders := map[string]struct{}{}
	for i, entry := range entries {
		// Ensure no locale headers conflict.
		mappedEntryHeaders := map[string]string{}
		for k := range entry.Values {
			l := strings.ToLower(k)
			if m, ok := mappedEntryHeaders[l]; ok {
				return fmt.Errorf("entry %d: value field %q conflicts with %q", i, k, m)
			}
			mappedEntryHeaders[l] = k
		}
		// Remap locale headers to lowercase and add to list of headers.
		for k, v := range entry.Values {
			l := strings.ToLower(k)
			if l != k {
				delete(entry.Values, k)
				entry.Values[l] = v
			}
			if _, ok := mappedHeaders[l]; !ok {
				mappedHeaders[l] = struct{}{}
				headers = append(headers, l)
			}
		}
		// Ensure no index headers conflict.
		if entry.Key == "" && entry.Source == "" {
			return fmt.Errorf("entry %d has empty key and source", len(entries))
		} else if entry.Key != "" {
			if j, ok := kIndex[entry.Key]; ok {
				return fmt.Errorf("entry %d conflicts with %d by key %q", len(entries), j, entry.Key)
			}
			kIndex[entry.Key] = len(entries) + 1
		} else {
			t := tuple{entry.Context, entry.Source}
			if j, ok := csIndex[t]; ok {
				return fmt.Errorf("entry %d conflicts with %d by (context,source) (%q, %q)", len(entries), j, t.context, t.source)
			}
			csIndex[t] = len(entries) + 1
		}
	}
	sort.Strings(headers[l10nIndexHeaders:])

	cw := csv.NewWriter(w)
	if err := cw.Write(headers); err != nil {
		return err
	}
	entry := make([]string, len(headers))
	for _, jentry := range entries {
		if jentry.Key == "" && jentry.Source == "" {
			continue
		}
		entry[0] = jentry.Key
		entry[1] = jentry.Context
		entry[2] = jentry.Examples
		entry[3] = jentry.Source
		for i := l10nIndexHeaders; i < len(headers); i++ {
			key := headers[i]
			entry[i] = jentry.Values[key]
		}
		if err := cw.Write(entry); err != nil {
			return err
		}
	}
	cw.Flush()
	if err := cw.Error(); err != nil {
		return err
	}
	return nil
}

func init() { register(CSV) }
func CSV() rbxmk.Format {
	return rbxmk.Format{
		Name:       "csv",
		MediaTypes: []string{"text/csv", "text/plain"},
		CanDecode: func(g rbxmk.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == "Array"
		},
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			cr := csv.NewReader(r)
			cr.ReuseRecord = true
			var vrecords rtypes.Array
			for {
				record, err := cr.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, fmt.Errorf("decode CSV: %w", err)
				}
				vrecord := make(rtypes.Array, len(record))
				for i, v := range record {
					vrecord[i] = types.String(v)
				}
				vrecords = append(vrecords, vrecord)
			}
			return vrecords, nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			if _, ok := v.(rtypes.Dictionary); ok {
				// Assume empty array, encode as no content.
				return nil
			}
			vrecords, ok := v.(rtypes.Array)
			if !ok {
				return cannotEncode(v)
			}
			cw := csv.NewWriter(w)
			var record []string
			for i, vrecord := range vrecords {
				vrecord, ok := vrecord.(rtypes.Array)
				if !ok {
					return fmt.Errorf("record %d: %w", i+1, cannotEncode(vrecord))
				}
				record = record[:0]
				for j, v := range vrecord {
					s, ok := v.(types.Stringlike)
					if !ok {
						return fmt.Errorf("record %d:%d: %w", i+1, j+1, cannotEncode(v))
					}
					record = append(record, s.Stringlike())
				}
				if err := cw.Write(record); err != nil {
					return fmt.Errorf("encode CSV: %w", err)
				}
			}
			cw.Flush()
			if err := cw.Error(); err != nil {
				return fmt.Errorf("encode CSV: %w", err)
			}
			return nil
		},
	}
}

func init() { register(L10nCSV) }
func L10nCSV() rbxmk.Format {
	return rbxmk.Format{
		Name:       "l10n.csv",
		MediaTypes: []string{"text/csv", "text/plain"},
		CanDecode:  canDecodeInstance,
		Decode: func(g rbxmk.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			b, err := decodeL10nCSV(r)
			if err != nil {
				return nil, fmt.Errorf("decode CSV: %w", err)
			}
			table := rtypes.NewInstance("LocalizationTable", nil)
			table.Set("Contents", types.String(b))
			return table, nil
		},
		Encode: func(g rbxmk.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			b, ok := rtypes.Stringable(v)
			if !ok {
				return cannotEncode(v)
			}
			if err := encodeL10nCSV(w, []byte(b)); err != nil {
				return fmt.Errorf("encode CSV: %w", err)
			}
			return nil
		},
	}
}
