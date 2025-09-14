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
      src = pkgs.lib.cleanGit {
        inherit (self) src;
        name = "serein";
      };
      version = pkgs.lib.escapeShellArg (pkgs.lib.fileContents "${src}/.version");

    in {
      packages = rec {
        default = pkgs.buildGoModule {
          pname = "serein";
          inherit version;
          # Point the src to the cleaned Git source
          src = src;
          vendorHash = "sha256-+gNaABMs7XZbOFlvLQA5KtnZrBHDWgBtH6W29KMeBU0=";
          ldflags = [
            "-s"
            "-w"
            "-X main.version=${version}"
          ];
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
