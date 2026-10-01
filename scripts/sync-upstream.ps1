param(
    [Parameter(Mandatory = $true)]
    [ValidatePattern('^v[0-9]+\.[0-9]+\.[0-9]+$')]
    [string]$Release
)

$ErrorActionPreference = 'Stop'
Set-Location (Split-Path -Parent $PSScriptRoot)

function Invoke-Git {
    & git @args
    if ($LASTEXITCODE -ne 0) { throw "Git failed: git $args" }
}

if (git status --porcelain) {
    throw 'Commit or stash your changes before starting an upstream update.'
}
if ((git remote get-url origin) -ne 'https://github.com/sitescomputer/SCR-Mailer.git') {
    throw 'The origin remote must point to sitescomputer/SCR-Mailer.'
}
if ((git remote) -notcontains 'upstream') {
    Invoke-Git remote add upstream https://github.com/knadh/listmonk.git
}
if ((git remote get-url upstream) -ne 'https://github.com/knadh/listmonk.git') {
    throw 'The upstream remote must point to knadh/listmonk.'
}
Invoke-Git remote set-url --push upstream DISABLED
Invoke-Git fetch origin
# Namespace upstream tags so importing them never creates an SCR release tag.
Invoke-Git fetch --no-tags upstream '+refs/heads/master:refs/remotes/upstream/master' '+refs/tags/*:refs/tags/upstream/*'
Invoke-Git rev-parse --verify "refs/tags/upstream/$Release^{commit}"
& git merge-base --is-ancestor "refs/tags/upstream/$Release" origin/master
if ($LASTEXITCODE -eq 0) {
    Write-Host "Upstream $Release is already included in origin/master. No integration branch is needed."
    return
}
if ($LASTEXITCODE -ne 1) { throw 'Unable to compare upstream release ancestry.' }
$updateBranch = "update/upstream-$Release"
Invoke-Git switch -c $updateBranch origin/master
& git merge --no-ff --no-edit "refs/tags/upstream/$Release"
if ($LASTEXITCODE -ne 0) {
    throw 'Resolve the conflicts, git add the resolved files, then run git merge --continue. Preserve SCR branding while integrating upstream functionality.'
}
Write-Host "Update prepared on $updateBranch. Build, test and review branding before pushing."
Write-Host "git push -u origin $updateBranch"
Write-Host 'Open a pull request into master. Merge it using a merge commit to preserve upstream ancestry.'
