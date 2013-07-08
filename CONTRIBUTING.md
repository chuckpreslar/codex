# contributing to codex

To make things easy, I suggest making a fork and installing a local copy of the codex source from it.

    $ go get -u github.com/<username>/codex

## Formatting

Standard `gofmt` formatting is used with one minor exception - a tab width of 8 is rediculous, please run `gofmt` with flag `-tabwidth=2` to adjust tabs for consistency across the project.

## Commiting

When making a commit aimed at an issue for a milestone (future release), please make the commit toward the branch for the milestone's version.

If the commit is targetting a bug or something that requires immediate attention for the current release, then apply the commit to the current version's branch.
