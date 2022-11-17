package formats

import (
	"archive/zip"
	"io"
	"strings"

	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/reflect"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/types"
)

const F_Zip = "zip"

func init() { register(Zip) }
func Zip() rbxmk.Format {
	return rbxmk.Format{
		MediaTypes: []string{"application/zip"},
		Name:       F_Zip,
		CanDecode: func(g rtypes.Global, f rbxmk.FormatOptions, typeName string) bool {
			return typeName == rtypes.T_Table
		},
		Decode: func(g rtypes.Global, f rbxmk.FormatOptions, r io.Reader) (v types.Value, err error) {
			var rat io.ReaderAt
			var size int64
			if srat, ok := r.(interface {
				io.Seeker
				io.ReaderAt
			}); ok {
				if size, err = srat.Seek(0, io.SeekEnd); err == nil {
					if _, err := srat.Seek(0, io.SeekStart); err == nil {
						rat = srat
					}
				}
			}
			if rat == nil {
				b, err := io.ReadAll(r)
				if err != nil {
					return nil, err
				}
				rat = strings.NewReader(string(b))
				size = int64(len(b))
			}
			z, err := zip.NewReader(rat, size)
			if err != nil {
				return nil, err
			}

			zip := make(rtypes.Dictionary, len(z.File))
			for _, file := range z.File {
				if strings.HasSuffix(file.Name, "/") {
					// Skip directory
					continue
				}
				f, err := file.Open()
				if err != nil {
					continue
				}
				b, err := io.ReadAll(f)
				f.Close()
				if err != nil {
					continue
				}
				zip[file.Name] = types.BinaryString(b)
			}
			return zip, nil
		},
		Encode: func(g rtypes.Global, f rbxmk.FormatOptions, w io.Writer, v types.Value) error {
			z := v.(rtypes.Dictionary)
			zw := zip.NewWriter(w)
			for path, value := range z {
				content, ok := value.(types.BinaryString)
				if !ok {
					continue
				}
				zf, err := zw.CreateHeader(&zip.FileHeader{
					Name:   path,
					Method: zip.Deflate,
				})
				if err != nil {
					return err
				}
				if _, err := zf.Write([]byte(content)); err != nil {
					return err
				}
			}
			return zw.Close()
		},
		Dump: func() dump.Format {
			return dump.Format{
				Hidden:      true,
				Summary:     "Formats/zip:Summary",
				Description: "Formats/zip:Description",
			}
		},
		Types: []func() rbxmk.Reflector{
			reflect.Dictionary,
		},
	}
}
