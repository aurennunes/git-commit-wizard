#!/bin/bash

REPO_OWNER="aurennunes"
REPO_NAME="git-commit-wizard"

latest_release=$(curl --silent "https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

os=""
arch=""
case "$(uname -s)" in
  Linux*) os=linux ;;
  Darwin*) os=macos ;;
  CYGWIN* | MINGW* | MSYS*) os=windows ;;
*)
  echo "Sistema operacional não suportado"
  exit 1
  ;;
esac

# Determina a arquitetura (amd64, arm64, arm)
case "$(uname -m)" in
  x86_64) arch=amd64 ;;
  arm64) arch=arm64 ;;
  armv7l) arch=arm ;;
*)
  echo "Arquitetura não suportada"
  exit 1
  ;;
esac

BINARY="commit-wizard-$os-$arch"
REPO_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$latest_release/$BINARY"

echo "Baixando CommitWizard para $OS..."
curl -L "$REPO_URL" -o /tmp/$BINARY

echo "Instalando CommitWizard em /usr/local/bin..."
sudo mv /tmp/$BINARY /usr/local/bin/git-commit-wizard

sudo chmod +x /usr/local/bin/git-commit-wizard

echo "CommitWizard instalado com sucesso!"
