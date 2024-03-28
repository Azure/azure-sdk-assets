param(
    [string] $MatrixFile,
    [string] $RepoWithTags, 
    [string] $Since = $null
)

[DateTime]$SinceDate = [DateTime]::UtcNow.AddMinutes(-20)
if (($Since -ne $null -or $Since -ne "") -and $Since -ne "<default to now() - 20>") {
    $SinceDate = [DateTime]::Parse($Since)
}

$success = $true
$MATRIX_BASE = [PSCustomObject]@{
    "matrix" = @{
        "Agent" = @{
            "ubuntu-20.04" = @{
                "OSVmImage" = "azsdk-pool-mms-ubuntu-2004-1espt"
                "Pool" = "azsdk-pool-mms-ubuntu-2004-general"
            }
        }
        "Tags" = @()
    }
}

try {
    Push-Location $RepoWithTags
    git fetch --tags

    if ($LASTEXITCODE -ne 0) {
        Write-Error "Unable to fetch tags for $RepoWithTags"
        $success = $false
    }

    $allTags = git log --date=iso --tags --simplify-by-decoration --pretty="format:%ai::%d"
    $tagObjects = @()
    $results = @()
  
    # parse the log output to get the latest tags
    foreach($tagObject in $allTags)
    {
      $parts = $tagObject.Split("::")
      foreach($part in $parts){ $part = $part.Trim() }
  
      $TagRegex = "tag\: (?<tag>[^,)]*)"
      $tagList = @()
  
      Select-String $TagRegex -input $parts[1] -AllMatches `
        | % { $_.matches } `
        | % { $tagList += $_.Groups['tag'].Value }
  
      $tagObjects += New-Object PSObject -Property @{
        Date = [DateTime]::Parse($parts[0])
        TagList = $tagList
      }
    }
    
    $filteredTagObjects = $tagObjects | Where-Object { $_.TagList.Length -gt 0 -and $_.Date -ge $SinceDate}
    $filteredTagObjects | Sort-Object -Property Date -Descending
    
    foreach($tagObject in $filteredTagObjects)
    {
      foreach($tag in $tagObject.TagList)
      {
        $results += $tag
      }
    }

    $MATRIX_BASE.matrix.Tags = $results | Sort-Object -Unique
}
finally {
    Pop-Location
}

if ($success) {
    $MATRIX_BASE | ConvertTo-Json -Depth 40 | Out-File $MatrixFile
}
else {
    Write-Error "Failed to generate tag matrix"
    exit 1
}