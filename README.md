# SSHX

This is a simple wrapper around ssh to simplify loading identity files without being in the directory where you keep the SSH keys.

# Installation



```shell
go get go.xitonix.io/sshx

# Add the following lines to your shell init script (ie. ~/.zshrc or ~/.bashrc)
export SSH_IDENTITY_HOME="~/.ssh" # Or any path where you keep your keys
alias ssh='sshx'
alias scp='sshx'
```



# Usage

Once you installed the tool, you can simply use `ssh` or `scp` the normal way.

```shell
ssh -i "identity.pem" user@host
scp -i "identity.pem" user@host:/src/file /dest/file
```

