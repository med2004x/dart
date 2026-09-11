param(
    [string]$DumpPath = ".\database-backups\go_course.dump",
    [string]$TargetDatabase = "go_course_restore"
)

$resolvedDump = Resolve-Path -LiteralPath $DumpPath

docker exec postgres-course `
    createdb -U student $TargetDatabase
if ($LASTEXITCODE -ne 0) {
    Write-Error "Target database must not already exist: $TargetDatabase"
    exit $LASTEXITCODE
}

docker cp $resolvedDump "postgres-course:/tmp/go_course_restore.dump"
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

docker exec postgres-course `
    pg_restore -U student -d $TargetDatabase /tmp/go_course_restore.dump
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

docker exec postgres-course `
    psql -X -U student -d $TargetDatabase `
    -c "SELECT count(*) AS task_count FROM app.tasks;"
