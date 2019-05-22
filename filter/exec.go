package filter

import (
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/types"
	"os/exec"
	"strings"
)

func init() {
	Filters.Register(
		rbxmk.Filter{Name: "exec", Func: Execute},
	)
}

func Execute(f rbxmk.FilterArgs, opt *rbxmk.Options, arguments []interface{}) (results []interface{}, err error) {
	data := arguments[0].(interface{})
	command := arguments[1].(string)
	args := make([]string, len(arguments)-2)
	for i, a := range arguments[2:] {
		args[i] = a.(string)
	}
	f.ProcessedArgs()

	result, err := types.AsString(func(s *types.Stringlike) error {
		cmd := exec.Command(command, args...)
		cmd.Stdin = strings.NewReader(s.GetString())
		var buf strings.Builder
		cmd.Stdout = &buf
		if err := cmd.Run(); err != nil {
			return err
		}
		s.SetFrom(buf.String())
		return nil
	}, data)

	if err != nil {
		return nil, err
	}
	results = append(results, result)
	return results, nil
}
