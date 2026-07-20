import React, { createContext, useContext, useState } from 'react';

interface AuthContextType {
  isAuthenticated: boolean;
  token: string | null;
  tenantId: string | null;
  login: (token: string, tenantId: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem('prahari-jwt-token'));
  const [tenantId, setTenantId] = useState<string | null>(() => localStorage.getItem('prahari-tenant-id'));

  const login = (jwt: string, tenant: string) => {
    localStorage.setItem('prahari-jwt-token', jwt);
    localStorage.setItem('prahari-tenant-id', tenant);
    setToken(jwt);
    setTenantId(tenant);
  };

  const logout = () => {
    localStorage.removeItem('prahari-jwt-token');
    localStorage.removeItem('prahari-tenant-id');
    setToken(null);
    setTenantId(null);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated: !!token, token, tenantId, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
