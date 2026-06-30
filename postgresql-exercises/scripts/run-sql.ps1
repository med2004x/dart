param(
    [Parameter(Mandatory = $true)]
    [string]$Path
)

$resolved = Resolve-Path -LiteralPath $Path
Get-Content -Raw -LiteralPath $resolved |
    docker exec -i postgres-course `
    psql -X -v ON_ERROR_STOP=1 -U student -d go_course

if ($LASTEXITCODE -ne 0) {
    exit $LASTEXITCODE
}
