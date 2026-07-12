# Getting Started

Build the CLI and create a generated environment:

```bash
cd cli
go build -o cf .
cd ..

./cli/cf project init demo
./cli/cf env create dev --cloud aws --orchestrator eks --region eu-central-1
./cli/cf generate dev
./cli/cf init dev
./cli/cf plan dev
```

No default command applies infrastructure automatically.

For a credential-free local starter project using an existing Kubernetes
context (including kind), use:

```bash
cf --non-interactive wizard --defaults
cf generate dev
cf init dev
cf plan dev --out .cf/plans/dev.tfplan
```

Run `cf local create kind` first only when you want the CLI to create a local
kind cluster. Review the generated provider context before applying workloads.
