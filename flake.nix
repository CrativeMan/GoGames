{
description = "Basic Go Flake";

inputs = {
  flake-compat = {
    url = "github:edolstra/flake-compat";
    flake = false;
  };
  flake-utils.url = "github:numtide/flake-utils";
  nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
};

outputs = {
  self,
  flake-utils,
  nixpkgs,
  ...
}:
  flake-utils.lib.eachDefaultSystem (system: let
    pkgs = nixpkgs.legacyPackages.${system};
  in {
    devShell = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        gofumpt
        goimports-reviser
        golines
        delve

        wayland
        wayland-protocols
        glew
        glfw       
        libxkbcommon
        xorg.libX11 
        xorg.libXcursor
        xorg.libXi
        xorg.libXrandr
        xorg.libXinerama
      ];
    };
  });
}
