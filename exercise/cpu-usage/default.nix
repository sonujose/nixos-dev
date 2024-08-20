{ pkgs ? import <nixpkgs> {} }:

let
  # Define your own maintainer object
  maintainers = {
    sonu = {
      name = "Sonu Jose";
      email = "sonu@something.com";
    };
  };
in

# Nix derivation
pkgs.stdenv.mkDerivation rec {
  pname = "cpuusage";
  version = "1.0";

  # The source code directory
  src = ./.;

  # Dependencies for building the application
  buildInputs = [ pkgs.go ];

  # Commands to build the application
  buildPhase = ''
    export HOME=$(pwd)
    export GOCACHE=$(pwd)/.cache/go-build
    mkdir -p $GOCACHE
    go build -o cpuusage main.go
  '';

  # Commands to install the built application
  installPhase = ''
    mkdir -p $out/bin
    cp cpuusage $out/bin/
  '';

  # Metadata about the package
  meta = {
    description = "Log system cpu usage";
    license = pkgs.lib.licenses.mit;
    maintainers = [ pkgs.lib.maintainers.sonu ];
  };
}