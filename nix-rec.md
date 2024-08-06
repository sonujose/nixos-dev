## Nix rec - recursive

In Nix, `rec` is a keyword that stands for "recursive" and is used to define a set of attributes where the attributes can reference each other. This is particularly useful in situations where you need to refer to other attributes within the same set.

When you see `rec` in the context of `pkgs.stdenv.mkDerivation`, it allows you to define a derivation where the attributes can be self-referential. This is useful for organizing and structuring complex derivations.

### Example: Understanding `rec`

Consider the following example without `rec`:

```s
{ stdenv, fetchurl }:

stdenv.mkDerivation {
  pname = "hello";
  version = "2.10";

  src = fetchurl {
    url = "http://ftp.gnu.org/gnu/hello/hello-${version}.tar.gz";
    sha256 = "0ssi1wpaf7plaswqqjwigppsg5fyh99vdlb9kzl7c9lng89ndssi";
  };

  meta = {
    description = "A program that prints 'Hello, world!'";
    license = stdenv.lib.licenses.gpl3;
  };
}
```

In this example, you cannot directly reference `version` inside the `src` attribute because `version` is not yet defined in the scope of `src`.

Now, consider the same example with `rec`:

```s
{ stdenv, fetchurl }:

stdenv.mkDerivation rec {
  pname = "hello";
  version = "2.10";

  src = fetchurl {
    url = "http://ftp.gnu.org/gnu/hello/hello-${version}.tar.gz";
    sha256 = "0ssi1wpaf7plaswqqjwigppsg5fyh99vdlb9kzl7c9lng89ndssi";
  };

  meta = {
    description = "A program that prints 'Hello, world!'";
    license = stdenv.lib.licenses.gpl3;
  };
}
```

In this example, the use of `rec` allows `src` to reference the `version` attribute. Here’s a breakdown of what happens:

- rec: The `rec` keyword makes the attribute set recursive, meaning you can reference any attribute within the same set.
- Self-Referential Attributes: Inside the `rec` set, you can reference other attributes defined in the same set. For instance, `version` is used inside `src` to construct the URL.

### Explanation in the Context of `mkDerivation`

When used with `pkgs.stdenv.mkDerivation`, `rec` allows you to:

1. Organize Complex Derivations: Reference attributes within the derivation, making it easier to structure complex derivations.
2. Self-Referential Logic: Use attributes to define other attributes, such as constructing URLs based on the `version` attribute.
3. Clean and Maintainable Code: Keep related configuration logic together and avoid duplication by reusing attributes.

### Practical Example

Here’s a practical example showing the benefits of using `rec` in a derivation:

```s
{ pkgs }:

pkgs.stdenv.mkDerivation rec {
  pname = "mysoftware";
  version = "1.0.0";

  src = pkgs.fetchurl {
    url = "https://example.com/${pname}-${version}.tar.gz";
    sha256 = "0mifnpwshf9xm9cw96qvmzp63k2n7mfh5l9i4lbf6bzjlzz3kxb3";
  };

  buildInputs = [ pkgs.makeWrapper ];

  buildPhase = ''
    echo "Building ${pname} version ${version}"
    # build commands here
  '';

  installPhase = ''
    mkdir -p $out/bin
    cp mysoftware $out/bin/
  '';

  meta = {
    description = "A custom software package";
    license = pkgs.lib.licenses.mit;
  };
}
```

### Summary

- `rec` Keyword: Stands for "recursive" and allows attributes within a set to reference each other.
- Usage in `mkDerivation`: Helps organize complex derivations and makes them easier to maintain by allowing self-referential attributes.
- Practical Benefits: Enables cleaner, more maintainable code by avoiding duplication and keeping related configuration logic together.

By using `rec` in your Nix expressions, you can write more flexible and maintainable configurations, especially when dealing with complex derivations.