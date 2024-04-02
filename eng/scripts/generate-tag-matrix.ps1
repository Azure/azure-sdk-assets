param(
    [string] $RepoWithTags, 
    [string] $Since = $null
)

Set-StrictMode -Version 4

$DisallowedCharactersRegex = (@('\', '/', "'", ':', '<', '>', '|', '*', '?') | ForEach-Object { [regex]::Escape($_) }) -join '|'

function SanitizeTagForMatrix {
  param(
    [string] $tag
  )
  $placeholder = '_'
  $artifactName = $tag -replace $pattern, $placeholder
  return $artifactName
}


[DateTime]$SinceDate = [DateTime]::UtcNow.AddMinutes(-305)
if (($Since -ne $null -or $Since -ne "") -and $Since -ne "<default to now() - 6 hours>") {
    $SinceDate = [DateTime]::Parse($Since)
}

$success = $true

# format is a dictionary with the following structure:
#{
#  "<tag_name_sanitized>": {
#    "Tag": "<raw_tag_name>" 
#  } 
#}

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

    $MATRIX_BASE.matrix.Tag = $results | Sort-Object -Unique
}
finally {
    Pop-Location
}
