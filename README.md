# git-remover

Simple wrapper around `git` to simplify the deletion of files in a git repository

## Usage

- `git-remover ls` to list files commited
- `git-remover rm` to remove one or several files

```bash
$ git-remover ls

e69de29bb2d1      0B README.md
2185d06fcfb7     96B main.go
ed8638311058     97B README.md
d6ace8ce3633    163B go.mod
f61f43467105    245B internal/exec/shell.go
8cfd8a611c15    393B cmd/list.go
071bca3f6ee9    406B internal/exec/shellout.go
cfbf5ac16a25    485B pkg/git/ls.go
94fc347c61c1    746B cmd/list.go
0cdd1894da3f    879B pkg/git/rm.go
b5c1c2cab53c  1.3KiB cmd/rm.go
b4584143d065  1.4KiB cmd/root.go
ffff916a4a42  3.9KiB go.sum
3e3af9eaed99   12MiB git-remover
5776cac057e4   50MiB data_50mb.txt
7d4602716c7d   50MiB data_50mb.txt

$ git-remover rm data_50mb.txt git-remover

Rewrite a551588fd0d6722d472d251446abea0cff6e7f76 (4/7) (0 seconds passed, remaining 0 predicted)    rm 'data_50mb.txt'
Rewrite 23b3abc1207cfdab41e49e59114d1761092072fa (5/7) (0 seconds passed, remaining 0 predicted)    rm 'data_50mb.txt'
Rewrite 813e33b4f00302dfa1da902cab701f456da15daa (6/7) (0 seconds passed, remaining 0 predicted)    rm 'data_50mb.txt'
rm 'data_5mb.txt'
Rewrite ede5356f191ac33a79deedbd3fdbafffeff08a90 (7/7) (0 seconds passed, remaining 0 predicted)    rm 'data_50mb.txt'
rm 'git-remover'
Ref 'refs/heads/master' was rewritten

Successful, please update your remote with `git push --all -f`

$ git-remover ls

e69de29bb2d1      0B README.md
2185d06fcfb7     96B main.go
ed8638311058     97B README.md
d6ace8ce3633    163B go.mod
f61f43467105    245B internal/exec/shell.go
8cfd8a611c15    393B cmd/list.go
071bca3f6ee9    406B internal/exec/shellout.go
cfbf5ac16a25    485B pkg/git/ls.go
94fc347c61c1    746B cmd/list.go
0cdd1894da3f    879B pkg/git/rm.go
b5c1c2cab53c  1.3KiB cmd/rm.go
b4584143d065  1.4KiB cmd/root.go
ffff916a4a42  3.9KiB go.sum
```