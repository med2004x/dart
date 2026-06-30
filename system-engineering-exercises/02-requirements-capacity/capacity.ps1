param(
    [int]$DailyUsers = 25000,
    [int]$RequestsPerUser = 30,
    [int]$PeakMultiplier = 12,
    [int]$WritesPerUser = 4,
    [int]$BytesPerWrite = 900
)

$dailyRequests = $DailyUsers * $RequestsPerUser
$averageRps = $dailyRequests / 86400
$peakRps = $averageRps * $PeakMultiplier
$dailyStorageBytes = $DailyUsers * $WritesPerUser * $BytesPerWrite
$yearlyStorageGB = $dailyStorageBytes * 365 / 1GB

[pscustomobject]@{
    DailyRequests  = $dailyRequests
    AverageRps     = [math]::Round($averageRps, 2)
    PeakRps        = [math]::Round($peakRps, 2)
    YearlyStorageGB = [math]::Round($yearlyStorageGB, 2)
}
