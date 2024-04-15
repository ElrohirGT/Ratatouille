{
  description = "Ratatouille flake for reproducible environments and builds!";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:nix-systems/default";
    devenv = {
      url = "github:cachix/devenv";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = {
    self,
    nixpkgs,
    devenv,
    systems,
  } @ inputs: let
    forEachSystem = nixpkgs.lib.genAttrs (import systems);
    dbImageName = "ratatouille_db_image";
    dbImageTag = "current";
    postgresPort = 5566;
    postgresHost = "127.0.0.1";
  in {
    packages = forEachSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        config.allowUnfree = true;
      };
      dervFromDBFile = file: pkgs.writeTextDir "docker-entrypoint-initdb.d/${file}" (builtins.readFile ./db/${file});
      dbInitFile = dervFromDBFile "tables.sql";
    in {
      devenv-up = self.devShells.${system}.default.config.procfileScript;
      dbDocker = pkgs.dockerTools.buildImage {
        name = dbImageName;
        tag = dbImageTag;
        fromImage = pkgs.dockerTools.pullImage {
          imageName = "postgres";
          # Obtained using `nix run nixpkgs#nix-prefetch-docker -- --image-name postgres --image-tag 16`
          imageDigest =
            if system == "aarch64-darwin"
            then "sha256-dXZo5CaKobuKRUFS3FUgjN2jnBrmQ9do7+815lEZ2mo=" # AARM Mac (M1 Mac)
            else "sha256:f58300ac8d393b2e3b09d36ea12d7d24ee9440440e421472a300e929ddb63460"; # x64 Mac and Linux
          sha256 = "1dpmibx8llrnsa9slq8cvx2r7ppsicxxf6kwaz00lnyvp9hwhs0k";
          finalImageTag = "16";
        };

        copyToRoot = pkgs.buildEnv {
          name = "image-root";
          paths = [dbInitFile];
          pathsToLink = ["/docker-entrypoint-initdb.d"];
        };

        config.Entrypoint = "/usr/local/bin/docker-entrypoint.sh";
        config.Cmd = ["postgres"];
        config.Env = [
          "POSTGRES_PASSWORD=myPassword"
        ];
      };
      restartServices = pkgs.writeShellApplication {
        name = "services_restarter";
        runtimeInputs = with pkgs; [ansi];
        text = ''
          echo -e "$(ansi yellow)"WARNING:"$(ansi reset)" This script must be run on the project root directory!

          echo "Trying to remove old .devenv..."
          rm ./.devenv/state/postgres/ || rm -r ./.devenv/state/postgres/ || true

          echo "Entering devshell..."
          nix develop --impure . -c devenv up
        '';
      };
    });

    devShells = forEachSystem (
      system: let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };
        strFromDBFile = file: builtins.readFile ./db/${file};
        dbInitFile = builtins.concatStringsSep "\n" [
          (strFromDBFile "setup.sql")
          (strFromDBFile "tables.sql")
          (strFromDBFile "inserts.sql")
          (strFromDBFile "triggers.sql")
          (strFromDBFile "indexes.sql")
        ];
      in {
        default = devenv.lib.mkShell {
          inherit inputs pkgs;
          modules = [
            {
              packages = with pkgs; [
                sqlfluff
                go
                gnumake
                sqlc
              ];

              # Enable .env integration
              dotenv.enable = true;

              services.postgres = {
                enable = true;
                listen_addresses = postgresHost;
                port = postgresPort;
                initialScript = dbInitFile;
                settings = {
                  log_connections = true;
                  log_statement = "all";
                  logging_collector = true;
                  log_disconnections = true;
                  log_destination = "stderr";
                };
              };
            }
          ];
        };
      }
    );
  };
}
