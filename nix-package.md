## Hacking Your First Package - NixOS

### Example: Overriding a Package

Overriding a package in NixOS allows you to customize the build of an existing package. This can be useful if you need to change the version, modify build inputs, or adjust configuration options.

Let's say you want to override the `hello` package to change its version and add a build input. Here’s a step-by-step guide:

### Step 1: Find the Original Package

First, locate the original package definition in the Nixpkgs repository. You can search for the package using:

```sh
nix search nixpkgs hello
```

### Step 2: Create an Override

You can create an override by using the `overrideAttrs` function to modify the package's attributes. Here’s how you can do it in a Nix expression.

#### Create a `default.nix` File

Create a `default.nix` file with the following content:

```nix
# default.nix
{ pkgs ? import <nixpkgs> {} }:

let
  # Original package
  originalHello = pkgs.hello;

  # Override the package
  myHello = originalHello.overrideAttrs (oldAttrs: rec {
    version = "2.10"; # Change the version

    src = pkgs.fetchurl {
      url = "http://ftp.gnu.org/gnu/hello/hello-${version}.tar.gz";
      sha256 = "0ssi1wpaf7plaswqqjwigppsg5fyh99vdlb9kzl7c9lng89ndssi";
    };

    buildInputs = oldAttrs.buildInputs ++ [ pkgs.curl ]; # Add a build input
  });
in
{
  hello = myHello;
}
```

### Step 3: Build the Overridden Package

To build the overridden package, run:

```sh
nix-build default.nix -A hello
```

This will build the `hello` package with the modifications specified in the override.

### Step 4: Use the Overridden Package in NixOS Configuration

If you want to use the overridden package in your NixOS configuration, you need to modify your `configuration.nix`.

#### Edit `configuration.nix`

Add the override to your NixOS configuration:

```nix
{ config, pkgs, ... }:

let
  # Import the custom override
  myPkgs = import /path/to/your/default.nix;
in
{
  # Use the overridden package
  environment.systemPackages = with pkgs; [
    myPkgs.hello
  ];
}
```

### Step 5: Apply the Configuration

To apply the NixOS configuration with the overridden package, run:

```sh
sudo nixos-rebuild switch
```

### Summary

1. **Locate the Original Package**: Find the package you want to override in the Nixpkgs repository.
2. **Create an Override**: Use the `overrideAttrs` function to modify the package's attributes in a `default.nix` file.
3. **Build the Overridden Package**: Use `nix-build` to build the package with the modifications.
4. **Use in NixOS Configuration**: Modify your `configuration.nix` to use the overridden package.
5. **Apply the Configuration**: Run `nixos-rebuild switch` to apply the changes.

This process allows you to customize Nix packages to fit your specific requirements.

