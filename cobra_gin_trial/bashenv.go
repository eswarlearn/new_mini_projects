package config

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ServerConfig holds the values our Gin server cares about.
// These are populated from whatever the sourced bash file exports.
type ServerConfig struct {
	Host    string
	Port    string
	Mode    string // gin.DebugMode / gin.ReleaseMode / gin.TestMode
	AppName string
	// Raw holds every env var found, in case you want to read
	// something beyond the four fields above.
	Raw map[string]string
}

// LoadFromBashFile sources the given bash file (e.g. ~/.bashrc) in a real
// bash subshell, dumps the resulting environment, and returns it parsed.
//
// We deliberately shell out to bash instead of hand-parsing the file:
// bashrc files can contain conditionals, sourced sub-files, command
// substitution, etc. Only bash itself can reliably evaluate that.
func LoadFromBashFile(path string) (map[string]string, error) {
	expanded, err := expandPath(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(expanded); err != nil {
		return nil, fmt.Errorf("bash file not found: %s", expanded)
	}

	// -i (interactive) makes bash actually process a .bashrc, which
	// otherwise skips itself in non-interactive shells. We still run
	// it non-login so it doesn't also pull in /etc/profile chains.
	script := fmt.Sprintf("source %s && env", shellQuote(expanded))
	cmd := exec.Command("bash", "-i", "-c", script)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to source %s: %w (stderr: %s)", expanded, err, stderr.String())
	}

	return parseEnvOutput(stdout.Bytes()), nil
}

// BuildServerConfig extracts the keys our server understands from the raw
// env map, applying sensible defaults for anything that wasn't exported.
func BuildServerConfig(env map[string]string) ServerConfig {
	cfg := ServerConfig{
		Host:    getOrDefault(env, "HOST", "0.0.0.0"),
		Port:    getOrDefault(env, "PORT", "8080"),
		Mode:    getOrDefault(env, "GIN_MODE", "debug"),
		AppName: getOrDefault(env, "APP_NAME", "bashenv-server"),
		Raw:     env,
	}
	return cfg
}

func getOrDefault(env map[string]string, key, fallback string) string {
	if v, ok := env[key]; ok && v != "" {
		return v
	}
	return fallback
}

func parseEnvOutput(out []byte) map[string]string {
	result := make(map[string]string)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	// env output can contain multi-line values (e.g. function exports);
	// we only care about simple KEY=VALUE lines, so skip anything that
	// doesn't look like one.
	for scanner.Scan() {
		line := scanner.Text()
		if idx := strings.Index(line, "="); idx > 0 {
			key := line[:idx]
			val := line[idx+1:]
			if isValidEnvKey(key) {
				result[key] = val
			}
		}
	}
	return result
}

func isValidEnvKey(key string) bool {
	if key == "" {
		return false
	}
	for i, r := range key {
		if r == '_' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			continue
		}
		if i > 0 && r >= '0' && r <= '9' {
			continue
		}
		return false
	}
	return true
}

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, strings.TrimPrefix(path, "~")), nil
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}
