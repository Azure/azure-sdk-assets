Set-Strict -Version 4

param(
    [string] $RepoWithTags, 
    [string] $Since = $null
)

[DateTime]$SinceDate = $null
if ($Since -ne $null) {
    $SinceDate = [DateTime]::Parse($Date)
}
else {
    $SinceDate = [DateTime]::UtcNow.AddMinutes(-20)
}

$MATRIX_BASE = [PSCustomObject]@{
    "matrix" = @{
        "Agent" = @{
            "ubuntu-20.04" = @{
                "OSVmImage" = "env:LINUXVMIMAGE"
                "Pool" = "env:LINUXPOOL"
                "os" = "linux"
            }
        }
        "Tag" = @()
    }
}

try {
    Push-Location $RepoWithTags


}
finally {
    Pop-Location
}