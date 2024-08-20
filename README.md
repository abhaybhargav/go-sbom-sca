# Go Vulnerability Demonstration Application

This project demonstrates supply-chain vulnerabilities in Go applications and showcases how to generate and scan a Software Bill of Materials (SBOM) for security issues.

## Prerequisites

- Docker
- Git (for cloning the repository)

## Setup


2. Build the Docker image:
   ```
   docker build -t go-vuln-demo .
   ```

3. Run the container:
   ```
   docker run -d --name go-vuln-demo-container go-vuln-demo
   ```

## Interacting with the Application

1. Exec into the running container:
   ```
   docker exec -it go-vuln-demo-container /bin/bash
   ```

2. Once inside the container, you can run the following commands:

   a. Generate the CycloneDX SBOM:
      ```
      cyclonedx-gomod mod -output bom.json -json=true
      ```
      This will generate a file named `bom.json` in the current directory.

   b. View the generated SBOM:
      ```
      cat bom.json
      ```

   c. Check the versions of dependencies:
      ```
      go list -m all
      ```

## Scanning for Vulnerabilities


1. Install Grype:
   
   ```
   curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sh -s -- -b /usr/local/bin
   ```

2. Use Grype to scan the SBOM for vulnerabilities:
   ```
   grype bom.json
   ```

## Understanding the Results

The Grype scan will show a list of vulnerabilities found in the dependencies. You should see entries for the intentionally vulnerable packages we included:

- github.com/gin-gonic/gin
- github.com/tidwall/gjson
- golang.org/x/text

Each vulnerability will have details such as severity, CVE ID, and a brief description.