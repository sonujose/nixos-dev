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
  pname = "logparserdir";
  version = "1.0";

  src = ./.;

  buildInputs = [ pkgs.go ];

  buildPhase = ''
    export GOCACHE=$(pwd)/.cache/go-build
    mkdir -p $GOCACHE
    go build -o logparserdir main.go
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp logparserdir $out/bin/
  '';

  meta = {
    description = "A simple log parser in Go";
    license = pkgs.lib.licenses.mit;
    maintainers = [ pkgs.lib.maintainers.your_name ];
  };
}
