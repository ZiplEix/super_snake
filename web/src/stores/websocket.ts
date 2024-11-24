import { writable } from "svelte/store";

export const socket = writable<WebSocket | null>(null);
export const gameID = writable<string>("");
export const connectionStatus = writable<string>("");

export function joinGame(gameID: string, baseApiUrl: string, messageHandler: ((event: MessageEvent) => any) | null) {
    const ws = new WebSocket(`${baseApiUrl}/ws/${gameID}`);

    ws.onopen = () => {
        connectionStatus.set("open");
        console.log("Connexion WebSoket établie");
    };

    // ws.onmessage = (event) => {
    //     console.log("Message reçu du serveur:", event.data);
    // };
    ws.onmessage = messageHandler;

    ws.onerror = (error) => {
        console.error("Erreur WebSocket:", error);
    };

    ws.onclose = () => {
        connectionStatus.set("closed");
        console.log("Connexion WebSocket fermée");
    };

    socket.set(ws);
}

export function closeSocket() {
    socket.update(($socket) => {
        if ($socket) {
            $socket.close();
        }
        return null;
    });
}
