# NixOs Experiments

## Setup NixOs on Local

```sh
curl -L https://nixos.org/nix/install | sh
nix --version
```

### Ad hoc shell environments

In a Nix shell environment, you can immediately use any program packaged with Nix, without installing it permanently.

```sh
nix-shell -p cowsay lolcat

cowsay Hello, Nix! | lolcat

```

Running programs once, You can go even faster, by running any program directly:

```sh
~/Documents/sonu/nixos » nix-shell -p cowsay --run "cowsay Nix"                                                                                                                                                                    sonaojus@Mac01
 _____
< Nix >
 -----
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

### Nested shell sessions

If you need an additional program temporarily, you can run a nested Nix shell. The programs provided by the specified packages will be added to the current environment.

### Towards reproducibility
These shell environments are very convenient, but the examples so far are not reproducible yet. Running these commands on another machine may fetch different versions of packages, depending on when Nix was installed there.

What do we mean by reproducible? A fully reproducible example would give exactly the same results no matter when or where you run the command. The environment provided would be identical each

```s
nix-shell -p git --run "git --version" --pure -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/2a601aafdc5605a5133a2ca506a34a3a73377247.tar.gz
```

If you’re done trying out Nix for now, you may want to free up some disk space occupied by the different versions of programs you downloaded by running the examples:

```
nix-collect-garbage
```

## Reproducible interpreted scripts

```s
#!/usr/bin/env nix-shell
#! nix-shell -i bash --pure
#! nix-shell -p bash cacert curl jq python3Packages.xmljson
#! nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/2a601aafdc5605a5133a2ca506a34a3a73377247.tar.gz

curl https://github.com/NixOS/nixpkgs/releases.atom | xml2json | jq .
```

The key point is that all the #! lines are processed by nix-shell before executing the script itself. The -I option sets the NIX_PATH before the environment is built, ensuring that the correct version of nixpkgs is used when resolving the packages listed with -p.

So, when nix-shell reads the script, it:

Sets the NIX_PATH with the -I option.
Fetches and evaluates the packages listed with -p from the specified nixpkgs archive URL.
Starts a bash shell in a pure environment with the specified packages available.

## Declarative shell environments with shell.nix

```s
let
  nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
  pkgs = import nixpkgs { config = {}; overlays = []; };
in

pkgs.mkShellNoCC {
  packages = with pkgs; [
    cowsay
    lolcat
  ];
}
```

## Startup commands
You may want to run some commands before entering the shell environment. These commands can be placed in the shellHook attribute provided to mkShellNoCC.

Set shellHook to output a colorful greeting:

```sh

 let
   nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-24.05";
   pkgs = import nixpkgs { config = {}; overlays = []; };
 in

 pkgs.mkShellNoCC {
   packages = with pkgs; [
     cowsay
     lolcat
   ];

   GREETING = "Hello, Nix!";

   shellHook = ''
    echo $GREETING | cowsay | lolcat
   '';
 }

```s
{ pkgs ? import <nixpkgs> {} }:

```
nix-shell

### Evaluating Nix files
Use nix-instantiate --eval to evaluate the expression in a Nix file.

echo 1 + 2 > file.nix
nix-instantiate --eval file.nix
3

### Attribute set

Attribute sets and let expressions are used to assign names to values. Assignments are denoted by a single equal sign (=).

Whenever you encounter an equal sign (=) in Nix language code:

On its left is the assigned name.

On its right is the value, delimited by a semicolon (;).

Attribute set { ... }
An attribute set is a collection of name-value-pairs, where names must be unique.

The following example shows all primitive data types, lists, and attribute sets.

## let ... in ...
Also known as “let expression” or “let binding”

let expressions allow assigning names to values for repeated use.

```sh
let
  b = a + 1;
  a = 1;
in
a + b
```

with ...; ...
The with expression allows access to attributes without repeatedly referencing their attribute set.

Example:
```sh
let
  a = {
    x = 1;
    y = 2;
    z = 3;
  };
in
with a; [ x y z ]
```

[ 1 2 3 ]
The expression

with a; [ x y z ]
is equivalent to

[ a.x a.y a.z ]

## inherit ...
inherit is shorthand for assigning the value of a name from an existing scope to the same name in a nested scope. It is for convenience to avoid repeating the same name multiple times.

## String interpolation ${ ... }
Previously known as “antiquotation”.

The value of a Nix expression can be inserted into a character string with the dollar-sign and braces (${ }).

Example:
```sh
let
  name = "Nix";
in
"hello ${name}"
```

## Lookup paths
Also known as “angle bracket syntax”.

Example:

<nixpkgs>
/nix/var/nix/profiles/per-user/root/channels/nixpkgs
The value of a lookup path is a file system path that depends on the value of builtins.nixPath.

In practice, <nixpkgs> points to the file system path of some revision of Nixpkgs.

For example, <nixpkgs/lib> points to the subdirectory lib of that file system path:

<nixpkgs/lib>
/nix/var/nix/profiles/per-user/root/channels/nixpkgs/lib

## Functions

```s
let
 f = x: x + 1;
 a = 1;
in [ f a ]
```

## Multiple arguments
Also known as “curried functions”.

Nix functions take exactly one argument. Multiple arguments can be handled by nesting functions.

Such a nested function can be used like a function that takes multiple arguments, but offers additional flexibility.

## Impurities
So far we have only covered what we call pure expressions: declaring data and transforming it with functions.

In practice, describing derivations requires observing the outside world.

There is only one impurity in the Nix language that is relevant here: reading files from the file system as build inputs

Build inputs are files that derivations refer to in order to describe how to derive new files. When run, a derivation will only have access to explicitly declared build inputs.

The only way to specify build inputs in the Nix language is explicitly with:

- File system paths
- Dedicated functions.

Nix and the Nix language refer to files by their content hash. If file contents are not known in advance, it’s unavoidable to read files during expression evaluation.

## Paths
Whenever a file system path is used in string interpolation, the contents of that file are copied to a special location in the file system, the Nix store, as a side effect.

The evaluated string then contains the Nix store path assigned to that file.

Example:
```sh
echo 123 > data
"${./data}"
"/nix/store/h1qj5h5n05b5dl5q4nldrqq8mdg7dhqk-data"

```

## Fetchers
Files to be used as build inputs do not have to come from the file system.

The Nix language provides built-in impure functions to fetch files over the network during evaluation:

```s
builtins.fetchurl

builtins.fetchTarball

builtins.fetchGit

builtins.fetchClosure
```
These functions evaluate to a file system path in the Nix store.

Example:

```sh
builtins.fetchurl "https://github.com/NixOS/nix/archive/7c3ab5751568a0bc63430b33a5169c5e4784a0ff.tar.gz"
"/nix/store/7dhgs330clj36384akg86140fqkgh8zf-7c3ab5751568a0bc63430b33a5169c5e4784a0ff.tar.gz"
```

Some of them add extra convenience, such as automatically unpacking archives.

Example:
```sh
builtins.fetchTarball "https://github.com/NixOS/nix/archive/7c3ab5751568a0bc63430b33a5169c5e4784a0ff.tar.gz"
"/nix/store/d59llm96vgis5fy231x6m7nrijs0ww36-source"
```

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

