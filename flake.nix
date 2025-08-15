{
  description = "A CLI tool for various utilities";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      packages = rec {
        default = pkgs.buildGoModule {
          pname = "serein";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-PYoO3JMlIbtF8sHm+pO2RQN6nJKIc001toGY7/b+t0I=";
        };
      };

      apps = rec {
        default = flake-utils.lib.mkApp { drv = self.packages.${system}.default; };
      };
    });
}
