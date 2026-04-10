package mcp

import (
	"strings"

	mcpServer "github.com/52dev/kratos-transport/transport/mcp"
	"github.com/mark3labs/mcp-go/server"

	conf "github.com/52dev/kratos-bootstrap/api/gen/go/conf/v1"
)

// NewMcpServer creates a new MCP server.
func NewMcpServer(cfg *conf.Server_MCP, opts ...mcpServer.ServerOption) *mcpServer.Server {
	if cfg == nil {
		return nil
	}

	var o []mcpServer.ServerOption

	if cfg.GetServerName() != "" {
		o = append(o, mcpServer.WithServerName(cfg.GetServerName()))
	}
	if cfg.GetServerVersion() != "" {
		o = append(o, mcpServer.WithServerVersion(cfg.GetServerVersion()))
	}
	if cfg.GetServerType() != "" {
		serverType := convertMcpServerType(cfg.GetServerType())
		o = append(o, mcpServer.WithMCPServeType(serverType))
	}
	if cfg.GetAddress() != "" {
		o = append(o, mcpServer.WithMCPServeAddress(cfg.GetAddress()))
	}

	var mcpOpts []server.ServerOption
	if cfg.GetRecovery() {
		mcpOpts = append(mcpOpts, server.WithRecovery())
	}
	if cfg.GetLogging() {
		mcpOpts = append(mcpOpts, server.WithLogging())
	}
	if cfg.GetToolCapabilities() {
		mcpOpts = append(mcpOpts, server.WithToolCapabilities(cfg.GetToolCapabilities()))
	}
	if cfg.GetPromptCapabilities() {
		mcpOpts = append(mcpOpts, server.WithPromptCapabilities(cfg.GetPromptCapabilities()))
	}
	if cfg.GetElicitation() {
		mcpOpts = append(mcpOpts, server.WithElicitation())
	}
	if cfg.GetRoots() {
		mcpOpts = append(mcpOpts, server.WithRoots())
	}

	if len(mcpOpts) > 0 {
		o = append(o, mcpServer.WithMCPServerOptions(mcpOpts...))
	}

	if opts != nil {
		o = append(o, opts...)
	}

	srv := mcpServer.NewServer(o...)

	return srv
}

func convertMcpServerType(serverType string) mcpServer.ServerType {
	switch strings.TrimSpace(strings.ToLower(serverType)) {
	case "http":
		return mcpServer.ServerTypeHTTP
	case "sse":
		return mcpServer.ServerTypeSSE
	case "stdio":
		return mcpServer.ServerTypeStdio
	case "in_process":
		return mcpServer.ServerTypeInProcess
	default:
		return mcpServer.ServerTypeHTTP
	}
}
