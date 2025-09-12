{
  description = "A CLI tool for various utilities";

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
    in {
      packages = rec {
        default = pkgs.buildGoModule {
          pname = "serein";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-+gNaABMs7XZbOFlvLQA5KtnZrBHDWgBtH6W29KMeBU0=";
          # Add installShellFiles to build inputs
          nativeBuildInputs = [pkgs.installShellFiles];
          # Install fish completion directly in flake
          postFixup = ''
            installShellCompletion --fish ${./completions/fish-completion.fish}
            installShellCompletion --zsh  ${./completions/zsh-completion.zsh}
            installShellCompletion --bash ${./completions/bash-completion.bash}
          '';
        };
      };

      apps = rec {
        default = flake-utils.lib.mkApp {drv = self.packages.${system}.default;};
      };
    });
}
