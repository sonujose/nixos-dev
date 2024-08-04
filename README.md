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

```

```
{ pkgs ? import <nixpkgs> {} }:

...
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

Multiple arguments
Also known as “curried functions”.

Nix functions take exactly one argument. Multiple arguments can be handled by nesting functions.

Such a nested function can be used like a function that takes multiple arguments, but offers additional flexibility.

