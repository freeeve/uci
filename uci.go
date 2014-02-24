package uci

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Options struct {
	MultiPV int
	Hash    int
	Ponder  bool
	OwnBook bool
}

type Results struct {
	BestMove string
	score    int
}

type Engine struct {
	cmd    *exec.Cmd
	stdout *bufio.Reader
	stdin  *bufio.Writer
}

func NewEngine(path string) (*Engine, error) {
	eng := Engine{}
	eng.cmd = exec.Command(path)
	stdin, err := eng.cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := eng.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := eng.cmd.Start(); err != nil {
		return nil, err
	}
	eng.stdin = bufio.NewWriter(stdin)
	eng.stdout = bufio.NewReader(stdout)
	return &eng, nil
}

func (eng *Engine) SetFEN(fen string) error {
	_, err := eng.stdin.WriteString(fmt.Sprintf("position fen %s\n", fen))
	if err != nil {
		return err
	}
	err = eng.stdin.Flush()
	return err
}

func (eng *Engine) GoDepth(depth int) (*Results, error) {
	res := Results{}
	_, err := eng.stdin.WriteString(fmt.Sprintf("go depth %d\n", depth))
	if err != nil {
		return nil, err
	}
	err = eng.stdin.Flush()
	if err != nil {
		return nil, err
	}
	for {
		line, err := eng.stdout.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.Trim(line, "\n")
		if strings.HasPrefix(line, "bestmove") {
			dummy := ""
			_, err := fmt.Sscanf(line, "%s %s", &dummy, &res.BestMove)
			if err != nil {
				return nil, err
			}
			break
		}
	}
	return &res, nil
}
