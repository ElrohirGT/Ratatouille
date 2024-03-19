{
  description = "Ratatouille flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = {
    self,
    nixpkgs,
  }: let
    forAllSystems = function:
      nixpkgs.lib.genAttrs [
        "x86_64-linux"
        "x86_64-macos"
        "aarch64-linux"
        "aarch64-darwin"
      ] (system:
        function {
          pkgs = import nixpkgs {
            inherit system;
            config.allowUnfree = true;
            overlays = [
              #inputs.something.overlays.default
            ];
          };
          system = system;
        });
    dbImageName = "ratatouille_db_image";
    dbImageTag = "current";
    dbContainerName = "ratatouille_postgres";
  in {
    packages = forAllSystems ({
      pkgs,
      system,
    }: {
      dbDocker = pkgs.dockerTools.buildImage {
        name = dbImageName;
        tag = dbImageTag;
        created = "now";
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

        copyToRoot = let
          makeDerFromFile = file: pkgs.writeTextDir "docker-entrypoint-initdb.d/${file}" (builtins.readFile ./db/${file});
          dbInitscript = makeDerFromFile "tables.sql";
        in
          pkgs.buildEnv {
            name = "image-root";
            paths = [dbInitscript];
            pathsToLink = ["/docker-entrypoint-initdb.d"];
          };

        config.Entrypoint = "/usr/local/bin/docker-entrypoint.sh";
        config.Cmd = ["postgres"];
        config.Env = [
          "POSTGRES_PASSWORD=myPassword"
        ];
      };
      restartDBDocker = let
        appName = "restartRatatouilleDBDocker";
      in
        pkgs.writeShellApplication {
          name = appName;
          text = ''
            echo WARNING: This command should be executed inside the root of the project!
            docker stop ${dbContainerName} || true # Ignore error, since we can't guarantee the container was running
            docker rm ${dbContainerName} || true # Ignore error, since we can't guarantee the container already existed before
            docker rmi ${dbImageName}:${dbImageTag} || true # Ignore error, since we can't guarantee the image exists

            if nix build .#dbDocker && docker load < result; then
            	set -o allexport
            	# shellcheck disable=SC1091
            	. .env
            	set +o allexport
            	docker run -d --name ${dbContainerName} -p "''$DB_PORT":5432 ${dbImageName}:${dbImageTag}
            fi
          '';
        };
    });

    devShells = forAllSystems ({
      pkgs,
      system,
    }: {
      default = pkgs.mkShell {
        packages = with pkgs; [
          sqlfluff
        ];
      };
    });
  };
}
