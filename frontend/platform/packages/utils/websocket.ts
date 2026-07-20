export class WebSocketManager {
  private socket: WebSocket | null = null;
  private url: string;
  private listeners: Set<(data: any) => void> = new Set();
  private autoReconnect: boolean = true;

  constructor(url: string) {
    this.url = url;
  }

  connect() {
    try {
      this.socket = new WebSocket(this.url);

      this.socket.onmessage = (event) => {
        const parsed = JSON.parse(event.data);
        this.listeners.forEach((listener) => listener(parsed));
      };

      this.socket.onclose = () => {
        if (this.autoReconnect) {
          setTimeout(() => this.connect(), 5000);
        }
      };
    } catch {
      // Catch connection failure in tests/node
    }
  }

  subscribe(listener: (data: any) => void): () => void {
    this.listeners.add(listener);
    return () => {
      this.listeners.delete(listener);
    };
  }

  send(data: any) {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(data));
    }
  }

  disconnect() {
    this.autoReconnect = false;
    this.socket?.close();
  }
}
export const ws = new WebSocketManager('ws://localhost:8122/ws');
