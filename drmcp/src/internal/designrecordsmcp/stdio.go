package designrecordsmcp

import (
	"bufio"
	"fmt"
	"io"
)

func (s *Server) ServeJSONRPCLines(in io.Reader, out io.Writer) error {
	if s == nil {
		return fmt.Errorf("design records mcp server is nil")
	}
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		response, ok := s.HandleJSONRPCLine(line)
		if !ok {
			continue
		}
		if _, err := out.Write(response); err != nil {
			return err
		}
		if _, err := out.Write([]byte("\n")); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
