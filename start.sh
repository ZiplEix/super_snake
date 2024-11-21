#! /bin/sh

dockerCompose() {
    docker compose up --build --detach
}

api() {
    cd api || exit 1
    air
}

web() {
    cd web || exit 1
    npm run dev -- --open
}

usage() {
    echo "Usage: $0 [docker|api|web]..."
    exit 1
}

if [ "$#" -eq 0 ]; then
    usage
fi

for arg in "$@"; do
    case $arg in
        docker)
            dockerCompose
            ;;
        api)
            api
            ;;
        web)
            web
            ;;
        *)
            echo "Commande inconnue : $arg"
            echo "Usage: $0 [docker|api|web] ..."
            exit 1
            ;;
    esac
done
