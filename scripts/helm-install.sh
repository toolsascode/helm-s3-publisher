#!/usr/bin/env sh

PROJECT_NAME="helm-s3-publisher"
PROJECT_GH="toolsascode/$PROJECT_NAME"
# GitHub's URL for the latest release, will redirect.
GITHUB_BASE_URL="https://github.com/${PROJECT_GH}"
# Run the script in a temporary directory that we know is empty.
SCRATCH=$(mktemp -d || mktemp -d -t 'tmp')

set -e
set -u

if [ -n "${HELM_S3_PLUGIN_NO_INSTALL_HOOK:-}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

validateCheckSum() {
    if ! grep -q "${1}" "${2}"; then
        echo "Invalid checksum" >/dev/stderr
        exit 1
    fi
    echo "Checksum is valid."
}

getArch() {
    arch=$(uname -m)

    case $arch in
    x86_64 | amd64)
        arch="amd64"
        ;;
    aarch64 | arm64)
        arch="arm64"
        ;;
    *)
        echo "Arch '$(uname -m)' not supported!" >&2
        exit 1
        ;;
    esac

}

getOS() {
    os=$(uname -s)
    case "$(uname)" in
    Darwin)
        os="darwin"
        ;;
    Linux)
        os="linux"
        ;;
    CYGWIN* | MINGW* | MSYS_NT*)
        os="windows"
        ;;
    *)
        echo "OS '$(uname)' not supported!" >&2
        exit 1
        ;;
    esac
}

onExit() {
    exit_code=$?
    if [ ${exit_code} -ne 0 ]; then
        echo "${PROJECT_NAME} install hook failed. Please remove the plugin using 'helm plugin remove s3' and install again." >/dev/stderr
    fi
    rm -rf "releases"
    # Delete the working directory when the install was successful.
    rm -r "$SCRATCH"
    exit ${exit_code}
}

trap onExit EXIT

version="$(grep "version" plugin.yaml | cut -d '"' -f 2)"
echo "Downloading and installing ${PROJECT_NAME} v${version} ..."

getArch
getOS

RELEASE_URL="${GITHUB_BASE_URL}/releases/download/v${version}/${PROJECT_NAME}_${version}_${os}_${arch}.tar.gz"
CHECKSUM_URL="${GITHUB_BASE_URL}/releases/download/v${version}/${PROJECT_NAME}_${version}_checksums.txt"

cd "$SCRATCH"

mkdir "releases"
RELEASE_FILE="$SCRATCH/releases/v${version}.tar.gz"
CHECKSUM_FILENAME="$SCRATCH/releases/v${version}_checksums.txt"

# Download binary and checksums files.
(
    if command -v curl >/dev/null 2>&1; then
        curl -sSL "${RELEASE_URL}" -o "${RELEASE_FILE}"
        curl -sSL "${CHECKSUM_URL}" -o "${CHECKSUM_FILENAME}"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "${RELEASE_URL}" -O "${RELEASE_FILE}"
        wget -q "${CHECKSUM_URL}" -O "${CHECKSUM_FILENAME}"
    else
        echo "ERROR: no curl or wget found to download files." >/dev/stderr
    fi
)

# Verify checksum.
(
    if command -v sha256sum >/dev/null 2>&1; then
        checksum=$(sha256sum "${RELEASE_FILE}" | awk '{ print $1 }')
        validateCheckSum "${checksum}" "${CHECKSUM_FILENAME}"
    elif command -v openssl >/dev/null 2>&1; then
        checksum=$(openssl dgst -sha256 "${RELEASE_FILE}" | awk '{ print $2 }')
        validateCheckSum "${checksum}" "${CHECKSUM_FILENAME}"
    else
        echo "WARNING: no tool found to verify checksum" >/dev/stderr
    fi
)

# Unpack the binary.
tar xzf "${RELEASE_FILE}" bin/helm-s3-publisher
