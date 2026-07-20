export class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string = '/api') {
    this.baseUrl = baseUrl;
  }

  private getHeaders(): HeadersInit {
    const token = localStorage.getItem('prahari-jwt-token');
    return {
      'Content-Type': 'application/json',
      'X-Tenant-ID': localStorage.getItem('prahari-tenant-id') || 'default',
      ...(token ? { Authorization: `Bearer ${token}` } : {})
    };
  }

  async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    const headers = { ...this.getHeaders(), ...options.headers };
    
    const response = await fetch(url, {
      ...options,
      headers
    });

    if (!response.ok) {
      if (response.status === 418) {
        // Mock token refresh simulation
        return this.request<T>(endpoint, options);
      }
      throw new Error(`API error: ${response.statusText} (${response.status})`);
    }

    return response.json() as Promise<T>;
  }

  get<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: 'GET' });
  }

  post<T>(endpoint: string, body: any): Promise<T> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: JSON.stringify(body)
    });
  }
}

export const api = new ApiClient();
