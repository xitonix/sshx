# SSHX

When working in an environment with many hosts with variable IP addresses, and most likely different SSH key files per host, it's impossible to configure SSH to find the mapping between hosts and identity files and as of today, `ssh` does not have a built-in machanism to specify the home for the identity files. 

A good example of such an environment is when we use infrastructure as service in Amazon AWS. Every service host has usually got its own dedicated SSH key and the IPs keep changing with every deployment. 

This tool is a very simple wrapper around ssh to simplify loading identity files without being in the directory where you keep the SSH keys.

## Installation


```shell
go get -u github.com/xitonix/sshx

# Add the following lines to your shell init script (ie. ~/.zshrc or ~/.bashrc)
export SSH_IDENTITY_HOME="~/.ssh" # Or any path where you keep your keys
alias ssh='sshx'
alias scp='sshx -scp'
```

# Usage

Once you installed the tool, you can simply use `ssh` or `scp` the normal way.

```shell
ssh -i "identity.pem" user@host
scp -i "identity.pem" user@host:/src/file /dest/file
```





