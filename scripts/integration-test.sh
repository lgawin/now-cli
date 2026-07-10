#!/usr/bin/env bash
set -euo pipefail

# 1. Resolve project paths dynamically
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
TARBALL="${REPO_ROOT}/dist/now-linux-x86_64.tar.gz"

echo "🧪 Starting End-to-End Release Artifact Validation..."

# 2. Guard clause: Ensure the tarball was actually built first
if [[ ! -f "$TARBALL" ]]; then
    echo "❌ Error: Release tarball not found at ${TARBALL}." >&2
    echo "Please build and package the project into dist/ first." >&2
    exit 1
fi

echo "🐳 Spinning up clean Ubuntu container for extraction testing..."

# 3. Run a disposable container, mount the tarball, and execute the test routine
docker run --rm -v "${TARBALL}:/tmp/now-linux-x86_64.tar.gz" ubuntu:latest /bin/bash -c "
    set -euo pipefail

    echo '🛠️  Setting up global ADB mock inside container...'
    cat << 'EOF' > /usr/local/bin/adb
#!/usr/bin/env bash
if [[ \"\$1\" == \"get-state\" ]]; then echo \"device\"; exit 0; fi
if [[ \"\$1\" == \"shell\" ]]; then echo \"Mock Android Shell Executed: \${*:2}\"; exit 0; fi
if [[ \"\$1\" == \"root\" ]]; then exit 0; fi
exit 0
EOF
    chmod +x /usr/local/bin/adb

    echo '📦 Extracting release artifact into isolated workspace...'
    TEST_DIR=\$(mktemp -d)
    tar -xzvf /tmp/now-linux-x86_64.tar.gz -C \"\$TEST_DIR\"

    cd \"\$TEST_DIR\"

    echo '🔍 Verifying artifact footprint and permissions...'
    ls -la

    echo '🚀 Executing co-located adb-time-sync orchestration script...'
    ./adb-time-sync

    echo '✅ Containerized execution sequence completed successfully!'
"

echo "🎉 All integration pipelines passed!"
