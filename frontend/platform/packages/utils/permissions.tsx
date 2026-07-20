import React, { createContext, useContext, useState } from 'react';

interface PermissionsContextType {
  roles: string[];
  hasPermission: (requiredRole: string) => boolean;
  setRoles: (roles: string[]) => void;
}

const PermissionsContext = createContext<PermissionsContextType | undefined>(undefined);

export const PermissionsProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [roles, setRolesState] = useState<string[]>([]);

  const hasPermission = (requiredRole: string) => {
    return roles.includes(requiredRole) || roles.includes('ADMIN');
  };

  return (
    <PermissionsContext.Provider value={{ roles, hasPermission, setRoles: setRolesState }}>
      {children}
    </PermissionsContext.Provider>
  );
};

export const usePermissions = () => {
  const context = useContext(PermissionsContext);
  if (!context) {
    throw new Error('usePermissions must be used within a PermissionsProvider');
  }
  return context;
};
