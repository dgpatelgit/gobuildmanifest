package internal

import (
	"io"
	"os/exec"
)

// GoListCmd ... Go list command structure.
type GoListCmd struct {
	cmd    *exec.Cmd
	output io.ReadCloser
}

func GetGoExecutable() (*GoListCmd, error) {
	goList := exec.Command("which", "go")
	output, err := goList.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = goList.Start(); err != nil {
		return nil, err
	}
	return &GoListCmd{cmd: goList, output: output}, nil
}

// RunGoList ... Actual function that executes go list command and returns output as string.
func RunGoList(cwd string, goExec string) (*GoListCmd, error) {
	/*	cmd, err := GetGoExecutable()
		if err != nil {
			log.Error().Err(err).Msg("`which go` failed")
			return nil, err
		}

		defer cmd.ReadCloser().Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(cmd.ReadCloser())
		s := strings.TrimSpace(buf.String())
		log.Info().Msg("Found go executable " + s)

		// Wait for the `go list` command to complete.
		//if err := cmd.Wait(); err != nil {
		//	return nil, fmt.Errorf("%v: `go list` failed, use `go mod tidy` to known more", err)
		//}
	*/
	goList := exec.Command(goExec, "list", "-json", "-deps", "-mod=readonly", "./...")
	goList.Dir = cwd
	output, err := goList.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = goList.Start(); err != nil {
		return nil, err
	}
	return &GoListCmd{cmd: goList, output: output}, nil
}

// ReadCloser implements internal.GoList
func (list *GoListCmd) ReadCloser() io.ReadCloser {
	return list.output
}

// Wait implements internal.GoList
func (list *GoListCmd) Wait() error {
	return list.cmd.Wait()
}
