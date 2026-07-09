[CmdletBinding()]
param(
    [string]$Root = (Resolve-Path (Join-Path $PSScriptRoot "..\..")).Path,
    [int]$Port = 8184,
    [string]$HostAddress = "127.0.0.1",
    [switch]$Stateless
)

$ErrorActionPreference = "Stop"

if ($Port -lt 1 -or $Port -gt 65535) {
    throw "Port must be between 1 and 65535."
}

$uv = (Get-Command uv -ErrorAction Stop).Source
$server = Join-Path $PSScriptRoot "server.py"

if (-not (Test-Path -LiteralPath $server -PathType Leaf)) {
    throw "MCP server not found: $server"
}

$rootPath = (Resolve-Path -LiteralPath $Root).Path

Write-Host "USDM MCP root: $rootPath"
Write-Host "Streamable HTTP: http://${HostAddress}:$Port/mcp"
Write-Host "SSE:             http://${HostAddress}:$Port/sse"
Write-Host "Status:          http://${HostAddress}:$Port/status"

$arguments = @(
    "run",
    "--project", $PSScriptRoot,
    "mcp-proxy",
    "--host=$HostAddress",
    "--port=$Port",
    "--cwd=$rootPath",
    "--pass-environment",
    "--env", "USDM_MCP_ROOT", $rootPath
)

if ($Stateless) {
    $arguments += "--stateless"
}

$arguments += @(
    "--",
    $uv,
    "run",
    "--project", $PSScriptRoot,
    "python", $server
)

& $uv @arguments
exit $LASTEXITCODE
