$ErrorActionPreference = 'Stop'

$assets = $PSScriptRoot
$package = Split-Path -Parent $assets
$amd64Object = Join-Path $package 'rsrc_windows_amd64.syso'
$arm64Object = Join-Path $package 'rsrc_windows_arm64.syso'
$coffFileHeaderSize = 20
$coffSectionHeaderSize = 40
$coffRelocationSize = 10
$machineAmd64 = 0x8664
$machineArm64 = 0xAA64
$relocationAmd64Addr32Nb = 0x0003
$relocationArm64Addr32Nb = 0x0002

$windres = Get-Command windres -ErrorAction SilentlyContinue
if (-not $windres) { throw 'GNU windres is required on PATH (MSYS2 UCRT64: pacman -S mingw-w64-ucrt-x86_64-binutils).' }

Push-Location -LiteralPath $assets
try {
    & $windres.Source --input launcher.rc --output-format=coff --target=pe-x86-64 --output $amd64Object
    if ($LASTEXITCODE -ne 0) { throw "windres failed with exit code $LASTEXITCODE." }
} finally {
    Pop-Location
}
Write-Output "wrote $amd64Object"

$coff = [System.IO.File]::ReadAllBytes($amd64Object)
$machine = [BitConverter]::ToUInt16($coff, 0)
if ($machine -ne $machineAmd64) { throw ('windres produced machine 0x{0:X4}, expected 0x{1:X4}.' -f $machine, $machineAmd64) }
[BitConverter]::GetBytes([uint16]$machineArm64).CopyTo($coff, 0)
$sectionCount = [BitConverter]::ToUInt16($coff, 2)
$sectionTable = $coffFileHeaderSize + [BitConverter]::ToUInt16($coff, 16)
for ($section = 0; $section -lt $sectionCount; $section++) {
    $sectionHeader = $sectionTable + $coffSectionHeaderSize * $section
    $relocationTable = [BitConverter]::ToUInt32($coff, $sectionHeader + 24)
    $relocationCount = [BitConverter]::ToUInt16($coff, $sectionHeader + 32)
    for ($relocation = 0; $relocation -lt $relocationCount; $relocation++) {
        $typeOffset = $relocationTable + $coffRelocationSize * $relocation + 8
        $type = [BitConverter]::ToUInt16($coff, $typeOffset)
        if ($type -ne $relocationAmd64Addr32Nb) { throw ('relocation type 0x{0:X4} has no ARM64 mapping; only IMAGE_REL_AMD64_ADDR32NB is produced for resources.' -f $type) }
        [BitConverter]::GetBytes([uint16]$relocationArm64Addr32Nb).CopyTo($coff, $typeOffset)
    }
}
[System.IO.File]::WriteAllBytes($arm64Object, $coff)
Write-Output "wrote $arm64Object"
