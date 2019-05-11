<h1 align="center">
 gobc - Go branch clean
</h1>

## Purpose

My typical workflow consists of making new branch with needed changes, then making PR with it and deleting it afterwards, when it is no logner needed.

Unfortunately it leaves me with useless local branches. That's why `gobc` was created - it checks which local branches are not on remote (origin) and deletes them. 

## Usage

Just type `gobc` to delete branches that are not on remote (origin). Use `-l` flag to list local and origin branches, without deleting them.