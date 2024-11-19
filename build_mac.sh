#!/bin/bash
# Script de build pour macOS et Windows
echo "Début des compilations..."

# Supprimer les binaires existants
rm -f app_mac app_mac_arm app_windows.exe

# Compilation pour macOS (Intel)
echo "Compilation pour macOS (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o app_mac main.go
if [ $? -ne 0 ]; then
  echo "Erreur lors de la compilation pour macOS (Intel)."
  exit 1
fi
echo "macOS (Intel) build terminé : app_mac"

# Compilation pour macOS (Apple Silicon)
echo "Compilation pour macOS (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o app_mac_arm main.go
if [ $? -ne 0 ]; then
  echo "Erreur lors de la compilation pour macOS (Apple Silicon)."
  exit 1
fi
echo "macOS (Apple Silicon) build terminé : app_mac_arm"
