param(
    [string]$Uri = "http://localhost:8080/tasks",
    [int]$Count = 20
)

$samples = 1..$Count | ForEach-Object {
    (Measure-Command {
        Invoke-WebRequest -Uri $Uri -Method Post -ContentType "application/json" `
            -Body '{"title":"measure notification"}' | Out-Null
    }).TotalMilliseconds
}

$samples | Measure-Object -Minimum -Maximum -Average
