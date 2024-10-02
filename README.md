# 💀 Dead Man Trigger or THE KILLA 
> Gist-based cloaked dead man trigger

## Why do you need it?
Mainly it's for security reasons: automatically shutting down, locking, or erasing sensitive data if regular check-ins or commands aren’t received, preventing unauthorized access.

## How to install it?
Easy way: 
```bash
curl https://i.jpillora.com/m1ome/killa@1.0! | bash
```

Simple, but you need to have `Go` installed:
```bash
go get github.com/m1ome/killa
```

Building from source:
```bash
mkdir -p /source/killa
cd /source/killa
git clone github.com/m1ome/killa .
go build -o killa .
```

## How to use it?
You can use it how you wanna to, but i recommend using bulletproof solution with **Crontab**, first open
crontab with `crontab -e` afterward, please add new rule:

```bash
* * * * * killa -gist d9ed18e162f33c64236dd4f936dfcfff -codeword Testing -cmd "/bin/echo 'Hallo!'" 2>&1
```

And thats it! Everyting will work now, crossing fingers!

