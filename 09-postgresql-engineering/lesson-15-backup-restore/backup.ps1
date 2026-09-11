param(
    [string]$OutputDirectory = ".\database-backups"
)

New-Item -ItemType Directory -Force $OutputDirectory | Out-Null
$outputPath = Join-Path $OutputDirectory "go_course.dump"

docker exec postgres-course `
    pg_dump -U student -d go_course -Fc -f /tmp/go_course.dump
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

docker cp "postgres-course:/tmp/go_course.dump" $outputPath
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Get-Item -LiteralPath $outputPath
