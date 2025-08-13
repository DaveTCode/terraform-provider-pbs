#!/bin/bash

set -e

echo "🐳 Building Docker image for docs generation..."
docker build -f Dockerfile.docs -t terraform-provider-pbs-docs .

echo "📝 Generating documentation in clean Docker environment..."
docker run --rm -v "$(pwd)/docs:/output" terraform-provider-pbs-docs sh -c "cp -r /docs/* /output/"

echo "✅ Documentation generated successfully!"
echo "📁 Check the docs/ directory for updated files."
