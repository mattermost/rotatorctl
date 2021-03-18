# rotatorctl (Rotator CLI)

rotatorctl (Rotator CLI) which uses as module the [Mattermost Rotator]("github.com/mattermost/rotator/rotator"). Is a tool meant to smooth and accelerate k8s cluster upgrades and node rotations. It offers automation on an autoscaling group recognition and flexibility on options such as, how fast to rotate nodes, drain retries, waiting time between rotations and drains as well as mater/worker node separation.

## How to use

Get or build the release binaries and use it as cli locally or in any pipeline to automate nodes rotation.

#### Building

Simply run the following:

```bash
go install ./cmd/rotatorctl
alias rotatorctl='$HOME/$GOROOT/bin/rotatorctl'
```
or use the `make` shorthands

For Linux binary:
`make build`

For MacOS (darwin) binary:
`make build-mac`

#### Running

To rotate a cluster:
```bash
rotatorctl rotate --cluster <cluster_name>  --wait-between-rotations 30 --wait-between-drains 60 --max-scaling 4 --evict-grace-period 30
```

### Other Setup

For the `rotatorctl` to run access to both the AWS account and the K8s cluster is required to be able to do actions such as, `DescribeInstances`, `DetachInstances`, `TerminateInstances`, `DescribeAutoScalingGroups`, as well as `drain`, `kill`, `evict` pods, etc.

The relevant AWS Access and Secret key pair should be exported and k8s access should be provided via a passed clientset or a locally exported k8s config. 
