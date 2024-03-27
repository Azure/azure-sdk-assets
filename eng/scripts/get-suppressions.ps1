Set-Strict -Version 4

param(
    [string] $CredScanSuppressionFolder
)

$ErrorActionPreference = 'Stop'
New-Item -ItemType Directory -Force eng

$req = Invoke-WebRequest https://raw.githubusercontent.com/Azure/azure-sdk-assets/main/eng/CredScanSuppression.json
$req.RawContent | Out-File (Join-Path -Path $CredScanSuppressionFolder -ChildPath CredScanSuppression.json)