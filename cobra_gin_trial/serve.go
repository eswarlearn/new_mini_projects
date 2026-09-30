package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"bashenv-server/internal/config"
)

var (
	bashFile   string
	portFlag   string
	quietStart bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Load config from a bash file and start the Gin server",
	RunE:  runServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&bashFile, "file", "f", "~/.bashrc",
		"bash file to source for config (e.g. ~/.bashrc, ~/.bash_profile)")
	serveCmd.Flags().StringVarP(&portFlag, "port", "p", "",
		"override the PORT value found in the bash file")
	serveCmd.Flags().BoolVarP(&quietStart, "quiet", "q", false,
		"suppress the startup config summary")
}

func runServe(_ *cobra.Command, _ []string) error {
	env, err := config.LoadFromBashFile(bashFile)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	cfg := config.BuildServerConfig(env)

	// CLI flag wins over whatever was in the bash file.
	if portFlag != "" {
		cfg.Port = portFlag
	}

	if !quietStart {
		fmt.Printf("Loaded config from %s\n", bashFile)
		fmt.Printf("  APP_NAME = %s\n", cfg.AppName)
		fmt.Printf("  HOST     = %s\n", cfg.Host)
		fmt.Printf("  PORT     = %s\n", cfg.Port)
		fmt.Printf("  GIN_MODE = %s\n", cfg.Mode)
	}

	gin.SetMode(normalizeGinMode(cfg.Mode))
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"app":    cfg.AppName,
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	router.GET("/config", func(c *gin.Context) {
		// Only show the values we actually consumed, not the whole
		// environment — no reason to leak everything from the shell.
		c.JSON(http.StatusOK, gin.H{
			"app_name": cfg.AppName,
			"host":     cfg.Host,
			"port":     cfg.Port,
			"gin_mode": cfg.Mode,
			"source":   bashFile,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s is running", cfg.AppName),
		})
	})

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Printf("Starting %s on %s (mode=%s)\n", cfg.AppName, addr, cfg.Mode)
	return router.Run(addr)
}

func normalizeGinMode(mode string) string {
	switch mode {
	case gin.ReleaseMode, gin.TestMode, gin.DebugMode:
		return mode
	default:
		return gin.DebugMode
	}
}
