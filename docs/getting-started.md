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
