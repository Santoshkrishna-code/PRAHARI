export const config = {
  apiBaseUrl: (import.meta as any).env?.VITE_API_BASE_URL || 'http://localhost:8100/api',
  websocketUrl: (import.meta as any).env?.VITE_WS_URL || 'ws://localhost:8122/ws',
  environment: (import.meta as any).env?.MODE || 'development'
};
