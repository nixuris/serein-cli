{
  description = "An opinionated CLI wrapper that replaces cryptic flags with self-explanatory, English-like sub-commands.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
      # Get the version from the Git tag or commit hash
      version = self.rev or "dirty";
    in {
      packages = rec {
        default = pkgs.buildGoModule {
          pname = "serein";
          inherit version; # Use the dynamically set version
          src = ./.;
          vendorHash = "sha256-+gNaABMs7XZbOFlvLQA5KtnZrBHDWgBtH6W29KMeBU0=";
          nativeBuildInputs = [pkgs.installShellFiles];
          postFixup = ''
            installShellCompletion --fish ${./completions/serein.fish}
            installShellCompletion --zsh  ${./completions/serein.zsh}
            installShellCompletion --bash ${./completions/serein.bash}
          '';
        };
      };

      apps = rec {
        default = flake-utils.lib.mkApp {drv = self.packages.${system}.default;};
      };
    });
}
