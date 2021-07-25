package shell

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Sourcer interface {
	Source(context.Context, string) error
}

func NewSourcer() Sourcer {
	return &sourcer{}
}

type sourcer struct{}

func (s *sourcer) Source(ctx context.Context, filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("the provided path is a directory (%s); expected to be a file", filename)
	}

	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", fmt.Sprintf("source %s; env", filename))
	cmd.Stderr = os.Stderr
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		return err
	}

	vv := bytes.Split(buf.Bytes(), []byte{'\n'})
	for _, v := range vv {
		kv := strings.SplitN(string(v), "=", 2)
		if len(kv) == 2 {
			fmt.Printf("export %s=\"%s\"\n", kv[0], kv[1])
		}
	}

	return nil
}
