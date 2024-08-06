# default.nix
{ pkgs ? import <nixpkgs> {} }:

let
  # Define your own maintainer object
  maintainers = {
    your_name = {
      name = "Sonu Jose";
      email = "sonujse@outlook.com";
    };
  };
in

pkgs.stdenv.mkDerivation rec {
  pname = "logparser";
  version = "1.0";

  src = ./.;

  buildInputs = [ pkgs.go ];

  buildPhase = ''
    export GOCACHE=$(pwd)/.cache/go-build
    mkdir -p $GOCACHE
    go build -o logparser main.go
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp logparser $out/bin/
  '';

  meta = {
    description = "A simple log parser in Go";
    license = pkgs.lib.licenses.mit;
    maintainers = [ pkgs.lib.maintainers.your_name ];
  };
}
