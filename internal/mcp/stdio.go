package mcp

import (
	"bufio"
	"fmt"
	"io"
)

func (s *Server) ServeJSONRPCLines(in io.Reader, out io.Writer) error {
	if s == nil {
		return fmt.Errorf("mcp server is nil")
	}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		response := s.HandleJSONRPCLine(line)
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
