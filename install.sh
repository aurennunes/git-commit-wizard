#!/bin/bash

OS=$(uname -s)
ARCH=$(uname -m)

if [[ "$OS" == "Linux" ]]; then
    BINARY="commit-wizard-linux-amd64"
elif [[ "$OS" == "Darwin" ]]; then
    BINARY="commit-wizard-macos-amd64"
elif [[ "$OS" == "MINGW"* || "$OS" == "CYGWIN"* ]]; then
    BINARY="commit-wizard-windows-amd64.exe"
else
    echo "Sistema não suportado!"
    exit 1
fi

REPO_URL="https://github.com/aurennunes/git-commit-wizard/releases/download/v1.0.0/$BINARY"

echo "Baixando CommitWizard para $OS..."
curl -L -o /usr/local/bin/commit-wizard "$REPO_URL"

chmod +x /usr/local/bin/commit-wizard

echo "CommitWizard instalado com sucesso!"
